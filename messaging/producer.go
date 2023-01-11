package messaging

import (
	"log"
	"strconv"

	"github.com/nsqio/go-nsq"
)

type Producer struct {
	producer *nsq.Producer
}

func NewProducer() Producer {
	config := nsq.NewConfig()
	nsqProducer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	producer := Producer{nsqProducer}
	return producer
}

func (producer Producer) Stop() {
	producer.producer.Stop()
}

func (producer Producer) PublishImageRequest() {
	messageBody := []byte("request image")
	topicName := "image_requests"

	err := producer.producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}
}

func (producer Producer) PublishImageResponse(id uint) {
	messageBody := []byte(strconv.FormatUint(uint64(id), 10))
	topicName := "image_responses"

	err := producer.producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}
}

func (producer Producer) PublishImageProcess(id uint) {
	messageBody := []byte(strconv.FormatUint(uint64(id), 10))
	topicName := "image_processes"

	err := producer.producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}
}
