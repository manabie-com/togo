package routers

import (
	"manabie/manabie/controllers"

	"github.com/gin-gonic/gin"
)

func UserAPIRoute(r *gin.RouterGroup) {
	r.GET("/login", controllers.Login)
}
