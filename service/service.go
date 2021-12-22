package service

import (
	"todo-api/config"
	"todo-api/logger"

	"gorm.io/gorm"
)

type Service struct {
	Logger logger.Logger
	Config *config.Config
	DB     *gorm.DB
}
