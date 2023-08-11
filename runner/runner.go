// Package runner
// @author uangi
// @date 2023/6/30 16:18
package runner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/real-uangi/cockc/client"
	"github.com/real-uangi/cockc/common/datasource"
	"github.com/real-uangi/cockc/common/plog"
	"github.com/real-uangi/cockc/common/rdb"
	"github.com/real-uangi/cockc/common/snowflake"
	"github.com/real-uangi/cockc/config"
	"sync"
	"time"
)

var logger = plog.New("runner")

type CockRunner struct {
	cockClient       client.CockClientService
	httpServerEnable bool
	engine           *gin.Engine
	once             sync.Once
}

func Prepare() *CockRunner {
	r := &CockRunner{}
	config.Reload()

	r.cockClient = &client.CockClient{}
	r.cockClient.Load()
	r.cockClient.Echo()
	r.cockClient.PullConfig()

	gin.SetMode(gin.ReleaseMode)
	r.engine = gin.New()
	formatter := func(param gin.LogFormatterParams) string {
		var msg = fmt.Sprintf("[%d] %s takes %dms", param.StatusCode, param.Path, param.Latency.Microseconds())
		return logger.GetLine(plog.LvInfo, msg, param.TimeStamp)
	}
	r.engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: formatter,
		Output:    nil,
		SkipPaths: nil,
	}))

	return r
}

func (r *CockRunner) EnableRedisAndSnowflake() {
	rdb.Init()
	snowflake.Init()
}

func (r *CockRunner) GetRouter() *gin.Engine {
	if !r.httpServerEnable {
		r.httpServerEnable = true
	}
	return r.engine
}

func (r *CockRunner) InitDatasource() {
	datasource.InitDataSource()
}

func (r *CockRunner) RunAsync() {
	r.once.Do(func() {
		if r.httpServerEnable {
			go func() {
				port := fmt.Sprintf(":%d", config.GetPropertiesRO().Cock.Port)
				logger.Info("server running on " + port)
				err := r.engine.Run(port)
				if err != nil {
					logger.Error(err.Error())
				}
			}()
		}
		time.Sleep(5 * time.Second)
		r.cockClient.StartHeartbeat()
		r.cockClient.Online()
	})
}

func (r *CockRunner) Run() {
	r.once.Do(func() {
		go func() {
			time.Sleep(5 * time.Second)
			r.cockClient.StartHeartbeat()
			r.cockClient.Online()
		}()
		if r.httpServerEnable {
			port := fmt.Sprintf(":%d", config.GetPropertiesRO().Cock.Port)
			logger.Info("server running on " + port)
			err := r.engine.Run(port)
			if err != nil {
				logger.Error(err.Error())
			}
		}
	})
}
