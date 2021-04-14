package conf

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)
type Server struct {
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
}

var DatabaseSetting = &Database{}
var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/conf.ini")
	if err != nil {
		log.Fatalf("failed to load config! :%v", err)
	}
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("failed to load config! MapTo: %s, :%v", section, err)
	}
}
