package team

import "context"

func (s *Service) Delete(ctx context.Context) error {
	return s.team.Delete(ctx)
}
