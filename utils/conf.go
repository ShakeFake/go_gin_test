package utils

import (
	"github.com/Unknwon/goconfig"
	log "github.com/cihub/seelog"
)

var (
	HttpPort string
)

func Config() {
	conf, err := goconfig.LoadConfigFile("./conf/conf.ini")
	if err != nil {
		panic(err)
	}

	HttpPort = conf.MustValue("Server", "http_port", "8091")

	log.Infof("HttpPort is:%v", HttpPort)
}
