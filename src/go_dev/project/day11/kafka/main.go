package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {

	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, conf)
	if err != nil {
		fmt.Println("producer is closed,err:", err)
		return
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is good test,my message is good")

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,err:", err)
		return
	}

	fmt.Printf("partition:%v offset:%v\n", partition, offset)

}
