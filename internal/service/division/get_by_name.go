package division

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/match_repo"
	"tournament_scoring/internal/repo/team_repo"
)

func (s *Service) GetByName(ctx context.Context, name domain.DivisionName) (division domain.Division, err error) {
	division, err = s.division.GetByName(ctx, name)
	if err != nil {
		return division, fmt.Errorf("get division by name: %w", err)
	}
	division.Matches, err = s.match.Find(ctx, match_repo.FindIn{
		DivisionID: division.ID,
	})
	if err != nil {
		return division, fmt.Errorf("find match: %w", err)
	}

	division.Teams, err = s.team.Find(ctx, team_repo.FindIn{
		DivisionID: division.ID,
	})
	if err != nil {
		return division, fmt.Errorf("find teams: %w", err)
	}

	return division, nil
}
