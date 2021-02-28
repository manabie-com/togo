package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToDos_CanAddNewTodo(t *testing.T) {
	t.Run("TestToDos_CanAddNewTodo", func(t *testing.T) {
		todos := NewToDos(2, 0)

		assert.True(t, todos.CanAddNewTodo())
		assert.True(t, todos.CanAddNewTodo())
	})

	t.Run("TestToDos_CanAddNewTodo_False", func(t *testing.T) {
		todos := NewToDos(1, 0)

		assert.True(t, todos.CanAddNewTodo())
		assert.False(t, todos.CanAddNewTodo())
	})
}
