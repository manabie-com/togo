package taskstorage

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/phathdt/libs/go-sdk/sdkcm"
)

type redisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *redisStore {
	return &redisStore{client: client}
}

func (s *redisStore) IncrBy(ctx context.Context, key string, number int) (int, error) {
	result, err := s.client.IncrBy(ctx, key, int64(number)).Result()
	if err == redis.Nil {
		return 0, sdkcm.ErrEntityNotFound("limit", err)
	} else if err != nil {
		return 0, sdkcm.ErrDB(err)
	}

	return int(result), nil
}

func (s *redisStore) Get(ctx context.Context, key string) (string, error) {
	result, err := s.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", sdkcm.RecordNotFound
	} else if err != nil {
		return "", sdkcm.ErrDB(err)
	}

	return result, nil
}
