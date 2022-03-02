package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/src/service"
)

/*
TaskController is the end of controller of task in user
 */

func (c *Controller) TaskController(g *gin.RouterGroup) {
	// all method of task will be implemented here
	g.Group(Task).POST(NewTask, c.CheckLimit, service.NewTask)
}

func(c *Controller) CheckLimit(g *gin.Context) {
	err, data := service.JsonBody(g)
	if err != nil {
		g.JSON(http.StatusServiceUnavailable, gin.H{"message": err})
		g.Abort()
		return
	}
	currentTaskCount := int(data["user_id"].(float64))
	// the limit default is 10, this number will be changed after we have an API create user
	//and set the limit for that.
	if c.Frequency[currentTaskCount] > 10 {
		g.JSON(http.StatusServiceUnavailable, gin.H{"message": "out of limit task"})
		g.Abort()
		return
	}
	c.Frequency[currentTaskCount] += 1
	g.Next()
}