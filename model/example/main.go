package main

import (
	"fmt"

	"github.com/moritztng/perceptionOS/model"
)

func main() {
	//capture := goperception.Camera("http://192.168.0.99:8080/shot.jpg")
	//capture.SaveImage("image.jpg")
	input := model.Preprocess("image.jpg", 416, 416)
	yoloModel := model.NewModel("../yolov4.onnx", []int{1, 416, 416, 3}, []int{1, 52, 52, 3, 85}, []string{"input_1:0"}, []string{"Identity:0"})
	modelOutput := yoloModel.Run(input)
	output := model.Postprocess(modelOutput)
	fmt.Print(output)
}
