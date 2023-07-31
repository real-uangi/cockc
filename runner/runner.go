// Package runner
// @author uangi
// @date 2023/6/30 16:18
package runner

import (
	"github.com/real-uangi/cockc/client"
	"github.com/real-uangi/cockc/common/datasource"
	"github.com/real-uangi/cockc/common/plog"
	"github.com/real-uangi/cockc/common/rdb"
	"github.com/real-uangi/cockc/common/snowflake"
	"github.com/real-uangi/cockc/config"
	"net/http"
)

var logger = plog.New("runner")

type CockRunner struct {
	cockClient       client.CockClientService
	httpServerEnable bool
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

func (r *CockRunner) EnableRedisAndSnowflake() {
	rdb.Init()
	snowflake.Init()
}

func (r *CockRunner) InitDatasource() {
	datasource.InitDataSource()
}

func (r *CockRunner) HttpHandle(pattern string, handler http.Handler) {
	if !r.httpServerEnable {
		r.httpServerEnable = true
	}
	http.Handle(pattern, handler)
}

func (r *CockRunner) HttpHandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	if !r.httpServerEnable {
		r.httpServerEnable = true
	}
	http.HandleFunc(pattern, handler)
}

func (r *CockRunner) Init() {
	if r.httpServerEnable {
		go func() {
			port := config.GetPropertiesRO().Server.Http.Port
			logger.Info("http server listen on port " + port)
			err := http.ListenAndServe(":"+port, nil)
			if err != nil {
				logger.Error(err.Error())
			}
		}()
	}
	r.cockClient.StartHeartbeat()
	r.cockClient.Online()
}
