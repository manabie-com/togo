package router

import (
	"github.com/gin-gonic/gin"
	"pt.example/grcp-test/http/events"
)

func New() (r *gin.Engine) {
	r = gin.Default()

	r.POST("/task", events.CreateTodoTask)

	return
}
