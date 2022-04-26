package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/database"
	"github.com/vitorscassiano/voting-app/wires"
	"golang.org/x/crypto/bcrypt"
)

/*
func Authorize(email string, password string) bool {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("something wrong with hashing password")
	}

	string(hashPassword)

	return true
}
*/

func userAuthenticated(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func Authorize(ctx *gin.Context) {
	var input wires.PostAuthentication

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := database.FindUserByEmail(input.Email)
	if err := userAuthenticated(user.Password, input.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email or password is wrong."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "OK"})
}
