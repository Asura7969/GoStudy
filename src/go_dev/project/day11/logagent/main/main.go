package main

import (
	"../kafka"
	"../tailf"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	fileName := "./conf/logagent.conf"
	err := loadConf("ini", fileName)
	if err != nil {
		fmt.Println("load config err:", err)
		panic("load config failed!")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed,err:%v\n", err)
		panic("load logger failed!")
		return
	}

	err = tailf.InitTail(appconfig.collectConf, appconfig.chanSize)
	if err != nil {
		logs.Error("init tail failed,err:%v\n", err)
		panic("init tail failed!")
		return
	}

	err = kafka.InitKafka(appconfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka_consumer failed,err:%v\n", err)
		panic("init kafka_consumer failed!")
		return
	}

	logs.Debug("initaiallze all succ")
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed,err:", err)
		return
	}

	logs.Info("program exited!")
}
