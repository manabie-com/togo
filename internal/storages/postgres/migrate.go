package postgres

import (
	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/storages/entities"
)

func Migrate(s *StorageManager) error {
	if err := s.db.AutoMigrate(&entities.Task{}, &entities.User{}, &entities.RateLimit{}); err != nil {
		return err
	}

	manabie := entities.User{
		ID:       "manabie",
		Password: common.HashPassword("manabie"),
	}
	if err := s.db.Save(&manabie).Error; err != nil {
		return err
	}
	return nil
}
