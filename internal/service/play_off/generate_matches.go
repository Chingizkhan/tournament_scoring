package play_off

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/match_repo"
)

func (s *Service) GenerateMatches(ctx context.Context, playOffID uuid.UUID, teams domain.Teams, iteration int) (matches domain.Matches, err error) {
	teams.SortDesc()

	matches = make(domain.Matches, 0, len(teams)/2)

	for i := 0; i < len(teams)/2; i++ {
		matches = append(matches, domain.Match{
			Team1: domain.TeamInMatch{
				Team: domain.Team{
					ID: teams[i].ID,
				},
			},
			Team2: domain.TeamInMatch{
				Team: domain.Team{
					ID: teams[len(teams)-(i+1)].ID,
				},
			},
		})
	}

	if matches, err = s.match.Create(ctx, match_repo.CreateIn{
		Matches:   matches,
		Iteration: iteration,
	}); err != nil {
		return nil, fmt.Errorf("create matches: %w", err)
	}

	if err = s.match.Bind(ctx, match_repo.BindIn{
		Matches:   matches,
		PlayOffID: playOffID,
	}); err != nil {
		return nil, fmt.Errorf("bind matches: %w", err)
	}

	if matches, err = s.match.Find(ctx, match_repo.FindIn{
		PlayOffID: playOffID,
	}); err != nil {
		return nil, fmt.Errorf("find matches: %w", err)
	}

	return matches, nil
}
