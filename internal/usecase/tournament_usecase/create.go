package tournament_usecase

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/dto"
	"tournament_scoring/internal/errs"
)

func (uc *UseCase) Create(ctx context.Context, in dto.CreateTournamentIn) (out dto.CreateTournamentOut, err error) {
	// check if tournament already exists
	exists, err := uc.tournament.Exists(ctx)
	if err != nil {
		return out, fmt.Errorf("tournament exists: %w", err)
	}
	if exists {
		return out, errs.TournamentAlreadyExists
	}

	var (
		tournament domain.Tournament
	)

	if err = uc.tx.Exec(ctx, func(txCtx context.Context) error {
		// create tournament in tables
		tournament, err = uc.tournament.Create(txCtx)
		if err != nil {
			return fmt.Errorf("create tournament: %w", err)
		}

		tournament.Divisions, err = uc.division.Create(txCtx, tournament.ID)
		if err != nil {
			return fmt.Errorf("create divisions: %w", err)
		}

		// randomly separate teams for these divisions and saves them with bindings
		if err = uc.division.SeparateAndSaveTeams(txCtx, tournament.Divisions, in.Teams); err != nil {
			return fmt.Errorf("separate teams by divisions via division service: %w", err)
		}

		if err = uc.division.GenerateSchedule(txCtx, tournament.Divisions); err != nil {
			return fmt.Errorf("generate schedule: %w", err)
		}

		return nil
	}); err != nil {
		return out, fmt.Errorf("exec tx: %w", err)
	}

	// return info about schedule for A & B divisions
	return out.ConvertResponse(
		tournament.Divisions,
	), nil
}
