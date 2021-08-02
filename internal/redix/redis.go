package redix

import "github.com/go-redis/redis/v8"

type redisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *redisStore {
	return &redisStore{client: client}
}
