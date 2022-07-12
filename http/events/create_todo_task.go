package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pt.example/grcp-test/http/actions"
	"pt.example/grcp-test/http/utils"
)

type CreateTodoParam struct {
	UserId string `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

func CreateTodoTask(c *gin.Context) {
	var r interface{}
	var p CreateTodoParam

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok := actions.ReduceRemainedTodoCountOfUser(&p)

	if !ok {
		return
	}

	r, ok = actions.SaveTodoTask(&p)

	c.JSON(http.StatusOK, utils.SuccessResponse(r))
}

func (p *CreateTodoParam) GetUserId() (r string) {
	r = p.UserId
	return
}

func (p *CreateTodoParam) GetTitle() (r string) {
	r = p.Title
	return
}

func (p *CreateTodoParam) GetTaskSavedCount() (r int) {
	i := 1
	r = i

	return
}
