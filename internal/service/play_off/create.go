package play_off

import (
	"context"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
)

func (s *Service) Create(ctx context.Context, tournamentID uuid.UUID) (domain.PlayOff, error) {
	return s.playOff.Create(ctx, tournamentID)
}
