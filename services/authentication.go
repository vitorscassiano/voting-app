package services

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vitorscassiano/voting-app/repositories"
	"golang.org/x/crypto/bcrypt"
)

// This value will be regenerated and moved to configurations
const SECRET_KEY = "Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9HixkmBhVrYaB0NhtHpHgAWeTnLZpTSxCKs0gigByk5SH9pmQ=="

type AuthenticationService interface{}

type aservice struct {
	userRepo repositories.UserRepository
}

func NewAuthenticationService(userRepo repositories.UserRepository) AuthenticationService {
	return &aservice{userRepo: userRepo}
}

func UserAuthenticated(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func HashPassword(password string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	return string(passwordHash)
}

func GenerateJWTToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRET_KEY))
}
