package main

import (
	"github.com/moritztng/perceptionOS/goperception"
)

func main() {
	capture := goperception.Camera("http://192.168.0.99:8080/shot.jpg")
	capture.SaveImage("image.jpg")
}
