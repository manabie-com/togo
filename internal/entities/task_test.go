package entities_test

import (
	"testing"

	"github.com/manabie-com/togo/internal/entities"
)

func TestTask(t *testing.T) {
	t.Run("The content of task should be normalized once its method called", func(t *testing.T) {
		task := entities.Task{Content: "              Hello world "}
		want := "Hello world"
		task.NormalizeContent()
		assertTaskContent(t, task, want)
	})

	t.Run("The valid content of task should remains consistent after being normalized", func(t *testing.T) {
		task := entities.Task{Content: "Hello world"}
		want := "Hello world"
		task.NormalizeContent()
		assertTaskContent(t, task, want)
	})
}

func assertTaskContent(t testing.TB, task entities.Task, want string) {
	t.Helper()
	if task.Content != want {
		t.Errorf("Task content is invalid, want: '%v', got: %v", task.Content, want)
	}
}
