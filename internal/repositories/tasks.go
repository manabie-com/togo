package repositories

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	TasksTableName = "tasks"
	fieldTaskUserIDName = "user_id"
	fieldCreatedAtName = "created_at"
)

func (Task) TableName() string {
	return TasksTableName
}

type TaskRepo interface {
	AddTask(task *Task) (*Task, error)
	ListTasks(userID, createdAt string) (*[]Task, error)
}

type TaskRepoImpl struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepoImpl {
	return &TaskRepoImpl{
		db: db,
	}
}

func (t *TaskRepoImpl) AddTask(task *Task) (*Task, error) {
	if err := t.db.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskRepoImpl) ListTasks(userID, createdAt string) (*[]Task, error) {
	var tasks []Task
	err := t.db.Where(fmt.Sprintf("%s = ? AND date(%s) = ?", fieldTaskUserIDName, fieldCreatedAtName), userID, createdAt).Find(&tasks).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &tasks, nil
		}
		return nil, err
	}
	return &tasks, nil
}