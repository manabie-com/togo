package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmramos02/akaru/internal/database"
	"github.com/jmramos02/akaru/internal/user"
)

type TaskRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

//piece it up all together
func CreateTask(c *gin.Context) {
	var request TaskRequest
	c.Bind(&request)
	db := database.Init()

	ctx := context.TODO()
	ctx = context.WithValue(ctx, "db", db)
	ctx = context.WithValue(ctx, "username", "username")

	//validation first
	userService := user.Initialize(ctx)
	canUserInsert := userService.CanUserInsert()

	if canUserInsert {
		//insert user

		c.JSON(200, map[string]interface{}{
			"success": true,
		})

		return
	}

	c.JSON(400, map[string]interface{}{
		"success": false,
		"error":   "User has reached max limit for the day",
	})
}
