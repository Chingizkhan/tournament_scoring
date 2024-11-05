package service

import (
	"context"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo/division_repo"
	"tournament_scoring/internal/repo/match_repo"
	"tournament_scoring/internal/repo/play_off_repo"
	"tournament_scoring/internal/repo/team_repo"
)

type (
	DivisionRepo interface {
		List(ctx context.Context) ([]domain.Division, error)
		GetByName(ctx context.Context, name domain.DivisionName) (out domain.Division, err error)
		Create(ctx context.Context, in division_repo.CreateIn) ([]domain.Division, error)
		Delete(ctx context.Context) error
	}

	TeamRepo interface {
		Save(ctx context.Context, in team_repo.SaveIn) (out []domain.Team, err error)
		Delete(ctx context.Context) error
		Find(ctx context.Context, in team_repo.FindIn) (out []domain.Team, err error)
		GetByIDs(ctx context.Context, ids []uuid.UUID) (domain.Teams, error)
		Update(ctx context.Context, in team_repo.UpdateIn) error
	}

	MatchRepo interface {
		GetLastIteration(ctx context.Context, in match_repo.GetLastIterationIn) (out int, err error)
		Create(ctx context.Context, in match_repo.CreateIn) (matches domain.Matches, err error)
		Bind(ctx context.Context, in match_repo.BindIn) error
		Find(ctx context.Context, in match_repo.FindIn) ([]domain.Match, error)
		Delete(ctx context.Context) error
		Update(ctx context.Context, in match_repo.UpdateIn) error
	}

	PlayOffRepo interface {
		Create(ctx context.Context, tournamentID uuid.UUID) (domain.PlayOff, error)
		Update(ctx context.Context, in play_off_repo.UpdateIn) error
		Get(ctx context.Context) (out domain.PlayOff, err error)
		Delete(ctx context.Context) error
	}

	TournamentRepo interface {
		Create(ctx context.Context) (out domain.Tournament, err error)
		Exists(ctx context.Context) (bool, error)
		Delete(ctx context.Context) error
		Update(ctx context.Context, winnerID uuid.UUID) error
	}
)
