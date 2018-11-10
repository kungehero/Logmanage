package logappend

import (
	"Logmanage/models"
	"time"

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

	models.InitKafka(config.Kafkapath)
}

func ReadLogWithTailf() {
	ch := make(chan bool)
	for _, LogpathStruct := range models.EtcdValue {
		models.NewTailf(LogpathStruct.LogPath)
		go LogRun(models.Tails, LogpathStruct.Topic, ch)
		ch <- true
	}
}

func LogRun(tails *tail.Tail, topic string, ch chan bool) {
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines

		if !ok {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		//{{  传入kafka函数   }}
		models.SendToKafka(msg.Text, topic)
	}
	<-ch
}
