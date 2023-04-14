package taskstorage

import (
	"context"
	"fmt"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"gorm.io/gorm"
	"task_service/modules/taskmodel"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

func (s *sqlStore) ListItem(ctx context.Context, filter *taskmodel.Filter, paging *sdkcm.Paging) ([]taskmodel.Task, error) {
	var tasks []taskmodel.Task

	db := s.db.Table(taskmodel.Task{}.TableName())

	if f := filter; f != nil {
		if v := f.IsDone; v != nil {
			db = db.Where("is_done = ?", v)
		}

		if userId := f.UserId; userId != 0 {
			db = db.Where("user_id = ?", userId)
		}
	}

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if paging.Cursor == nil {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	} else {
		db = db.Where("id < ?", paging.Cursor)
	}

	if err := db.Select("*").
		Order("id desc").
		Limit(paging.Limit).
		Find(&tasks).Error; err != nil {

		return nil, sdkcm.ErrDB(err)
	}

	if len(tasks) > 0 {
		paging.NextCursor = fmt.Sprintf("%d", tasks[len(tasks)-1].ID)
	}

	paging.HasNext = len(tasks) >= paging.Limit

	return tasks, nil
}
