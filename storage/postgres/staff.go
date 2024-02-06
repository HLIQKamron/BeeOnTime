package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/BeeOntime/models"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func (s *postgresRepo) CreateStaff(ctx context.Context, req models.Staff) (models.Staff, error) {
	res := models.Staff{}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return res, err
	}

	query := s.Db.Builder.Insert("staff").Columns(
		`id,
		name,
		email,
		phone,
		password,
		lastname,
		branch_name`,
	).Values(uuid.String(), req.Name, req.Email, req.Phone, req.Password, req.Lastname, req.BranchName).Suffix(
		"RETURNING id, name, email, phone, lastname, created_at")
	err = query.RunWith(s.Db.Db).Scan(
		&res.Id,
		&res.Name,
		&res.Email,
		&res.Phone,
		&res.Lastname,
		&res.CreatedAt,
	)
	if err != nil {
		fmt.Println("err", err)
		return res, err
	}
	return res, nil
}
func (s *postgresRepo) GetByLogin(ctx context.Context, login string) (models.Staff, error) {
	res := models.Staff{}
	query := s.Db.Builder.Select("id,name,email,phone").From("staff").Where("email = ?", login)
	err := query.RunWith(s.Db.Db).Scan(
		&res.Id,
		&res.Name,
		&res.Email,
		&res.Phone,
		// &res.Password,
		// &res.Lastname,
		// &res.BranchName,
		// &res.CreatedAt,
	)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *postgresRepo) GetStaffs(ctx context.Context, req models.GetStaffs) ([]models.Staff, error) {
	res := []models.Staff{}
	query := s.Db.Builder.Select(`id,name,email,phone,lastname,branch_name,blocked,blocked_date,created_at,updated_at`).From("staff")

	query = query.Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(s.Db.Db).Query()
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		staff := models.Staff{}
		var blockedData sql.NullString
		err = rows.Scan(
			&staff.Id,
			&staff.Name,
			&staff.Email,
			&staff.Phone,
			&staff.Lastname,
			&staff.BranchName,
			&staff.Blocked,
			&blockedData,
			&staff.CreatedAt,
			&staff.UpdatedAt,
		)
		if err != nil {
			return res, err
		}
		staff.BlockedAt = blockedData.String
		res = append(res, staff)
	}

	return res, nil
}
func (s *postgresRepo) DeleteStaff(ctx context.Context, id string) error {
	query := s.Db.Builder.Delete("staff").Where("id = ?", id)
	resp, err := query.RunWith(s.Db.Db).Exec()
	if err != nil {
		return err
	}
	ok, err := resp.RowsAffected()
	if err != nil {
		return err
	}
	if ok == 0 {
		return fmt.Errorf("staff not found")
	}
	return nil
}
func (s *postgresRepo) UpdateStaff(ctx context.Context, req models.Staff) (models.Staff, error) {
	var (
		mp             = make(map[string]interface{})
		whereCondition = squirrel.And{squirrel.Eq{"id": req.Id}}
	)

	mp["name"] = req.Name
	mp["email"] = req.Email
	mp["phone"] = req.Phone
	mp["lastname"] = req.Lastname
	mp["branch_name"] = req.BranchName
	mp["blocked"] = req.Blocked
	if req.Blocked {
		mp["blocked_date"] = time.Now()
	}
	mp["password"] = req.Password
	mp["updated_at"] = time.Now()

	if req.Id == "" {
		return models.Staff{}, fmt.Errorf("id is required")
	}
	whereCondition = append(whereCondition, squirrel.Eq{"id": req.Id})

	query := s.Db.Builder.Update("staff").SetMap(mp).
		Where(whereCondition).
		Suffix("RETURNING id, name,email,branch_name,blocked,blocked_date,created_at,updated_at,lastname,phone")

	res := models.Staff{}
	var blockedDate sql.NullString
	err := query.RunWith(s.Db.Db).QueryRow().Scan(
		&res.Id, &res.Name,
		&res.Email, &res.BranchName,
		&res.Blocked, &blockedDate,
		&CreatedAt, &UpdatedAt, &res.Lastname,
		&res.Phone,
	)
	res.BlockedAt = blockedDate.String

	if err != nil {
		return res, fmt.Errorf("staff - Update - QueryRow: %w", err)
	}

	res.CreatedAt = CreatedAt.Format(time.RFC3339)
	res.UpdatedAt = UpdatedAt.Format(time.RFC3339)

	return res, nil
}
