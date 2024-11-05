package tournament

import "tournament_scoring/internal/service"

type (
	Service struct {
		tournament service.TournamentRepo
	}
)

func New(tournament service.TournamentRepo) *Service {
	return &Service{
		tournament: tournament,
	}
}
