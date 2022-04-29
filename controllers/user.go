package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/wires"
)

func CreateUser(ctx *gin.Context) {
	var input wires.PostUser

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
}
