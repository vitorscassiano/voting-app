package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/controllers/wires"
	"github.com/vitorscassiano/voting-app/services"
)

type Authentication struct {
	aservice AuthenticationService
	uservice UserService
}

func NewAuthenticationController(aservice AuthenticationService, uservice UserService) *Authentication {
	return &Authentication{aservice: aservice, uservice: uservice}
}

func (a *Authentication) Authorize(ctx *gin.Context) {
	var input wires.PostAuthentication

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := a.uservice.FindUserByEmail(input.Email)
	if err := services.UserAuthenticated(user.Password, input.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email or password is wrong."})
		return
	}

	jwtString, err := services.GenerateJWTToken(user.UserID)
	if err != nil {
		fmt.Println("something wrong with jwt generation", err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jwtString})

}
