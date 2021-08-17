package storage

import (
	"context"
	"github.com/manabie-com/togo/task/model"
)

func (s *taskStorage) Count(ctx context.Context, conditions map[string]interface{}) (*int64, error) {
	db := s.db

	db = db.Table(model.Task{}.TableName())
	db = db.Select("select count(*) where date(created_date) = ? and user_id = ?", conditions["created_at"], conditions["user_id"])
	var count int64

	if err := db.Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}
