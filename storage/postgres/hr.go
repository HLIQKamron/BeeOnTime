package postgres

import (
	"context"
	"fmt"

	"github.com/BeeOntime/models"
	"github.com/google/uuid"
)

func (s *postgresRepo) CreateHr(ctx context.Context, req models.Hr) (models.Hr, error) {
	var res models.Hr

	uuid, err := uuid.NewUUID()
	if err != nil {
		return res, err
	}

	if req.Login == "" || req.Password == "" {
		return res, fmt.Errorf("login and password are required")
	}
	var count int
	queryCount := s.Db.Builder.Select("COUNT(*)").From("hr").Where("login = ?", req.Login)

	err = queryCount.RunWith(s.Db.Db).QueryRow().Scan(&count)
	if err != nil {
		return res, err
	}

	if count > 0 {
		return res, fmt.Errorf("login already exists")
	}

	query := s.Db.Builder.Insert("hr").Columns(
		`id,
		login,
		password`,
	).Values(uuid.String(), req.Login, req.Password).Suffix(
		"RETURNING id, created_at,updated_at")
	err = query.RunWith(s.Db.Db).Scan(
		&res.Id,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		fmt.Println("err", err)
		return res, err
	}
	return res, nil

}
func (s *postgresRepo) GetHrs(ctx context.Context, req models.GetHrs) ([]models.Hr, error) {
	var res []models.Hr
	where := `(id = $1 or $1 = '')`
	query := s.Db.Builder.Select("id,login,created_at,updated_at").From("hr").Where(where, req.Id)
	rows, err := query.RunWith(s.Db.Db).Query()
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var r models.Hr
		err = rows.Scan(&r.Id, &r.Login, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return res, err
		}
		res = append(res, r)
	}
	return res, nil
}
func (s *postgresRepo) DeleteHr(ctx context.Context, id string) error {
	query := s.Db.Builder.Delete("hr").Where("id = ?", id)
	resp, err := query.RunWith(s.Db.Db).Exec()
	if err != nil {
		return err
	}
	ok, err := resp.RowsAffected()
	if err != nil {
		return err
	}
	if ok == 0 {
		return fmt.Errorf("hr not found")
	}
	return nil
}
