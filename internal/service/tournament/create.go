package tournament

import (
	"context"
	"tournament_scoring/internal/domain"
)

func (s *Service) Create(ctx context.Context) (out domain.Tournament, err error) {
	return s.tournament.Create(ctx)
}
