package handler

import (
	"togo/constants"
	"togo/internal/common"
	"togo/internal/response"
	"togo/internal/task/dto"
	"togo/internal/task/service"
	"togo/pkg/logger"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(
	router *echo.Group,
	taskService service.TaskService,
) *TaskHandler {
	handler := &TaskHandler{taskService}
	router.GET("", handler.GetList)
	router.POST("", handler.Create)

	return handler
}

func (h *TaskHandler) Create(c echo.Context) error {
	userID := c.Get(string(constants.ContextUserID)).(int)
	createTaskDto := new(dto.CreateTaskDto)

	if errBindValidate := common.BindValidate(c, createTaskDto); errBindValidate != nil {
		//logger.L.Sugar().Errorf("[TaskHandler] Create errBindValidate: %s", errBindValidate)
		return response.Error(c, errBindValidate)
	}

	task, err := h.taskService.Create(createTaskDto, userID)
	if err != nil {
		//logger.L.Sugar().Errorf("[TaskHandler] Create errCreateTask: %s", err)
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
