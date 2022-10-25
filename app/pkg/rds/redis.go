package rds

import (
	"ansidev.xyz/pkg/log"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05"
)

func NewRedisClient(config RedisConfig) *redis.Client {
	address := fmt.Sprintf(DsnFormat, config.RedisHost, config.RedisPort)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.RedisPassword,
		DB:       DefaultDb,
	})

	return redisClient
}

func NewRedisDB(context context.Context, redisClient *redis.Client) *RedisDB {
	_, err := redisClient.Ping(context).Result()
	log.FatalfIf(err, "Could not connect with Redis")

	redisStorage := &RedisDB{
		context: context,
		storage: redisClient,
	}

	return redisStorage
}

type RedisDB struct {
	context context.Context
	storage *redis.Client
}

func (s *RedisDB) Exists(key string) bool {
	exists, err := s.storage.Exists(s.context, key).Result()

	if err != nil {
		log.Error(err)
		return false
	}

	if exists == 0 {
		return false
	} else {
		return true
	}
}

func (s *RedisDB) Expire(key string, timeToLive time.Duration) {
	log.Debug("Set TTL for key ", key, ", TTL = ", timeToLive)
	s.storage.Expire(s.context, key, timeToLive)
}

func (s *RedisDB) ExpireAt(key string, time time.Time) {
	log.Debug("Set expire time for key ", key, ", time = ", time.Format(DateTimeFormat))
	s.storage.ExpireAt(s.context, key, time)
}

func (s *RedisDB) GetAsString(key string) (string, error) {
	v := s.storage.Get(s.context, key)
	return v.Result()
}

func (s *RedisDB) GetAsBytes(key string) ([]byte, error) {
	return s.storage.Get(s.context, key).Bytes()
}

func (s *RedisDB) Set(key string, value interface{}, timeToLive time.Duration) (string, error) {
	log.Debug("Save key ", key, ", value = ", value, ", TTL = ", timeToLive)
	return s.storage.Set(s.context, key, value, timeToLive).Result()
}

func (s *RedisDB) Delete(key string) (int64, error) {
	log.Debug("Delete key ", key)
	return s.storage.Del(s.context, key).Result()
}
