package match_repo

import (
	"tournament_scoring/internal/repo"
	"tournament_scoring/pkg/postgres"
)

type Repository struct {
	*repo.DefaultRepo
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{repo.NewDefaultRepo(pg.Pool)}
}
