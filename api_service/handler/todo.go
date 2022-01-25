package handler

import (
	"log"
	"net/http"
	"strconv"

	"api_service/connection"
	"api_service/proto"

	"github.com/gin-gonic/gin"
)

//CreateTodo ...
func CreateTodo(ctx *gin.Context) {

	var createTodoRequest proto.CreateTodoRequest
	if err := ctx.ShouldBindJSON(&createTodoRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sc := connection.DialToTodoServiceServer()
	response, err := sc.ClientTodoService.CreateTodo(ctx, &createTodoRequest)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

//GetTodo ...
func GetTodo(ctx *gin.Context) {

	var getTodoRequest proto.GetTodoRequest
	todoId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	getTodoRequest.TodoId = todoId

	sc := connection.DialToTodoServiceServer()
	response, err := sc.ClientTodoService.GetTodo(ctx, &getTodoRequest)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"title": response.Title, "description": response.Description, "Status": response.Status})
}

//UpdateTodo ...
func UpdateTodo(ctx *gin.Context) {

	var updateTodoRequest proto.UpdateTodoRequest
	todoId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	updateTodoRequest.TodoId = todoId
	if err := ctx.ShouldBindJSON(&updateTodoRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sc := connection.DialToTodoServiceServer()
	response, err := sc.ClientTodoService.UpdateTodo(ctx, &updateTodoRequest)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

//DeleteTodo ...
func DeleteTodo(ctx *gin.Context) {

	var deleteTodoRequest proto.DeleteTodoRequest
	todoId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	deleteTodoRequest.TodoId = todoId
	if err := ctx.ShouldBindJSON(&deleteTodoRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	sc := connection.DialToTodoServiceServer()
	response, err := sc.ClientTodoService.DeleteTodo(ctx, &deleteTodoRequest)
	if err != nil {
		log.Print(err)
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
