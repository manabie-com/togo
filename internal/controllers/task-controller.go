package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/models"
	"time"
)

func (c *Controller) Login(ctx *gin.Context) {
	id := ctx.PostForm("user_id")
	password := ctx.PostForm("password")
	existUser := c.taskService.GetAuthToken(ctx, id, password)
	if !existUser {
		c.ErrorResponse(errors.New("incorrecr usr/pwd"), ctx)
	} else {
		token, err := c.utils.CreateToken(id)
		if err != nil {
			c.ErrorResponse(err, ctx)
		} else {
			c.Response(token, ctx)
		}
	}
}

func (c *Controller) ListTask(ctx *gin.Context) {
	userId, ok := c.utils.UserIDFromCtx(ctx)
	createdDate := ctx.PostForm("created_date")
	if !ok {
		c.ErrorResponse(errors.New("token is required"), ctx)
	} else {
		tasks, err := c.taskService.ListTasks(ctx, userId, createdDate)
		if err != nil {
			c.ErrorResponse(err, ctx)
		} else {
			c.Response(tasks, ctx)
		}
	}
}

func (c *Controller) AddTask(ctx *gin.Context) {
	userID, ok := c.utils.UserIDFromCtx(ctx)
	if !ok {
		c.ErrorResponse(errors.New("token is required"), ctx)
	} else {
		t := &models.Task{}
		ctx.BindJSON(t)
		now := time.Now()
		t.ID = uuid.New().String()
		t.UserID = userID
		t.CreatedDate = now.Format("2006-01-02")
		err := c.taskService.AddTask(ctx, t)
		if err != nil {
			c.ErrorResponse(err, ctx)
		} else {
			c.Response(t, ctx)
		}
	}
}
