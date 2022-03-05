package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/khangjig/togo/repository/todo"
	"github.com/khangjig/togo/repository/user"
)

type Repository struct {
	User      user.Repository
	UserCache user.CacheRepository
	Todo      todo.Repository
	TodoCache todo.CacheRepository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		User:      user.NewPG(getSQLClient),
		UserCache: user.NewCacheRepository(getRedisClient, getSQLClient),
		Todo:      todo.NewPG(getSQLClient),
		TodoCache: todo.NewCacheRepository(getRedisClient),
	}
}
