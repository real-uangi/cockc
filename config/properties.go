// Package config
// @author uangi
// @date 2023/6/30 10:35
package config

import (
	"encoding/json"
	"fmt"
	"github.com/real-uangi/cockc/common/plog"
	"os"
	"sync"
)

var p Properties

type Properties struct {
	Cock   Cock   `json:"cock"`
	Server Server `json:"server"`
	mu     sync.Mutex
}

type Cock struct {
	AppName  string   `json:"app_name"`
	Register Register `json:"register"`
}

type Register struct {
	Address   string    `json:"address"`
	Version   int       `json:"version"`
	Heartbeat Heartbeat `json:"heartbeat"`
}

type Heartbeat struct {
	Interval int `json:"interval"`
	Timeout  int `json:"timeout"`
	Offline  int `json:"offline"`
}

type Server struct {
	Loaded     bool         `json:"loaded"`
	Http       http         `json:"http"`
	Datasource []datasource `json:"datasource"`
	Redis      redis        `json:"rdb"`
	Snowflake  snowflake    `json:"snowflake"`
}

type http struct {
	Port string `json:"port"`
	Log  bool   `json:"log"`
}

type datasource struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
	PoolMin  int    `json:"poolMin"`
	PoolMax  int    `json:"poolMax"`
}

type snowflake struct {
	Interval int `json:"interval"`
}

func Reload() {
	f, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dc := json.NewDecoder(f)
	err = dc.Decode(&p)
	if err != nil {
		panic(err)
	}
}

func GetPropertiesRO() Properties {
	if p.Cock.AppName == "" {
		fmt.Println(" [Warning] Empty Config !")
		Reload()
	}
	if p.Cock.AppName == "" {
		panic("Empty Config")
	}
	return Properties{
		Cock:   p.Cock,
		Server: p.Server,
	}
}

func UpdateServerConfig(cs string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	err := json.Unmarshal([]byte(cs), &p.Server)
	if err != nil {
		plog.Error(err.Error())
	}
}
