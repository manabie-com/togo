package postgresql

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/internal/storages/models"
)

type TaskPostgreSQL interface {
	CreateTask(task *models.Task)
	ListTasks(Uid float64) ([]models.Task, error)
}

type taskPostgreSQL struct {
	ctx *gin.Context
}

func NewTaskPostgreSQL(ctx *gin.Context) TaskPostgreSQL {
	return &taskPostgreSQL{ctx}
}

func (tpsql *taskPostgreSQL) CreateTask(task *models.Task) {
	db, _ := GetDb(tpsql.ctx)
	db.Database.Create(task)
}

func (tpsql *taskPostgreSQL) ListTasks(uid float64) ([]models.Task, error) {
	db, _ := GetDb(tpsql.ctx)
	tasks := []models.Task{}
	err := db.Database.Find(&tasks, "uid", uid).Error
	return tasks, err
}
