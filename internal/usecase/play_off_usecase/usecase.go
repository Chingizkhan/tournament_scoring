package play_off_usecase

import (
	"tournament_scoring/internal/usecase"
)

type UseCase struct {
	tournament usecase.TournamentService
	playOff    usecase.PlayOffService
	division   usecase.DivisionsService
	match      usecase.MatchService
	tx         usecase.Transactional
}

func New(
	tournament usecase.TournamentService,
	playOff usecase.PlayOffService,
	division usecase.DivisionsService,
	match usecase.MatchService,
	tx usecase.Transactional,
) *UseCase {
	return &UseCase{
		tournament: tournament,
		playOff:    playOff,
		division:   division,
		match:      match,
		tx:         tx,
	}
}
