// Package api @author uangi 2023-05
package api

import (
	"encoding/json"
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

func Fail(err error, message string) *Result {
	if err != nil {
		if message == "" {
			return newResult(0, message, err.Error())
		}
		return newResult(500, "failed", err.Error())
	}
	if message == "" {
		return newResult(0, message, message)
	}
	return newResult(500, "failed", nil)
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
	bs, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return bs
}
