package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"togo/pkg/random"
)

func TestNewTask(t *testing.T) {
	id, _ := uuid.NewUUID()
	userID, _ := uuid.NewUUID()
	title := random.RandomQuote()
	description := random.RandomQuote()
	dueDate := time.Now()
	task := NewTask(id, userID, title, description, dueDate)
	assert.Equal(t, id, task.ID)
	assert.Equal(t, userID, task.UserID)
	assert.Equal(t, title, task.Title)
	assert.Equal(t, description, task.Description)
	assert.Equal(t, dueDate, task.DueDate)
}
