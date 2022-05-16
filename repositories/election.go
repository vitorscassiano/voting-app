package repositories

import (
	"github.com/vitorscassiano/voting-app/entities"
	"gorm.io/gorm"
)

type Election struct {
	db *gorm.DB
}

func NewElectionRepository(db *gorm.DB) *Election {
	return &Election{db: db}
}

func (e *Election) CreateElection(election *entities.Election) error {
	tx := e.db.Create(election)
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
