package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BeeOntime/models"

	"github.com/google/uuid"
)

func (s *postgresRepo) CreateLeaveRequest(ctx context.Context, req models.LeaveRequest) (models.LeaveRequest, error) {
	var (
		res          models.LeaveRequest
		readTime     sql.NullString
		approvedTime sql.NullString
	)

	uuid, err := uuid.NewUUID()
	if err != nil {
		return res, err
	}
	fmt.Println("reason :", req.Reason)
	query := s.Db.Builder.Insert("leave_request").Columns(
		`id,
		staff_id,
		reason
		`).Values(uuid.String(), req.StaffId, req.Reason).Suffix(`
		RETURNING id,staff_id,reason,read,created_at,updated_at,read_time,approved,approved_time`)
	err = query.RunWith(s.Db.Db).Scan(
		&res.Id, &res.StaffId, &res.Reason, &res.Read, &res.CreatedAt, &res.UpdatedAt, &readTime, &res.Approved, &approvedTime,
	)
	res.ReadTime = readTime.String
	res.ApprovedTime = approvedTime.String
	if err != nil {
		return res, err
	}
	return res, nil
}
