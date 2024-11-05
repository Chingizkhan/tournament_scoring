package tournament

import "context"

func (s *Service) Exists(ctx context.Context) (bool, error) {
	return s.tournament.Exists(ctx)
}
