package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorscassiano/voting-app/controllers/wires"
	"github.com/vitorscassiano/voting-app/entities"
)

type CandidateHandler struct {
	candidateService CandidateService
}

func NewCandidateHandler(candidateHandler CandidateService) *CandidateHandler {
	return &CandidateHandler{
		candidateService: candidateHandler,
	}
}

func (h *CandidateHandler) PostCandidateHandler(ctx *gin.Context) {
	var input wires.PostCandidateIn
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"errors": err.Error()})
		return
	}

	candidate := entities.Candidate{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	h.candidateService.CreateCandidate(&candidate)

	ctx.JSON(http.StatusAccepted, map[string]entities.Candidate{"candidate": candidate})
}

func (h *CandidateHandler) GetCandidateHandler(ctx *gin.Context) {
	var candidateID = ctx.Param("id")

	candidate, err := h.candidateService.GetCandidate(candidateID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]entities.Candidate{"candidate": *candidate})
}
