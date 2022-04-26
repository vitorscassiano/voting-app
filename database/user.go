package database

import (
	"github.com/vitorscassiano/voting-app/models"
)

func CreateUser(user *models.User) {
	models.DB.Create(&user)
}

func FindUserByEmail(email string) models.User {
	var output models.User

	models.DB.Where("email = ?", email).Find(&output)

	return output
}

func FindUserByCPF()  {}
func FindUserByName() {}
