// Package runner
// @author uangi
// @date 2023/6/30 16:18
package runner

import (
	"github.com/real-uangi/cockc/client"
	"github.com/real-uangi/cockc/config"
)

type CockRunner struct {
}

var cockClientService client.CockClientService

func Prepare() *CockRunner {
	runner := &CockRunner{}
	config.Reload()

	cockClientService = &client.CockClient{}
	cockClientService.Load()
	cockClientService.Connect()
	cockClientService.PullConfig()

	return runner
}
