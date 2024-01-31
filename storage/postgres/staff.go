package postgres

import (
	"context"
	"fmt"

	"github.com/BeeOntime/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type staffRepo struct {
	db *pgxpool.Pool
}

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
