package config

import "github.com/jinzhu/configor"

var Config = struct {
	SampleMssql struct {
		Conn string `json:"Conn"`
	} `json:"SampleMssql"`
	SampleMysql struct {
		Conn string `json:"Conn"`
	} `json:"SampleMysql"`
}{}

func InitConfig(cfg string) {
	// configor.Load(&Config, "config/config.json")
	configor.Load(&Config, cfg)
}
