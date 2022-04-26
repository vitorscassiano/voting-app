package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Candidate struct {
	CandidateID string `json:"candidateID"`
	Name        string `json:"name"`
}

type Submit struct {
	SubmitID  string    `json:"submitID"`
	User      User      `json:"user"`
	Candidate Candidate `json:"candidate"`
}

type Voting struct {
	VotingID    string      `json:"votingID"`
	Description string      `json:"description"`
	Candidates  []Candidate `json:"candidates"`
	Submits     []Submit    `json:"submits"`
}

func (v *Voting) BeforeCreate(scope *gorm.DB) (err error) {
	v.VotingID = uuid.NewString()
	return
}

func (c *Candidate) BeforeCreate(scope *gorm.DB) (err error) {
	c.CandidateID = uuid.NewString()
	return
}

func (s *Submit) BeforeCreate(scope *gorm.DB) (err error) {
	s.SubmitID = uuid.NewString()
	return
}
