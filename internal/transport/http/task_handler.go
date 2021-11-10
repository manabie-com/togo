package http

import (
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/helper"
	"github.com/manabie-com/togo/internal/transport/http/middleware"
	"github.com/manabie-com/togo/internal/transport/http/request"
	"github.com/manabie-com/togo/internal/transport/http/response"
	"github.com/manabie-com/togo/pkg/token"
	"net/http"
)

// TaskHandler represent task http handler
type TaskHandler struct {
	TaskUseCase domain.TaskUseCase
}

// NewTaskHandler initialize new tasks endpoint
func NewTaskHandler(e *echo.Echo, taskUseCase domain.TaskUseCase, middleware middleware.AuthMiddleWare) {
	taskHandler := &TaskHandler{
		TaskUseCase: taskUseCase,
	}
	group := e.Group("", middleware.UseAuthMiddleWare)
	group.POST("/tasks", taskHandler.CreateTask)
	group.GET("/tasks", taskHandler.GetTasks)
}

// CreateTask create new task by a user
func (t *TaskHandler) CreateTask(ctx echo.Context) error {
	createTaskRequest := new(request.CreateTaskRequest)
	if err := ctx.Bind(createTaskRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
	}
	if createTaskRequest.Content == "" {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "content must be not null or empty",
		})
	}
	username := ctx.Get(middleware.AuthorizationPayload).(*token.Payload).Username
	err := t.TaskUseCase.CreateTask(ctx.Request().Context(), createTaskRequest.Content, username)
	if err != nil {
		return ctx.JSON(response.GetStatusCode(err), response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse{
		Message: "create task success",
	})
}

// GetTasks get task by username from token and created date
func (t *TaskHandler) GetTasks(ctx echo.Context) error {
	createdDate := ctx.QueryParam("created_date")
	if createdDate == "" || !helper.DateValidator(createdDate) {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "date must not be null or wrong format yyyy-mm-dddd",
		})
	}
	username := ctx.Get(middleware.AuthorizationPayload).(*token.Payload).Username
	tasks, err := t.TaskUseCase.GetTask(ctx.Request().Context(), username, createdDate)
	if err != nil {
		return ctx.JSON(response.GetStatusCode(err), response.ErrorResponse{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, tasks)

}
