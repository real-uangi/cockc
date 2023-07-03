// Package runner
// @author uangi
// @date 2023/6/30 16:18
package runner

import (
	"github.com/real-uangi/cockc/client"
	"github.com/real-uangi/cockc/config"
)

type CockRunner struct {
	cockClient client.CockClientService
}

func Prepare() *CockRunner {
	r := &CockRunner{}
	config.Reload()

	r.cockClient = &client.CockClient{}
	r.cockClient.Load()
	r.cockClient.Echo()
	r.cockClient.PullConfig()

	return r
}

func (r *CockRunner) Init() {

	r.cockClient.StartHeartbeat()
	r.cockClient.Online()

}
