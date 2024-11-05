package domain

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
)

const (
	maxGoals = 10
)

type (
	Match struct {
		ID        uuid.UUID   `json:"id"`
		Passed    bool        `json:"passed"`
		Team1     TeamInMatch `json:"team_1"`
		Team2     TeamInMatch `json:"team_2"`
		Iteration int         `json:"iteration"`
	}

	Matches []Match

	TeamInMatch struct {
		Team  `json:"team"`
		Goals int `json:"goals"`
	}
)

func (m *Match) GetWinner() (TeamInMatch, bool) {
	if m.Team1.Goals == m.Team2.Goals {
		return TeamInMatch{}, false
	}

	if m.Team1.Goals > m.Team2.Goals {
		return m.Team1, true
	}
	return m.Team2, true
}

func (m *Match) GetLooser() (TeamInMatch, bool) {
	if m.Team1.Goals == m.Team2.Goals {
		return TeamInMatch{}, false
	}

	if m.Team1.Goals > m.Team2.Goals {
		return m.Team2, true
	}
	return m.Team1, true
}

func (m *Match) GetScore() string {
	return fmt.Sprintf("%d:%d", m.Team1.Goals, m.Team2.Goals)
}

func (m *Match) Play() {
	if m.Passed {
		return
	}

	m.Team1.GenerateMatchResult()
	m.Team2.GenerateMatchResult()

	if _, ok := m.GetWinner(); !ok {
		log.Println("play match again: ", m.Team1.Goals, ":", m.Team2.Goals)
		m.Play()
		return
	}

	m.Team1.CalcRating(m.Team2.Goals)
	m.Team2.CalcRating(m.Team1.Goals)

	m.Team1.SetTeamStatus()
	m.Team2.SetTeamStatus()
}

func (t *TeamInMatch) GenerateMatchResult() {
	t.Goals = rand.Intn(maxGoals)
}

func (t *TeamInMatch) CalcRating(missedGoals int) {
	t.Rating = t.Rating + t.Goals - missedGoals
}

func (t *TeamInMatch) SetTeamStatus() {
	if t.TeamStatus == TeamStatusWinning || t.TeamStatus == TeamStatusLoosing {
		return
	}

	if t.Rating > 0 {
		t.TeamStatus = TeamStatusWinning
		return
	}
	t.TeamStatus = TeamStatusLoosing
}

func (matches Matches) GetWinners(iteration int) []TeamInMatch {
	winners := make([]TeamInMatch, 0, len(matches)/2)
	for _, match := range matches {
		if match.Iteration != iteration {
			continue
		}

		winner, ok := match.GetWinner()
		if ok {
			winners = append(winners, winner)
		}
	}
	return winners
}

func (matches Matches) GetLoosers(iteration int) []TeamInMatch {
	loosers := make([]TeamInMatch, 0, len(matches)/2)
	for _, match := range matches {
		if match.Iteration != iteration {
			continue
		}

		looser, ok := match.GetLooser()
		if ok {
			loosers = append(loosers, looser)
		}
	}
	return loosers
}

func (matches Matches) GetNotPassed() Matches {
	results := make(Matches, 0, len(matches))
	for _, match := range matches {
		if !match.Passed {
			results = append(results, match)
		}
	}
	return results
}
