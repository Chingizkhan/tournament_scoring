package v1

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
	"tournament_scoring/config"
	custom_middleware "tournament_scoring/internal/controller/http/middleware"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/dto"
	"tournament_scoring/pkg/logger"
)

type (
	Handler struct {
		l          logger.ILogger
		cfg        *config.Config
		tournament TournamentUseCase
		division   DivisionUseCase
		playOff    PlayOffUseCase
	}

	TournamentUseCase interface {
		Create(ctx context.Context, in dto.CreateTournamentIn) (dto.CreateTournamentOut, error)
		Delete(ctx context.Context) error
	}

	DivisionUseCase interface {
		Result(ctx context.Context, in dto.DivisionResultIn) (dto.DivisionResultOut, error)
	}

	PlayOffUseCase interface {
		GenerateSchedule(ctx context.Context) (domain.Team, error)
	}
)

func NewHandler(
	l logger.ILogger,
	cfg *config.Config,
	tournament TournamentUseCase,
	division DivisionUseCase,
	playOff PlayOffUseCase,
) *Handler {
	return &Handler{
		l:          l,
		cfg:        cfg,
		tournament: tournament,
		division:   division,
		playOff:    playOff,
	}
}

func (h *Handler) Register(r *chi.Mux, timeout time.Duration) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(timeout))
	r.Use(custom_middleware.Cors)
	r.Use(custom_middleware.Logging(h.l))

	h.routes(r)
}
