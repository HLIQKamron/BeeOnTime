package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/BeeOntime/models"
	"github.com/Masterminds/squirrel"

	"github.com/google/uuid"
)

func (s *postgresRepo) CreateLeaveRequest(ctx context.Context, req models.LeaveRequest) (models.LeaveRequest, error) {
	var (
		res          models.LeaveRequest
		readTime     sql.NullString
		approvedTime sql.NullString
	)

	if req.Reason == "" {
		return res, fmt.Errorf("reason is required")
	}

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

	var (
		res   models.StaffLeaveList
		count int
	)
	where := `(staff_id = $1 or $1 = '')
		and (id = $2 or $2 = '')
		and (case 
			when $3 = '' then true
			else created_at >= $3::timestamp
		end)
	and (
			case
				when $4 = '' then true
				else created_at <= $4::timestamp
			end
		)`

	query := s.Db.Builder.Select(`id,staff_id,reason,read,read_time,created_at,updated_at,approved,approved_time`).From("leave_request").Where(where, req.StaffID, req.Id, req.From, req.To)

	query = query.Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(s.Db.Db).Query()
	if err != nil {
		log.Println("err", err)
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		leave := models.LeaveRequest{}
		var readTime, approvedTime sql.NullString
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

	query = s.Db.Builder.Select("count(*) ").From("leave_request").Where(where, req.StaffID, req.Id, req.From, req.To)
	err = query.RunWith(s.Db.Db).Scan(&count)
	if err != nil {
		log.Println("err", err)
		return res, err
	}
	res.Count = count
	return res, nil
}
func (s *postgresRepo) UpdateLeaveRequest(ctx context.Context, req models.LeaveRequest) (models.LeaveRequest, error) {

	var (
		updateData     = make(map[string]interface{})
		whereCondition = squirrel.And{squirrel.Eq{"id": req.Id}}
		res            = models.LeaveRequest{}
	)

	if req.Id == "" {
		return res, fmt.Errorf("id is required")
	}

	if req.Read {
		updateData["read"] = req.Read
		updateData["read_time"] = squirrel.Expr("now()")
		updateData["updated_at"] = squirrel.Expr("now()")
	}
	if req.Approved {
		updateData["approved"] = req.Approved
		updateData["approved_time"] = squirrel.Expr("now()")
		updateData["updated_at"] = squirrel.Expr("now()")
	}

	query := s.Db.Builder.Update("leave_request").SetMap(updateData).
		Where(whereCondition).
		Suffix("RETURNING id,read,read_time,updated_at,approved,approved_time")

	var (
		readTime,updatedAt,approvedAt sql.NullString
	)

	err := query.RunWith(s.Db.Db).QueryRow().Scan(
		&res.Id,
		&res.Read,
		&readTime,
		&updatedAt,
		&res.Approved,
		&approvedAt,
	)
	res.ReadTime = readTime.String
	res.UpdatedAt = updatedAt.String
	res.ApprovedTime = approvedAt.String

	if err != nil {
		return res, fmt.Errorf("staff - Update - QueryRow: %w", err)
	}

	return res, nil
}
