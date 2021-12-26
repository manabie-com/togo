package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/helpers"
	"github.com/manabie-com/togo/services"
)

type ITaskController interface {
	GetAllTask(context *gin.Context)
}

type TaskController struct {
	TaskService services.ITaskService
}

func NewTaskController(taskService services.ITaskService) ITaskController {
	return &TaskController{
		TaskService: taskService,
	}
}

func (taskController *TaskController) GetAllTask(context *gin.Context) {
	tasks, resErr := taskController.TaskService.GetAllTask()
	if resErr != nil {
		context.JSON(http.StatusInternalServerError, *resErr)
		return
	}

	context.JSON(http.StatusOK, *helpers.BuildResponse(true, "Success", tasks))
}
