package task

import (
	"github.com/manabie-com/togo/internal/storages"

	"github.com/jinzhu/gorm"
)

type taskStorage struct {
	db *gorm.DB
}

func NewTaskStorage(db *gorm.DB) TaskStorageInterface {
	return &taskStorage{
		db: db,
	}
}

func (s *taskStorage) CreateTask(task *storages.Task) error {
	tx := s.db.Begin()
	if err := tx.Create(task).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *taskStorage) RetrieveTasks(userID, createdDate string) ([]*storages.Task, error) {
	tasks := []*storages.Task{}
	tx := s.db.Begin()
	err := tx.Scopes(filterCreatedDate(createdDate)).
		Where(storages.Task{UserID: userID}).
		Find(&tasks).Error

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return tasks, nil
}

func filterCreatedDate(date string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if date == "" {
			return db
		}

		return db.Where("created_date = ?", date)
	}
}
