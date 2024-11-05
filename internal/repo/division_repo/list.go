package division_repo

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

func (r *Repository) List(ctx context.Context) (out []domain.Division, err error) {
	tx := repo.GetTX(ctx)

	sql := "select * from division;"
	if err = r.Query(ctx, tx, &out, sql); err != nil {
		return out, fmt.Errorf("query: %w", err)
	}
	return out, nil
}
