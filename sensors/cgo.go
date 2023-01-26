package sensors

// #cgo LDFLAGS: -L. -lstdc++ -lopencv_core -lopencv_videoio -lopencv_imgcodecs
// #cgo CXXFLAGS: -std=c++14 -I. -I/usr/include/opencv2
// #include "camera.h"
import "C"
