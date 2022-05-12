package services

import "github.com/vitorscassiano/voting-app/entities"

type UserRepository interface {
	CreateUser(user *entities.User)
	FindUserByEmail(email string) entities.User
	FindUserByCPF(cpf string) entities.User
	FindUserByName(name string) entities.User
}

type CandidateRepository interface {
	Candidates(email string) ([]entities.Candidate, error)
	CreateCandidate(candidate *entities.Candidate) error
}
