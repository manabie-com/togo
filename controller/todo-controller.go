package controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/qgdomingo/todo-app/model"
)

type TaskDB struct {
	DBPoolConn *pgxpool.Pool
}

func (db *TaskDB) GetTasks(c *gin.Context) {
	allTasks, errMessage := model.GetTasksDB(db.DBPoolConn, nil)
	
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

func (db *TaskDB) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	
	if id, err := strconv.Atoi(id); err == nil  {
		allTasks, errMessage := model.GetTasksDB(db.DBPoolConn, id)

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

func (db *TaskDB) GetTaskByUser(c *gin.Context) {
	user := c.Param("user")
	
	if user != "" {
		allTasks, errMessage := model.GetTasksDB(db.DBPoolConn, user)

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

func (db *TaskDB) CreateTask(c *gin.Context) {
	var taskDetails model.TaskUserEnteredDetails

	err := c.ShouldBindJSON(&taskDetails)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "Error on binding data from request",
			"error"   : err.Error() })
	}

	isTaskCreated, errMessage := model.InsertTaskDB(db.DBPoolConn, &taskDetails)

	if errMessage != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	if isTaskCreated {
		c.IndentedJSON(http.StatusOK, gin.H{ "message" : "New task has been created successfully" })
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task has not been created due to the user reaching the new task limit today" })
	}
}

func (db *TaskDB) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	
	if id, err := strconv.Atoi(id); err == nil {
		var taskDetails model.TaskUserEnteredDetails

		err := c.ShouldBindJSON(&taskDetails)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message" : "Error on binding data from request",
				"error"   : err.Error() })
		}
	
		isTaskUpdated, errMessage := model.UpdateTaskDB(db.DBPoolConn, &taskDetails, id)

		if errMessage != nil {
			c.IndentedJSON(http.StatusInternalServerError, errMessage)
			return
		}
	
		if isTaskUpdated {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task has been updated successfully" })
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task was not updated, task with the provided id and/or username is not found." })
		}
	
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Invalid ID was entered on the request URL" })
	}

}

func (db *TaskDB) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	
	if id, err := strconv.Atoi(id); err == nil {
	
		isTaskDeleted, errMessage := model.DeleteTaskDB(db.DBPoolConn, id)

		if errMessage != nil {
			c.IndentedJSON(http.StatusInternalServerError, errMessage)
			return
		}
	
		if isTaskDeleted {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task has been removed successfully" })
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task was not removed, task with the provided id is not found." })
		}
		
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Invalid ID was entered on the request URL" })
	}

}