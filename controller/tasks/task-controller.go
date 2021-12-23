package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/backend/entity"
	"github.com/manabie-com/backend/repository"
	taskservice "github.com/manabie-com/backend/services/task"
	"github.com/manabie-com/backend/utils"

	userServiceValidate "github.com/manabie-com/backend/services/user"
)

type Controller struct {
	Service taskservice.I_TaskService
}

type I_TaskController interface {
	GetTaskAll(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

func NewTaskController(serv taskservice.I_TaskService) I_TaskController {

	return &Controller{
		Service: serv,
	}
}

func (ct *Controller) GetTaskAll(c *gin.Context) {
	tasks, err := ct.Service.GetTaskAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (ct *Controller) UpdateTask(c *gin.Context) {
	repo := repository.NewRepository()
	taskservice := taskservice.NewTaskService(repo)

	id := c.Param("id")
	var task entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": utils.ErrBadRequest("Invalid body json")})
		return
	}

	task.ID = id

	result, err := taskservice.UpdateTask(&task)
	if err != nil {
		c.JSON(err.Code, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})

}

func (ct *Controller) DeleteTask(c *gin.Context) {
	repo := repository.NewRepository()
	taskservice := taskservice.NewTaskService(repo)

	id := c.Param("id")

	err := taskservice.DeleteTask(id)
	if err != nil {
		c.JSON(err.Code, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "ok"})

}

func (ct *Controller) CreateTask(c *gin.Context) {
	var task entity.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": utils.ErrBadRequest("Invalid body json")})

		return
	}

	repo := repository.NewRepository()
	taskServiceValidate := taskservice.NewTaskServiceValidate(repo)

	if err := taskServiceValidate.Validate(&task); err != nil {

		c.JSON(err.Code, gin.H{"data": err})

		return
	}

	userValidate, err := userServiceValidate.NewUserServiceValidate(repo, task.UserID)
	if err != nil {

		c.JSON(err.Code, gin.H{"data": err})

		return
	}

	if err := userValidate.IsAllowedAddTask(); err != nil {

		c.JSON(err.Code, gin.H{"data": err})

		return
	}

	savedTask, savedErr := ct.Service.CreateTask(&task)

	if savedErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedTask})
}
