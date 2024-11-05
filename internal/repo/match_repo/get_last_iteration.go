package match_repo

import (
	"context"
	sql2 "database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/repo"
)

type GetLastIterationIn struct {
	DivisionID uuid.UUID
	PlayOffID  uuid.UUID
}

func (r *Repository) GetLastIteration(ctx context.Context, in GetLastIterationIn) (int, error) {
	tx := repo.GetTX(ctx)
	var (
		lastIteration sql2.NullInt32
		condition     sq.Eq
		join          string
	)

	switch {
	case in.DivisionID != uuid.Nil:
		condition = sq.Eq{"division_id": in.DivisionID}
		join = "match_division md on m.id = md.match_id"
	case in.PlayOffID != uuid.Nil:
		condition = sq.Eq{"play_off_id": in.PlayOffID}
		join = "match_play_off mp on m.id = mp.match_id"
	}

	sql, args, err := sq.
		Select("max(iteration)").
		From("match m").
		Join(join).
		Where(condition).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("to_sql: %w", err)
	}

	if err = r.QueryRow(ctx, tx, &lastIteration, sql, args...); err != nil {
		return 0, fmt.Errorf("query_row: %w", err)
	}

	if lastIteration.Valid {
		return int(lastIteration.Int32), nil
	}

	return 0, nil
}
