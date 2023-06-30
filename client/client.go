// Package client
// @author uangi
// @date 2023/6/30 10:00
package client

import (
	"github.com/real-uangi/cockc/common/plog"
	"github.com/real-uangi/cockc/config"
	"log"
	"math/rand"
	"net/rpc"
	"strconv"
	"sync"
	"time"
)

const (
	rpcMode         = "tcp"
	registerService = "CockServerService"
)

const (
	rpcEcho       = registerService + ".Echo"
	rpcPullConfig = registerService + ".PullConfig"
	rpcHeartbeat  = registerService + ".Heartbeat"
	rpcOnline     = registerService + ".Online"
	rpcOffline    = registerService + ".Offline"
)

type CockClientService interface {
	Load()
	Connect()
	PullConfig()
	StartHeartbeat()
	Online()
	Offline()
}

type CockClient struct {
	Up            bool
	Config        config.Cock
	Client        *rpc.Client
	HeartbeatLock sync.Mutex
}

func (c *CockClient) Load() {
	c.Config = config.GetPropertiesRO().Cock
	c.Up = false
}

func (c *CockClient) Connect() {
	if c.Client == nil {
		client, err := rpc.Dial(rpcMode, c.Config.Register.Address)
		if err != nil {
			plog.Error(err.Error())
			return
		}
		c.Client = client
		var request = strconv.Itoa(rand.Intn(100000))
		var reply string
		err = c.Client.Call(rpcEcho, request, &reply)
		if err != nil {
			plog.Error(err.Error())
			return
		}
		if reply == request {
			log.Printf("connect to %s %s successfully \n", registerService, c.Config.Register.Address)
		}
	}

}

func (c *CockClient) PullConfig() {
	var cs string
	c.Connect()
	err := c.Client.Call(rpcPullConfig, c.Config.Register.Address, &cs)
	if err != nil {
		plog.Error(err.Error())
	}
	config.UpdateServerConfig(cs)
}

func (c *CockClient) StartHeartbeat() {

}

func (c *CockClient) heartbeat() {
	if c.HeartbeatLock.TryLock() {
		defer c.HeartbeatLock.Unlock()
		for {
			//
			time.Sleep(time.Duration(c.Config.Register.Heartbeat.Interval) * time.Millisecond)
		}
	}
}

func beat() {

}

func (c *CockClient) Online() {
	c.Up = true
}

func (c *CockClient) Offline() {
	c.Up = false
}
