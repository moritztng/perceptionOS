package goperception

/*
#include <stdlib.h>
#include "camera.h"
*/
import "C"

import (
	"unsafe"
)

type VideoCapture struct {
	p C.VideoCapture
}

func Camera(url string) (capture *VideoCapture) {
	var curl *C.char = C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	capture = &VideoCapture{p: C.camera(curl)}
	return
}

func (capture *VideoCapture) SaveImage(filename string) {
	var cfilename *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	C.save_image(capture.p, cfilename)
	return
}
