package division

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/team_repo"
)

func (s *Service) SeparateAndSaveTeams(ctx context.Context, divisions []domain.Division, teams []domain.Team) error {
	s.separateTeams(divisions, teams)
	if err := s.saveTeams(ctx, divisions); err != nil {
		return err
	}

	return nil
}

func (s *Service) separateTeams(divisions []domain.Division, teams domain.Teams) {
	divisionsCount := len(divisions)

	teams.Shuffle()

	for i, team := range teams {
		divisions[i%divisionsCount].Teams = append(divisions[i%divisionsCount].Teams, team)
	}
}

func (s *Service) saveTeams(ctx context.Context, divisions []domain.Division) error {
	for i, division := range divisions {

		teams, err := s.team.Save(ctx, team_repo.SaveIn{
			Teams: s.prependSaveTeams(division.Teams, division.ID),
		})
		if err != nil {
			return fmt.Errorf("save teams: %w", err)
		}

		divisions[i].Teams = teams
	}

	return nil
}

func (s *Service) prependSaveTeams(teams []domain.Team, divisionID uuid.UUID) []team_repo.SaveParams {
	in := make([]team_repo.SaveParams, 0, len(teams))

	for _, team := range teams {
		in = append(in, team_repo.SaveParams{
			Name:       team.Name,
			DivisionID: divisionID,
		})
	}

	return in
}
