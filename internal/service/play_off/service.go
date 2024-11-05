package play_off

import (
	"tournament_scoring/internal/service"
	"tournament_scoring/internal/usecase"
)

type Service struct {
	team              service.TeamRepo
	match             service.MatchRepo
	playOff           service.PlayOffRepo
	matchService      usecase.MatchService
	tournamentService usecase.TournamentService
}

func New(
	team service.TeamRepo,
	match service.MatchRepo,
	playOff service.PlayOffRepo,
	matchService usecase.MatchService,
	tournamentService usecase.TournamentService,
) *Service {
	return &Service{
		team:              team,
		match:             match,
		playOff:           playOff,
		matchService:      matchService,
		tournamentService: tournamentService,
	}
}
