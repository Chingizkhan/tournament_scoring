package match_repo

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"tournament_scoring/internal/domain"
)

type (
	Team struct {
		ID    uuid.UUID
		Goals int
	}

	MatchIn struct {
		DivisionID uuid.UUID
		FirstTeam  Team
		SecondTeam Team
	}

	MatchesIn []MatchIn

	Match struct {
		ID              uuid.UUID `json:"id"`
		Passed          bool      `json:"passed"`
		FirstTeamID     uuid.UUID `json:"first_team_id"`
		SecondTeamID    uuid.UUID `json:"second_team_id"`
		GoalsFirstTeam  int       `json:"goals_first_team"`
		GoalsSecondTeam int       `json:"goals_second_team"`
		Iteration       int       `json:"iteration"`
	}

	Matches []Match

	MatchFullInfo struct {
		Match
		DivisionID          uuid.UUID         `json:"division_id"`
		PlayOffID           pgtype.UUID       `json:"play_off_id"`
		FirstTeamName       string            `json:"first_team_name"`
		FirstTeamDivisionID uuid.UUID         `json:"first_team_division_id"`
		FirstTeamRating     int               `json:"first_team_rating"`
		FirstTeamStatus     domain.TeamStatus `json:"first_team_status"`

		SecondTeamName       string            `json:"second_team_name"`
		SecondTeamDivisionID uuid.UUID         `json:"second_team_division_id"`
		SecondTeamRating     int               `json:"second_team_rating"`
		SecondTeamStatus     domain.TeamStatus `json:"second_team_status"`
	}

	MatchesFullInfo []MatchFullInfo
)

func (matchesOut MatchesFullInfo) convToDomain() []domain.Match {
	out := make([]domain.Match, 0, len(matchesOut))

	for _, match := range matchesOut {
		out = append(out, domain.Match{
			ID:        match.ID,
			Passed:    match.Passed,
			Iteration: match.Iteration,
			Team1: domain.TeamInMatch{
				Team: domain.Team{
					ID:         match.FirstTeamID,
					Name:       match.FirstTeamName,
					DivisionID: match.FirstTeamDivisionID,
					Rating:     match.FirstTeamRating,
					TeamStatus: match.FirstTeamStatus,
				},
				Goals: match.GoalsFirstTeam,
			},
			Team2: domain.TeamInMatch{
				Team: domain.Team{
					ID:         match.SecondTeamID,
					Name:       match.SecondTeamName,
					DivisionID: match.SecondTeamDivisionID,
					Rating:     match.SecondTeamRating,
					TeamStatus: match.SecondTeamStatus,
				},
				Goals: match.GoalsSecondTeam,
			},
		})
	}

	return out
}

func (matches Matches) convToDomain() domain.Matches {
	res := make(domain.Matches, 0, len(matches))
	for _, m := range matches {
		res = append(res, domain.Match{
			ID:     m.ID,
			Passed: m.Passed,
			Team1: domain.TeamInMatch{
				Team: domain.Team{
					ID: m.FirstTeamID,
				},
				Goals: m.GoalsFirstTeam,
			},
			Team2: domain.TeamInMatch{
				Team: domain.Team{
					ID: m.SecondTeamID,
				},
				Goals: m.GoalsSecondTeam,
			},
			Iteration: m.Iteration,
		})
	}
	return res
}
