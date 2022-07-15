package cache

import (
	"context"

	"github.com/datshiro/togo-manabie/internal/interfaces/models"
)

func NewService(redis Redis) CacheService {
	return &cacheService{redisClient: redis}
}

type CacheService interface {
	IncreaseQuota(context.Context, *models.User) error
	ValidateQuota(context.Context, *models.User) (bool, error)
}
