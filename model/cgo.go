package model

// #cgo LDFLAGS: -lonnxruntime
// #cgo CFLAGS: -std=c11 -I. -I/usr/include/onnxruntime
// #include "model.h"
import "C"
