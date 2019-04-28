package main

import (
	"../tailf"
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)

type Config struct {
	logLevel    string
	logPath     string
	chanSize    int
	kafkaAddr   string
	collectConf []tailf.CollectConf
}

var (
	appconfig *Config
)

func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::Topic")
		return
	}

	appconfig.collectConf = append(appconfig.collectConf, cc)
	return
}

func loadConf(fileType string, fileName string) (err error) {
	conf, err := config.NewConfig(fileType, fileName)
	if err != nil {
		fmt.Println("new config failed,err:", err)
		return
	}
	appconfig = &Config{}
	appconfig.logLevel = conf.String("log::log_level")
	if len(appconfig.logLevel) == 0 {
		appconfig.logLevel = "info"
	}
	appconfig.logPath = conf.String("log::log_path")
	if len(appconfig.logPath) == 0 {
		appconfig.logPath = "./logs"
	}
	appconfig.chanSize, err = conf.Int("collect::chan_size")
	if err != nil {
		appconfig.chanSize = 100
	}

	appconfig.kafkaAddr = conf.String("kafka_consumer::address")
	if len(appconfig.kafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka_consumer address")
		return
	}

	err = loadCollectConf(conf)
	if err != nil {
		fmt.Println("load collect conf failed,err:%v\n", err)
	}

	return
}
