package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/controllers"
	"github.com/vitorscassiano/voting-app/repositories"
	"github.com/vitorscassiano/voting-app/services"
)

func main() {
	r := gin.Default()

	databaseConnection := repositories.InitializeRepositories()

	userRepository := repositories.NewUserRepository(databaseConnection)

	userService := services.NewUserService(userRepository)
	authenticationService := services.NewAuthenticationService(userRepository)

	authorizationHandler := controllers.NewAuthenticationController(authenticationService, userService)
	userHandler := controllers.NewUserHandler(userService)

	r.POST("api/v1/users", userHandler.CreateUser)
	r.POST("api/v1/login", authorizationHandler.Authorize)

	r.Run()
}
