package play_off_usecase

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/errs"
)

const (
	_defaultTeamQuantity = 4
)

func (uc *UseCase) GenerateSchedule(ctx context.Context) (out domain.Team, err error) {
	exists, err := uc.tournament.Exists(ctx)
	if err != nil {
		return out, fmt.Errorf("tournament exists: %w", err)
	}
	if !exists {
		return out, errs.TournamentNotExists
	}

	divisions, err := uc.division.List(ctx)
	if err != nil {
		return out, fmt.Errorf("list divisions: %w", err)
	}

	bestTeams := make([]domain.Team, 0, _defaultTeamQuantity*len(divisions))

	for _, division := range divisions {
		if !division.IsFinished() {
			return out, errs.DivisionsNotFinished
		}

		teams, err := uc.division.GetBestTeams(ctx, division.ID, _defaultTeamQuantity)
		if err != nil {
			return out, fmt.Errorf("get best teams in division: %w", err)
		}

		bestTeams = append(bestTeams, teams...)
	}

	var (
		playOff domain.PlayOff
	)

	if err = uc.tx.Exec(ctx, func(txCtx context.Context) error {
		// create play off
		playOff, err = uc.playOff.Create(txCtx, divisions[0].TournamentID)
		if err != nil {
			return fmt.Errorf("create play off: %w", err)
		}

		// bind teams for play_off
		if err = uc.playOff.BindTeams(txCtx, playOff.ID, bestTeams); err != nil {
			return fmt.Errorf("bind teams: %w", err)
		}

		// create matches for play off
		if playOff.Matches, err = uc.playOff.GenerateMatches(txCtx, playOff.ID, bestTeams, 0); err != nil {
			return fmt.Errorf("generate matches for play off: %w", err)
		}

		if out, err = uc.playOff.PlayGames(txCtx); err != nil {
			return fmt.Errorf("play off games: %w", err)
		}

		return nil
	}); err != nil {
		return out, fmt.Errorf("exec tx: %w", err)
	}

	return out, nil
}
