package data

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Filename string
}

type Detection struct {
	gorm.Model
	Person  float32
	ImageID uint
	Image   Image
}
