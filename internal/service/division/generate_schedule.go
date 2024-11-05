package division

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/match_repo"
)

func (s *Service) GenerateSchedule(ctx context.Context, divisions []domain.Division) (err error) {
	for i, division := range divisions {
		pairs := s.getPairs(division.Teams)

		matches, err := s.match.Create(ctx, match_repo.CreateIn{
			Matches: pairs,
		})
		if err != nil {
			return fmt.Errorf("create matches via repo: %w", err)
		}

		if err = s.match.Bind(ctx, match_repo.BindIn{
			Matches:    matches,
			DivisionID: division.ID,
		}); err != nil {
			return fmt.Errorf("bind matches via repo: %w", err)
		}

		if divisions[i].Matches, err = s.match.Find(ctx, match_repo.FindIn{
			DivisionID: division.ID,
		}); err != nil {
			return fmt.Errorf("find matches: %w", err)
		}
	}

	return nil
}

func (s *Service) getPairs(teams []domain.Team) []domain.Match {
	pairs := make([]domain.Match, 0, len(teams)/2)

	for i := 0; i <= len(teams)-1; i = i + 2 {
		pairs = append(pairs, domain.Match{
			Team1: domain.TeamInMatch{
				Team: teams[i],
			},
			Team2: domain.TeamInMatch{
				Team: teams[i+1],
			},
		})
	}

	return pairs
}
