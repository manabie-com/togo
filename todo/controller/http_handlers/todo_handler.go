package http_handlers

import (
	"encoding/json"
	"errors"

	"github.com/labstack/echo"
	"github.com/triet-truong/todo/domain"
	"github.com/triet-truong/todo/todo/dto"
	"github.com/triet-truong/todo/utils"
)

type TodoHandler struct {
	usecase domain.TodoUseCase
}

func NewTodoHandler(usecase domain.TodoUseCase) TodoHandler {
	return TodoHandler{
		usecase: usecase,
	}
}

func (h *TodoHandler) Add(ctx echo.Context) error {
	var bodyObject dto.TodoDto
	err := json.NewDecoder(ctx.Request().Body).Decode(&bodyObject)
	if err != nil {
		utils.ErrorLog(errors.New("malformed JSON body"))
		ctx.Error(err)
		return err
	}
	err = h.usecase.AddTodo(bodyObject)
	return err
}
