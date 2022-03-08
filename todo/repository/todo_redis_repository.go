package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/triet-truong/todo/todo/model"
	"github.com/triet-truong/todo/utils"
)

type TodoRedisRepository struct {
	rdb *redis.Client
}

func NewTodoRedisRepository(opts redis.Options) *TodoRedisRepository {
	client := redis.NewClient(&opts)
	err := client.Ping(context.TODO()).Err()
	utils.FatalLog(err)
	return &TodoRedisRepository{
		rdb: client,
	}
}

func (r *TodoRedisRepository) SetUser(user model.UserRedisModel) error {
	json, _ := json.Marshal(user)

	return r.rdb.Set(context.TODO(), fmt.Sprint(user.ID), string(json), time.Until(utils.EndOfCurrentDate())).Err()

}
func (r *TodoRedisRepository) GetCachedUser(id uint) (model.UserRedisModel, error) {
	var res model.UserRedisModel
	val, err := r.rdb.Get(context.TODO(), fmt.Sprint(id)).Result()
	json.Unmarshal(bytes.NewBufferString(val).Bytes(), &res)
	utils.ErrorLog(err)
	return res, err
}
