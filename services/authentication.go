package services

import (
	"errors"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vitorscassiano/voting-app/repositories"
	"golang.org/x/crypto/bcrypt"
)

// This value will be regenerated and moved to configurations
const SECRET_KEY = "Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9HixkmBhVrYaB0NhtHpHgAWeTnLZpTSxCKs0gigByk5SH9pmQ=="

type AuthenticationService interface {
	AuthorizeByEmail(email, password string) (string, error)
	AuthorizeByCPF(cpf, password string)
}

type aservice struct {
	userRepo repositories.UserRepository
}

func NewAuthenticationService(userRepo repositories.UserRepository) AuthenticationService {
	return &aservice{userRepo: userRepo}

}

func (s *aservice) AuthorizeByEmail(email, password string) (string, error) {
	user := s.userRepo.FindUserByEmail(email) // Refatorar, se não encontrar user retornar error
	if err := userAuthenticated(user.Password, password); err != nil {
		return "", errors.New("email or password is wrong")
	}

	jwtString, err := generateJWTToken(user.UserID)
	if err != nil {
		log.Println("something wrong with jwt generation:", err.Error())
	}

	return jwtString, nil
}

func (s *aservice) AuthorizeByCPF(cpf, password string) {}

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

/*
func Authorize(ctx *gin.Context) {
	var input wires.PostAuthentication

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := repositories.FindUserByEmail(input.Email)
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
*/
