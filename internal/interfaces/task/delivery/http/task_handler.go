package http

import (
	"net/http"

	"github.com/datshiro/togo-manabie/internal/infras/errors"
	"github.com/datshiro/togo-manabie/internal/interfaces/domain"
	"github.com/datshiro/togo-manabie/internal/interfaces/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type TaskHandler struct {
	TUseCase       domain.TaskUseCase
	UUseCase       domain.UserUseCase
	NewPostRequest func(c echo.Context) PostRequest
}

func (h TaskHandler) HandlePost(c echo.Context) error {
	request := h.NewPostRequest(c)
	if err := request.Bind(); err != nil {
		return errors.InvalidRequest
	}
	if err := request.Validate(); err != nil {
		return err
	}

	ctx := c.Request().Context()
	user, err := h.UUseCase.GetUser(ctx, request.GetUserID())
	if err != nil {
		return err
	}

	task := &models.Task{
		UserID:      user.ID,
		Title:       request.GetTitle(),
		Description: null.StringFrom(request.GetDescription()),
		Priority:    request.GetPriority(),
	}

	if err := h.TUseCase.CreateTask(ctx, task, user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, task)
}
