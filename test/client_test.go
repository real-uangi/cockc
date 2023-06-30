// Package test
// @author uangi
// @date 2023/6/30 13:33
package test

import (
	"github.com/real-uangi/cockc/client"
	"github.com/real-uangi/cockc/config"
	"testing"
	"time"
)

var cockClientService client.CockClientService

func TestClient(t *testing.T) {
	config.Reload()
	go cockRun()
	time.Sleep(time.Minute)
}

func cockRun() {
	cockClientService = &client.CockClient{}
	cockClientService.Load()
	cockClientService.Connect()
}
