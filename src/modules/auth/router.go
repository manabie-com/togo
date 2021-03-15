package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	users := app.Group("/auth")
	{
		users.POST("/login", Login)
	}
}
