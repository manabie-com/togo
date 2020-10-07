package task

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/middleware"
)

// Controller interface
type Controller interface {
	AddTask(c echo.Context) error
	RetrieveTasks(c echo.Context) error
}

// NewTaskController func
func NewTaskController(taskService Service) (Controller, error) {
	return &controller{taskService: taskService}, nil
}

type controller struct {
	taskService Service
}

func (controller *controller) AddTask(c echo.Context) error {
	userID := middleware.GetUserIDFromContext(c)

	addTaskDTO := new(dto.AddTaskDTO)
	if err := c.Bind(addTaskDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if err := c.Validate(addTaskDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	task, err := controller.taskService.AddTask(userID, addTaskDTO.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": task,
	})
}

func (controller *controller) RetrieveTasks(c echo.Context) error {

	userID := middleware.GetUserIDFromContext(c)
	createdDate := c.Param("created_date")

	tasks, err := controller.taskService.RetrieveTasks(userID, createdDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tasks,
	})
}
