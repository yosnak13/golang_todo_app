package config

import (
	"github.com/go-ini/ini"
	"log"
)

type AppConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config AppConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = AppConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
