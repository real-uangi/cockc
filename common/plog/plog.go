// Package plog @author uangi 2023-05
package plog

import (
	"errors"
	"fmt"
	"time"
)

const (
	LvInfo  string = "INFO "
	LvWarn  string = "WARN "
	LvError string = "ERROR"
)

func GetLine(lv string, msg string, t time.Time) string {
	return fmt.Sprintf("%s [%s] :  %s\n", time.Now().Format("2006-01-02 15:04:05"), lv, msg)
}

func Info(s string) {
	fmt.Print(GetLine(LvInfo, s, time.Now()))
}

func Warn(s string) {
	fmt.Print(GetLine(LvWarn, s, time.Now()))
}

func Error(s string) {
	fmt.Print(GetLine(LvError, s, time.Now()))
}

func TryThrow(err error) {
	if err != nil {
		Error(err.Error())
		panic(err)
	}
}

func Panic(msg string) {
	Error(msg)
	panic(errors.New(msg))
}
