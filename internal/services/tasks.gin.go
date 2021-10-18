package services

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jericogantuangco/togo/internal/storages/postgres"
	"github.com/jericogantuangco/togo/internal/token"
)

type listTasksRequest struct {
	CreatedDate string `form:"created_date" binding:"required"`
}

type addTaskRequest struct {
	Content string `json:"content" binding:"required"`
}

func (server *Server) listTasks(ctx *gin.Context) {
	var req listTasksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayload).(*token.Payload)
	
	arg := postgres.RetrieveTasksParams{
		UserID:      authPayload.Username,
		CreatedDate: req.CreatedDate,
	}
	
	tasks, err := server.Store.RetrieveTasks(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (server *Server) addTask(ctx *gin.Context) {
	var req addTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayload).(*token.Payload)
	currentTime := time.Now().Format("2006-01-02")

	arg := postgres.CreateTaskParams{
		UserID:      authPayload.Username,
		Content:     req.Content,
		CreatedDate: currentTime,
	}

	taskArg := postgres.RetrieveTasksParams{
		UserID:      authPayload.Username,
		CreatedDate: currentTime,
	}

	user, err := server.Store.RetrieveUser(ctx, arg.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	tasks, err := server.Store.RetrieveTasks(ctx, taskArg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	currentTodo := len(tasks)
	message := ResponseMessage{
		Message: "You are not allowed to create anymore tasks for the day.",
	}

	if currentTodo >= int(user.MaxTodo) {
		ctx.JSON(http.StatusMethodNotAllowed, message)
		return
	}
	task, err := server.Store.CreateTask(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, task)
}
