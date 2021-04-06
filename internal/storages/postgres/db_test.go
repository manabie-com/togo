package postgres

import (
	"context"
	"database/sql"
	"testing"

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
			createDate: util.GetDate(),
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
			tasks, err := testPostgres.RetrieveTasks(context.Background(), tc.userId,
				tc.createDate, 100, 1)

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

			isValid := testPostgres.ValidateUser(context.Background(), tc.id, tc.password)

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
				return util.RandomTask()
			},
			check: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "Fail due to invalid user",
			getTask: func() entities.Task {
				task := util.RandomTask()
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
				task := util.RandomTask()
				task.UserID = "fourthUser"
				task.CreatedDate = util.GetDate()
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
