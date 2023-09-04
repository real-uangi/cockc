// Package plog @author uangi 2023-05
package plog

import (
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

func (l *Logger) GetLine(lv string, msg string, t time.Time) string {
	return fmt.Sprintf("%s [%s] %s :  %s\n", t.Format("2006-01-02 15:04:05"), lv, l.PrettyName, msg)
}

func (l *Logger) Info(s string) {
	fmt.Print(l.GetLine(LvInfo, s, time.Now()))
}

func (l *Logger) Warn(s string) {
	fmt.Print(l.GetLine(LvWarn, s, time.Now()))
}

func (l *Logger) Error(s string) {
	fmt.Print(l.GetLine(LvError, s, time.Now()))
}

func (l *Logger) Panic(err error) {
	if err != nil {
		l.Error(err.Error())
		panic(err)
	}
}
