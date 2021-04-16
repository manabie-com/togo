package postgres

import (
	"github.com/manabie-com/togo/internal/storages/entities"
)

func Migrate(s *StorageManager) error {
	return s.db.AutoMigrate(&entities.Task{}, &entities.User{}, &entities.RateLimit{})
}
