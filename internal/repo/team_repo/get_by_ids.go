package team_repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"tournament_scoring/internal/domain"
	"tournament_scoring/internal/repo"
)

func (r *Repository) GetByIDs(ctx context.Context, ids []uuid.UUID) (out domain.Teams, err error) {
	tx := repo.GetTX(ctx)

	sb := strings.Builder{}
	args := make([]interface{}, 0, len(ids))

	sb.WriteString("select * from team where id in (")

	for i, id := range ids {
		args = append(args, id)
		if i > 0 {
			sb.WriteString(fmt.Sprintf(", $%d", len(args)))
		} else {
			sb.WriteString(fmt.Sprintf("$%d", len(args)))
		}
	}

	sb.WriteString(");")

	query := sb.String()

	if err = r.Query(ctx, tx, &out, query, args...); err != nil {
		return nil, fmt.Errorf("query teams: %w", err)
	}

	return out, nil
}
