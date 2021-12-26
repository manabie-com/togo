package repositories

import (
	"github.com/manabie-com/togo/entities"
	"github.com/manabie-com/togo/helpers"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAllTask() ([]entities.Task, *helpers.Response)
}

type taskConnection struct {
	connection *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskConnection{
		connection: db,
	}
}

func (taskConn *taskConnection) GetAllTask() ([]entities.Task, *helpers.Response) {
	var tasks []entities.Task
	err := taskConn.connection.Find(&tasks).Error
	if err != nil {
		return nil, helpers.BuildErrorResponse("An Error occurred!", err.Error(), nil)
	}
	return tasks, nil
}
