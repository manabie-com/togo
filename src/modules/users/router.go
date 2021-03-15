package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {
	users := app.Group("/users")
	{
		users.POST("/", Create)
	}
}
