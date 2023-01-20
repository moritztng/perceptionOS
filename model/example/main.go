package main

import (
	"fmt"
	_ "image/jpeg"

	"github.com/moritztng/perceptionOS/model"
)

func main() {
	input := model.Preprocess("person.jpg")
	yoloModel := model.NewModel("../yolov4.onnx", []int{1, 416, 416, 3}, []int{1, 52, 52, 3, 85}, []string{"input_1:0"}, []string{"Identity:0"})
	modelOutput := yoloModel.Run(input)
	output := model.Postprocess(modelOutput)
	fmt.Print(output)
}
