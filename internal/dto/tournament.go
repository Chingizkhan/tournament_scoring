package dto

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"io"
	"net/http"
	"tournament_scoring/internal/domain"
)

const numberOfTeams = 16

type CreateTournamentIn struct {
	Teams []domain.Team `json:"teams"`
}

func (in *CreateTournamentIn) Parse(req io.ReadCloser) error {
	teamNames := make([]string, 0, numberOfTeams)
	if err := json.NewDecoder(req).Decode(&teamNames); err != nil {
		return fmt.Errorf("could not decode request body: %w", err)
	}

	for _, name := range teamNames {
		in.Teams = append(in.Teams, domain.Team{Name: name})
	}

	return nil
}

func (in *CreateTournamentIn) Validate() error {
	if len(in.Teams) != numberOfTeams {
		return errors.New("must be 16 teams")
	}
	return nil
}

type (
	Team struct {
		ID     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		Rating int       `json:"rating"`
		Goals  int       `json:"goals"`
	}

	Match struct {
		ID         uuid.UUID `json:"id"`
		Passed     bool      `json:"passed"`
		FirstTeam  Team      `json:"first_team"`
		SecondTeam Team      `json:"second_team"`
	}

	Division struct {
		ID      uuid.UUID           `json:"id"`
		Name    domain.DivisionName `json:"name"`
		Matches []Match             `json:"matches"`
	}

	CreateTournamentOut struct {
		TournamentID uuid.UUID  `json:"tournament_id"`
		Divisions    []Division `json:"divisions"`
	}
)

func (out *CreateTournamentOut) ConvertResponse(divisions []domain.Division) CreateTournamentOut {
	return CreateTournamentOut{
		TournamentID: divisions[0].TournamentID,
		Divisions:    out.convDivisions(divisions),
	}
}

func (out *CreateTournamentOut) convDivisions(divisions []domain.Division) []Division {
	res := make([]Division, 0, len(divisions))

	for _, division := range divisions {
		res = append(res, Division{
			ID:      division.ID,
			Name:    division.Name,
			Matches: out.convMatches(division.Matches),
		})
	}

	return res
}

func (out *CreateTournamentOut) convMatches(matches []domain.Match) []Match {
	res := make([]Match, 0, len(matches))

	for _, match := range matches {
		res = append(res, Match{
			ID:     match.ID,
			Passed: match.Passed,
			FirstTeam: Team{
				ID:     match.Team1.ID,
				Name:   match.Team1.Name,
				Rating: match.Team1.Rating,
				Goals:  match.Team1.Goals,
			},
			SecondTeam: Team{
				ID:     match.Team2.ID,
				Name:   match.Team2.Name,
				Rating: match.Team2.Rating,
				Goals:  match.Team2.Goals,
			},
		})
	}

	return res
}

type DivisionResultIn struct {
	DivisionName domain.DivisionName
}

func (in *DivisionResultIn) Parse(r *http.Request) error {
	in.DivisionName = domain.DivisionName(chi.URLParam(r, "divisionName"))
	return nil
}

type DivisionResultOut struct {
	TournamentID uuid.UUID `json:"tournament_id"`
	Division     Division  `json:"division"`
}

func (dto *DivisionResultOut) ConvertResponse(division domain.Division) DivisionResultOut {
	return DivisionResultOut{
		TournamentID: division.TournamentID,
		Division: Division{
			ID:      division.ID,
			Name:    division.Name,
			Matches: dto.convMatches(division.Matches),
		},
	}
}

func (dto *DivisionResultOut) convMatches(matches []domain.Match) []Match {
	res := make([]Match, 0, len(matches))

	for _, match := range matches {
		res = append(res, Match{
			ID:     match.ID,
			Passed: match.Passed,
			FirstTeam: Team{
				ID:     match.Team1.ID,
				Name:   match.Team1.Name,
				Rating: match.Team1.Rating,
				Goals:  match.Team1.Goals,
			},
			SecondTeam: Team{
				ID:     match.Team2.ID,
				Name:   match.Team2.Name,
				Rating: match.Team2.Rating,
				Goals:  match.Team2.Goals,
			},
		})
	}

	return res
}
