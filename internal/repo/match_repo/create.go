package match_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

type CreateIn struct {
	Matches   []domain.Match
	Iteration int
}

func (r *Repository) Create(ctx context.Context, in CreateIn) (matches domain.Matches, err error) {
	if len(in.Matches) == 0 {
		return matches, nil
	}

	tx := repo.GetTX(ctx)

	builder := sq.Insert(repo.TableMatch).
		Columns("first_team_id", "second_team_id", "goals_first_team", "goals_second_team", "iteration").
		PlaceholderFormat(sq.Dollar)

	for _, match := range in.Matches {
		builder = builder.Values(match.Team1.ID, match.Team2.ID, match.Team1.Goals, match.Team2.Goals, in.Iteration)
	}

	sql, args, err := builder.
		Suffix("returning *").
		ToSql()
	if err != nil {
		return matches, fmt.Errorf("to sql: %w", err)
	}

	var res Matches

	if err = r.Query(ctx, tx, &res, sql, args...); err != nil {
		return matches, fmt.Errorf("query: %w", err)
	}

	return res.convToDomain(), nil
}
