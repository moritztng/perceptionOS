package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

type ImageWithDetection struct {
	ID        uint
	Filename  string
	Detection *Detection
}

func Open(filename string) Database {
	db, err := gorm.Open(sqlite.Open("images.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Image{}, &Detection{})
	database := Database{db: db}
	return database
}

func (database *Database) AddImage(filename string) Image {
	image := Image{Filename: filename}
	database.db.Create(&image)
	return image
}

func (database *Database) AddDetection(imageID uint, person float32) Detection {
	detection := Detection{ImageID: imageID, Person: person}
	database.db.Create(&detection)
	return detection
}

func (database *Database) GetAllImages() []*ImageWithDetection {
	result := []struct {
		ID       uint
		Filename string
		Person   *float32
	}{}
	database.db.Model(&Image{}).Select("images.id, images.filename, detections.person").Joins("left join detections on images.id = detections.image_id").Scan(&result)
	output := make([]*ImageWithDetection, len(result))
	for index, image := range result {
		var detection *Detection
		if image.Person == nil {
			detection = nil
		} else {
			detection = &Detection{Person: *image.Person}
		}
		output[index] = &ImageWithDetection{ID: image.ID, Filename: image.Filename, Detection: detection}
	}
	return output
}

func (database *Database) GetImage(imageID uint) *ImageWithDetection {
	image := struct {
		ID       uint
		Filename string
		Person   *float32
	}{ID: imageID}
	database.db.Model(&Image{}).Select("images.id, images.filename, detections.person").Joins("left join detections on images.id = detections.image_id").First(&image, imageID)
	var detection *Detection
	if image.Person == nil {
		detection = nil
	} else {
		detection = &Detection{Person: *image.Person}
	}
	output := &ImageWithDetection{ID: image.ID, Filename: image.Filename, Detection: detection}
	return output
}

/*func (database *Database) GetUnprocessedImages() []*ImageWithFaceDetected {
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
*/
