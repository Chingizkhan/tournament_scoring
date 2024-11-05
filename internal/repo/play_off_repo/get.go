package play_off_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

func (r *Repository) Get(ctx context.Context) (out domain.PlayOff, err error) {
	tx := repo.GetTX(ctx)

	sql, args, err := sq.
		Select("*").
		PlaceholderFormat(sq.Question).
		From(repo.TablePlayOff).
		ToSql()
	if err != nil {
		return out, fmt.Errorf("to_sql: %w", err)
	}

	if err = r.QueryRow(ctx, tx, &out, sql, args...); err != nil {
		return out, fmt.Errorf("query: %w", err)
	}

	return out, nil
}
