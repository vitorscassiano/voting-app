package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/controllers"
	"github.com/vitorscassiano/voting-app/entities"
	"github.com/vitorscassiano/voting-app/infrastructure/database"
	"github.com/vitorscassiano/voting-app/repositories"
	"github.com/vitorscassiano/voting-app/services"
)

func main() {
	r := gin.Default()

	db := database.InitializeRepositories()
	db.AutoMigrate(entities.Authentication{}, entities.User{})

	userRepository := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepository)
	authenticationService := services.NewAuthenticationService(userRepository)

	authorizationHandler := controllers.NewAuthenticationController(authenticationService, userService)
	userHandler := controllers.NewUserHandler(userService)

	r.POST("api/v1/users", userHandler.CreateUser)
	r.POST("api/v1/login", authorizationHandler.Authorize)

	r.POST("api/v1/candidates")

	r.Run()
}
