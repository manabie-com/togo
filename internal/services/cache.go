package services

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/google/martian/log"
)

type ICache interface {
	SetMaxTodo(userId string, maxTodo int32) error
	GetMaxTodo(userId string) (int32, error)
	GetNumberOfTasks(userId string, createdDate string) (int32, error)
	SetNumberOfTasks(userId string, createdDate string, count int32) error
	IncTask(userId string, createdDate string) error
}

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
	if data == nil {
		return -1, nil
	}
	result, ok := data.(int32)
	if !ok {
		err := errors.New("cannot cast interface{} into int32")
		log.Errorf("cannot cast interface{} into int32")
		return -1, err
	}
	return result, nil
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
	result, ok := data.(int32)
	if !ok {
		log.Errorf("cannot cast from interface{} to int32")
		return -1, nil
	}
	return result, nil
}
