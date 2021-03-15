package tasks

import (
	AuthMiddleware "togo/src/middleware/auth"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	users := app.Group("/tasks")
	{
		users.POST("/", AuthMiddleware.Authorized, Create)
	}
}
