package tournament_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/repo"
)

func (r *Repository) Update(ctx context.Context, winnerID uuid.UUID) error {
	tx := repo.GetTX(ctx)

	sql, args, err := sq.
		Update(repo.TableTournament).
		PlaceholderFormat(sq.Dollar).
		Set("winner", winnerID).
		ToSql()
	if err != nil {
		return fmt.Errorf("to_sql: %w", err)
	}

	if err = r.Exec(ctx, tx, sql, args...); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
