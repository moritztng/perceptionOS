package main

import (
	"github.com/moritztng/perceptionOS/sensors"
)

func main() {
	capture := sensors.Camera("http://192.168.0.99:8080/shot.jpg")
	capture.SaveImage("image.jpg")
}
