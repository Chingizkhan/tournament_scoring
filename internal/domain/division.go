package domain

import (
	"github.com/google/uuid"
)

type (
	Division struct {
		ID           uuid.UUID    `json:"id"`
		Name         DivisionName `json:"name"` // A or B
		Teams        []Team       `json:"teams"`
		TournamentID uuid.UUID    `json:"tournament_id"`
		Matches      Matches      `json:"matches"`
	}

	DivisionName string
)

const (
	DivisionA DivisionName = "A"
	DivisionB DivisionName = "B"
)

var (
	DivisionsList = []Division{
		{
			Name: DivisionA,
		},
		{
			Name: DivisionB,
		},
	}
)

func (d *Division) GetTeamsLength() int {
	return len(d.Teams)
}

func (d *Division) IsFinished() bool {
	for _, match := range d.Matches {
		if !match.Passed {
			return false
		}
	}
	return true
}
