package main

import (
	"fmt"

	"github.com/moritztng/perceptionOS/messaging"
)

func main() {
	producer := messaging.NewProducer()
	producer.PublishImageRequest()
	//producer.PublishImageProcess(1234)
	consumer := messaging.NewCameraConsumer(func(msg string) { fmt.Println(msg) })
	consumer.Listen()
}
