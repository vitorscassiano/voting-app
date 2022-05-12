package controllers

import "github.com/vitorscassiano/voting-app/entities"

type UserService interface {
	FindUserByEmail(email string) entities.User
	CreateUser(user *entities.User)
}

type AuthenticationService interface{}
