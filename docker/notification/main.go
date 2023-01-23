package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/perceptionOS/messaging"
	"github.com/moritztng/perceptionOS/qlient"
	"github.com/moritztng/perceptionOS/storage"
)

var storageAddress = os.Getenv("STORAGE_ADDRESS")
var accessKeyID = os.Getenv("STORAGE_ACCESS_KEY_ID")
var secretAccessKey = os.Getenv("STORAGE_SECRET_ACCESS_KEY")
var useSSL = os.Getenv("STORAGE_USE_SSL") == "true"
var bucketName = os.Getenv("STORAGE_BUCKET_NAME")
var apiUrl = os.Getenv("QLIENT_API_URL")
var consumerAddress = os.Getenv("MESSAGING_CONSUMER_ADDRESS")
var producerAddress = os.Getenv("MESSAGING_PRODUCER_ADDRESS")
var tempDir = os.TempDir()
var storageClient = storage.NewStorage(storageAddress, accessKeyID, secretAccessKey, useSSL)
var apiClient = qlient.NewClient(apiUrl)

func handler(message string) {
	ctx := context.Background()
	detectionId, _ := strconv.Atoi(message)
	response, _ := qlient.Detection(ctx, apiClient, detectionId)
	fmt.Println(response.Detection.Person)
}

func main() {
	messageConsumer := messaging.NewNotificationConsumer(handler)
	messageConsumer.Listen(consumerAddress)
}
