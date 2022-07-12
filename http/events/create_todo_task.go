package events

import (
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

	p := &CreateTodoParam{}

	ok := actions.ReduceRemainedTodoCountOfUser(p)

	if !ok {
		return
	}

	r, ok = actions.SaveTodoTask(p)

	c.JSON(200, utils.SuccessResponse(r))
}

func (p *CreateTodoParam) GetUserId() (r *string) {
	r = &p.UserId
	return
}

func (p *CreateTodoParam) GetTitle() (r *string) {
	r = &p.Title
	return
}

func (p *CreateTodoParam) GetTaskSavedCount() (r *int) {
	i := 1
	r = &i

	return
}
