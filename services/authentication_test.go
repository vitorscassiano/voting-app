package services_test

import (
	"testing"

	"github.com/vitorscassiano/voting-app/repositories"
	"github.com/vitorscassiano/voting-app/services"
)

func TestAuthentication(t *testing.T) {
	authService := services.AuthenticationService{}
	userRepository := repositories.UserRepository{}

}
