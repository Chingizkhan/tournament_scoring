package match

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/match_repo"
	"tournament_scoring/internal/repo/team_repo"
)

func (s *Service) Update(ctx context.Context, matches []domain.Match) error {

	for _, match := range matches {
		if match.Passed {
			continue
		}

		if err := s.match.Update(ctx, match_repo.UpdateIn{
			ID:              match.ID,
			Passed:          true,
			GoalsFirstTeam:  match.Team1.Goals,
			GoalsSecondTeam: match.Team2.Goals,
		}); err != nil {
			return fmt.Errorf("update match(id:%s): %w", match.ID, err)
		}

		var (
			team1 = match.Team1
			team2 = match.Team2
		)
		if err := s.team.Update(ctx, team_repo.UpdateIn{
			ID:         team1.ID,
			Rating:     team1.Rating,
			TeamStatus: team1.TeamStatus,
		}); err != nil {
			return fmt.Errorf("update team1(id:%s;name:%s): %w", team1.ID, team1.Name, err)
		}

		if err := s.team.Update(ctx, team_repo.UpdateIn{
			ID:         team2.ID,
			Rating:     team2.Rating,
			TeamStatus: team2.TeamStatus,
		}); err != nil {
			return fmt.Errorf("update team2(id:%s;name:%s): %w", team2.ID, team2.Name, err)
		}
	}

	return nil
}
