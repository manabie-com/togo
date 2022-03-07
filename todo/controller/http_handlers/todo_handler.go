package http_handlers

import (
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/triet-truong/todo/todo"
	"github.com/triet-truong/todo/todo/dto"
)

type TodoHandler struct {
	usecase todo.TodoUseCase
}

func NewTodoHandler(usecase todo.TodoUseCase) TodoHandler {
	return TodoHandler{
		usecase: usecase,
	}
}

func (h *TodoHandler) Add(ctx echo.Context) error {
	var bodyObject dto.TodoDto
	err := json.NewDecoder(ctx.Request().Body).Decode(&bodyObject)
	if err != nil {
		logrus.Error("malformed JSON body")
		ctx.Error(err)
		return err
	}
	return h.usecase.AddTodo(bodyObject)
}
