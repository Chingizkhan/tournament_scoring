package division

import (
	"context"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/division_repo"
)

func (s *Service) Create(ctx context.Context, tournamentID uuid.UUID) ([]domain.Division, error) {
	var createIn = division_repo.CreateIn{
		Divisions: make([]division_repo.CreateParams, 0, len(domain.DivisionsList)),
	}

	for _, division := range domain.DivisionsList {
		createIn.Divisions = append(createIn.Divisions, division_repo.CreateParams{
			Name:         division.Name,
			TournamentID: tournamentID,
		})
	}

	return s.division.Create(ctx, createIn)
}
