package domain

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/storages/models"
	"github.com/shanenoi/togo/internal/storages/postgresql"
	"github.com/shanenoi/togo/internal/storages/redis"
)

type TaskDomain interface {
	GetlistTask(uid float64) ([]models.Task, error)
	CreateOneTask(uid float64, data []byte) (message string, ok bool)
}

func NewTaskDomain(ctx *gin.Context) TaskDomain {
	return &taskDomain{ctx}
}

type taskDomain struct {
	ctx *gin.Context
}

func (td *taskDomain) GetlistTask(uid float64) ([]models.Task, error) {
	taskPostgreSQL := postgresql.NewTaskPostgreSQL(td.ctx)
	return taskPostgreSQL.ListTasks(uid)
}

func (td *taskDomain) CreateOneTask(uid float64, data []byte) (message string, ok bool) {
	task := &models.Task{Uid: uid, Data: data}

	taskRedis := redis.NewTaskRedis(td.ctx)
	taskPostgreSQL := postgresql.NewTaskPostgreSQL(td.ctx)

	userSlot, _ := taskRedis.GetSlotByUid(uid)
	ok = true

	if userSlot < config.MAX_TASK_PER_DAY {
		message = fmt.Sprintf(config.RESP_REMAINING, config.MAX_TASK_PER_DAY - userSlot)
		duration, _ := time.ParseDuration(config.LIMIT_DURATION)
		taskRedis.UpdateSlotByUid(uid, userSlot + 1, duration)
		go taskPostgreSQL.CreateTask(task)
	} else {
		message = config.RESP_OUT_OF_SLOT
		ok = false

	}
	return
}
