package handler

import (
	"togo/constants"
	"togo/internal/dto"
	"togo/internal/response"
	"togo/internal/service"
	"togo/pkg/logger"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(
	router *echo.Group,
	taskService *service.TaskService,
) {
	handler := &TaskHandler{taskService}
	router.GET("", handler.GetList)
	router.POST("", handler.Create)
}

func (h *TaskHandler) Create(c echo.Context) error {
	userID := c.Get(string(constants.ContextUserID)).(int)
	createTaskDto := new(dto.CreateTaskDto)

	if errBindDto := c.Bind(createTaskDto); errBindDto != nil {
		logger.L.Sugar().Errorf("[TaskHandler] Create errBindDto: %s", errBindDto)
		return response.Error(c, errBindDto)
	}

	task, err := h.taskService.Create(createTaskDto, userID)
	if err != nil {
		logger.L.Sugar().Errorf("[TaskHandler] Create errCreateTask: %s", err)
		return response.Error(c, err)
	}
	return response.Success(c, task)
}

func (h *TaskHandler) GetList(c echo.Context) error {
	userID := c.Get(string(constants.ContextUserID)).(int)

	tasks, err := h.taskService.GetListByUserID(userID)
	if err != nil {
		logger.L.Sugar().Errorf("[TaskHandler] GetList errGetListByUserID: %s", err)
		return response.Error(c, err)
	}

	return response.Success(c, tasks)
}
