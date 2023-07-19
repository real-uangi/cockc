// Package test
// @author uangi
// @date 2023/7/19 13:55
package test

import (
	"github.com/real-uangi/cockc/runner"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	r := runner.Prepare()
	r.EnableRedisAndSnowflake()
	r.InitDatasource()
	r.Init()
	time.Sleep(time.Minute)
}
