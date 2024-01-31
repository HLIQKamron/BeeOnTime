package postgres

import (
	"context"

	"github.com/BeeOntime/models"
)

type PostgresI interface {
	// common
	CreateStaff(ctx context.Context, req models.Staff) (models.Staff, error)
	// Don't delete this line, it is used to modify the file automatically
}
