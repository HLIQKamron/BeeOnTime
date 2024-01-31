package postgres

import (
	"time"

	"github.com/BeeOntime/config"
	"github.com/BeeOntime/pkg/db"
)

var (
	CreatedAt time.Time
	UpdatedAt time.Time
)

type postgresRepo struct {
	Db  *db.Postgres
	Cfg config.Config
}

func New(db *db.Postgres, cfg config.Config) PostgresI {
	return &postgresRepo{
		Db:  db,
		Cfg: cfg,
	}
}
