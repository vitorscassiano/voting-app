package services

import (
	"github.com/vitorscassiano/voting-app/entities"
)

type CandidateService struct {
	candidateRepo CandidateRepository
}

func NewCandidateService(candidateRepo CandidateRepository) *CandidateService {
	return &CandidateService{candidateRepo: candidateRepo}
}

func (s *UserService) Candidates(email string) entities.User {
	return s.userRepo.FindUserByEmail(email)
}

func (s *UserService) CreateCandidate(user *entities.User) {
	s.userRepo.CreateUser(user)
}
