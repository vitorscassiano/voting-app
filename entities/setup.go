package entities

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initializerepositories() {
	repositories, err := gorm.Open(sqlite.Open("voting-app.db"), &gorm.Config{})
	if err != nil {
		panic("repositories connection does not start")
	}

	repositories.AutoMigrate(&Authentication{}, &User{})

	DB = repositories
}
