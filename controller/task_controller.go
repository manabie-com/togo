package controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/qgdomingo/todo-app/interfaces"
	"github.com/qgdomingo/todo-app/model"
)

type TaskController struct {
	TaskRepo interfaces.ITaskRepository
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	allTasks, errMessage := tc.TaskRepo.GetTasksDB(nil)
	
	if errMessage != nil {
		c.IndentedJSON(http.StatusInternalServerError, errMessage)
		return
	}

	if len(allTasks) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{ "message" : "No tasks data were found" })
	} else {
		c.IndentedJSON(http.StatusOK, allTasks)
	}
}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	
	if id, err := strconv.Atoi(id); err == nil  {
		allTasks, errMessage := tc.TaskRepo.GetTasksDB(id)

		if errMessage != nil {
			c.IndentedJSON(http.StatusInternalServerError, errMessage)
			return
		}

		if len(allTasks) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{ "message" : "No data was found for the specified id" })
		} else {
			c.IndentedJSON(http.StatusOK, allTasks[0])
		}

	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Invalid ID was entered on the request URL" })
	}
}

func (tc *TaskController) GetTaskByUser(c *gin.Context) {
	user := c.Param("user")
	
	if user != "" {
		allTasks, errMessage := tc.TaskRepo.GetTasksDB(user)

		if errMessage != nil {
			c.IndentedJSON(http.StatusInternalServerError, errMessage)
			return
		}

		if len(allTasks) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{ "message" : "No data was found for the specified username" })
		} else {
			c.IndentedJSON(http.StatusOK, allTasks)
		}

	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "No username was entered on the request URL" })
	}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var taskDetails model.TaskUserEnteredDetails

	err := c.ShouldBindJSON(&taskDetails)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "Error on binding data from request",
			"error"   : err.Error() })
		return
	}

	if taskDetails.Title != "" && taskDetails.Description != "" && taskDetails.Username != "" {
		isTaskCreated, errMessage := tc.TaskRepo.InsertTaskDB(&taskDetails)

		if errMessage != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
	
		if isTaskCreated {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "New task has been created successfully" })
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task has not been created due to the user reaching the new task limit today" })
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Either one or more of the required data on the request is empty" })
	}

}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	
	if id, err := strconv.Atoi(id); err == nil {
		var taskDetails model.TaskUserEnteredDetails

		err := c.ShouldBindJSON(&taskDetails)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message" : "Error on binding data from request",
				"error"   : err.Error() })
			return
		}
		
		if taskDetails.Title != "" && taskDetails.Description != "" && taskDetails.Username != "" {
			isTaskUpdated, errMessage := tc.TaskRepo.UpdateTaskDB(&taskDetails, id)

			if errMessage != nil {
				c.IndentedJSON(http.StatusInternalServerError, errMessage)
				return
			}
		
			if isTaskUpdated {
				c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task has been updated successfully" })
			} else {
				c.IndentedJSON(http.StatusNotFound, gin.H{ "message" : "Task was not updated, task with the provided id and/or username is not found." })
			}
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Either one or more of the required data on the request is empty" })
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Invalid ID was entered on the request URL" })
	}

}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	
	if id, err := strconv.Atoi(id); err == nil {
	
		isTaskDeleted, errMessage := tc.TaskRepo.DeleteTaskDB(id)

		if errMessage != nil {
			c.IndentedJSON(http.StatusInternalServerError, errMessage)
			return
		}
	
		if isTaskDeleted {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task has been removed successfully" })
		} else {
			c.IndentedJSON(http.StatusNotFound, gin.H{ "message" : "Task was not removed, task with the provided id is not found." })
		}
		
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Invalid ID was entered on the request URL" })
	}

}