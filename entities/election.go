package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Election struct {
	ElectionID  string    `json:"electionId" gorm:"primaryKey"`
	Description string    `json:"description"`
	CandidateID []string  `json:"candidates" gorm:"foreignKey:"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

func (v *Election) BeforeCreate(scope *gorm.DB) (err error) {
	v.ElectionID = uuid.NewString()
	return
}
