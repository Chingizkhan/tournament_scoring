package match_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/errs"
	"tournament_scoring/internal/repo"
)

type BindIn struct {
	Matches    []domain.Match
	DivisionID uuid.UUID
	PlayOffID  uuid.UUID
}

func (r *Repository) Bind(ctx context.Context, in BindIn) error {
	if len(in.Matches) == 0 {
		return nil
	}
	
	tx := repo.GetTX(ctx)

	var (
		tableName  string
		columnName string
		value      uuid.UUID
	)
	switch {
	case in.DivisionID != uuid.Nil:
		tableName = repo.TableMatchDivision
		columnName = "division_id"
		value = in.DivisionID
	case in.PlayOffID != uuid.Nil:
		tableName = repo.TableMatchPlayOff
		columnName = "play_off_id"
		value = in.PlayOffID
	default:
		return errs.UnexpectedArgument
	}

	builder := sq.
		Insert(tableName).
		Columns("match_id", columnName).
		PlaceholderFormat(sq.Dollar)
	for _, match := range in.Matches {
		builder = builder.Values(match.ID, value)
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("to_sql: %w", err)
	}

	if err = r.Exec(ctx, tx, sql, args...); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
