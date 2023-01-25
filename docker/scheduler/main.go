package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/perceptionOS/messaging"
)

func main() {
	interval, _ := time.ParseDuration(os.Getenv("INTERVAL"))
	messageProducer := messaging.NewProducer(os.Getenv("MESSAGING_PRODUCER_ADDRESS"))
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		time := <-ticker.C
		fmt.Println("Current time: ", time)
		messageProducer.PublishImageRequest()
	}
}
