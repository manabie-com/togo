package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmramos02/akaru/internal/database"
	"github.com/jmramos02/akaru/internal/task"
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
	ctx = context.WithValue(ctx, "username", request.Username)

	//validation first
	userService := user.Initialize(ctx)
	canUserInsert := userService.CanUserInsert()

	//get user id
	id := userService.GetUserID()
	ctx = context.WithValue(ctx, "userID", id)

	if canUserInsert {
		//insert task

		taskService := task.Initalize(ctx)
		task, err := taskService.CreateTask(request.Name)

		if err != nil {
			c.JSON(500, map[string]interface{}{
				"error": "Something went wrong",
			})
		}

		c.JSON(200, map[string]interface{}{
			"task": task,
		})

		return
	}

	c.JSON(400, map[string]interface{}{
		"error": "User has reached max limit for the day",
	})
}
