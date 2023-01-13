package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/perceptionOS/goperception"
	"github.com/moritztng/perceptionOS/messaging"
	"github.com/moritztng/perceptionOS/qlient"
	"github.com/moritztng/perceptionOS/storage"
)

var cameraUrl = os.Getenv("CAMERA_URL")
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
var messageProducer = messaging.NewProducer(producerAddress)

const contentType = "image/jpeg"

func handler(message string) {
	ctx := context.Background()
	camera := goperception.Camera(cameraUrl)
	id := uuid.New().String()
	fileName := id + ".jpg"
	filePath := filepath.Join(tempDir, fileName)
	camera.SaveImage(filePath)
	storageClient.Store(ctx, bucketName, fileName, filePath, contentType)
	response, _ := qlient.AddImage(ctx, apiClient, fileName)
	imageId := uint(response.AddImage.GetId())
	messageProducer.PublishImageResponse(imageId)
}

func main() {
	messageConsumer := messaging.NewCameraConsumer(handler)
	messageConsumer.Listen(consumerAddress)
}
