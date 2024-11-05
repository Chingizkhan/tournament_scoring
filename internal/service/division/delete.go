package division

import "context"

func (s *Service) Delete(ctx context.Context) error {
	return s.division.Delete(ctx)
}
