package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Election struct {
	VotingID    string      `json:"votingId" gorm:"primaryKey"`
	Description string      `json:"description"`
	Candidates  []Candidate `json:"candidates"`
	Submits     []Vote      `json:"vote"`
	CreatedAt   time.Time   `json:"createdAt,omitempty"`
}

func (v *Election) BeforeCreate(scope *gorm.DB) (err error) {
	v.VotingID = uuid.NewString()
	return
}
