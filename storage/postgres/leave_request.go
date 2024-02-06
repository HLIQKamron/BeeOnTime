package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
func (s *postgresRepo) GetStaffLeaves(ctx context.Context, req models.GetStaffLeavesRequest) (models.StaffLeaveList, error) {

	fmt.Println("req", req)
	var res models.StaffLeaveList
	where := `(staff_id = $1 or $1 = '')
		and (id = $2 or $2 = '')`

	query := s.Db.Builder.Select(`id,staff_id,reason,read,read_time,created_at,updated_at,approved,approved_time`).From("leave_request").Where(where, req.StaffID, req.Id)

	query = query.Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(s.Db.Db).Query()
	if err != nil {
		log.Println("err", err)
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		leave := models.LeaveRequest{}
		var readTime,approvedTime sql.NullString
		err := rows.Scan(
			&leave.Id,
			&leave.StaffId,
			&leave.Reason,
			&leave.Read,
			&readTime,
			&leave.CreatedAt,
			&leave.UpdatedAt,
			&leave.Approved,
			&approvedTime,
		)
		if err != nil {
			return models.StaffLeaveList{}, err
		}
		leave.ReadTime = readTime.String
		leave.ApprovedTime = approvedTime.String
		res.Leave = append(res.Leave, leave)
	}

	return res, nil
}
