// Package test
// @author uangi
// @date 2023/7/19 13:55
package test

import (
	"github.com/real-uangi/cockc/common/api"
	"github.com/real-uangi/cockc/runner"
	"net/http"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	r := runner.Prepare()
	r.EnableRedisAndSnowflake()
	r.InitDatasource()

	r.HttpHandleFunc("/hello", handleHttp)

	r.Init()
	time.Sleep(time.Hour)
}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(api.Success("hi!").JsonBytes())
	if err != nil {
		return
	}
}
