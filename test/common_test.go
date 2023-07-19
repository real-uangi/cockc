// Package test
// @author uangi
// @date 2023/7/19 10:04
package test

import (
	"github.com/real-uangi/cockc/common/plog"
	"testing"
)

var logger = plog.New("test")

func TestLog(t *testing.T) {
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")

}
