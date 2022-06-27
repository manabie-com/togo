package transport

import (
	"github.com/gin-gonic/gin"
	"manabieAssignment/internal/core/usecase"
	"manabieAssignment/internal/todo/transport/model"
	"net/http"
)

type TodoHandler struct {
	todoUC usecase.TodoUseCase
}

func NewTodoHandler(engine *gin.Engine, todoUC usecase.TodoUseCase) {
	handler := TodoHandler{
		todoUC: todoUC,
	}
	engine.POST("/todo", handler.CreateTodo)
}

func (t *TodoHandler) CreateTodo(c *gin.Context) {
	var todo model.TodoModel
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = t.todoUC.CreateTodo(todo.ToDomainModel())
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Success")
}
