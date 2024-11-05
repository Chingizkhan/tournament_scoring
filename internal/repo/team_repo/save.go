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
	SaveParams struct {
		Name       string
		DivisionID uuid.UUID
	}

	SaveIn struct {
		Teams []SaveParams
	}
)

func (r *Repository) Save(ctx context.Context, in SaveIn) (out []domain.Team, err error) {
	tx := repo.GetTX(ctx)

	builder := sq.
		Insert(repo.TableTeam).
		Columns("name", "division_id")

	for _, team := range in.Teams {
		builder = builder.Values(team.Name, team.DivisionID)
	}

	sql, args, err := builder.
		PlaceholderFormat(sq.Dollar).
		Suffix("returning *").
		ToSql()
	if err != nil {
		return out, fmt.Errorf("to sql: %w", err)
	}

	if err = r.Query(ctx, tx, &out, sql, args...); err != nil {
		return out, fmt.Errorf("exec: %w", err)
	}

	return out, nil
}
