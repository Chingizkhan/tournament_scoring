package division

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/match_repo"
)

func (s *Service) CreateMatches(ctx context.Context, division domain.Division) (err error) {
	var (
		pairs         []domain.Match
		lastIteration int
	)

	if lastIteration, err = s.match.GetLastIteration(ctx, match_repo.GetLastIterationIn{
		DivisionID: division.ID,
	}); err != nil {
		return fmt.Errorf("get last iteration: %w", err)
	}

	pairs = s.getWinnersPairs(division, lastIteration)
	winnerMatches, err := s.match.Create(ctx, match_repo.CreateIn{
		Matches:   pairs,
		Iteration: lastIteration + 1,
	})
	if err != nil {
		return fmt.Errorf("create winners matches via repo: %w", err)
	}

	if err = s.match.Bind(ctx, match_repo.BindIn{
		Matches:    winnerMatches,
		DivisionID: division.ID,
	}); err != nil {
		return fmt.Errorf("bind winner matches via repo: %w", err)
	}

	pairs = s.getLooserPairs(division, lastIteration)
	looserMatches, err := s.match.Create(ctx, match_repo.CreateIn{
		Matches:   pairs,
		Iteration: lastIteration + 1,
	})
	if err != nil {
		return fmt.Errorf("create loosers matches via repo: %w", err)
	}

	if err = s.match.Bind(ctx, match_repo.BindIn{
		Matches:    looserMatches,
		DivisionID: division.ID,
	}); err != nil {
		return fmt.Errorf("bind looser matches via repo: %w", err)
	}

	return nil
}

func (s *Service) getWinnersPairs(division domain.Division, lastIteration int) []domain.Match {
	winnerTeams := division.Matches.GetWinners(lastIteration)
	// filter for team status
	res := make([]domain.Team, 0, len(winnerTeams)/2)
	for _, t := range winnerTeams {
		if t.TeamStatus == domain.TeamStatusWinning {
			res = append(res, t.Team)
		}
	}
	if len(res) == 1 {
		return []domain.Match{}
	}

	return s.getPairs(res)
}

func (s *Service) getLooserPairs(division domain.Division, lastIteration int) []domain.Match {
	looserTeams := division.Matches.GetLoosers(lastIteration)

	res := make([]domain.TeamInMatch, 0, len(looserTeams)/2)
	for _, t := range looserTeams {
		if t.TeamStatus == domain.TeamStatusLoosing {
			res = append(res, t)
		}
	}

	if len(res) == 1 {
		return []domain.Match{}
	}

	teams := make([]domain.Team, 0, len(res))

	for _, looserTeam := range res {
		teams = append(teams, looserTeam.Team)
	}

	return s.getPairs(teams)
}
