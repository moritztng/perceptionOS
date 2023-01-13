package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/perceptionOS/goperception"
	"github.com/moritztng/perceptionOS/storage"
)

var cameraUrl = os.Getenv("CAMERA_URL")
var storageAddress = os.Getenv("STORAGE_ADDRESS")
var accessKeyID = os.Getenv("STORAGE_ACCESS_KEY_ID")
var secretAccessKey = os.Getenv("STORAGE_SECRET_ACCESS_KEY")
var useSSL = os.Getenv("STORAGE_USE_SSL") == "true"
var bucketName = os.Getenv("STORAGE_BUCKET_NAME")
var tempDir = os.TempDir()
var storageClient = storage.NewStorage(storageAddress, accessKeyID, secretAccessKey, useSSL)

const contentType = "image/jpeg"

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	camera := goperception.Camera(cameraUrl)
	id := uuid.New().String()
	fileName := id + ".jpg"
	filePath := filepath.Join(tempDir, fileName)
	camera.SaveImage(filePath)
	storageClient.Store(ctx, bucketName, fileName, filePath, contentType)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
