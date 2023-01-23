package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
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

func (database *Database) AddImage(filename string) *Image {
	image := Image{Filename: filename}
	database.db.Create(&image)
	return &image
}

func (database *Database) AddDetection(imageId uint, person float32) *Detection {
	detection := Detection{ImageID: imageId, Person: person}
	database.db.Create(&detection)
	return &detection
}

func (database *Database) GetImage(id uint) *Image {
	image := Image{}
	database.db.First(&image, id)
	return &image
}

func (database *Database) GetDetection(id uint) *Detection {
	detection := Detection{}
	database.db.First(&detection, id)
	return &detection
}

func (database *Database) GetDetectionOfImage(id uint) *Detection {
	detection := Detection{}
	database.db.Where(&Detection{ImageID: id}).First(&detection)
	return &detection
}
