package router

import (
	"github.com/gin-gonic/gin"
	"togo/api"
)

func initAuthenRouter(Router *gin.RouterGroup) {
	AuthenRouter := Router.Group("auth")
	{
		AuthenRouter.POST("login", api.Login) //login
	}

}
