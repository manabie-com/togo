package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter(env string) *gin.Engine {
	var Router = gin.Default()
	apiGroup := Router.Group(fmt.Sprintf("%s/api", env))
	initAuthenRouter(apiGroup)
	initTaskRouter(apiGroup)

	return Router
}
