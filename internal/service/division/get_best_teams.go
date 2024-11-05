package division

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/team_repo"
)

const (
	ratingColumn = "rating"
)

func (s *Service) GetBestTeams(ctx context.Context, divisionID uuid.UUID, n int) ([]domain.Team, error) {
	teams, err := s.team.Find(ctx, team_repo.FindIn{
		DivisionID: divisionID,
		OrderBy: team_repo.OrderBy{
			Column: ratingColumn,
			Desc:   true,
		},
		Limit: n,
	})
	if err != nil {
		return nil, fmt.Errorf("find teams: %w", err)
	}

	return teams, nil
}
