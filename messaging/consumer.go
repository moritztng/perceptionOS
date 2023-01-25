package messaging

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

type Consumer struct {
	consumer *nsq.Consumer
}

func newConsumer(topic string, channel string, messageTimeout time.Duration, handler func(string)) Consumer {
	config := nsq.NewConfig()
	nsqConsumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}
	nsqConsumer.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		messageTime := time.Unix(0, m.Timestamp)
		nowTime := time.Now()
		messageAge := nowTime.Sub(messageTime)
		if messageAge < messageTimeout {
			handler(string(m.Body))
		}
		return nil
	}))
	consumer := Consumer{nsqConsumer}
	return consumer
}

func NewCameraConsumer(messageTimeout time.Duration, handler func(string)) Consumer {
	consumer := newConsumer("image_requests", "camera", messageTimeout, handler)
	return consumer
}

func NewProcessConsumer(messageTimeout time.Duration, handler func(string)) Consumer {
	consumer := newConsumer("image_responses", "process", messageTimeout, handler)
	return consumer
}

func NewNotificationConsumer(messageTimeout time.Duration, handler func(string)) Consumer {
	consumer := newConsumer("image_processes", "notification", messageTimeout, handler)
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
