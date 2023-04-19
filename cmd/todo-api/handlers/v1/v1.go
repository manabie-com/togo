// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/cmd/todo-api/handlers/v1/todogrp"
	"github.com/manabie-com/togo/cmd/todo-api/handlers/v1/usergrp"
	"github.com/manabie-com/togo/internal/features/authenticate"
	authenticateStore "github.com/manabie-com/togo/internal/features/authenticate/store"
	createtodo "github.com/manabie-com/togo/internal/features/create_todo"
	createtodoStore "github.com/manabie-com/togo/internal/features/create_todo/store"
	"github.com/manabie-com/togo/internal/web/auth"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log  *zap.SugaredLogger
	DB   *sqlx.DB
	Auth *auth.Auth
}

// Routes binds all the version 1 routes.
func Routes(app *echo.Echo, cfg Config) {
	v1Grp := app.Group("v1")

	createtodoFeature := createtodo.NewFeature(createtodoStore.NewStore(cfg.Log, cfg.DB))
	authenticateFeature := authenticate.NewFeature(authenticateStore.NewStore(cfg.Log, cfg.DB))

	authen := cfg.Auth.EchoMiddleware()

	usergrph := usergrp.Handlers{
		Auth:                cfg.Auth,
		AuthenticateFeature: authenticateFeature,
	}
	v1Grp.POST("/users/auth", usergrph.Authenticate)

	todogrph := todogrp.Handlers{
		Auth:              cfg.Auth,
		CreateTodoFeature: createtodoFeature,
	}
	v1Grp.POST("/todos", todogrph.Create, authen)
}
