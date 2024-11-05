package division

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/match_repo"
	"tournament_scoring/internal/repo/team_repo"
)

func (s *Service) List(ctx context.Context) ([]domain.Division, error) {
	divisions, err := s.division.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("division.List: %w", err)
	}

	for i, d := range divisions {
		if divisions[i].Teams, err = s.team.Find(ctx, team_repo.FindIn{DivisionID: d.ID}); err != nil {
			return nil, fmt.Errorf("s.team.Find: %w", err)
		}

		if divisions[i].Matches, err = s.match.Find(ctx, match_repo.FindIn{DivisionID: d.ID}); err != nil {
			return nil, fmt.Errorf("s.match.Find: %w", err)
		}
	}

	return divisions, nil
}
