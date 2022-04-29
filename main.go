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

	controller := controllers.NewAuthenticationController(authenticationService, userService)

	r.POST("api/v1/users", controllers.CreateUser)
	r.POST("api/v1/login", controller.Authorize)

	r.Run()
}
