package repositories

import (
	"github.com/vitorscassiano/voting-app/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeRepositories() *gorm.DB {
	repositories, err := gorm.Open(sqlite.Open("voting-app.db"), &gorm.Config{})
	if err != nil {
		panic("repositories connection does not start")
	}

	repositories.AutoMigrate(&entities.Authentication{}, &entities.User{})

	return repositories
}
