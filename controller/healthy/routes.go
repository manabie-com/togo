package healthy

import (
	"github.com/gin-gonic/gin"
	"manabie-com/togo/global"
)

func RegisterRoutes(route *gin.Engine) {
	healthyGroup := route.Group(global.Config.Prefix + "/v1")
	{
		healthyGroup.GET("/status", Healthy)
		healthyGroup.GET("/info", Info)
	}
}
