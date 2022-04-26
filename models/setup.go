package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	database, err := gorm.Open(sqlite.Open("voting-app.db"), &gorm.Config{})
	if err != nil {
		panic("database connection does not start")
	}

	database.AutoMigrate(&Authentication{}, &User{})

	DB = database
}
