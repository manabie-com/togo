package user

import (
	"github.com/gin-gonic/gin"
	"manabie-com/togo/entity"
	"manabie-com/togo/query"
	"manabie-com/togo/util"
	"net/http"
)

func ListTasks(c *gin.Context) {
	id := util.TokenUserID(c)

	task, err := query.TaskByUserID(id)
	if err != nil {
		task = []entity.Task{}
	}
	c.JSON(http.StatusOK, task)
}