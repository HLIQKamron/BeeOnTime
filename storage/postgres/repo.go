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
	GetStaffEntries(ctx context.Context, req models.GetStaffEntries) (models.GetEntryResponse, error)
	DeleteStaffEntry(ctx context.Context, id string) error
	UpdateStaffEntry(ctx context.Context, req models.Entry) error

	CreateLeaveRequest(ctx context.Context, req models.LeaveRequest) (models.LeaveRequest, error)
	GetStaffLeaves(ctx context.Context, req models.GetStaffLeavesRequest) (models.StaffLeaveList, error)
	UpdateLeaveRequest(ctx context.Context, req models.LeaveRequest) (models.LeaveRequest, error)

	//hr
	CreateHr(ctx context.Context, req models.Hr) (models.Hr, error)
	GetHrs(ctx context.Context, req models.GetHrs) ([]models.Hr, error)
	DeleteHr(ctx context.Context, id string) error

	//Auth
	LoginCheck(ctx context.Context, username string, password string) (string, error)

	// Don't delete this line, it is used to modify the file automatically
}
