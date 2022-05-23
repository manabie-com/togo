package taskstorage

import (
	"context"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/modules/task/taskmodel"
	"time"
)

func (s *sqlStore) CreateTask(_ context.Context, data *taskmodel.TaskCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) CountUserDailyTask(_ context.Context, createdBy int) (int, error) {
	db := s.db
	layoutISO := "2006-01-02"
	todayDate := time.Now().Format(layoutISO)
	tomorrowDate := time.Now().AddDate(0, 0, 1).Format(layoutISO)
	var count int64

	db = db.Table(taskmodel.Task{}.TableName()).
		Where("created_by = ? ", createdBy).
		Where("created_at >= ? and created_at < ?", todayDate, tomorrowDate).
		Count(&count)

	return int(count), nil
}
