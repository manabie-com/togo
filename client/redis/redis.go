package redis

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/khangjig/togo/client/logger"
	"github.com/khangjig/togo/config"
)

var rd *redis.Client

func init() {
	cfg := config.GetConfig()

	rd = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Pass,
		DB:           cfg.Redis.DB,
		MaxRetries:   3,
		PoolSize:     10 * runtime.NumCPU(),
		DialTimeout:  time.Second * time.Duration(cfg.Redis.Timeout),
		ReadTimeout:  time.Second * time.Duration(cfg.Redis.Timeout),
		WriteTimeout: time.Second * time.Duration(cfg.Redis.Timeout),
	})

	pong, err := rd.Ping(context.Background()).Result()
	if err != nil {
		logger.GetLogger().Fatal(fmt.Sprintf("Fail to connect to redis %v", err))
	}

	logger.GetLogger().Info(fmt.Sprintf("Connected redis %v", pong))
}

func GetClient(ctx context.Context) *redis.Client {
	return rd
}
