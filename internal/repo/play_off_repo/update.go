package play_off_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/repo"
)

type UpdateIn struct {
	WinnerID uuid.UUID
}

func (r *Repository) Update(ctx context.Context, in UpdateIn) error {
	tx := repo.GetTX(ctx)

	sql, args, err := sq.
		Update(repo.TablePlayOff).
		PlaceholderFormat(sq.Dollar).
		SetMap(map[string]interface{}{
			"winner": in.WinnerID,
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("to_sql: %w", err)
	}

	if err = r.Exec(ctx, tx, sql, args...); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
