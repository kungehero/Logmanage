package logcollect

import (
	"Logmanage/models"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
)

type LogAgent struct {
	Tailfs []*tail.Tail
	Topic  []string
}

var config models.LogConfig

func init() {
	config = models.NewConfig()

	models.EtcdNew(config.EtcdPath)

	models.EtcdGet(config.EtcdKey)

	models.InitKafkaConsumer(config.Kafkapath)
	models.NewEs(config.EsPath)

}

func ReadLogWithEs() {
	ch := make(chan bool)
	for _, LogpathStruct := range models.EtcdValue {
		Run(LogpathStruct.Topic, ch)
	}

	for v := range ch {
		fmt.Println(v)
	}
}

func Run(topic string, ch chan bool) {
	partition_consumer, err := models.Consumerclient.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}

	//defer partition_consumer.Close()
	for true {
		select {
		case msg := <-partition_consumer.Messages():
			content := fmt.Sprintf("msg offset: %d, partition: %d, timestamp: %s, value: %s,topic:%s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value), msg.Topic)
			fmt.Println(content)
			/* 		exists, err := models.EsClient.IndexExists("twitter").Do(context.Background())

			if err != nil {
				// Handle error
				panic(err)
			}
			if !exists {
				// Create a new index.
				mapping := `fedbgfrthty`
				createIndex, err := models.EsClient.CreateIndex("twitter").BodyString(mapping).Do(context.Background())
				if err != nil {
					// Handle error
					panic(err)
				}
				if !createIndex.Acknowledged {
					// Not acknowledged
				}
			}

			// Index a tweet (using JSON serialization)

			tweet1 := Tweet{User: "olivere", Message: content, Retweets: 0}
			put1, err := models.EsClient.Index().
				Index("twitter").
				Type("doc").
				Id("1").
				BodyJson(tweet1).
				Do(context.Background())
			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("Indexed tweet %s to index %s, type %s ,result:%s\n", put1.Id, put1.Index, put1.Type, put1.Result)*/
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
	ch <- true
}

type Tweet struct {
	User     string
	Message  string
	Retweets int
}
