package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"pt.example/grcp-test/http/actions"
	"pt.example/grcp-test/http/utils"
)

type CreateTodoParam struct {
	AssigneeEmail string `json:"assignee_email" binding:"required"`
	Title         string `json:"title" binding:"required"`
}

func CreateTodoTask(c *gin.Context) {
	var p CreateTodoParam

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := actions.ReduceRemainedTodoCountOfUser(c, &p); err != nil {

		switch err.Error() {
		case mongo.ErrNoDocuments.Error():
			c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found"))
		default:
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		}

		return
	}

	if _, err := actions.SaveTodoTask(c, &p); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse("Added"))
}

func (p *CreateTodoParam) GetAssigneeEmail() (r string) {
	r = p.AssigneeEmail
	return
}

func (p *CreateTodoParam) GetTitle() (r string) {
	r = p.Title
	return
}

func (p *CreateTodoParam) GetTaskSavedCount() (r uint8) {
	var i uint8 = 1
	r = i

	return
}
