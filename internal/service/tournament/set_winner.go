package tournament

import (
	"context"
	"github.com/google/uuid"
)

func (s *Service) SetWinner(ctx context.Context, winnerID uuid.UUID) error {
	return s.tournament.Update(ctx, winnerID)
}
