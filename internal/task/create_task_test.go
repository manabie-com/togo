package task

import (
	"context"
	"testing"

	"github.com/jmramos02/akaru/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	ctx := context.TODO()
	db := test.InitTestDB()
	taskName := "Eat Lunch"
	userID := 5

	ctx = context.WithValue(ctx, "db", db)
	ctx = context.WithValue(ctx, "userID", userID)

	taskService := Initalize(ctx)
	task, err := taskService.CreateTask(taskName)

	assert.Nil(t, err, "Should have no errors")
	assert.Equal(t, task.Name, taskName, "Should have the same task name")
	assert.Equal(t, task.UserID, userID, "Should have the same user id")

	db.Delete(&task)
}
