package match_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

type FindIn struct {
	DivisionID uuid.UUID
	PlayOffID  uuid.UUID
	Passed     *bool
}

func (r *Repository) Find(ctx context.Context, in FindIn) ([]domain.Match, error) {
	tx := repo.GetTX(ctx)

	var (
		id   string
		join string
	)
	switch {
	case in.DivisionID != uuid.Nil:
		id = "md.division_id"
		join = "match_division md on m.id = md.match_id"
	case in.PlayOffID != uuid.Nil:
		id = "mp.play_off_id"
		join = "match_play_off mp on m.id = mp.match_id"
	}

	builder := sq.Select(
		"m.id",
		"passed",
		"first_team_id",
		"second_team_id",
		"goals_first_team",
		"goals_second_team",
		id,
		"m.iteration",
		"t1.name as first_team_name",
		"t1.division_id as first_team_division_id",
		"t1.rating as first_team_rating",
		"t1.team_status as first_team_status",
		"t2.name as second_team_name",
		"t2.division_id as second_team_division_id",
		"t2.rating as second_team_rating",
		"t2.team_status as second_team_status",
	).
		From("match as m").
		Join(join).
		Join("team t1 on m.first_team_id = t1.id").
		Join("team t2 on m.second_team_id = t2.id")

	if in.DivisionID != uuid.Nil {
		builder = builder.Where(sq.Eq{"md.division_id": in.DivisionID})
	}

	sql, args, err := builder.
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql: %w", err)
	}

	var matchesOut MatchesFullInfo

	if err = r.Query(ctx, tx, &matchesOut, sql, args...); err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return matchesOut.convToDomain(), nil
}
