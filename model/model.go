package model

/*
#include <stdlib.h>
#include "model.h"
*/
import "C"

import (
	"reflect"
	"unsafe"
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
