package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/khangjig/togo/delivery/http/auth"
	"github.com/khangjig/togo/delivery/http/healthcheck"
	"github.com/khangjig/togo/delivery/http/user"
	md "github.com/khangjig/togo/middleware"
	"github.com/khangjig/togo/repository"
	"github.com/khangjig/togo/usecase"
)

func NewHTTPHandler(useCase *usecase.UseCase, repo *repository.Repository) *echo.Echo {
	var (
		e         = echo.New()
		loggerCfg = middleware.DefaultLoggerConfig
	)

	loggerCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.Path == "/health-check"
	}

	e.Use(middleware.LoggerWithConfig(loggerCfg))
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"}, // Allow any request
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))

	// Health check
	healthcheck.Init(e.Group("/health-check"))

	// public APIs
	publicAPI := e.Group("/auth")
	auth.Init(publicAPI.Group(""), useCase)

	// APIs
	api := e.Group("/api")
	api.Use(md.Auth(repo))
	user.Init(api.Group("/users"), useCase)

	return e
}
