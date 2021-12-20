package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/api/model"
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
func Todo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusCreated, model.Todo{
			ID: 1, Title: "Sample TODO", Description: "A sample todo descriotion",
		})
	}
}
