package todogrp

import (
	createtodo "github.com/manabie-com/togo/internal/features/create_todo"
)

// =============================================================================

// AppNewTodo contains information needed to create a new todo.
type AppNewTodo struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content"`
}

func toFeatureNewTodo(app AppNewTodo) createtodo.NewTodo {
	return createtodo.NewTodo{
		Title:   app.Title,
		Content: app.Content,
	}

}
