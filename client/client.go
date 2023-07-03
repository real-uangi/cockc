// Package client
// @author uangi
// @date 2023/6/30 10:00
package client

import (
	"encoding/json"
	"fmt"
	"github.com/real-uangi/cockc/config"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

type CockMsg struct {
	Operation string      `json:"operation"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
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
}

func (c *CockClient) Load() {
	c.Config = config.GetPropertiesRO().Cock
	c.Up = false
}

func (c *CockClient) Echo() {
	response := c.dial(echo, strconv.Itoa(rand.Intn(1000000)))
	fmt.Printf("Cock server dealy %d ms msg: %s data: %s\n", time.Now().UnixMilli()-response.Timestamp, response.Msg, response.Data.(string))
}

func (c *CockClient) PullConfig() {
	var cs string

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
			time.Sleep(time.Duration(c.Config.Register.Heartbeat.Interval) * time.Millisecond)
		}
	}
}

func beat(c *CockClient) {
	c.dial(heartbeat, c.Config)
}

func (c *CockClient) Online() {
	c.Up = true
}

func (c *CockClient) Offline() {
	c.Up = false
}

func (c *CockClient) dial(operation string, data interface{}) CockMsg {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	udp, err := net.ResolveUDPAddr("udp", c.Config.Register.Address)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, udp)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	msg := CockMsg{
		Operation: operation,
		Data:      data,
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
