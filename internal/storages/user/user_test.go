package user_test

import (
	"testing"

	testfixture "github.com/manabie-com/togo/internal/database/testfixtures"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/user"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

var fixturePath = "../../database/testfixtures/fixtures"

func TestGetUser(t *testing.T) {
	cases := []struct {
		Context  string
		ID       string
		Password string
		ErrStr   string
	}{
		{
			Context:  "success",
			ID:       "1000",
			Password: "password",
		},
		{
			Context: "id not exists",
			ID:      "0000",
			ErrStr:  "record not found",
		},
		{
			Context:  "user_id null",
			ID:       "1001",
			Password: "abcd",
			ErrStr:   "record not found",
		},
	}

	for _, c := range cases {
		t.Run(c.Context, func(t *testing.T) {
			db := testfixture.SetupRepo(fixturePath)
			userStore := user.NewUserStorage(db)
			err := userStore.GetUser(c.ID, c.Password)
			if c.ID == "1000" {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.ErrStr, err.Error())
			}
		})
	}
}

var (
	task1 = storages.Task{
		ID:     "10001001",
		UserID: "1000",
	}
)

func TestGetUsersTasks(t *testing.T) {
	cases := []struct {
		Context     string
		ID          string
		CreatedDate string
		ErrStr      string
		Expected    []*storages.Task
	}{
		{
			Context:  "success",
			ID:       "1000",
			Expected: []*storages.Task{&task1, &task1},
		},
		{
			Context:  "user_id not exists",
			ID:       "abcd",
			Expected: []*storages.Task{},
			ErrStr:   "record not found",
		},
		{
			Context:     "created date invalid",
			ID:          "1000",
			CreatedDate: "1998-03-19",
			Expected:    []*storages.Task{},
		},
	}

	for _, c := range cases {
		t.Run(c.Context, func(t *testing.T) {
			db := testfixture.SetupRepo(fixturePath)
			userStore := user.NewUserStorage(db)
			user, err := userStore.GetUsersTasks(c.ID, c.CreatedDate)
			assert.Equal(t, len(c.Expected), len(user.Tasks))

			if c.ID == "1000" {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, c.ErrStr, err.Error())
			}
		})
	}
}
