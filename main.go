package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/models"
	"github.com/vitorscassiano/voting-app/services"
)

func main() {
	r := gin.Default()

	models.InitializeDatabase()

	r.POST("api/v1/users", services.CreateUser)
	r.POST("api/v1/login", services.Authorize)

	r.Run()
}
