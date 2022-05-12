package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nvhai245/togo/internal/domain"
)

type taskController struct {
	taskService domain.TaskService
}

// TaskController represents the controller for task
type TaskController interface {
	CreateTask(c *gin.Context)
}

// NewTaskController creates a new task controller and inject task service
func NewTaskController(s domain.TaskService) TaskController {
	return &taskController{taskService: s}
}

// CreateTask creates a new task for a user
func (t *taskController) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.String(400, "error: %s", errors.New("invalid user_id"))
		return
	}

	task := c.Query("task")
	if task == "" {
		c.String(400, "error: %s", errors.New("task is empty"))
		return
	}

	_, err = t.taskService.CreateTask(ctx, int64(userID), task)
	if err != nil {
		c.String(400, "error: %s", err.Error())
		return
	}

	userTasks, err := t.taskService.GetAllTaskByUserID(ctx, int64(userID))
	if err != nil {
		c.String(400, "error: %s", err.Error())
		return
	}

	c.JSON(200, userTasks)
	return
}
