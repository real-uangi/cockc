// Package config
// @author uangi
// @date 2023/7/19 13:39
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

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
	p.Server.Loaded = true
	if err != nil {
		panic(err)
	}
}
