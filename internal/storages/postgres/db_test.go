package postgres

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages/entities"
	"github.com/manabie-com/togo/internal/util"
	"github.com/stretchr/testify/require"
)

func TestRetrieveTasks(t *testing.T) {
	testCases := []struct {
		name       string
		userId     string
		createDate string
		check      func(t *testing.T, tasks []*entities.Task, err error)
	}{
		{
			name:       "Success",
			userId:     "fourthUser",
			createDate: "2021-04-03",
			check: func(t *testing.T, tasks []*entities.Task, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, tasks)
				task := tasks[0]
				require.NotEmpty(t, task)
				require.Equal(t, "fourthUser", task.UserID)
			},
		},
		{
			name:       "Fail valid user, invalid date",
			userId:     "firstUser",
			createDate: "1998-04-03",
			check: func(t *testing.T, tasks []*entities.Task, err error) {
				require.NoError(t, err)
				require.Empty(t, tasks)
			},
		},
		{
			name:       "Fail invalid user, valid date",
			userId:     "firstUser123",
			createDate: "2021-04-03",
			check: func(t *testing.T, tasks []*entities.Task, err error) {
				require.NoError(t, err)
				require.Empty(t, tasks)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			tasks, err := testPostgres.RetrieveTasks(context.Background(), valueString(tc.userId),
				valueString(tc.createDate), valueInt(100), valueInt(1))

			tc.check(t, tasks, err)
		})
	}
}

func TestValidateUser(t *testing.T) {
	tesCases := []struct {
		name     string
		id       string
		password string
		check    func(t *testing.T, value bool)
	}{
		{
			name:     "Valid User",
			id:       "firstUser",
			password: "example",
			check: func(t *testing.T, value bool) {
				require.Equal(t, true, value)
			},
		},

		{
			name:     "Invalid User",
			id:       "firstUser",
			password: "example123",
			check: func(t *testing.T, value bool) {
				require.Equal(t, false, value)
			},
		},
	}

	for i := range tesCases {
		tc := tesCases[i]
		t.Run(tc.name, func(t *testing.T) {

			id := valueString(tc.id)
			password := valueString(tc.password)

			isValid := testPostgres.ValidateUser(context.Background(), id, password)

			tc.check(t, isValid)
		})
	}
}

func TestAddTask(t *testing.T) {

	testCases := []struct {
		name    string
		getTask func() entities.Task
		check   func(t *testing.T, err error)
	}{
		{
			name: "Success",
			getTask: func() entities.Task {
				return randomTask()
			},
			check: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},

		{
			name: "Fail due to invalid user",
			getTask: func() entities.Task {
				task := randomTask()
				task.UserID = "123"
				return task
			},
			check: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},

		{
			name: "Fail due to maxtodo",
			getTask: func() entities.Task {
				task := randomTask()
				task.UserID = "fourthUser"
				task.CreatedDate = "2021-04-03"
				return task
			},
			check: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			task := tc.getTask()
			err := testPostgres.AddTask(context.Background(), &task)
			tc.check(t, err)
		})
	}
}

func randomTask() entities.Task {
	user := []string{
		"firstUser",
		"secondUser",
		"thirdUser",
	}
	task := entities.Task{
		ID:          uuid.New().String(),
		Content:     util.RandomString(8),
		CreatedDate: time.Now().Format("2006-01-02"),
		UserID:      util.RandomStringArray(user),
	}

	return task
}

func valueString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func valueInt(value int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: value,
		Valid: true,
	}
}
