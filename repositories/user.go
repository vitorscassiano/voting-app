package repositories

import (
	"github.com/vitorscassiano/voting-app/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(connection *gorm.DB) *UserRepository {
	return &UserRepository{db: connection}
}

func (r *UserRepository) FindUserByEmail(email string) entities.User {
	var output entities.User
	r.db.Where("email = ?", email).Find(&output)

	return output
}

func (r *UserRepository) CreateUser(user *entities.User) {
	r.db.Create(&user)
}

func (r *UserRepository) FindUserByCPF(cpf string) entities.User {
	return entities.User{}
}

func (r *UserRepository) FindUserByName(name string) entities.User {
	return entities.User{}
}
