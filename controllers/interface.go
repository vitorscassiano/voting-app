package controllers

import "github.com/vitorscassiano/voting-app/entities"

type UserService interface {
	FindUserByEmail(email string) entities.User
	CreateUser(user *entities.User)
}

type AuthenticationService interface{}

type CandidateService interface {
	CreateCandidate(user *entities.Candidate) error
	GetCandidate(candidateID string) (*entities.Candidate, error)
	UpdateCandidate(candidateID string, candidate *entities.Candidate) (*entities.Candidate, error)
}

type ElectionService interface {
	CreateElection(election *entities.Election) error
}
