package tournament_repo

import (
	"context"
	"fmt"
	"tournament_scoring/internal/repo"
)

func (r *Repository) Exists(ctx context.Context) (out bool, err error) {
	tx := repo.GetTX(ctx)

	var total int
	sql := "select count(*) from tournament;"

	if err = r.QueryRow(ctx, tx, &total, sql); err != nil {
		return false, fmt.Errorf("query row: %w", err)
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}
