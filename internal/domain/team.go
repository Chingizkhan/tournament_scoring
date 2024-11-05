package domain

import (
	"github.com/google/uuid"
	"math/rand"
	"sort"
	"time"
)

type (
	TeamStatus string

	Team struct {
		ID         uuid.UUID  `json:"id"`
		Name       string     `json:"name"`
		DivisionID uuid.UUID  `json:"division_id"`
		PlayOffID  uuid.UUID  `json:"play_off_id"`
		Rating     int        `json:"rating"`
		TeamStatus TeamStatus `json:"team_status"`
	}

	Teams []Team
)

var (
	TeamStatusWinning TeamStatus = "winning"
	TeamStatusLoosing TeamStatus = "loosing"
	TeamStatusPrepare TeamStatus = "prepare"
)

func (t Teams) SortDesc() {
	sort.Slice(t, func(i, j int) bool {
		return t[i].Rating > t[j].Rating
	})
}

func (t Teams) Shuffle() {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(t), func(i, j int) { t[i], t[j] = t[j], t[i] })
}
