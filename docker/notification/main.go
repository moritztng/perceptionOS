package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/perceptionOS/messaging"
	"github.com/moritztng/perceptionOS/qlient"
)

var apiUrl = os.Getenv("QLIENT_API_URL")
var consumerAddress = os.Getenv("MESSAGING_CONSUMER_ADDRESS")
var apiClient = qlient.NewClient(apiUrl)
var bot, _ = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))

func handler(message string) {
	ctx := context.Background()
	detectionId, _ := strconv.Atoi(message)
	response, _ := qlient.Detection(ctx, apiClient, detectionId)
	personDetection := response.Detection.Person
	fmt.Println(personDetection)
	msg := tgbotapi.NewMessage(5519266765, strconv.FormatFloat(personDetection, 'E', -1, 64))
	bot.Send(msg)
}

func main() {
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	messageConsumer := messaging.NewNotificationConsumer(handler)
	messageConsumer.Listen(consumerAddress)
}
