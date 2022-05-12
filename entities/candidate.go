package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Candidate struct {
	CandidateID string    `json:"candidateId" gorm:"primaryKey"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

func (c *Candidate) BeforeCreate(scope *gorm.DB) (err error) {
	c.CandidateID = uuid.NewString()
	return
}
