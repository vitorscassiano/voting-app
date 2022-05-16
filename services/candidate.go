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

func (c *CandidateService) CreateCandidate(candidate *entities.Candidate) error {
	err := c.candidateRepo.CreateCandidate(candidate)
	if err != nil {
		return err
	}

	return nil
}

func (c *CandidateService) GetCandidate(candidateID string) (*entities.Candidate, error) {
	candidate, err := c.candidateRepo.GetCandidate(candidateID)
	if err != nil {
		return &entities.Candidate{}, err
	}

	return candidate, nil
}
func (c *CandidateService) UpdateCandidate(candidateID string, candidate *entities.Candidate) (*entities.Candidate, error) {
	return &entities.Candidate{}, nil
}
