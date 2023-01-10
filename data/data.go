package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

type ImageWithFaceDetected struct {
	ID           uint
	Filename     string
	FaceDetected *FaceDetected
}

func Open(filename string) Database {
	db, err := gorm.Open(sqlite.Open("images.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Image{}, &FaceDetected{})
	database := Database{db: db}
	return database
}

func (database *Database) AddImage(filename string) Image {
	image := Image{Filename: filename}
	database.db.Create(&image)
	return image
}

func (database *Database) AddFaceDetected(imageID uint, faceDetected bool) FaceDetected {
	detected := FaceDetected{ImageID: imageID, FaceDetected: faceDetected}
	database.db.Create(&detected)
	return detected
}

func (database *Database) GetAllImages() []*ImageWithFaceDetected {
	result := []struct {
		ID           uint
		Filename     string
		FaceDetected *bool
	}{}
	database.db.Model(&Image{}).Select("images.id, images.filename, face_detecteds.face_detected").Joins("left join face_detecteds on images.id = face_detecteds.image_id").Scan(&result)
	output := make([]*ImageWithFaceDetected, len(result))
	for index, image := range result {
		var faceDetected *FaceDetected
		if image.FaceDetected == nil {
			faceDetected = nil
		} else {
			faceDetected = &FaceDetected{FaceDetected: *image.FaceDetected}
		}
		output[index] = &ImageWithFaceDetected{ID: image.ID, Filename: image.Filename, FaceDetected: faceDetected}
	}
	return output
}

func (database *Database) GetUnprocessedImages() []*ImageWithFaceDetected {
	result := []struct {
		ID           uint
		Filename     string
		FaceDetected *bool
	}{}
	database.db.Model(&Image{}).Select("images.id, images.filename, face_detecteds.face_detected").Joins("left join face_detecteds on images.id = face_detecteds.image_id").Where("face_detecteds.id is NULL").Scan(&result)
	output := make([]*ImageWithFaceDetected, len(result))
	for index, image := range result {
		var faceDetected *FaceDetected
		if image.FaceDetected == nil {
			faceDetected = nil
		} else {
			faceDetected = &FaceDetected{FaceDetected: *image.FaceDetected}
		}
		output[index] = &ImageWithFaceDetected{ID: image.ID, Filename: image.Filename, FaceDetected: faceDetected}
	}
	return output
}
