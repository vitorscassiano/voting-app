package repositories

import (
	"github.com/vitorscassiano/voting-app/entities"
	"gorm.io/gorm"
)

type CandidateRepository struct {
	db *gorm.DB
}

func NewCandidateRepository(connection *gorm.DB) *CandidateRepository {
	return &CandidateRepository{db: connection}
}

func (r *CandidateRepository) Candidates(email string) ([]entities.Candidate, error) {
	return []entities.Candidate{}, nil
}

func (r *CandidateRepository) CreateCandidate(candidate *entities.Candidate) error {
	tx := r.db.Create(&candidate)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *CandidateRepository) GetCandidate(candidateID string) (*entities.Candidate, error) {
	var output entities.Candidate

	tx := r.db.Where("candidate_id = ?", candidateID).Find(&output)
	if err := tx.Error; err != nil {
		return &entities.Candidate{}, err
	}

	return &output, nil
}
