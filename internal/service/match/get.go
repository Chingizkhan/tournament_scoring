package match

import (
	"context"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/match_repo"
)

func (s *Service) Get(ctx context.Context, divisionID uuid.UUID) ([]domain.Match, error) {
	return s.match.Find(ctx, match_repo.FindIn{
		DivisionID: divisionID,
	})
}
