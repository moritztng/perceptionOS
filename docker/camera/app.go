package main


import (
    "fmt"
    "log"
    "net/http"
)

// #cgo LDFLAGS: -L. -lstdc++ -lopencv_core -lopencv_videoio -lopencv_imgcodecs
// #cgo CXXFLAGS: -std=c++14 -I. -I/usr/include/opencv2
// #include "camera.h"
import "C"
type VideoCapture struct {
	p C.VideoCapture
}
func handler(w http.ResponseWriter, r *http.Request) {
    capture := VideoCapture{p: C.camera(C.CString("hi"))} 
    C.save_image(capture.p, C.CString("hi"));
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
