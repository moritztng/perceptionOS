package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/perceptionOS/messaging"
	"github.com/moritztng/perceptionOS/qlient"
)

var apiUrl = os.Getenv("QLIENT_API_URL")
var consumerAddress = os.Getenv("MESSAGING_CONSUMER_ADDRESS")
var apiClient = qlient.NewClient(apiUrl)
var detectionSensitivity, _ = strconv.ParseFloat(os.Getenv("DETECTION_SENSITIVITY"), 64)
var notificationInterval, _ = time.ParseDuration(os.Getenv("NOTIFICATION_INTERVAL"))
var lastDetectionTime time.Time
var bot, _ = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))

func handler(message string) {
	ctx := context.Background()
	detectionId, _ := strconv.Atoi(message)
	response, _ := qlient.Detection(ctx, apiClient, detectionId)
	personDetection := response.Detection.Person
	fmt.Println(personDetection)
	if personDetection > detectionSensitivity {
		detectionTime := time.Now()
		lastDetectionInterval := detectionTime.Sub(lastDetectionTime)
		if lastDetectionInterval > notificationInterval {
			botMessage := tgbotapi.NewMessage(5519266765, fmt.Sprintf("Person detected with %.6f propability!", personDetection))
			bot.Send(botMessage)
		}
		lastDetectionTime = detectionTime
	}
}

func main() {
	messageConsumer := messaging.NewNotificationConsumer(handler)
	messageConsumer.Listen(consumerAddress)
}
