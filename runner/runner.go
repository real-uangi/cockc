// Package runner
// @author uangi
// @date 2023/6/30 16:18
package runner

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/real-uangi/cockc/client"
	"github.com/real-uangi/cockc/common/api"
	"github.com/real-uangi/cockc/common/datasource"
	"github.com/real-uangi/cockc/common/plog"
	"github.com/real-uangi/cockc/common/rdb"
	"github.com/real-uangi/cockc/common/snowflake"
	"github.com/real-uangi/cockc/config"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var logger = plog.New("runner")
var ginLogger = plog.New("gin")

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
		return ginLogger.GetLine(plog.LvInfo, msg, param.TimeStamp)
	}
	r.engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: formatter,
		Output:    nil,
		SkipPaths: nil,
	}))
	r.engine.Use(api.Auth())
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

func (r *CockRunner) EnableAuthentication() {

}

func (r *CockRunner) Run() {
	r.once.Do(func() {
		go func() {
			time.Sleep(5 * time.Second)
			r.cockClient.StartHeartbeat()
			r.cockClient.Online()
		}()
		if r.httpServerEnable {
			serveOnly4(r.engine, config.GetPropertiesRO().Cock.Port)
		}
	})
}

func serveOnly4(r *gin.Engine, port int) {
	//use ctx
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	//start http server
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	logger.Info("server running on " + addr)
	server := &http.Server{Addr: addr, Handler: r}
	ln, err := net.Listen("tcp4", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		err = server.Serve(ln.(*net.TCPListener))
		if err != nil {
			logger.Error(err.Error())
		}
	}()
	//graceful shutdown
	<-ctx.Done()
	logger.Info("shutting down ...")
	stop()
	//wait for 15 sec
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()
	if err := server.Shutdown(timeoutCtx); err != nil {
		logger.Error(err.Error())
	}
}
