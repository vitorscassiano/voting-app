package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/entities"
	"github.com/vitorscassiano/voting-app/services"
	"github.com/vitorscassiano/voting-app/wires"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	var input wires.PostUser

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	user := entities.User{
		Name:     input.Name,
		CPF:      input.CPF,
		Email:    input.Email,
		Password: services.HashPassword(input.Password),
	}
	u.userService.CreateUser(&user)

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
