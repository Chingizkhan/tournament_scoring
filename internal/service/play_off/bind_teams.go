package play_off

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/team_repo"
)

func (s *Service) BindTeams(ctx context.Context, playOffID uuid.UUID, teams domain.Teams) error {
	for _, team := range teams {
		if err := s.team.Update(ctx, team_repo.UpdateIn{
			PlayOffID: playOffID,
			ID:        team.ID,
		}); err != nil {
			return fmt.Errorf("update team: %w", err)
		}
	}

	return nil
}
