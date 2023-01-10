package storage

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
	ImageID      int
	Image        Image `gorm:"embedded"`
}
