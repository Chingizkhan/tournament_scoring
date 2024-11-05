package play_off

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/match_repo"
	"tournament_scoring/internal/repo/team_repo"
)

func (s *Service) Get(ctx context.Context) (out domain.PlayOff, err error) {
	out, err = s.playOff.Get(ctx)
	if err != nil {
		return out, fmt.Errorf("get play off: %w", err)
	}
	out.Matches, err = s.match.Find(ctx, match_repo.FindIn{
		PlayOffID: out.ID,
	})
	out.Teams, err = s.team.Find(ctx, team_repo.FindIn{
		PlayOffID: out.ID,
	})
	return out, nil
}
