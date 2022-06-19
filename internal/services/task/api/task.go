package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
	"togo/internal/pkg/http"
	"togo/internal/services/task/application"
)

type TaskAPI struct {
	taskService *application.TaskService
}

func NewTaskAPI(
	taskService *application.TaskService,
) *TaskAPI {
	return &TaskAPI{taskService}
}

func (api *TaskAPI) AddRoutes(g *gin.Engine) {
	g.POST("/tasks", api.create)
}

func (api *TaskAPI) create(c *gin.Context) {
	var req CreateTaskRequest
	if http.HandleError(c, c.BindJSON(&req)) {
		return
	}
	logrus.Info(req.DueDate, " - ", time.Now())
	if req.DueDate.Before(time.Now()) {
		http.ErrBadRequest(c, "due date must be in the future")
		return
	}
	userID, _ := uuid.Parse(req.UserID)
	addTaskCommand := application.AddTaskCommand{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
	}
	newUser, err := api.taskService.CreateTask(addTaskCommand)
	if http.HandleError(c, err) {
		return
	}
	http.Success(c, newUser, "create task successfully")
}

type CreateTaskRequest struct {
	UserID      string    `json:"userId" binding:"required,uuid"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate" time_format:"2006-01-02T15:04:05Z" binding:"required"`
}
