package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/moritztng/perceptionOS/goperception"
)

var camera_url = os.Getenv("CAMERA_URL")
var image_dir = os.Getenv("IMAGE_DIR")

func handler(w http.ResponseWriter, r *http.Request) {
	camera := goperception.Camera(camera_url)
	id := uuid.New().String()
	filename := id + ".jpg"
	camera.SaveImage(filepath.Join(image_dir, filename))
	io.WriteString(w, id)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
