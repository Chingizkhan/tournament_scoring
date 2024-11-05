package tournament_repo

import (
	"context"
	"fmt"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

func (r *Repository) Create(ctx context.Context) (out domain.Tournament, err error) {
	tx := repo.GetTX(ctx)

	sql := "insert into tournament (winner) values (null) returning *;"

	if err = r.QueryRow(ctx, tx, &out, sql); err != nil {
		return out, fmt.Errorf("query_row: %w", err)
	}

	return out, nil
}
