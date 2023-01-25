package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/perceptionOS/messaging"
	"github.com/moritztng/perceptionOS/model"
	"github.com/moritztng/perceptionOS/qlient"
	"github.com/moritztng/perceptionOS/storage"
)

var yoloModel = model.NewModel(os.Getenv("MODEL_PATH"), []int{1, 416, 416, 3}, []int{1, 52, 52, 3, 85}, []string{"input_1:0"}, []string{"Identity:0"})

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

func handler(message string) {
	ctx := context.Background()
	fmt.Println(message)
	imageId, _ := strconv.Atoi(message)
	fmt.Println(imageId)
	filenameResponse, _ := qlient.ImageFilename(ctx, apiClient, imageId)
	fmt.Println(filenameResponse)
	fileName := filenameResponse.Image.Filename
	storageClient.Get(ctx, bucketName, fileName, tempDir)
	filePath := filepath.Join(tempDir, fileName)
	fmt.Println(filePath)
	modelInput := model.Preprocess(filePath, 416, 416)
	modelOutput := yoloModel.Run(modelInput)
	personDetection := model.Postprocess(modelOutput)
	idResponse, _ := qlient.AddDetection(ctx, apiClient, imageId, float64(personDetection))
	fmt.Print(idResponse.AddDetection.Id)
	messageProducer.PublishImageProcess(uint(idResponse.AddDetection.Id))
}

func main() {
	messageConsumer := messaging.NewProcessConsumer(handler)
	messageConsumer.Listen(consumerAddress)
}
