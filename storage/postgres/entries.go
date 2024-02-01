package postgres

import (
	"context"

	"github.com/BeeOntime/models"
)

func (s *postgresRepo) CreateStaffEntry(ctx context.Context, req models.Entry) (models.Entry, error) {
	return models.Entry{}, nil
}
