package storage

import (
	"context"
	"github.com/manabie-com/togo/task/model"
)

func (s *taskStorage) Create(ctx context.Context, data *model.Task) error {
	db := s.db.Begin()

	if err := db.Table("tasks").Create(data).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}

	return nil
}
