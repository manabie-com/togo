package todo

import "github.com/triet-truong/todo/todo/dto"

type TodoUseCase interface {
	AddTodo(newAlert dto.TodoDto) error
}
