package router

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/handler/users"
)

// Init init router
func Init() *gin.Engine {
	r := gin.Default()

	// any custom for router

	userHandler := users.UserHandler{}
	r.POST("/users/:id/tasks", userHandler.AssignTasks)

	return r
}
