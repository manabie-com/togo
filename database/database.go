package database

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DBConnection struct {
	SQLdb *gorm.DB
	Rdb   *redis.Client
}
