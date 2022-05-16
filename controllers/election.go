package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/controllers/wires"
	"github.com/vitorscassiano/voting-app/entities"
)

type ElectionHandler struct {
	electionService ElectionService
}

func NewElectionHandler(electionService ElectionService) *ElectionHandler {
	return &ElectionHandler{
		electionService: electionService,
	}
}

func (e *ElectionHandler) CreateElection(ctx *gin.Context) {
	var input wires.PostElectionIn

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	var candidates []entities.Candidate
	for _, candidate := range input.Candidates {
		candidates = append(candidates, entities.Candidate{CandidateID: candidate})
	}

	election := entities.Election{
		Description: input.Description,
		Candidates:  candidates,
	}

	e.electionService.CreateElection(&election)

	ctx.JSON(http.StatusAccepted, map[string]entities.Election{"election": election})
}
