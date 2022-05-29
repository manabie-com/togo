package httpInterface

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"togo/domain/model"
	"togo/domain/service"
)

type taskController struct {
	svc service.TaskService
}

func NewTaskController(svc service.TaskService) *taskController {
	return &taskController{svc: svc}
}

type createTaskrequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func buildResponse(rc int, rd string) interface{} {
	return gin.H{
		"rc": rc,
		"rd": rd,
	}
}

func (this *taskController) Create(c *gin.Context) {
	var req createTaskrequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, buildResponse(1, err.Error()))
		return
	}
	u, ok := c.Get("userInfo")
	if !ok {
		c.JSON(http.StatusUnauthorized, buildResponse(1, "relogin"))
		return
	}
	userinfo, ok := u.(model.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, buildResponse(1, "relogin"))
		return
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	_, err := this.svc.CreateTask(ctx, &userinfo, &model.Task{
		Title:       req.Title,
		Description: req.Description,
		CreatedBy:   userinfo.Id,
	})
	if err != nil {
		c.JSON(http.StatusOK, buildResponse(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, buildResponse(0, "Create Task Successfully"))
	return
}
