package repository

import (
	"log"

	mysqlDb "github.com/manabie-com/backend/datasource"
	"github.com/manabie-com/backend/entity"
	"github.com/manabie-com/backend/utils"
)

type I_Repository interface {
	GetTaskAll() ([]entity.Task, *utils.ErrorRest)
	GetTask(id string) (*entity.Task, *utils.ErrorRest)
	UpdateTask(*entity.Task) *utils.ErrorRest
	DeleteTask(id string) *utils.ErrorRest
	CreateTask(*entity.Task) *utils.ErrorRest
	FindTaskByUserIdAndDay(string, string) ([]entity.Task, *utils.ErrorRest)
	FindTaskByContent(string) (*entity.Task, *utils.ErrorRest)
}

func NewRepository() I_Repository {
	conn, err := mysqlDb.Conn()
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}
	return &mysqlRepository{
		mysqlDb: conn,
	}
}
