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
	db.AutoMigrate(
		entities.Authentication{},
		entities.User{},
		entities.Candidate{},
		entities.Election{},
	)

	// Users
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := controllers.NewUserHandler(userService)

	// Authorizers
	authenticationService := services.NewAuthenticationService(userRepository)
	authorizationHandler := controllers.NewAuthenticationController(authenticationService, userService)

	// Candidates
	candidateRepository := repositories.NewCandidateRepository(db)
	candidateServices := services.NewCandidateService(candidateRepository)
	candidateHandler := controllers.NewCandidateHandler(candidateServices)

	// Elections
	electionRepository := repositories.NewElectionRepository(db)
	electionService := services.NewElectionService(electionRepository)
	electionHandler := controllers.NewElectionHandler(electionService)

	r.POST("api/v1/users", userHandler.CreateUser)
	r.POST("api/v1/login", authorizationHandler.Authorize)

	r.POST("api/v1/candidates", candidateHandler.PostCandidateHandler)
	r.GET("api/v1/candidates/:id", candidateHandler.GetCandidateHandler)

	r.POST("api/v1/elections", electionHandler.CreateElection)

	// r.POST("api/v1/elections/candidates/:candidate-id", func(ctx *gin.Context) {})
	// r.GET("api/v1/elections/candidates", func(ctx *gin.Context) {})

	// r.POST("api/v1/elections/candidates/:candidate-id/votes", func(ctx *gin.Context) {})

	r.Run()
}
