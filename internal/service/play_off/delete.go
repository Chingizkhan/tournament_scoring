package play_off

import (
	"context"
)

func (s *Service) Delete(ctx context.Context) error {
	return s.playOff.Delete(ctx)
}
