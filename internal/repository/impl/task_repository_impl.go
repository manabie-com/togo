package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/model"
)

type TaskRepositoryImpl struct{
	db *gorm.DB
}

func NewTaskRepositoryImpl(db *gorm.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{db}
}

func (r *TaskRepositoryImpl) GetByIdAndCreateDate(id string, createdDate string) (model.TaskList, error) {
	var taskList model.TaskList

	if err := r.db.Where("user_id = ? AND created_date = ?", id, createdDate).Find(&taskList).Error; err != nil {
		// error handling..
		return nil, err
	}
	return taskList, nil
}

func (r *TaskRepositoryImpl) CountByIdAndCreateDate(id string, createdDate string) (int, error) {
	var results = -1

	if err := r.db.Table("tasks").Where("user_id = ? AND created_date = ?", id, createdDate).Count(&results).Error; err != nil {
		// error handling..
		return results, err
	}
	return results, nil
}

func (r *TaskRepositoryImpl) Save(task *model.Task) (*model.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}