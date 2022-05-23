package taskstorage

import (
	"context"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/modules/task/taskmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindTaskByCondition(
	_ context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*taskmodel.Task, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var task taskmodel.Task

	if err := db.
		Where(conditions).
		First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &task, nil
}
