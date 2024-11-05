package match

import "context"

func (s *Service) Delete(ctx context.Context) error {
	return s.match.Delete(ctx)
}
