package task

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/middleware"
	"github.com/manabie-com/togo/internal/module/user"
)

// Controller interface
type Controller interface {
	AddTask(c echo.Context) error
	RetrieveTasks(c echo.Context) error
	AddManyTasks(c echo.Context) error
}

// NewTaskController func
func NewTaskController(taskService Service, userService user.Service) (Controller, error) {
	return &controller{taskService: taskService, userService: userService}, nil
}

type controller struct {
	taskService Service
	userService user.Service
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

	user, err := controller.userService.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	numTasksToday, err := controller.taskService.NumTasksToday(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	if int(numTasksToday) >= user.MaxTodo {
		return c.JSON(http.StatusTooManyRequests, map[string]string{
			"error": "Limited create task for today",
		})
	}
	fmt.Println(numTasksToday, int(numTasksToday), user.MaxTodo)

	task, err := controller.taskService.AddTask(userID, addTaskDTO.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	remainTaskToday := user.MaxTodo - int(numTasksToday) - 1
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":            task,
		"remainTaskToday": remainTaskToday,
	})
}

func (controller *controller) RetrieveTasks(c echo.Context) error {

	userID := middleware.GetUserIDFromContext(c)
	createdDate := c.QueryParam("created_date")

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

func (controller *controller) AddManyTasks(c echo.Context) error {
	userID := middleware.GetUserIDFromContext(c)

	addTasksDTO := new(dto.AddTasksDTO)
	if err := c.Bind(addTasksDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	if err := c.Validate(addTasksDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	user, err := controller.userService.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	numTasksToday, err := controller.taskService.NumTasksToday(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	fmt.Println(numTasksToday, len(addTasksDTO.Contents), user.MaxTodo)
	if len(addTasksDTO.Contents)+int(numTasksToday) >= user.MaxTodo {
		return c.JSON(http.StatusTooManyRequests, map[string]string{
			"error": "Limited create task for today",
		})
	}

	tasks, err := controller.taskService.AddManyTasks(userID, addTasksDTO.Contents)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	remainTaskToday := user.MaxTodo - int(numTasksToday) - len(addTasksDTO.Contents)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":            tasks,
		"remainTaskToday": remainTaskToday,
	})
}
