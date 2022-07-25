package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/trangmaiq/togo/internal/model"
	"github.com/trangmaiq/togo/pkg/uuidx"
)

type (
	CreateTaskRequest struct {
		UserID uuidx.UUID `json:"user_id,omitempty"`
		Title  string     `json:"title,omitempty"`
		Note   string     `json:"note,omitempty"`
	}
	RegularErrorResponse struct {
		Error            string `json:"message,omitempty"`
		ErrorDescription string `json:"error_description"`
	}
	CreateTaskResponse struct {
		ID uuidx.UUID `json:"id,omitempty"`
	}
)

func (r *CreateTaskRequest) validate() error {
	if r == nil {
		return ErrEmptyRequest
	}

	if r.UserID.Validate() != nil {
		return ErrInvalidUserID
	}

	if r.Title == "" {
		return ErrEmptyTitle
	}

	return nil
}

func (h *Handler) CreateTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			now   = time.Now()
			input CreateTaskRequest
		)

		err := c.Bind(&input)
		if err != nil {
			// TODO: Handle error send JSON response error
			_ = c.JSON(http.StatusBadRequest, RegularErrorResponse{
				Error:            "bind request failed",
				ErrorDescription: err.Error(),
			})
			return err
		}

		err = input.validate()
		if err != nil {
			_ = c.JSON(http.StatusBadRequest, RegularErrorResponse{
				Error:            "invalid request",
				ErrorDescription: err.Error(),
			})
			return fmt.Errorf("invalid request: %w", err)
		}

		if !h.d.RateLimiter().AllowN(now, input.UserID.String(), 1) {
			_ = c.JSON(http.StatusTooManyRequests, RegularErrorResponse{
				Error:            "too many requests",
				ErrorDescription: "a lot of requests are made in a short period of time",
			})

			return ErrTooManyRequests
		}

		id := uuid.NewString()
		err = h.d.Persister().CreateTask(
			c.Request().Context(),
			&model.Task{
				ID:        id,
				UserID:    input.UserID.String(),
				Title:     input.Title,
				Note:      input.Note,
				Status:    model.StatusInProgress,
				CreatedAt: now,
				UpdatedAt: now,
			})
		if err != nil {
			_ = c.JSON(http.StatusInternalServerError, RegularErrorResponse{
				Error:            "create task failed",
				ErrorDescription: "cannot store task to persistence storage",
			})
			return fmt.Errorf("create task failed: %w", err)
		}

		_ = c.JSON(http.StatusOK, CreateTaskResponse{ID: uuidx.UUID(id)})
		return nil
	}
}
