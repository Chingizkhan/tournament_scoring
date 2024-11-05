package division_usecase

import (
	"context"
	"fmt"
	"tournament_scoring/internal/dto"
	"tournament_scoring/internal/errs"
)

func (uc *UseCase) Result(ctx context.Context, in dto.DivisionResultIn) (out dto.DivisionResultOut, err error) {
	exists, err := uc.tournament.Exists(ctx)
	if err != nil {
		return out, fmt.Errorf("tournament exists: %w", err)
	}
	if !exists {
		return out, errs.TournamentNotExists
	}

	// Сервис обновляет результаты матчей и сохраняет победителей.
	division, err := uc.division.GetByName(ctx, in.DivisionName)
	if err != nil {
		return out, fmt.Errorf("get division by name: %w", err)
	}

	if division.IsFinished() {
		return out.ConvertResponse(division), nil
	}

	for i, _ := range division.Matches {
		division.Matches[i].Play()
	}

	if err = uc.tx.Exec(ctx, func(txCtx context.Context) error {
		if err = uc.match.Update(txCtx, division.Matches); err != nil {
			return fmt.Errorf("update matches and teams: %w", err)
		}

		if err = uc.division.CreateMatches(txCtx, division); err != nil {
			return fmt.Errorf("create new matches: %w", err)
		}

		return nil
	}); err != nil {
		return out, fmt.Errorf("exec tx: %w", err)
	}

	return uc.Result(ctx, in)
}
