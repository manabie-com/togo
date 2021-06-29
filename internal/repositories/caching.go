package repositories

import (
	"github.com/go-redis/redis"
	"time"
)

type CacheManagerImpl struct {
	Client *redis.Client
}

type CachingRepo interface {
	Get(key string) (string, error)
	Increase(key string, expireTime int64) error
}


func NewCacheManager(r *redis.Client) *CacheManagerImpl {
	return &CacheManagerImpl{
		Client: r,
	}
}

func (c *CacheManagerImpl) Get(key string) (string, error) {
	return c.Client.Get(key).Result()
}

func (c *CacheManagerImpl) Increase(key string, expireTime int64) error {
	_, err := c.Client.Incr(key).Result()
	if err != nil {
		return err
	}
	if expireTime > 0 {
		c.Client.Expire(key, time.Duration(expireTime)*time.Second)
	}
	return nil
}
