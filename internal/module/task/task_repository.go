package task

import "github.com/jinzhu/gorm"

// Repository interface
type Repository interface {
	AddTask(userID uint64, content string) (Task, error)
	RetrieveTasks(userID uint64, createdDate string) ([]Task, error)
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
	return tasks, err
}
