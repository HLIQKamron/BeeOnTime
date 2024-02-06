package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/BeeOntime/models"
	"github.com/google/uuid"
)

func (s *postgresRepo) CreateStaffEntry(ctx context.Context, req models.Entry) (models.Entry, error) {
	var res models.Entry

	uuid, err := uuid.NewUUID()
	if err != nil {
		return res, err
	}

	query := s.Db.Builder.Insert("entries").Columns(
		`id,
		staff_id,
		activity_type,
		city`).Values(uuid.String(), req.StaffId, req.ActivityType, req.City).Suffix(`
		RETURNING staff_id, activity_type, city,created_at`)
	err = query.RunWith(s.Db.Db).Scan(
		&res.StaffId, &res.ActivityType, &res.City, &res.Date,
	)
	if err != nil {
		fmt.Println("err", err)
		return res, err
	}
	return res, nil
}
func (s *postgresRepo) GetStaffEntries(ctx context.Context, req models.GetStaffEntries) (models.GetEntryResponse, error) {
	var res []models.Entry
	query := s.Db.Builder.Select("id,staff_id,activity_type,city,(select count(*) from entries) as count,created_at").From("entries").Where(`(staff_id = $1 or $1 = '')
	and (id = $2 or $2 = '')`, req.StaffID, req.Id, )

	query = query.Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(s.Db.Db).Query()
	if err != nil {
		return models.GetEntryResponse{}, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		var entry models.Entry
		err = rows.Scan(
			&entry.Id,
			&entry.StaffId,
			&entry.ActivityType,
			&entry.City,
			&count,
			&entry.Date,
		)
		if err != nil {
			return models.GetEntryResponse{}, err
		}
		res = append(res, entry)
	}
	if len(res) == 0 {
		return models.GetEntryResponse{}, fmt.Errorf("no data found")
	}
	return models.GetEntryResponse{Entry: res, Count: count}, nil
}
func (s *postgresRepo) DeleteStaffEntry(ctx context.Context, id string) error {
	query := s.Db.Builder.Delete("entries").Where("id = ?", id)
	_, err := query.RunWith(s.Db.Db).Exec()
	if err != nil {
		return err
	}
	return nil
}
func (s *postgresRepo) UpdateStaffEntry(ctx context.Context, req models.Entry) error {

	var (
		mp = make(map[string]interface{})
	)
	mp["activity_type"] = req.ActivityType
	mp["city"] = req.City
	mp["staff_id"] = req.StaffId
	mp["updated_at"] = time.Now().Format(time.RFC3339)

	query := s.Db.Builder.Update("entries").SetMap(mp).Where("id = ?", req.Id)
	_, err := query.RunWith(s.Db.Db).Exec()
	if err != nil {
		return err
	}

	return nil
}
