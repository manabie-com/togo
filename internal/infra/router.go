package infra

import (
	"net/http"

	read_side "github.com/manabie-com/togo/internal/services/read-side"

	user_tasks "github.com/manabie-com/togo/internal/services/user-tasks"

	"github.com/manabie-com/togo/internal/services/auth"

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func ProvideRestAPIHandler(
	authSrv auth.Service,
	userTaskSrv user_tasks.Service,
	readRepo read_side.ReadRepo,
) RestAPIHandler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(ginlogrus.Logger(logrus.StandardLogger()))
	router.Use(healthcheck.Default())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	auth.Handler(router, authSrv)

	authorizedV1 := router.Group("/v1")
	{
		authorizedV1.Use(auth.Middleware(authSrv))

		user_tasks.Handler(authorizedV1, userTaskSrv)
		read_side.Handler(authorizedV1, readRepo)
	}

	return router
}
