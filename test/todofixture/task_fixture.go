package todofixture

import (
	"errors"

	"github.com/laghodessa/togo/domain"
	"github.com/laghodessa/togo/domain/todo"
)

// NewTask returns a user fixture with optional override.
//
// Use default
//  task := NewTask()
//
// Override default
//  task := NewTask(func(task *todo.Task) {
//  	task.Message = "stop coding"
//  })
func NewTask(override ...func(*todo.Task)) todo.Task {
	user := todo.Task{
		ID:      domain.NewID(),
		UserID:  domain.NewID(),
		Message: "take out the garbage",
	}

	// override if neccessary
	if len(override) > 1 {
		panic(errors.New("override must be exactly one function"))
	}
	if len(override) > 0 {
		override[0](&user)
	}
	return user
}
