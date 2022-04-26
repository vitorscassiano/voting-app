package repositories

import (
	"github.com/vitorscassiano/voting-app/entities"
)

func CreateUser(user *entities.User) {
	entities.DB.Create(&user)
}

func FindUserByEmail(email string) entities.User {
	var output entities.User

	entities.DB.Where("email = ?", email).Find(&output)

	return output
}

func FindUserByCPF()  {}
func FindUserByName() {}
