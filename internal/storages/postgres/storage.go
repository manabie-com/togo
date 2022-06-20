package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/storages/entities"
)

var (
	storageManager *StorageManager
	once           sync.Once
	queryLimit     = 1000
)

// GetStorageManager init singleton of Postgres Impl for StorageManager
func GetStorageManager(host, port, username, pwd, dbname string) *StorageManager {
	once.Do(func() {
		psqlInfo := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, username, pwd, dbname,
		)

		retry := 0
		for {
			// Connect to db, retry 3 times
			db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
			if err != nil {
				if retry >= 3 {
					log.Fatalln("Failed to connect to db after 3 retries")
				}

				retry++
				time.Sleep(time.Second * time.Duration(retry+1))
				continue
			}

			if strings.EqualFold(os.Getenv("DB_DEBUG"), "true") {
				db = db.Debug()
			}

			storageManager = &StorageManager{
				db: db,
			}
			break
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

func (s *StorageManager) RetrieveTasks(ctx context.Context, userID string, date time.Time) ([]*entities.Task, error) {
	var tasks []*entities.Task
	result := s.db.WithContext(ctx).
		Where(&entities.Task{
			UserID:      userID,
			CreatedDate: date.Format("2006-01-02"),
		}).
		Limit(queryLimit).Find(&tasks)
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

func (s *StorageManager) ValidateUser(ctx context.Context, userID, pwd string) bool {
	var user entities.User
	result := s.db.WithContext(ctx).First(&user, &entities.User{ID: userID})
	if result.Error != nil {
		return false
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd)); err != nil {
		return false
	}

	return true
}

func (s *StorageManager) AddUser(ctx context.Context, userID, pwd string) error {
	result := s.db.WithContext(ctx).Where(&entities.User{ID: userID})
	if result.Error != nil {
		return result.Error
	}

	// User existed
	if result.RowsAffected > 0 {
		return fmt.Errorf("user is already existed")
	}

	return s.db.WithContext(ctx).Create(&entities.User{
		ID:       userID,
		Password: common.HashPassword(pwd),
	}).Error
}
