package utils

import (
	"github.com/ansi13/doctor-api-service/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func InitDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.Doctor{}, &models.Patient{})

	DB = db
}
