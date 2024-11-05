package division_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

func (r *Repository) GetByName(ctx context.Context, name domain.DivisionName) (out domain.Division, err error) {
	tx := repo.GetTX(ctx)

	sql, args, err := sq.
		Select("*").
		From(repo.TableDivision).
		Where(sq.Eq{"name": name}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return out, fmt.Errorf("to_sql: %w", err)
	}

	if err = r.QueryRow(ctx, tx, &out, sql, args...); err != nil {
		return out, fmt.Errorf("query_row: %w", err)
	}

	return out, nil
}
