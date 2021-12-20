package route

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/api/handler"
)

func RegisterTodo(router *gin.RouterGroup) {
	router.POST(todo, handler.Todo())
}
