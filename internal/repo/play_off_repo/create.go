package play_off_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

func (r *Repository) Create(ctx context.Context, tournamentID uuid.UUID) (out domain.PlayOff, err error) {
	tx := repo.GetTX(ctx)

	sql, args, err := sq.Insert(repo.TablePlayOff).
		PlaceholderFormat(sq.Dollar).
		Columns("tournament_id").
		Values(tournamentID).
		Suffix("returning *").
		ToSql()
	if err != nil {
		return out, fmt.Errorf("to_sql: %w", err)
	}

	if err = r.QueryRow(ctx, tx, &out, sql, args...); err != nil {
		return out, fmt.Errorf("query_row: %w", err)
	}

	return out, nil
}
