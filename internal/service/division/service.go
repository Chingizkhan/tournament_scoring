package division

import (
	"tournament_scoring/internal/service"
)

type Service struct {
	division service.DivisionRepo
	team     service.TeamRepo
	match    service.MatchRepo
}

func New(division service.DivisionRepo, team service.TeamRepo, match service.MatchRepo) *Service {
	return &Service{
		division: division,
		team:     team,
		match:    match,
	}
}
