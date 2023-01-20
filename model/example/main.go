package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"

	"github.com/moritztng/perceptionOS/model"
)

func preprocess(image image.Image) (output []float32) {
	bounds := image.Bounds()
	output = make([]float32, bounds.Dx()*bounds.Dy()*3)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := image.At(x, y).RGBA()
			pixelIndex := (y*bounds.Max.X + x) * 3
			output[pixelIndex] = (float32)(r) / 0xffff
			output[pixelIndex+1] = (float32)(g) / 0xffff
			output[pixelIndex+2] = (float32)(b) / 0xffff
		}
	}
	return
}

func main() {
	reader, err := os.Open("person.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	image, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	input := preprocess(image)
	model := model.NewModel("../yolov4.onnx", []int{1, 416, 416, 3}, []int{1, 52, 52, 3, 85}, []string{"input_1:0"}, []string{"Identity:0"})
	output := model.Run(input)
	var maxPersonConfidence float32 = 0
	nBoxes := 1 * 52 * 52 * 3
	for i := 0; i < nBoxes; i++ {
		personConfidenceIndex := i*85 + 5
		if output[personConfidenceIndex] > maxPersonConfidence {
			maxPersonConfidence = output[personConfidenceIndex]
		}
	}
	fmt.Print(maxPersonConfidence)
	return
}
