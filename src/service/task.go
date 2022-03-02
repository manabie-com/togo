package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/src/utils"
)

/*
All services of task will be represent here
*/

func NewTask(g *gin.Context) {
	err, data := JsonBody(g)
	if err != nil {
		g.JSON(http.StatusServiceUnavailable, gin.H{"message": err})
		return
	}

	//	Check tasks and the relationships existed or not, it will be 3 cases:
	//add both, add relations or skip. These cases will be implemented in real project.
	//Currently, the application will save the relationships of users and tasks in a file
	//without checking duplicated tasks or exist users
	utils.WriteFile("task-user.txt", fmt.Sprintf("%v - %v", data["user_id"], data["task_name"]))
	g.JSON(http.StatusOK, gin.H{"message": "add new task success"})
}
