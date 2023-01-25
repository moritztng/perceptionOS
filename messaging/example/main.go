package main

import (
	"github.com/moritztng/perceptionOS/messaging"
)

func main() {
	producer := messaging.NewProducer("127.0.0.1:4150")
	producer.PublishImageRequest()
	//producer.PublishImageProcess(1234)
	//consumer := messaging.NewCameraConsumer(time.Duration(time.Hour), func(msg string) { fmt.Println(msg) })
	//consumer.Listen("localhost:4161")
}
