package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/storages/entities"
)

var (
	storageManager *StorageManager
	once           sync.Once
)

// GetStorageManager init singleton of Postgres Impl for StorageManager
func GetStorageManager(host, port, username, pwd, dbname string) *StorageManager {
	once.Do(func() {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, username, pwd, dbname,
		)

		db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
		if err != nil {
			log.Fatalln("Init DB connection failed", err)
		}

		if strings.EqualFold(os.Getenv("DB_DEBUG"), "true") {
			db = db.Debug()
		}

		storageManager = &StorageManager{
			db: db,
		}

		if err := Migrate(storageManager); err != nil {
			log.Fatalln("Database migration failed", err)
		}
	})
	return storageManager
}

type StorageManager struct {
	db *gorm.DB
}

func (s *StorageManager) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*entities.Task, error) {
	var tasks []*entities.Task
	result := s.db.WithContext(ctx).
		Where("user_id = ? AND created_date = ?", userID, createdDate).
		Limit(1000).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (s *StorageManager) AddTask(ctx context.Context, task *entities.Task) error {
	result := s.db.WithContext(ctx).Table("tasks").Create(task)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *StorageManager) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	user := entities.User{}
	result := s.db.WithContext(ctx).First(&user, "id = ?", userID)
	if result.Error != nil {
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd.String)); err != nil {
		return false
	}

	return true
}

func (s *StorageManager) AddUser(ctx context.Context, userID, pwd string) error {
	result := s.db.WithContext(ctx).Where("id = ?", userID)
	if result.Error != nil {
		return result.Error
	}

	// User exist
	if result.RowsAffected > 0 {
		return fmt.Errorf("user is already existed")
	}

	return s.db.WithContext(ctx).Create(&entities.User{
		ID:       userID,
		Password: common.HashPassword(pwd),
	}).Error
}
