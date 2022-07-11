package events

import (
	"github.com/gin-gonic/gin"
	"pt.example/grcp-test/http/actions"
	"pt.example/grcp-test/http/utils"
)

type CreateTodoParam struct {
}

func CreateTodoTask(c *gin.Context) {
	p := CreateTodoParam{}

	result := actions.SaveTodoTask(p)

	c.JSON(200, utils.SuccessResponse(result))
}
