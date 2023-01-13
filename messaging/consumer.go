package messaging

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

type Consumer struct {
	consumer *nsq.Consumer
}

func newConsumer(topic string, channel string, handler func(string)) Consumer {
	config := nsq.NewConfig()
	nsqConsumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}
	nsqConsumer.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		handler(string(m.Body))
		return nil
	}))
	consumer := Consumer{nsqConsumer}
	return consumer
}

func NewCameraConsumer(handler func(string)) Consumer {
	consumer := newConsumer("image_requests", "camera", handler)
	return consumer
}

func NewProcessConsumer(handler func(string)) Consumer {
	consumer := newConsumer("image_responses", "process", handler)
	return consumer
}

func (consumer *Consumer) Listen(address string) {
	err := consumer.consumer.ConnectToNSQLookupd(address)
	if err != nil {
		log.Fatal(err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	consumer.consumer.Stop()
}
