// Package api @author uangi 2023-05
package api

import (
	"github.com/real-uangi/cockc/common/convert"
	"net/http"
	"time"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Time    string      `json:"time"`
}

func newResult(code int, message string, data interface{}) *Result {
	return &Result{
		Code:    code,
		Data:    data,
		Message: message,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
}

func Success(data interface{}) *Result {
	return newResult(200, "success", data)
}

func Fail(message string, data interface{}) *Result {
	return newResult(500, message, data)
}

func NotFound(message string) *Result {
	if message == "" {
		message = "404 Not Found"
	}
	return newResult(http.StatusNotFound, message, message)
}

func UnAuthorized(message string) *Result {
	if message == "" {
		message = "410 Unauthorized"
	}
	return newResult(http.StatusUnauthorized, message, message)
}

func (r *Result) JsonBytes() []byte {
	return convert.ToJsonBytes(r)
}
