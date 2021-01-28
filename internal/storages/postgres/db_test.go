package postgres

import (
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/utils"

	"context"
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

const createdDate string = "2021-01-28"

var user = storages.User{
	ID:       "firstUser",
	Password: "example",
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
	Convey("Test Get User", t, func() {
		userID := utils.NullString(user.ID)
		returnUser, err := testQueries.GetUser(context.Background(), userID)

		assert.NoError(t, err)
		assert.Equal(t, user.ID, returnUser.ID)
		assert.NotEmpty(t, returnUser.Password)
		assert.NotEmpty(t, returnUser.MaxTodo)
	})
}

func TestValidateUser(t *testing.T) {
	Convey("Test Validate User", t, func() {
		userID := utils.NullString(user.ID)
		password := utils.NullString(user.Password)
		user := testQueries.ValidateUser(context.Background(), userID, password)

		assert.Equal(t, true, user)
	})
}

// User not found
func TestValidateUserErr(t *testing.T) {
	Convey("Test Validate User Error", t, func() {
		userID := utils.NullString("id")
		password := utils.NullString("password")
		user := testQueries.ValidateUser(context.Background(), userID, password)

		assert.Equal(t, false, user)
	})
}

func TestAddTask(t *testing.T) {
	Convey("Test Add Random Task", t, func() {
		createRandomTask(t)
	})
}
func TestAddTaskErr(t *testing.T) {
	Convey("Test Add Task Error", t, func() {
		// Empty foreign key
		param := &storages.Task{
			ID:          faker.FirstName(),
			Content:     faker.Sentence(2),
			UserID:      "",
			CreatedDate: createdDate,
		}

		err := testQueries.AddTask(context.Background(), param)
		assert.Error(t, err)
	})
}

func TestRetrieveTasks(t *testing.T) {
	Convey("Test Retrieve Tasks", t, func() {
		userID := utils.NullString(user.ID)
		date := utils.NullString(createdDate)

		tasks, err := testQueries.RetrieveTasks(context.Background(), userID, date)

		assert.NoError(t, err)
		for _, task := range tasks {
			assert.Equal(t, user.ID, task.UserID)
			assert.Equal(t, createdDate, task.CreatedDate)
		}
	})
}

func TestCountTasks(t *testing.T) {
	Convey("Test Count Tasks", t, func() {
		userID := utils.NullString(user.ID)
		date := utils.NullString(createdDate)

		totalTasks, err := testQueries.CountTasks(context.Background(), userID, date)
		assert.NoError(t, err)

		tasks, err := testQueries.RetrieveTasks(context.Background(), userID, date)
		assert.NoError(t, err)
		assert.Equal(t, totalTasks, len(tasks))
	})
}
