package models

import (
	"beego-master/logs"
	"fmt"

	"github.com/Shopify/sarama"
)

var (
	Producerclient sarama.SyncProducer

	Consumerclient sarama.Consumer
)

func InitKafka(addr []string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	Producerclient, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		logs.Error("init kafka producer failed, err:", err)
		return
	}

	logs.Debug("init kafka succ")
	return
}

func SendToKafka(data, topic string) (err error) {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	_, _, err = Producerclient.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err:%v data:%v topic:%v", err, data, topic)
		return
	}
	return
}

func InitKafkaConsumer(addr []string) (err error) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	//初始"地址 localhost:9092"
	Consumerclient, err = sarama.NewConsumer(addr, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	logs.Debug("init kafka succ")
	return
}
