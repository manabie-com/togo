package controllers

import (
	"time"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	models "github.com/manabie-com/togo/internal/models"
	services "github.com/manabie-com/togo/internal/services"
	resources "github.com/manabie-com/togo/internal/resources"
)

type TaskController struct {
	TaskService services.TaskService
}

func ProvideTaskController(service services.TaskService) TaskController {
	return TaskController{TaskService: service}
}

func (ctrl *TaskController) FindAll(c *gin.Context) {
	tasks := ctrl.TaskService.FindAll()
	ResponseJSON(c, resources.TaskResources(tasks))
}

func (ctrl *TaskController) FindByID(c *gin.Context, id string) {
	task := ctrl.TaskService.FindByID(id)
	ResponseJSON(c, resources.ToTaskResource(task))
}

func (ctrl *TaskController) Create(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		ResponseError(c, 500, err.Error())
	}

	ctxUser, _ := c.Get("user")
	user := UserFromCtx(ctxUser)
	
	task.ID = uuid.Must(uuid.NewUUID())
	task.UserID = user.ID
	task.CreatedDate = time.Now().Format("2006-01-02")
	taskCreated := ctrl.TaskService.Create(task)
	ResponseJSON(c, resources.ToTaskResource(taskCreated))
}

func (ctrl *TaskController) Update(c *gin.Context, id string) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		ResponseError(c, 500, err.Error())
	}

	currentTask := ctrl.TaskService.FindByID(id)
	if currentTask == (models.Task{}) {
		ResponseError(c, 400, "Bad request")
		return
	}
	ctrl.TaskService.Create(task)
	ResponseJSON(c, task)
}

func (ctrl *TaskController) Delete(c *gin.Context, id string) {
	task := ctrl.TaskService.FindByID(id)
	if task == (models.Task{}) {
		ResponseError(c, 400, "Bad request")
		return
	}

	ctrl.TaskService.Delete(task)
	ResponseJSON(c, Empty{})
}