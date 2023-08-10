// Package test
// @author uangi
// @date 2023/7/19 13:55
package test

import (
	"github.com/gin-gonic/gin"
	"github.com/real-uangi/cockc/common/api"
	"github.com/real-uangi/cockc/runner"
	"net/http"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	r := runner.Prepare()
	//r.EnableRedisAndSnowflake()
	//r.InitDatasource()

	r.GetRouter().GET("/handle", handleHttp)

	r.Init()
	time.Sleep(time.Hour)
}

func handleHttp(c *gin.Context) {
	c.JSON(http.StatusOK, api.Success("yeah"))
}
