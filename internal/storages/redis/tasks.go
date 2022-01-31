package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TaskRedis interface {
	GetSlotByUid(uid float64) (int, error)
	UpdateSlotByUid(uid float64, value interface{}, expiration time.Duration) error
}

type taskRedis struct {
	ctx *gin.Context
}

func NewTaskRedis(ctx *gin.Context) TaskRedis {
	return &taskRedis{ctx}
}

func (trs *taskRedis) GetSlotByUid(uid float64) (int, error) {
	redis, _ := GetRedis(trs.ctx)
	rawData, err := redis.Database.Get(*redis.Contexy, fmt.Sprintf("%f", uid)).Result()
	userSlot := 0

	if err == nil {
		userSlot, err = strconv.Atoi(rawData)
	}

	return userSlot, err
}

func (trs *taskRedis) UpdateSlotByUid(uid float64, value interface{}, expiration time.Duration) error {
	redis, _ := GetRedis(trs.ctx)
	return redis.Database.Set(*redis.Contexy, fmt.Sprintf("%f", uid), value, expiration).Err()
}
