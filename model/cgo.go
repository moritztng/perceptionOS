package model

// #cgo LDFLAGS: -L/usr/lib/onnxruntime -lonnxruntime
// #cgo CFLAGS: -std=c11 -I. -I/usr/include/onnxruntime
// #include "model.h"
import "C"
