package usecase

import (
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/storages/entities"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/util"
	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	testCases := []struct {
		name   string
		userId string
		check  func(t *testing.T, token string, err error)
	}{
		{
			name:   "Success",
			userId: "firstUser",
			check: func(t *testing.T, token string, err error) {
				require.NotEmpty(t, token)
				require.NoError(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			err := util.LoadConfig("../../configs")
			require.NoError(t, err)

			testStore := postgres.NewPostgres()
			require.NotNil(t, testStore)
			todo := NewToDoUsecase(testStore)
			require.NotNil(t, todo)
			token, err := todo.createToken(tc.userId)

			tc.check(t, token, err)
		})
	}
}

func TestValidToken(t *testing.T) {
	testCases := []struct {
		name   string
		userId string
		check  func(t *testing.T, userId string, status bool)
	}{
		{
			name:   "Success",
			userId: "firstUser",
			check: func(t *testing.T, userId string, status bool) {
				require.NotEmpty(t, userId)
				require.Equal(t, true, status)
				require.Equal(t, "firstUser", userId)
			},
		},
		{
			name:   "Fail: userId incorrect",
			userId: "firstUser",
			check: func(t *testing.T, userId string, status bool) {
				require.NotEmpty(t, userId)
				require.Equal(t, true, status)
				require.NotEqual(t, "firstUser123", userId)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			err := util.LoadConfig("../../configs")
			require.NoError(t, err)

			testStore := postgres.NewPostgres()
			require.NotNil(t, testStore)
			todo := NewToDoUsecase(testStore)
			require.NotNil(t, todo)
			token, err := todo.createToken(tc.userId)
			require.NotEmpty(t, token)
			require.NoError(t, err)

			userId, status := todo.ValidToken(token)
			tc.check(t, userId, status)
		})
	}
}

func TestGetToken(t *testing.T) {
	testCases := []struct {
		name     string
		userId   string
		password string
		check    func(t *testing.T, token string, err error)
	}{
		{
			name:     "Success",
			userId:   "firstUser",
			password: "example",
			check: func(t *testing.T, token string, err error) {
				require.NotEmpty(t, token)
				require.NoError(t, err)
			},
		},
		{
			name:     "Fail",
			userId:   "firstUser123",
			password: "example",
			check: func(t *testing.T, token string, err error) {
				require.Empty(t, token)
				require.Error(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			err := util.LoadConfig("../../configs")
			require.NoError(t, err)

			testStore := postgres.NewPostgres()
			require.NotNil(t, testStore)
			todo := NewToDoUsecase(testStore)
			require.NotNil(t, todo)

			token, err := todo.GetToken(tc.userId, tc.password)
			tc.check(t, token, err)
		})
	}
}

func TestAddTask(t *testing.T) {
	testCases := []struct {
		name    string
		content string
		userId  string
		check   func(t *testing.T, err error)
	}{
		{
			name:    "Success",
			content: "test",
			userId:  "secondUser",
			check: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name:    "Fail due to invalid user",
			content: "test",
			userId:  "88988",
			check: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
		{
			name:    "Fail due to maxtodo",
			content: "test",
			userId:  "fourthUser",
			check: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			err := util.LoadConfig("../../configs")
			require.NoError(t, err)

			testStore := postgres.NewPostgres()
			require.NotNil(t, testStore)
			todo := NewToDoUsecase(testStore)
			require.NotNil(t, todo)

			err = todo.AddTask(tc.content, tc.userId)
			tc.check(t, err)
		})
	}
}

func TestListTask(t *testing.T) {
	testCases := []struct {
		name       string
		createDate string
		userId     string
		check      func(t *testing.T, tasks []*entities.Task, err error)
	}{
		{
			name:       "Success",
			userId:     "fourthUser",
			createDate: time.Now().Format(util.Conf.FormatDate),
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
			err := util.LoadConfig("../../configs")
			require.NoError(t, err)

			testStore := postgres.NewPostgres()
			require.NotNil(t, testStore)
			todo := NewToDoUsecase(testStore)
			require.NotNil(t, todo)

			tasks, err := todo.ListTask(tc.createDate, tc.userId, 100, 1)
			tc.check(t, tasks, err)
		})
	}
}
