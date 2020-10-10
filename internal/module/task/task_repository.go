package task

import (
	"time"

	"gorm.io/gorm"
)

// Repository interface
type Repository interface {
	AddTask(userID uint64, content string) (Task, error)
	AddManyTasks(userID uint64, contents []string) ([]Task, error)
	RetrieveTasks(userID uint64, createdDate string) ([]Task, error)
	NumTasksToday(userID uint64) (int64, error)
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

	err := repo.db.Create(&task).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (repo *repository) AddManyTasks(userID uint64, contents []string) ([]Task, error) {

	// WARNING: Max parameters PostgreSQL supports is 65535

	tasks := []Task{}
	for _, content := range contents {
		tasks = append(tasks, Task{Content: content, UserID: userID, Status: StatusActive})
	}
	err := repo.db.Create(&tasks).Error
	if err != nil {
		return []Task{}, err
	}
	return tasks, nil
}

func (repo *repository) RetrieveTasks(userID uint64, createdDate string) ([]Task, error) {
	var tasks []Task
	err := repo.db.Where("user_id = ? AND to_char(created_date,'YYYY-MM-DD') = ?", userID, createdDate).Find(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (repo *repository) NumTasksToday(userID uint64) (int64, error) {
	currentTime := time.Now()

	var count int64
	err := repo.db.Raw("SELECT count(*) FROM tasks WHERE status = ? AND user_id = ? AND to_char(created_date,'YYYY-MM-DD') = ?", StatusActive, userID, currentTime.Format("2006-01-02")).Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
