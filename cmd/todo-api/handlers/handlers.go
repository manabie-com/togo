// Package handlers manages the different versions of the API.
package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manabie-com/togo/cmd/todo-api/handlers/v1"
	"github.com/manabie-com/togo/internal/web/auth"
	"go.uber.org/zap"
)

// AppConfig contains all the mandatory systems required by handlers.
type AppConfig struct {
	Log  *zap.SugaredLogger
	DB   *sqlx.DB
	Auth *auth.Auth
}

// NewApp constructs an Echo app with all application routes defined.
func NewApp(cfg AppConfig) *echo.Echo {
	app := echo.New()
	app.HideBanner = true
	app.HidePort = true

	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				cfg.Log.Errorw("request",
					"method", v.Method,
					"uri", v.URI,
					"status", v.Status,
					"error", v.Error,
				)
			} else {
				cfg.Log.Infow("request",
					"method", v.Method,
					"uri", v.URI,
					"status", v.Status,
				)
			}

			return nil
		},
	}))

	v1.Routes(app, v1.Config{
		Log:  cfg.Log,
		DB:   cfg.DB,
		Auth: cfg.Auth,
	})

	return app
}
