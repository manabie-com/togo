package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/togo/api/dto"
	"github.com/kier1021/togo/api/services"
)

type UserTaskController struct {
	taskSrv *services.UserTaskService
}

func NewUserTaskController(taskSrv *services.UserTaskService) *UserTaskController {
	return &UserTaskController{
		taskSrv: taskSrv,
	}
}

func (ctrl *UserTaskController) AddTaskToUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Bind the body param to CreateTaskDTO
		var createTaskDto dto.CreateTaskDTO
		if err := c.Bind(&createTaskDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Bad Request",
				"error":   err.Error(),
			})
			return
		}
		// Set the InsDay field to date today
		createTaskDto.InsDay = time.Now().Format("2006-01-02")

		// Request the task service
		results, err := ctrl.taskSrv.AddTaskToUser(createTaskDto)
		if err != nil {
			makeErrResponse(err, c)
			return
		}

		c.JSON(201, map[string]interface{}{
			"message": "User created successfully.",
			"data":    results,
		})
	}
}

func (ctrl *UserTaskController) GetTasksOfUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Set the GetTaskOfUserDTO from the query param
		getTaskDto := dto.GetTaskOfUserDTO{
			UserName: c.Query("user_name"),
			InsDay:   c.Query("ins_day"),
		}

		// Request the task service
		results, err := ctrl.taskSrv.GetTasksOfUser(getTaskDto)
		if err != nil {
			makeErrResponse(err, c)
			return
		}

		c.JSON(201, map[string]interface{}{
			"message": "User tasks fetched successfully.",
			"data":    results,
		})
	}
}
