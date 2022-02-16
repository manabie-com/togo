package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kier1021/togo/api/services"
)

type TaskController struct {
	taskSrv *services.TaskService
}

func NewTaskController(taskSrv *services.TaskService) *TaskController {
	return &TaskController{
		taskSrv: taskSrv,
	}
}

func (ctrl *TaskController) CreateTasks() gin.HandlerFunc {
	return func(c *gin.Context) {

		results, err := ctrl.taskSrv.CreateTasks()
		if err != nil {
			c.AbortWithStatusJSON(500, map[string]interface{}{
				"message": "Internal server error occured.",
				"error":   err.Error(),
			})
		}

		c.JSON(201, map[string]interface{}{
			"message": "Tasks successfully created",
			"data":    results,
		})
	}
}
