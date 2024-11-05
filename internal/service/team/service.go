package team

import (
	"tournament_scoring/internal/service"
)

type Service struct {
	team service.TeamRepo
}

func New(team service.TeamRepo) *Service {
	return &Service{
		team: team,
	}
}
