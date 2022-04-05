package api

import (
	"net/http"
	db "togo/db/sqlc"
	"togo/token"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type listTasksRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=20"`
}

func (server *Server) listTasks(ctx *gin.Context) {
	var req listTasksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var tasks []db.Task
	var err error
	if authPayload.Username == "admin" {
		arg := db.ListTasksByOwnerParams{
			Limit:  req.PageSize,
			Offset: (req.PageID - 1) * req.PageSize,
			Owner:  authPayload.Username,
		}
		tasks, err = server.store.ListTasksByOwner(ctx, arg)
	} else {
		arg := db.ListTasksByOwnerParams{
			Limit:  req.PageSize,
			Offset: (req.PageID - 1) * req.PageSize,
			Owner:  authPayload.Username,
		}
		tasks, err = server.store.ListTasksByOwner(ctx, arg)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

type createTaskRequest struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (server *Server) createTask(ctx *gin.Context) {
	var req createTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateTaskTxParams{
		User:    user,
		Name:    req.Name,
		Content: req.Content,
	}
	result, err := server.store.CreateTaskTx(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// func (server *Server) editTask(ctx *gin.Context) {}

// func (server *Server) deleteTask(ctx *gin.Context) {}
