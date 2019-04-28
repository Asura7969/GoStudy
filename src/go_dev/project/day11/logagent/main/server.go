package main

import (
	"../kafka"
	"../tailf"
	"github.com/astaxie/beego/logs"
	"time"
)

func serverRun() (err error) {

	for true {
		msg := tailf.GetOneLine()
		err := sendKafka(msg)
		if err != nil {
			logs.Error("send to kafka_consumer failed,err:", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendKafka(msg *tailf.TextMsg) (err error) {
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	//fmt.Println("read msg:%s,topic:%s\n",msg.Msg,msg.Topic)
	return
}
