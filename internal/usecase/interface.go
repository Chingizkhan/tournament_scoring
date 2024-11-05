package usecase

import (
	"context"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
)

type (
	DivisionsService interface {
		List(ctx context.Context) ([]domain.Division, error)
		SeparateAndSaveTeams(ctx context.Context, divisions []domain.Division, teams []domain.Team) error
		GenerateSchedule(ctx context.Context, divisions []domain.Division) error
		CreateMatches(ctx context.Context, division domain.Division) (err error)
		GetByName(ctx context.Context, name domain.DivisionName) (domain.Division, error)
		Create(ctx context.Context, tournamentID uuid.UUID) ([]domain.Division, error)
		Delete(ctx context.Context) error
		GetBestTeams(ctx context.Context, divisionID uuid.UUID, n int) ([]domain.Team, error)
	}

	TournamentService interface {
		Create(ctx context.Context) (domain.Tournament, error)
		Exists(ctx context.Context) (bool, error)
		Delete(ctx context.Context) error
		SetWinner(ctx context.Context, winnerID uuid.UUID) error
	}

	TeamService interface {
		Delete(ctx context.Context) error
	}

	MatchService interface {
		Get(ctx context.Context, divisionID uuid.UUID) ([]domain.Match, error)
		Delete(ctx context.Context) error
		Update(ctx context.Context, matches []domain.Match) error
	}

	PlayOffService interface {
		Create(ctx context.Context, tournamentID uuid.UUID) (domain.PlayOff, error)
		GenerateMatches(ctx context.Context, playOffID uuid.UUID, teams domain.Teams, iteration int) (domain.Matches, error)
		BindTeams(ctx context.Context, playOffID uuid.UUID, teams domain.Teams) error
		PlayGames(ctx context.Context) (out domain.Team, err error)
		Delete(ctx context.Context) error
	}

	Transactional interface {
		Exec(ctx context.Context, fn func(txCtx context.Context) error) error
	}
)
