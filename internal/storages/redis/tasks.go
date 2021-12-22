package redis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"strconv"
	"time"
)

type TaskRedis interface {
	RetrieveUserSlotAvailable(uid float64) (int, bool)
}

type taskRedis struct {
	ctx *gin.Context
}

func NewTaskRedis(ctx *gin.Context) TaskRedis {
	return &taskRedis{ctx}
}

func (trs *taskRedis) RetrieveUserSlotAvailable(uid float64) (int, bool) {
	redis, _ := GetRedis(trs.ctx)

	userSlot := 0
	rawData, err := redis.Database.Get(*redis.Contexy, fmt.Sprintf("%f", uid)).Result()

	if err != nil {
		userSlot = 0
	} else {
		userSlot, err = strconv.Atoi(rawData)
	}

	if userSlot >= config.MAX_TASK_PER_DAY {
		return 0, false
	}

	userSlot += 1
	duration, _ := time.ParseDuration(config.LIMIT_DURATION)
	err = redis.Database.Set(*redis.Contexy, fmt.Sprintf("%f", uid), userSlot, duration).Err()

	if err != nil {
		panic(err)
	}

	return userSlot, true
}
