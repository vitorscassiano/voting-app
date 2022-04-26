package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/entities"
	"github.com/vitorscassiano/voting-app/services"
)

func main() {
	r := gin.Default()

	entities.Initializerepositories()

	r.POST("api/v1/users", services.CreateUser)
	r.POST("api/v1/login", services.Authorize)

	r.Run()
}
