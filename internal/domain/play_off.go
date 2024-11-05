package domain

import "github.com/google/uuid"

type (
	PlayOff struct {
		ID           uuid.UUID `json:"id"`
		Teams        []Team    `json:"teams"`
		TournamentID uuid.UUID `json:"tournament_id"`
		Matches      Matches   `json:"matches"`
		Winner       uuid.UUID `json:"winner"`
	}
)

func (p *PlayOff) IsFinished() bool {
	for _, match := range p.Matches {
		if !match.Passed {
			return false
		}
	}
	return true
}
