package play_off_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"tournament_scoring/internal/repo"
)

func (r *Repository) Delete(ctx context.Context) error {
	tx := repo.GetTX(ctx)

	sql, args, err := sq.
		Delete(repo.TablePlayOff).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("to_sql: %w", err)
	}

	if err = r.Exec(ctx, tx, sql, args...); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
