package tournament

import "context"

func (s *Service) Delete(ctx context.Context) error {
	return s.tournament.Delete(ctx)
}
