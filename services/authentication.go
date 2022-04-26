package services

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/database"
	"github.com/vitorscassiano/voting-app/wires"
	"golang.org/x/crypto/bcrypt"
)

// This value will be regenerated and moved to configurations
const SECRET_KEY = "Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9HixkmBhVrYaB0NhtHpHgAWeTnLZpTSxCKs0gigByk5SH9pmQ=="

func userAuthenticated(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func generateJWTToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET_KEY))
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

	jwtString, err := generateJWTToken(user.UserID)
	if err != nil {
		fmt.Println("something wrong with jwt generation", err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jwtString})
}
