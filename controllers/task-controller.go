package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/entities"
	"github.com/manabie-com/togo/helpers"
	"github.com/manabie-com/togo/services"
)

type ITaskController interface {
	GetAllTask(context *gin.Context)
	CreateTask(context *gin.Context)
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
	tasks, err := taskController.TaskService.GetAllTask()

	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("An Error occurred", err.Error(), nil))
		return
	}

	context.JSON(http.StatusOK, *helpers.BuildResponse(true, "Get all tasks successfully", tasks))
}

func (taskController *TaskController) CreateTask(context *gin.Context) {
	var task entities.Task

	task.CreatedAt = helpers.GetDateNow()

	if err := context.ShouldBindJSON(&task); err != nil {
		context.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid body json", err.Error(), nil))

		return
	}

	internalErr, userErr := taskController.TaskService.CreateTask(&task)

	if userErr != nil {
		context.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("An user error occurred", userErr.Error(), nil))

		return
	}

	if internalErr != nil {
		context.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("An error occurred", internalErr.Error(), nil))

		return
	}

	context.JSON(http.StatusCreated, helpers.BuildResponse(true, "Task created successfully", task))
}
