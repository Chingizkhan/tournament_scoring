package play_off

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"tournament_scoring/internal/repo/play_off_repo"
)

func (s *Service) SetWinner(ctx context.Context, winnerID uuid.UUID) error {
	if err := s.playOff.Update(ctx, play_off_repo.UpdateIn{
		WinnerID: winnerID,
	}); err != nil {
		return fmt.Errorf("set play off winner: %w", err)
	}

	if err := s.tournamentService.SetWinner(ctx, winnerID); err != nil {
		return fmt.Errorf("set tournament winner: %w", err)
	}

	return nil
}
