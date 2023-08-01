// Package test
// @author uangi
// @date 2023/7/19 10:04
package test

import (
	"github.com/real-uangi/cockc/common/bar"
	"github.com/real-uangi/cockc/common/plog"
	"testing"
	"time"
)

var logger = plog.New("test")

func TestLog(t *testing.T) {
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")

}

func TestProgressBar(t *testing.T) {
	b := bar.NewProgressBar(100)
	for !b.IsFinish() {
		b.Add(1)
		time.Sleep(100 * time.Millisecond)
	}
}
