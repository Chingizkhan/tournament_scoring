package tournament_usecase

import (
	"context"
	"fmt"
)

func (uc *UseCase) Delete(ctx context.Context) error {
	if err := uc.tx.Exec(ctx, func(txCtx context.Context) error {
		if err := uc.match.Delete(txCtx); err != nil {
			return fmt.Errorf("delete matches: %w", err)
		}

		if err := uc.team.Delete(txCtx); err != nil {
			return fmt.Errorf("delete team: %w", err)
		}

		if err := uc.division.Delete(txCtx); err != nil {
			return fmt.Errorf("delete division: %w", err)
		}

		if err := uc.playOff.Delete(txCtx); err != nil {
			return fmt.Errorf("delete playoff: %w", err)
		}

		if err := uc.tournament.Delete(txCtx); err != nil {
			return fmt.Errorf("delete tournament: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("exec tx: %w", err)
	}

	return nil
}
