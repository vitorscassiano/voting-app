package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vote struct {
	VoteID    string    `json:"voteId" gorm:"primaryKey"`
	User      User      `json:"user"`
	Candidate Candidate `json:"candidate"`
	Election  Election  `json:"election"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (s *Vote) BeforeCreate(scope *gorm.DB) (err error) {
	s.VoteID = uuid.NewString()
	return
}
