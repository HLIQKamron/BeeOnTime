package storage

import (
	"github.com/BeeOntime/config"
	"github.com/BeeOntime/pkg/db"
	"github.com/BeeOntime/storage/postgres"
)

type StorageI interface {
	Postgres() postgres.PostgresI
}

type StoragePg struct {
	postgres postgres.PostgresI
}

// NewStoragePg
func New(db *db.Postgres, cfg config.Config) StorageI {
	return &StoragePg{
		postgres: postgres.New(db, cfg),
	}
}

func (s *StoragePg) Postgres() postgres.PostgresI {
	return s.postgres
}
