package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/entities"
	"github.com/vitorscassiano/voting-app/repositories"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("something wrong with hashing password")
	}

	return string(hashPassword)
}

func CreateUser(ctx *gin.Context) {
	var input entities.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		Name:     input.Name,
		Email:    input.Email,
		CPF:      input.CPF,
		Password: hashPassword(input.Password),
	}

	repositories.CreateUser(&user)

	ctx.JSON(http.StatusAccepted, gin.H{"user": user})
}
