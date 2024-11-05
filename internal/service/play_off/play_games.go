package play_off

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/match_repo"
)

func (s *Service) PlayGames(ctx context.Context) (out domain.Team, err error) {
	playOff, err := s.Get(ctx)
	if err != nil {
		return out, fmt.Errorf("get play off: %w", err)
	}

	if playOff.IsFinished() {
		winners, err := s.team.GetByIDs(ctx, []uuid.UUID{playOff.Winner})
		if err != nil {
			return out, fmt.Errorf("get play off winner: %w", err)
		}
		return winners[0], nil
	}

	for i, _ := range playOff.Matches {
		playOff.Matches[i].Play()
	}

	if err = s.matchService.Update(ctx, playOff.Matches); err != nil {
		return out, fmt.Errorf("update matches and teams: %w", err)
	}

	iteration, err := s.match.GetLastIteration(ctx, match_repo.GetLastIterationIn{
		PlayOffID: playOff.ID,
	})
	if err != nil {
		return out, fmt.Errorf("get last iteration: %w", err)
	}

	winners := playOff.Matches.GetWinners(iteration)

	if len(winners) == 1 {
		if err = s.SetWinner(ctx, winners[0].ID); err != nil {
			return out, fmt.Errorf("set winner: %w", err)
		}
	}

	winnerTeams, err := s.team.GetByIDs(ctx, getWinnerIDs(winners))
	if err != nil {
		return out, fmt.Errorf("get teams by ids: %w", err)
	}

	_, err = s.GenerateMatches(ctx, playOff.ID, winnerTeams, iteration+1)
	if err != nil {
		return out, fmt.Errorf("generate matches: %w", err)
	}

	return s.PlayGames(ctx)
}

func getWinnerIDs(teams []domain.TeamInMatch) []uuid.UUID {
	res := make([]uuid.UUID, 0, len(teams))

	for _, t := range teams {
		res = append(res, t.ID)
	}

	return res
}
