package repositories

import (
	"github.com/vitorscassiano/voting-app/entities"
	"gorm.io/gorm"
)

type repo struct {
	Connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &repo{Connection: connection}
}

func (r *repo) FindUserByEmail(email string) entities.User {
	var output entities.User
	r.Connection.Where("email = ?", email).Find(&output)

	return output
}

func (r *repo) CreateUser(user *entities.User) {
	r.Connection.Create(&user)
}

func (r *repo) FindUserByCPF(cpf string) entities.User {
	return entities.User{}
}

func (r *repo) FindUserByName(name string) entities.User {
	return entities.User{}
}
