package team_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

type (
	OrderBy struct {
		Column string
		Desc   bool
	}

	FindIn struct {
		PlayOffID  uuid.UUID
		DivisionID uuid.UUID
		Limit      int
		OrderBy    OrderBy
	}
)

func (r *Repository) Find(ctx context.Context, in FindIn) (out []domain.Team, err error) {
	tx := repo.GetTX(ctx)

	var condition sq.Eq
	switch {
	case in.PlayOffID != uuid.Nil:
		condition = sq.Eq{"play_off_id": in.PlayOffID}
	case in.DivisionID != uuid.Nil:
		condition = sq.Eq{"division_id": in.DivisionID}
	}

	builder := sq.
		Select("*").
		From(repo.TableTeam).
		Where(condition)

	if in.OrderBy.Column != "" {
		by := repo.Desc
		if !in.OrderBy.Desc {
			by = repo.Asc
		}

		builder = builder.OrderBy(
			fmt.Sprintf("%s %s", in.OrderBy.Column, by),
		)
	}

	if in.Limit > 0 {
		builder = builder.
			Limit(uint64(in.Limit))
	}

	sql, args, err := builder.
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("to_sql: %w", err)
	}

	if err = r.Query(ctx, tx, &out, sql, args...); err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return out, nil
}
