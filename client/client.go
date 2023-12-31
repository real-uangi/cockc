// Package client
// @author uangi
// @date 2023/6/30 10:00
package client

import (
	"encoding/json"
	"fmt"
	"github.com/real-uangi/cockc/common/plog"
	"github.com/real-uangi/cockc/config"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

var logger = plog.New("client")

type CockMsg struct {
	Operation string      `json:"operation"`
	Msg       string      `json:"msg"`
	Data      config.Cock `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

const (
	echo       = "echo"
	pullConfig = "pullConfig"
	heartbeat  = "heartbeat"
	online     = "online"
	offline    = "offline"
)

type CockClientService interface {
	Load()
	Echo()
	PullConfig()
	StartHeartbeat()
	Online()
	Offline()
}

type CockClient struct {
	Up            bool
	Config        config.Cock
	HeartbeatLock sync.Mutex
	UdpLock       sync.Mutex
}

func (c *CockClient) Load() {
	c.Config = config.GetPropertiesRO().Cock
	c.Up = false
}

func (c *CockClient) Echo() {
	response := c.dial(echo, c.Config, strconv.Itoa(rand.Intn(1000000)), false)
	logger.Info(fmt.Sprintf("Cock server dealy %d ms msg: %s \n", time.Now().UnixMilli()-response.Timestamp, response.Msg))
}

func (c *CockClient) PullConfig() {
	var cs string
	response := c.dial(pullConfig, c.Config, pullConfig, false)
	if response.Msg == "" {
		logger.Warn("Failed to pull config")
		return
	}
	cs = response.Msg
	config.UpdateServerConfig(cs)
}

func (c *CockClient) StartHeartbeat() {
	go c.heartbeat()
}

func (c *CockClient) heartbeat() {
	if c.HeartbeatLock.TryLock() {
		defer c.HeartbeatLock.Unlock()
		for {
			go beat(c)
			//
			time.Sleep(time.Duration(c.Config.Register.Heartbeat.Interval) * time.Second)
		}
	}
}

func beat(c *CockClient) {
	response := c.dial(heartbeat, c.Config, strconv.Itoa(rand.Intn(1000000)), true)
	if response.Msg != "ok" {
		logger.Warn("failed to send heartbeat")
	}
}

func (c *CockClient) Online() {
	c.dial(online, c.Config, online, false)
	c.Up = true
}

func (c *CockClient) Offline() {
	c.dial(offline, c.Config, offline, false)
	c.Up = false
}

func (c *CockClient) dial(operation string, data config.Cock, serial string, ignoreWhenBlock bool) CockMsg {

	if ignoreWhenBlock {
		if c.UdpLock.TryLock() {
			defer c.UdpLock.Unlock()
		} else {
			return CockMsg{}
		}
	} else {
		c.UdpLock.Lock()
		defer c.UdpLock.Unlock()
	}
	defer func() {
		r := recover()
		if r != nil {
			logger.Error(fmt.Sprintf("%v", r))
		}
	}()
	remoteAddr, err := net.ResolveUDPAddr("udp", c.Config.Register.Address)
	if err != nil {
		panic(err)
	}
	localAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(c.Config.Port))
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		logger.Error(err.Error())
		return CockMsg{}
	}
	msg := CockMsg{
		Operation: operation,
		Data:      data,
		Msg:       serial,
		Timestamp: time.Now().UnixMilli(),
	}
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(b)
	if err != nil {
		panic(err)
	}
	response := make([]byte, 1024)
	size, err := conn.Read(response)
	if err != nil {
		panic(err)
	}
	msg = CockMsg{}
	err = json.Unmarshal(response[:size], &msg)
	if err != nil {
		panic(err)
	}
	return msg
}
