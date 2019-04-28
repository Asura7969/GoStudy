package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	producer sarama.SyncProducer
)

func InitKafka(address string) (err error) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.Return.Successes = true

	producer, err = sarama.NewSyncProducer([]string{address}, conf)
	if err != nil {
		logs.Error("init kafka_consumer producer failed,err:", err)
		return
	}
	logs.Debug("init kafka_consumer success")
	return
	//defer producer.Close()
}

func SendToKafka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed,err:%v data:%v topic:%v\n", err, data, topic)
		return
	}
	logs.Debug("send success,pid:%v offest:%v,topic:%v\n", partition, offset, topic)
	return
}
