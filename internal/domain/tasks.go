package domain

import (
	"fmt"
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

func(td *taskDomain) GetlistTask(uid float64) ([]models.Task, error) {
	taskPostgreSQL := postgresql.NewTaskPostgreSQL(td.ctx)
	return taskPostgreSQL.ListTasks(uid)
}

func(td *taskDomain) CreateOneTask(uid float64, data []byte) (message string, ok bool) {
	task := &models.Task{Uid: uid, Data: data}

	taskRedis := redis.NewTaskRedis(td.ctx)
	taskPostgreSQL := postgresql.NewTaskPostgreSQL(td.ctx)

	var remaining int
	remaining, ok = taskRedis.RetrieveUserSlotAvailable(uid)
	message = fmt.Sprintf(config.RESP_REMAINING, config.MAX_TASK_PER_DAY-remaining)

	if ok {
		go taskPostgreSQL.CreateTask(task)
	} else {
		message = config.RESP_OUT_OF_SLOT
	}
	return
}
