package data

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Filename string
}

type FaceDetected struct {
	gorm.Model
	FaceDetected bool
	ImageID      uint
	Image        Image
}
