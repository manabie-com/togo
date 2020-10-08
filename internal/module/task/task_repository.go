package task

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Repository interface
type Repository interface {
	AddTask(userID uint64, content string) (Task, error)
	RetrieveTasks(userID uint64, createdDate string) ([]Task, error)
	NumTasksToday(userID uint64) (int, error)
}

// NewTaskRepository func
func NewTaskRepository(db *gorm.DB) (Repository, error) {
	return &repository{
		db: db,
	}, nil
}

type repository struct {
	db *gorm.DB
}

func (repo *repository) AddTask(userID uint64, content string) (Task, error) {
	task := Task{Content: content, UserID: userID, Status: StatusActive}
	repo.db.Create(&task)
	return task, nil
}

func (repo *repository) RetrieveTasks(userID uint64, createdDate string) ([]Task, error) {
	var tasks []Task
	err := repo.db.Where("user_id = ? AND to_char(created_date,'YYYY-MM-DD') = ?", userID, createdDate).Find(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (repo *repository) NumTasksToday(userID uint64) (int, error) {
	var result int
	currentTime := time.Now()

	err := repo.db.Table("tasks").Where("user_id = ? AND to_char(created_date,'YYYY-MM-DD') = ?", userID, currentTime.Format("2006-01-02")).Count(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}
