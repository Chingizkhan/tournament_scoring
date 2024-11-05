package match_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/repo"
)

type UpdateIn struct {
	ID              uuid.UUID
	Passed          bool
	GoalsFirstTeam  int
	GoalsSecondTeam int
}

func (r *Repository) Update(ctx context.Context, in UpdateIn) error {
	tx := repo.GetTX(ctx)

	sql, args, err := sq.
		Update(repo.TableMatch).
		SetMap(map[string]interface{}{
			"passed":            in.Passed,
			"goals_first_team":  in.GoalsFirstTeam,
			"goals_second_team": in.GoalsSecondTeam,
		}).
		Where(sq.Eq{"id": in.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("to_sql: %w", err)
	}

	if err = r.Exec(ctx, tx, sql, args...); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
