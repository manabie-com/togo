package clients

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConf struct {
	URL                   string `envconfig:"REDIS_URL" required:"true" default:"redis://localhost:6379"`
	MaxRetries            int    `envconfig:"REDIS_MAX_RETRIES" default:"3"`
	IdleTimeoutSecond     int64  `envconfig:"REDIS_IDLE_TIMEOUT_SECOND" default:"300"`
	TaskPerDayExpiredTime int64  `envconfig:"TASK_PER_DAY_EXPIRED_TIME_SECOND" default:"86400"`
}

func InitRedisClient(c RedisConf) (*redis.Client, error) {
	options, err := redis.ParseURL(c.URL)
	if err != nil {
		return nil, err
	}
	options.MaxRetries = c.MaxRetries
	options.IdleTimeout = time.Duration(c.IdleTimeoutSecond) * time.Second
	redisClient := redis.NewClient(options)
	return redisClient, nil
}
