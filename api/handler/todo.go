package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/api/model"
	"github.com/manabie-com/togo/internal/service"
)

// Handler godoc
// @Summary endpoint to create todo item
// @Description API to add an item to todo list
// @ID create-todo
// @Param Authorization header string true "Authorization"
// @Param body body model.Todo true "todo information"
// @Produce json
// @Success 200 {object} model.Todo
// @Router /v1/todo [POST]
func AddTodo(s service.TodoService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req model.Todo
		err := context.BindJSON(&req)
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Error Encountered",
			})
		}
		id, err := s.Add(req)
		if err != nil {
			context.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Error Encountered",
			})
		} else {
			req.ID = id
			context.JSON(http.StatusCreated, req)
		}
	}
}

// Handler godoc
// @Summary endpoint to get todo item
// @Description API to add an item to todo list
// @ID get-todo
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} model.Todo
// @Router /v1/todo [GET]
func GetTodo(s service.TodoService) gin.HandlerFunc {
	return func(context *gin.Context) {
		d, err := s.Get([]string{})
		if err != nil {
			context.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Error Encountered",
			})
		}
		context.JSON(http.StatusOK, d)
	}
}
