package redis

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/shanenoi/togo/config"
	"log"
)

type RedisWrapper struct {
	Database *redis.Client
	Contexy  *context.Context
}

func Connect(addr string, password string, db int) *RedisWrapper {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	ctx := context.Background()

	log.Printf(config.CONNECTED_TO, "Redis Server")
	return &RedisWrapper{rdb, &ctx}
}

func GetRedis(ctx *gin.Context) (*RedisWrapper, bool) {
	res, ok := ctx.Get(config.REDIS_DB)
	return res.(*RedisWrapper), ok
}
