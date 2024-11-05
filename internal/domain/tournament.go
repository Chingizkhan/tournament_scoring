package domain

import "github.com/google/uuid"

type (
	Tournament struct {
		ID             uuid.UUID  `json:"id"`
		Divisions      []Division `json:"divisions"`
		PlayOffMatches []Match    `json:"play_off_matches"`
		Winner         *Team      `json:"winner"`
	}
)
