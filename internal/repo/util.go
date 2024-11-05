package repo

import (
	"context"
	"github.com/jackc/pgx/v5"
	"tournament_scoring/internal/service/transactional"
)

const (
	Desc = "desc"
	Asc  = "asc"
)

func GetTX(ctx context.Context) pgx.Tx {
	txVal := ctx.Value(transactional.TxKey)
	tx, ok := txVal.(pgx.Tx)
	if ok {
		return tx
	}
	return nil
}
