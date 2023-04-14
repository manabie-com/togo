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

func (s *sqlStore) CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return sdkcm.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return sdkcm.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) GetTask(ctx context.Context, cond map[string]interface{}) (*taskmodel.Task, error) {
	var data taskmodel.Task

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sdkcm.RecordNotFound
		}

		return nil, sdkcm.ErrDB(err)
	}

	return &data, nil
}

func (s *sqlStore) UpdateTask(ctx context.Context, cond map[string]interface{}, dataUpdate *taskmodel.TaskUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DeleteTask(ctx context.Context, cond map[string]interface{}) error {
	var task taskmodel.Task
	if err := s.db.Table(taskmodel.Task{}.TableName()).Where(cond).Delete(task).Error; err != nil {
		return sdkcm.ErrDB(err)
	}

	return nil
}
