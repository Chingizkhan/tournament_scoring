package tournament_usecase

import (
	"tournament_scoring/internal/usecase"
)

type UseCase struct {
	division   usecase.DivisionsService
	tournament usecase.TournamentService
	team       usecase.TeamService
	match      usecase.MatchService
	playOff    usecase.PlayOffService
	tx         usecase.Transactional
}

func New(
	tournament usecase.TournamentService,
	division usecase.DivisionsService,
	team usecase.TeamService,
	match usecase.MatchService,
	playOff usecase.PlayOffService,
	tx usecase.Transactional,
) *UseCase {
	return &UseCase{
		division:   division,
		tournament: tournament,
		team:       team,
		match:      match,
		playOff:    playOff,
		tx:         tx,
	}
}
