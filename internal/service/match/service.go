package match

import (
	"tournament_scoring/internal/service"
)

type Service struct {
	match service.MatchRepo
	team  service.TeamRepo
}

func New(match service.MatchRepo, team service.TeamRepo) *Service {
	return &Service{
		match: match,
		team:  team,
	}
}
