package services_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
	"github.com/vitorscassiano/voting-app/repositories"
	"github.com/vitorscassiano/voting-app/services"
)

func TestAuthentication(t *testing.T) {
	authService := services.AuthenticationService{}
	userRepository := repositories.UserRepository{}

	assert.Equal(t, 1, 1)
}
