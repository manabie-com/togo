package user

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/storages"
)

type userStorage struct {
	db *gorm.DB
}

func NewUserStorage(db *gorm.DB) UserStorageInterface {
	return &userStorage{
		db: db,
	}
}

func (s *userStorage) GetUser(id, password string) error {
	user := storages.User{}
	return s.db.Where(storages.User{ID: id, Password: password}).Find(&user).Error
}

func (s *userStorage) GetUsersTasks(userID string, createdDate string) (storages.User, error) {
	fmt.Println("userID", userID)
	user := storages.User{}

	tx := s.db.Begin()
	err := tx.Preload("Tasks", filterCreatedDate(createdDate)).
		Where(storages.User{ID: userID}).
		Find(&user).Error

	if err != nil {
		tx.Rollback()
		return storages.User{}, err
	}

	return user, err
}

func filterCreatedDate(date string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if date == "" {
			return db
		}

		return db.Where("created_date = ?", date)
	}
}
