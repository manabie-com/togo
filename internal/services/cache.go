package services

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/google/martian/log"
	"strconv"
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
	return r.getNumberFromRedis(data)
}

func (r *RedisCache) getNumberFromRedis(data interface{}) (int32, error) {
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
	return r.getNumberFromRedis(data)
}
