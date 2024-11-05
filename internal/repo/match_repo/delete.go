package match_repo

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"tournament_scoring/internal/repo"
)

func (r *Repository) Delete(ctx context.Context) error {
	tx := repo.GetTX(ctx)

	sql, args, err := sq.
		Delete(repo.TableMatchDivision).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("to_sql match_division: %w", err)
	}

	if err = r.Exec(ctx, tx, sql, args...); err != nil {
		return fmt.Errorf("exec delete match_division: %w", err)
	}

	sql, args, err = sq.
		Delete(repo.TableMatchPlayOff).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("to_sql match_play_off: %w", err)
	}

	if err = r.Exec(ctx, tx, sql, args...); err != nil {
		return fmt.Errorf("exec delete match_play_off: %w", err)
	}

	sql, args, err = sq.
		Delete(repo.TableMatch).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("to_sql match: %w", err)
	}

	if err = r.Exec(ctx, tx, sql, args...); err != nil {
		return fmt.Errorf("exec delete matches: %w", err)
	}

	return nil
}
