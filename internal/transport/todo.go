package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginParams struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) login(ctx *gin.Context) {
	var params loginParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	token, err := s.todo.GetToken(params.Id, params.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(token))
}

type createTaskParams struct {
	Content string `json:"content" binding:"required"`
}

func (s *Server) addTask(ctx *gin.Context) {
	var params createTaskParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	id := ctx.MustGet(authorizationPayloadKey).(string)
	err = s.todo.AddTask(params.Content, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("OK"))
}

type listTaskParams struct {
	CreatedDate string `uri:"created_date" binding:"required"`
	Total       int    `uri:"total" binding:"min=1"`
	Page        int    `uri:"page" binding:"min=1"`
}

func (s *Server) listTasks(ctx *gin.Context) {
	var params listTaskParams
	err := ctx.ShouldBindUri(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userId := ctx.MustGet(authorizationPayloadKey).(string)
	tasks, err := s.todo.ListTask(params.CreatedDate, userId, params.Total, params.Page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse(tasks))
}
