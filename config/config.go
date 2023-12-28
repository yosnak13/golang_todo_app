package config

import (
	"github.com/go-ini/ini"
	"log"
	"todo_app/utils"
)

type AppConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config AppConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
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
		Static:    cfg.Section("web").Key("static").String(),
	}
}
