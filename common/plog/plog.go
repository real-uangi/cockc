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

type Logger struct {
	Name       string `json:"name"`
	PrettyName string `json:"prettyName"`
}

func New(name string) *Logger {
	var pn string
	if len(name) > 10 {
		pn = name[0:9]
	} else {
		pn = name
		for len(pn) < 10 {
			pn = " " + pn
		}
	}
	return &Logger{
		Name:       name,
		PrettyName: pn,
	}
}

func (l *Logger) getLine(lv string, msg string, t time.Time) string {
	return fmt.Sprintf("%s [%s] %s :  %s\n", time.Now().Format("2006-01-02 15:04:05"), lv, l.PrettyName, msg)
}

func (l *Logger) Info(s string) {
	fmt.Print(l.getLine(LvInfo, s, time.Now()))
}

func (l *Logger) Warn(s string) {
	fmt.Print(l.getLine(LvWarn, s, time.Now()))
}

func (l *Logger) Error(s string) {
	fmt.Print(l.getLine(LvError, s, time.Now()))
}

func (l *Logger) TryThrow(err error) {
	if err != nil {
		l.Panic(err.Error())
	}
}

func (l *Logger) Panic(msg string) {
	l.Error(msg)
	panic(errors.New(msg))
}
