package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type Client struct {
	Store *redis.Client
}

func Init() Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return Client{Store: rdb}
}

func (r *Client) SetNX(ctx context.Context, key string, value int32) error {
	err := r.Store.SetNX(ctx, key, value, time.Hour * 24).Err()
	if err != nil {
		return errors.New("Redis SetEX Failed")
	}
	return nil
}

func (r *Client) Incr(ctx context.Context, key string) (int64, error) {
	incr := r.Store.Incr(ctx, key)
	if incr.Err() != nil {
		return 0, errors.New("Redis Incr Failed")
	}
	return incr.Val(), incr.Err()
}