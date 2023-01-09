package goperception

/*
#include "camera.h"
*/
import "C"

type VideoCapture struct {
	p C.VideoCapture
}

func (VideoCapture capture) SaveImage(string filename) {
    var cfilename *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	capture.p.save_image(cfilename)
    return
}
