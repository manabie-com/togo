package sqllite

import (
	"github.com/manabie-com/togo/internal/storages"

	"context"
	"database/sql"
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

const createdDate string = "2021-01-11"

var user = storages.User{
	ID:       "firstUser",
	Password: "example",
}

func nullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func createRandomTask(t *testing.T) {
	param := &storages.Task{
		ID:          faker.FirstName(),
		Content:     faker.Sentence(2),
		UserID:      user.ID,
		CreatedDate: createdDate,
	}

	err := testQueries.AddTask(context.Background(), param)
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	userID := nullString(user.ID)
	returnUser, err := testQueries.GetUser(context.Background(), userID)

	assert.NoError(t, err)
	assert.Equal(t, user.ID, returnUser.ID)
	assert.NotEmpty(t, returnUser.Password)
	assert.NotEmpty(t, returnUser.MaxTodo)
}

func TestValidateUser(t *testing.T) {
	userID := nullString(user.ID)
	password := nullString(user.Password)
	user := testQueries.ValidateUser(context.Background(), userID, password)

	assert.Equal(t, true, user)
}

// User not found
func TestValidateUserErr(t *testing.T) {
	userID := nullString("id")
	password := nullString("password")
	user := testQueries.ValidateUser(context.Background(), userID, password)

	assert.Equal(t, false, user)
}

func TestAddTask(t *testing.T) {
	createRandomTask(t)
}
func TestAddTaskErr(t *testing.T) {
	// Empty foreign key
	param := &storages.Task{
		ID:          faker.FirstName(),
		Content:     faker.Sentence(2),
		UserID:      "",
		CreatedDate: createdDate,
	}

	err := testQueries.AddTask(context.Background(), param)
	assert.Error(t, err)
}

func TestRetrieveTasks(t *testing.T) {
	userID := nullString(user.ID)
	date := nullString(createdDate)

	tasks, err := testQueries.RetrieveTasks(context.Background(), userID, date)

	assert.NoError(t, err)
	for _, task := range tasks {
		assert.Equal(t, user.ID, task.UserID)
		assert.Equal(t, createdDate, task.CreatedDate)
	}
}

func TestCountTasks(t *testing.T) {
	userID := nullString(user.ID)
	date := nullString(createdDate)

	totalTasks, err := testQueries.CountTasks(context.Background(), userID, date)
	assert.NoError(t, err)

	tasks, err := testQueries.RetrieveTasks(context.Background(), userID, date)
	assert.NoError(t, err)
	assert.Equal(t, totalTasks, len(tasks))
}
