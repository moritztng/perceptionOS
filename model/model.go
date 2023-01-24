package model

/*
#include <stdlib.h>
#include "model.h"
*/
import "C"

import (
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"reflect"
	"unsafe"

	"golang.org/x/image/draw"
)

type Model struct {
	cModel       C.Model
	outputLength int
}

func NewModel(filepath string, inputShape []int, outputShape []int, inputNames []string, outputNames []string) (model Model) {
	cFilepath := C.CString(filepath)
	cInputShape := (*C.int64_t)(unsafe.Pointer(&inputShape[0]))
	cInputShapeLen := (C.size_t)(len(inputShape))
	cOutputNames := make([]*C.char, len(outputNames))
	cInputNames := make([]*C.char, len(inputNames))
	for i, name := range inputNames {
		cName := C.CString(name)
		cInputNames[i] = cName
	}
	for i, name := range outputNames {
		cName := C.CString(name)
		cOutputNames[i] = cName
	}
	cInputNamesPointer := (**C.char)(unsafe.Pointer(&cInputNames[0]))
	cOutputNamesPointer := (**C.char)(unsafe.Pointer(&cOutputNames[0]))
	C.NewModel(cFilepath, cInputNamesPointer, cOutputNamesPointer, cInputShape, cInputShapeLen, &model.cModel)
	model.outputLength = 1
	for _, dim := range outputShape {
		model.outputLength *= dim
	}
	return
}

func (model Model) Run(input []float32) (output []float32) {
	var outputPointer (*C.float)
	C.Run(model.cModel, (*C.float)(&input[0]), (**C.float)(&outputPointer))
	outputSliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&output))
	outputSliceHeader.Data = (uintptr)(unsafe.Pointer(outputPointer))
	outputSliceHeader.Len = model.outputLength
	outputSliceHeader.Cap = model.outputLength
	return
}

func Preprocess(filename string, width int, height int) (output []float32) {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	input, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	resized := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(resized, resized.Rect, input, input.Bounds(), draw.Over, nil)
	bounds := resized.Bounds()
	output = make([]float32, bounds.Dx()*bounds.Dy()*3)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := resized.At(x, y).RGBA()
			pixelIndex := (y*bounds.Max.X + x) * 3
			output[pixelIndex] = (float32)(r) / 0xffff
			output[pixelIndex+1] = (float32)(g) / 0xffff
			output[pixelIndex+2] = (float32)(b) / 0xffff
		}
	}
	return
}

func Postprocess(input []float32) (output float32) {
	var maxPersonConfidence float32 = 0
	nBoxes := 1 * 52 * 52 * 3
	for i := 0; i < nBoxes; i++ {
		personConfidence := input[i*85+4] * input[i*85+5]
		if personConfidence > maxPersonConfidence {
			maxPersonConfidence = personConfidence
		}
	}
	output = maxPersonConfidence
	return
}
