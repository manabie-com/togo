package users

import (
	"database/sql"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/manabie-com/togo/internal/model"
	r_users "github.com/manabie-com/togo/internal/reqres/users"
	"github.com/manabie-com/togo/pkg/database"
	"github.com/manabie-com/togo/pkg/seeding"
	"github.com/stretchr/testify/assert"
)

func init() {
	// load env
	err := godotenv.Load("../../../test.env")
	if err != nil {
		panic("load env error")
	}
	database.Init()
}

func TestAssignUserTasks(t *testing.T) {
	seeding.Truncate()
	seeding.SeedUsers(3)
	seeding.SeedTasks(1, sql.NullInt16{
		Int16: 1,
		Valid: true,
	})
	seeding.SeedTasks(5, sql.NullInt16{
		Valid: false,
	})
	assert := assert.New(t)
	testCases := []testAssignUserTasks{
		{
			name:    "user id not found",
			userID:  10,
			request: &r_users.AssignTaskRequest{},
			wantFunc: func(t *testing.T, err error) {
				assert.Equal(gorm.ErrRecordNotFound, err)
			},
		},
		{
			name:   "user id not found",
			userID: 1,
			request: &r_users.AssignTaskRequest{
				TaskIDs: []int16{1},
			},
			wantFunc: func(t *testing.T, err error) {
				assert.Equal(model.ErrSomeTasksAreNotSatisfying, err)
			},
		},
		{
			name:   "exceeding task limit",
			userID: 1,
			request: &r_users.AssignTaskRequest{
				TaskIDs: []int16{2},
			},
			wantFunc: func(t *testing.T, err error) {
				assert.Equal(model.ErrExceedingTaskLimit, err)
			},
		},
		{
			name:   "success",
			userID: 2,
			request: &r_users.AssignTaskRequest{
				TaskIDs: []int16{2},
			},
			wantFunc: func(t *testing.T, err error) {
				assert.Equal(nil, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := AssignUserTasks(testCase.userID, testCase.request)
			testCase.wantFunc(t, err)
		})

	}
}

type testAssignUserTasks struct {
	name     string
	userID   int16
	request  *r_users.AssignTaskRequest
	wantFunc func(t *testing.T, err error)
}
