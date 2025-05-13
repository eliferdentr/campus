package db

import (
	"campus-project.com/study-service/internal/models"
	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("user-service.db"), &gorm.Config{})
	if err != nil {
		panic ("There was an error while connecting to the database. Error message: " + err.Error())
	}
	DB.AutoMigrate(&models.Group{})
	DB.AutoMigrate()
}