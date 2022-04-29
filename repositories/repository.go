package repositories

import "github.com/vitorscassiano/voting-app/entities"

type UserRepository interface {
	CreateUser(user *entities.User)
	FindUserByEmail(email string) entities.User
	FindUserByCPF(cpf string) entities.User
	FindUserByName(name string) entities.User
}
