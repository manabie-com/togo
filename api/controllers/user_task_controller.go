package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kier1021/togo/api/apierrors.go"
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

func (ctrl *UserTaskController) CreateUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Bind the body param to CreateUserDTO
		var createUserDto dto.CreateUserDTO
		if err := c.Bind(&createUserDto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Bad Request",
				"error":   err.Error(),
			})
			return
		}

		// Request the task service
		results, err := ctrl.taskSrv.CreateUser(createUserDto)
		if err != nil {

			// Check if the error is UserAlreadyExists
			if errors.Is(err, apierrors.UserAlreadyExists) {
				c.AbortWithStatusJSON(422, map[string]interface{}{
					"message": "Error in data input",
					"error":   err.Error(),
				})
				return
			}

			c.AbortWithStatusJSON(500, map[string]interface{}{
				"message": "Internal server error occurred.",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(201, map[string]interface{}{
			"message": "User created successfully.",
			"data":    results,
		})
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

			// Check if the error is UserDoesNotExists or MaxTasksReached
			if errors.Is(err, apierrors.UserDoesNotExists) || errors.Is(err, apierrors.MaxTasksReached) {
				c.AbortWithStatusJSON(422, map[string]interface{}{
					"message": "Error in data input",
					"error":   err.Error(),
				})
				return
			}

			c.AbortWithStatusJSON(500, map[string]interface{}{
				"message": "Internal server error occurred.",
				"error":   err.Error(),
			})
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

			// Check if the error is UserDoesNotExists
			if errors.Is(err, apierrors.UserDoesNotExists) {
				c.AbortWithStatusJSON(422, map[string]interface{}{
					"message": "Error in data input",
					"error":   err.Error(),
				})
				return
			}

			c.AbortWithStatusJSON(500, map[string]interface{}{
				"message": "Internal server error occurred.",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(201, map[string]interface{}{
			"message": "User tasks fetched successfully.",
			"data":    results,
		})
	}
}
