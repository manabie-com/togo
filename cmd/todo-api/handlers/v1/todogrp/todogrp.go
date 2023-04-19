// Package todogrp maintains the group of handlers for todo access.
package todogrp

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	createtodo "github.com/manabie-com/togo/internal/features/create_todo"
	"github.com/manabie-com/togo/internal/web/auth"
	"github.com/manabie-com/togo/platform/validate"
)

// Handlers manages the set of todo endpoints.
type Handlers struct {
	Auth              *auth.Auth
	CreateTodoFeature *createtodo.Feature
}

// Create adds a new user to the system.
func (h *Handlers) Create(c echo.Context) error {
	var payload AppNewTodo
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).WithInternal(err)
	}

	if err := validate.Check(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err).WithInternal(err)
	}

	nt := toFeatureNewTodo(payload)

	if userID, err := h.Auth.GetUserID(c); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized).WithInternal(err)
	} else {
		nt.UserID = userID
	}

	fmt.Println(nt)

	ctx := c.Request().Context()

	todo, err := h.CreateTodoFeature.CreateTodo(ctx, nt)
	if err != nil {
		if errors.Is(err, createtodo.ErrExceededDailyMaximumTodos) {
			return echo.NewHTTPError(http.StatusConflict, createtodo.ErrExceededDailyMaximumTodos.Error()).WithInternal(err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError).WithInternal(err)
	}

	return c.JSON(http.StatusCreated, todo)
}
