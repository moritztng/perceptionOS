package main

import (
	"fmt"

	"github.com/moritztng/perceptionOS/messaging"
)

func main() {
	producer := messaging.NewProducer("127.0.0.1:4150")
	producer.PublishImageRequest()
	//producer.PublishImageProcess(1234)
	consumer := messaging.NewProcessConsumer(func(msg string) { fmt.Println(msg) })
	consumer.Listen("localhost:4161")
}
