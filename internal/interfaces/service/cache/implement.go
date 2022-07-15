package cache

import (
	"context"

	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/go-redis/redis/v9"
)

func CreateRedisClient(location string) (CacheService, error) {
	options, err := redis.ParseURL(location)
	if err != nil {
		return nil, err
	}

	options.Username = ""
	client := redis.NewClient(options)

	return &cacheService{redisClient: Redis{client: client}}, nil

}

type cacheService struct {
	redisClient Redis
}

func (c *cacheService) IncreaseQuota(ctx context.Context, user *models.User) error {
	addedTaskCount, err := c.redisClient.GetUserQuota(ctx, user.ID)
	if err != nil {
		return err
	}
	return c.redisClient.SetUserQuota(ctx, user.ID, addedTaskCount+1)
}

func (c *cacheService) ValidateQuota(ctx context.Context, user *models.User) (bool, error) {
	addedTaskCount, err := c.redisClient.GetUserQuota(ctx, user.ID)
	if err != nil {
		return false, err
	}
	return addedTaskCount < user.Quota, nil
}
