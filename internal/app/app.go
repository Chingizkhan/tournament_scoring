package app

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"tournament_scoring/config"
	v1 "tournament_scoring/internal/controller/http/v1"
	"tournament_scoring/internal/repo/division_repo"
	"tournament_scoring/internal/repo/match_repo"
	"tournament_scoring/internal/repo/play_off_repo"
	"tournament_scoring/internal/repo/team_repo"
	"tournament_scoring/internal/repo/tournament_repo"
	"tournament_scoring/internal/service/division"
	"tournament_scoring/internal/service/match"
	"tournament_scoring/internal/service/play_off"
	"tournament_scoring/internal/service/team"
	"tournament_scoring/internal/service/tournament"
	"tournament_scoring/internal/service/transactional"
	"tournament_scoring/internal/usecase/division_usecase"
	"tournament_scoring/internal/usecase/play_off_usecase"
	"tournament_scoring/internal/usecase/tournament_usecase"
	"tournament_scoring/pkg/httpserver"
	"tournament_scoring/pkg/logger"
	"tournament_scoring/pkg/postgres"
)

func Run(cfg *config.Config, l *logger.Logger) {
	pg, err := postgres.New(
		cfg.PG.DSN(),
		postgres.MaxPoolSize(cfg.PG.PoolMax),
		postgres.MaxConnLifetime(cfg.PG.MaxConnLifetime),
		postgres.MaxConnIdleTime(cfg.PG.MaxConnIdleTime),
	)
	if err != nil {
		l.Error("app - Run - postgres.New:", logger.Err(err))
		os.Exit(1)
	}
	defer pg.Close()

	tournamentRepo := tournament_repo.New(pg)
	divisionRepo := division_repo.New(pg)
	teamRepo := team_repo.New(pg)
	matchRepo := match_repo.New(pg)
	playOffRepo := play_off_repo.New(pg)

	txService := transactional.New(pg)
	teamService := team.New(teamRepo)
	divisionService := division.New(divisionRepo, teamRepo, matchRepo)
	matchService := match.New(matchRepo, teamRepo)
	tournamentService := tournament.New(tournamentRepo)
	playOffService := play_off.New(teamRepo, matchRepo, playOffRepo, matchService, tournamentService)

	// get usecases
	tournamentUseCase := tournament_usecase.New(tournamentService, divisionService, teamService, matchService, playOffService, txService)
	divisionUseCase := division_usecase.New(tournamentService, divisionService, matchService, txService)
	playOffUseCase := play_off_usecase.New(tournamentService, playOffService, divisionService, matchService, txService)

	// start http server
	router := chi.NewRouter()
	handler := v1.NewHandler(l, cfg, tournamentUseCase, divisionUseCase, playOffUseCase)
	handler.Register(router, cfg.HTTP.Timeout)
	httpServer := httpserver.New(
		router,
		httpserver.Port(cfg.HTTP.Port),
	)
	l.Info("http server started", slog.String("env", cfg.Log.Level), slog.String("port", cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal:", slog.String("signal", s.String()))
	case err = <-httpServer.Notify():
		l.Error("app - Run - http_server.Notify:", logger.Err(err))
	}

	// shutdown
	if err = httpServer.Shutdown(); err != nil {
		l.Error("app - Run - httpServer.Shutdown:", logger.Err(err))
		return
	}
}
