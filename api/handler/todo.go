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
// @Param body body model.TodoRequest true "todo information"
// @Produce json
// @Success 200 {object} model.Todo
// @Router /v1/todo [POST]
func AddTodo(s service.TodoService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req model.TodoRequest
		err := context.BindJSON(&req)
		if err != nil {
			log.Println(err)
			context.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Error Encountered",
			})
		}
		res, err := s.Add(model.Todo{
			Title:       req.Title,
			Description: req.Description,
		})
		if err != nil {
			context.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Error Encountered",
			})
		} else {
			context.JSON(http.StatusCreated, res)
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
