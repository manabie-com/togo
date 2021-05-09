package services

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/google/martian/log"
	"github.com/manabie-com/togo/internal/config"
	"strconv"
	"time"
)

type ICache interface {
	SetMaxTodo(userId string, maxTodo int32) error
	GetMaxTodo(userId string) (int32, error)
	GetNumberOfTasks(userId string, createdDate string) (int32, error)
	SetNumberOfTasks(userId string, createdDate string, count int32) error
	IncTask(userId string, createdDate string) error
}

var (
	ErrCastValue = errors.New("cannot cast interface{} value from redis")

)

type RedisCache struct {
	redisPool *redis.Pool
}

func (r *RedisCache) key(userId, createdDate string) string {
	return fmt.Sprintf("%s_%s", userId, createdDate)
}

func (r *RedisCache) command(command string, params ...interface{}) (interface{}, error) {
	return r.redisPool.Get().Do(command, params...)
}

func (r *RedisCache) GetNumberOfTasks(userId string, createdDate string) (int32, error) {
	data, err := r.command("GET", r.key(userId, createdDate))
	if err != nil {
		log.Errorf("error while getting number of task from cache - %s", err.Error())
		return -1, err
	}
	return getNumberFromRedis(data)
}

func getNumberFromRedis(data interface{}) (int32, error) {
	if data == nil {
		return -1, nil
	}
	// cast to string
	result, ok := data.([]byte)
	if !ok {
		log.Errorf(ErrCastValue.Error())
		return -1, ErrCastValue
	}
	if string(result) == "" {
		return -1, nil
	}
	value, err := strconv.Atoi(string(result))
	if err != nil {
		log.Errorf(err.Error())
		return -1, err
	}
	return int32(value), nil
}

func (r *RedisCache) IncTask(userId string, createdDate string) error {
	_, err := r.command("INCR", r.key(userId, createdDate))
	if err != nil {
		log.Errorf("error while increasing number of task - %s", err.Error())
	}
	return err
}

func (r *RedisCache) SetNumberOfTasks(userId string, createdDate string, count int32) error {
	_, err := r.command("SET", r.key(userId, createdDate), count)
	if err != nil {
		log.Errorf("error while setting number of task of key:%s into task - %s", r.key(userId, createdDate), err.Error())
	}
	return err
}

func (r *RedisCache) SetMaxTodo(userId string, maxTodo int32) error {
	_, err := r.command("SET", userId, maxTodo)
	if err != nil {
		log.Errorf("error while setting maxTodo for userId:%s - %s", userId, err.Error())
	}
	return err
}

func (r *RedisCache) GetMaxTodo(userId string) (int32, error) {
	data, err := r.command("GET", userId)
	if err != nil {
		log.Errorf("error while getting maxTodo from userId:%s - %s", userId, err.Error())
		return -1, err
	}
	return getNumberFromRedis(data)
}

func newRedisPool(cfg *config.Config) *redis.Pool {
	return &redis.Pool{
		// Other pool configuration not shown in this example.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Redis.Address)
			if err != nil {
				return nil, err
			}
			if cfg.Redis.Password != "" {
				if _, err := c.Do("AUTH", cfg.Redis.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", cfg.Redis.DatabaseNum); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		MaxIdle:         cfg.Redis.MaxIdle,
		MaxActive:       cfg.Redis.MaxActive,
		IdleTimeout:     time.Duration(cfg.Redis.MaxIdleTimeout) * time.Second,
		Wait:            cfg.Redis.Wait,
		MaxConnLifetime: time.Duration(cfg.Redis.ConnectTimeout) * time.Second,
	}
}
