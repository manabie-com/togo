package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	db "task-manage/internal/db/sqlc"
	"time"
)

type createTaskRequest struct {
	UserId int32  `json:"user_id"`
	Title  string `json:"title"`
}

func (server *Server) createTask(context *gin.Context) {
	var req createTaskRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	currentUser, err := server.queries.GetUserById(context, req.UserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	listTask, err := server.queries.SelectTaskByUserId(context, req.UserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	if int(currentUser.MaximumTaskInDay) <= len(listTask) {
		message := errors.New(fmt.Sprintf("This user %s has maximum task", currentUser.UserName))
		context.JSON(http.StatusBadRequest, errResponse(message))
		return
	}
	arg := db.CreateTaskParams{
		Title:     req.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    req.UserId,
	}
	tasks, err := server.queries.CreateTask(context, arg)
	if err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	context.JSON(http.StatusOK, tasks)
}
