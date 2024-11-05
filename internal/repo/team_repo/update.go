package team_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

type UpdateIn struct {
	PlayOffID  uuid.UUID
	ID         uuid.UUID
	TeamStatus domain.TeamStatus
	Rating     int
}

func (r *Repository) Update(ctx context.Context, in UpdateIn) error {
	tx := repo.GetTX(ctx)

	sql, args, err := sq.
		Update(repo.TableTeam).
		SetMap(r.prependUpdate(in)).
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

func (r *Repository) prependUpdate(in UpdateIn) map[string]interface{} {
	data := make(map[string]interface{}, 3)

	if in.Rating != 0 {
		data["rating"] = in.Rating
	}

	if in.TeamStatus != "" {
		data["team_status"] = in.TeamStatus
	}

	if in.PlayOffID != uuid.Nil {
		data["play_off_id"] = in.PlayOffID
	}

	return data
}
