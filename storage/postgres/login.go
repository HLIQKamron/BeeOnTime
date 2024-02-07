package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/BeeOntime/models"
	"github.com/BeeOntime/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

func (s *postgresRepo) LoginCheck(ctx context.Context, username string, password string) (string, error) {
	res := models.Staff{}
	query := s.Db.Builder.Select("id,password").From("staff").Where("email = ?", username)
	err := query.RunWith(s.Db.Db).Scan(
		&res.Id,
		&res.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found")
		}
		return "", err
	}
	err = validator.VerifyPassword(password, res.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := validator.GenerateToken(res.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}
