package controller

import (
	"github.com/gin-gonic/gin"
)

/*
UserController will have many children controllers such as task controller
*/

func (c *Controller) UserController(g *gin.Engine) {
	// all children controllers will be implemented here
	c.TaskController(g.Group(User))
}
