package postgres

import (
	"context"

	"github.com/BeeOntime/models"
)

type PostgresI interface {
	// common
	CreateStaff(ctx context.Context, req models.Staff) (models.Staff, error)
	GetByLogin(ctx context.Context, login string) (models.Staff, error)
	GetStaffs(ctx context.Context, req models.GetStaffs) ([]models.Staff, error)
	DeleteStaff(ctx context.Context, id string) error
	UpdateStaff(ctx context.Context, req models.Staff) (models.Staff, error)

	CreateStaffEntry(ctx context.Context, req models.Entry) (models.Entry, error)
	// Don't delete this line, it is used to modify the file automatically
}
