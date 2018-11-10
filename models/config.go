package models

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LogConfig struct {
	Kafkapath []string
	EtcdKey   string
	EtcdPath  []string
	EsPath    string
}

func NewConfig() LogConfig {
	kafka, err := beego.AppConfig.GetSection("kafka")
	etcd, err := beego.AppConfig.GetSection("etcd")
	es, err := beego.AppConfig.GetSection("es")
	if err != nil {
		logs.Debug("读取配置错误：%s", err)
	}
	LogConf := LogConfig{}
	LogConf.Kafkapath = strings.Split(kafka["kafkapath"], ",")
	LogConf.EtcdPath = strings.Split(etcd["path"], ",")
	LogConf.EtcdKey = etcd["key"]
	LogConf.EsPath = es["path"]
	return LogConf
}
