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
)

func TestRunner(t *testing.T) {
	r := runner.Prepare()
	//r.EnableRedisAndSnowflake()
	//r.InitDatasource()

	r.GetRouter().GET("/handle", handleHttp)

	//r.RunAsync()
	r.Run()
}

func handleHttp(c *gin.Context) {
	c.JSON(http.StatusOK, api.Success("yeah"))
}
