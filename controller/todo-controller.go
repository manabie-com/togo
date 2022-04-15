package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jackc/pgx/v4/pgxpool"
	"context"
	"github.com/qgdomingo/todo-app/model"
)

type TaskDB struct {
	DBPoolConn *pgxpool.Pool
}

func (db *TaskDB) GetTasks(c *gin.Context) {
	var allTasks []model.Task



	rows, err := db.DBPoolConn.Query(context.Background(), "SELECT id, title, description, username, create_date from tasks")
	
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{ 
			"message" : "Unable to fetch data from tasks table",
			"error"   : err.Error() })
		return
	}

	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Username, &task.CreateDate)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{ 
				"message" : "Error encountered when row data is being fetched",
				"error"   : err.Error() })
			return
		}
		allTasks = append(allTasks, task)
	}

	if rows.Err() != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message" : "Error encountered when accesssing query results",
			"error"   : rows.Err().Error() })
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
	
	if id != "" {
		var task model.Task

		err := db.DBPoolConn.QueryRow(context.Background(), "SELECT id, title, description, username, create_date from tasks where id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Username, &task.CreateDate)

		if err != nil && err.Error() == "no rows in result set" {
			c.IndentedJSON(http.StatusNotFound, gin.H{ "message" : "No data was found for the specified id" })
			return

		} else if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{ 
				"message" : "Unable to fetch data from tasks table",
				"error"   : err.Error() })
			return
			
		}

		c.IndentedJSON(http.StatusOK, task)

	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "No ID was entered on the request URL" })
	}

}

func (db *TaskDB) GetTaskByUser(c *gin.Context) {
	user := c.Param("user")
	
	if user != "" {
		var allTasks []model.Task

		rows, err := db.DBPoolConn.Query(context.Background(), "SELECT id, title, description, username, create_date from tasks where username = $1", user)
	
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{ 
				"message" : "Unable to fetch data from tasks table" ,
				"error"   : err.Error() })
			return
		}
	
		defer rows.Close()
	
		for rows.Next() {
			var task model.Task
			err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Username, &task.CreateDate)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{ 
					"message" : "Error encountered when row data is being fetched",
					"error"   : err.Error() })
				return
			}
			allTasks = append(allTasks, task)
		}
	
		if rows.Err() != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message" : "Error encountered when accesssing query results",
				"error"   : rows.Err().Error() })
			return
		}
	
		if len(allTasks) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{ "message" : "No tasks data were found for the specified username" })
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

	row, insertErr := db.DBPoolConn.Query(context.Background(), "INSERT INTO tasks (title, description, username) SELECT $1, $2, $3::VARCHAR FROM task_config WHERE name = 'task_limit' AND value::INTEGER > (SELECT COUNT(id) FROM tasks WHERE username = $3 AND create_date = current_date) RETURNING id", taskDetails.Title, taskDetails.Description, taskDetails.Username)

	if insertErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{ 
			"message" : "Unable to insert data into the tasks table",
			"error"   : insertErr.Error() })
		return
	}

	if row.Next() {
		c.IndentedJSON(http.StatusOK, gin.H{ "message" : "New task has been created successfully" })
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task has not been created due to the user reaching the new task limit today" })
	}

	row.Close()
}

func (db *TaskDB) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	
	if id != "" {
		var taskDetails model.TaskUserEnteredDetails

		err := c.ShouldBindJSON(&taskDetails)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message" : "Error on binding data from request",
				"error"   : err.Error() })
		}
	
		row, updateErr := db.DBPoolConn.Query(context.Background(), "UPDATE tasks SET title = $1, description = $2 WHERE username = $3 AND id = $4 RETURNING id", taskDetails.Title, taskDetails.Description, taskDetails.Username, id)

		if updateErr != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{ 
				"message" : "Unable to update data into the tasks table",
				"error"   : updateErr.Error() })
			return
		}
	
		if row.Next() {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task has been updated successfully" })
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task was not updated, task with the provided id and/or username is not found." })
		}

		row.Close()
	
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "No ID was entered on the request URL" })
	}

}

func (db *TaskDB) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	
	if id != "" {
	
		row, err := db.DBPoolConn.Query(context.Background(), "DELETE FROM tasks WHERE id = $1 RETURNING id", id)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{ 
				"message" : "Unable to delete data from the tasks table",
				"error"   : err.Error() })
			return
		}
	
		if row.Next() {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task has been removed successfully" })
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Task was not removed, task with the provided id is not found." })
		}
		
		row.Close()
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "No ID was entered on the request URL" })
	}

}