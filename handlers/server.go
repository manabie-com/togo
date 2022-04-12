package handlers

import "github.com/gin-gonic/gin"

type Handlers struct {
	Task *TaskHandler
}

// SetupRouter provides endpoint's handler functions
func (h *Handlers) SetupRouter(ginEngine *gin.Engine) {
	apiGroup := ginEngine.Group("/api")
	{
		// Tasks route
		apiGroup.PUT("/tasks", h.Task.UpdateUserTask)
	}
}
