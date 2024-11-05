package division_usecase

import (
	"tournament_scoring/internal/usecase"
)

type UseCase struct {
	tournament usecase.TournamentService
	division   usecase.DivisionsService
	match      usecase.MatchService
	tx         usecase.Transactional
}

func New(
	tournament usecase.TournamentService,
	division usecase.DivisionsService,
	match usecase.MatchService,
	tx usecase.Transactional,
) *UseCase {
	return &UseCase{
		tournament: tournament,
		division:   division,
		match:      match,
		tx:         tx,
	}
}
