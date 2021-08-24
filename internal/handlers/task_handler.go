package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/constants"
	"github.com/manabie-com/togo/internal/dtos"
	"github.com/manabie-com/togo/internal/helpers"
	"github.com/manabie-com/togo/internal/services"
	"github.com/sirupsen/logrus"
	"net/http"
)

type TaskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(injectedTaskService services.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: injectedTaskService,
	}
}

// GetListTask godoc
// @Summary Get List Tasks
// @Description Get List Tasks
// @Param created_date query string true "Created Date"
// @Tags Task
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.GetListTaskResponse
// @Router /tasks [get]
func (h *TaskHandler) GetListTask(ctx *gin.Context) {
	createdDate := ctx.Query("created_date")
	userID, ok := helpers.GetUserIdFromContext(ctx)
	if !ok {
		logrus.Errorf("Get List Task Get User from context failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	response, err := h.taskService.GetListTask(ctx, userID, createdDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// CreateTask godoc
// @Summary Create New Task
// @Description Get New Task
// @Param CreateTaskRequest body dtos.CreateTaskRequest true "Information of CreateTaskRequest"
// @Tags Task
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.CreateTaskResponse
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	var request = &dtos.CreateTaskRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewError(err))
		return
	}

	userID, ok := helpers.GetUserIdFromContext(ctx)
	if !ok {
		logrus.Errorf("Create Task Get User from context failed")
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	request.UserID = userID
	response, err := h.taskService.CreateTask(ctx, request)
	if errors.Is(err, constants.ErrMaximumCreatedTask) {
		ctx.JSON(http.StatusConflict, dtos.NewError(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewError(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
