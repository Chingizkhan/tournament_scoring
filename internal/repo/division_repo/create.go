package division_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

type (
	CreateParams struct {
		Name         domain.DivisionName
		TournamentID uuid.UUID
	}

	CreateIn struct {
		Divisions []CreateParams
	}
)

func (r *Repository) Create(ctx context.Context, in CreateIn) (out []domain.Division, err error) {
	tx := repo.GetTX(ctx)

	builder := sq.
		Insert(repo.TableDivision).Columns("name", "tournament_id")

	for _, division := range in.Divisions {
		builder = builder.Values(division.Name, division.TournamentID)
	}

	sql, args, err := builder.Suffix("returning *").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return out, fmt.Errorf("to sql: %w", err)
	}

	if err = r.Query(ctx, tx, &out, sql, args...); err != nil {
		return out, fmt.Errorf("query: %w", err)
	}

	return out, nil
}
