package services

import "github.com/vitorscassiano/voting-app/entities"

type Election struct {
	electionRepo ElectionRepository
}

func NewElectionService(electionRepo ElectionRepository) *Election {
	return &Election{
		electionRepo: electionRepo,
	}
}

func (e *Election) CreateElection(election *entities.Election) error {
	err := e.electionRepo.CreateElection(election)
	if err != nil {
		return err
	}

	return nil
}
