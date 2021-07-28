package task_test

import (
	"testing"

	testfixture "github.com/manabie-com/togo/internal/database/testfixtures"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/task"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
)

var fixturePath = "../../database/testfixtures/fixtures"

func TestCreateTask(t *testing.T) {
	cases := []struct {
		Context string
		Task    storages.Task
		ErrStr  string
	}{
		{
			Context: "success",
			Task:    storages.Task{ID: "11001000", UserID: "1000", Content: "123", CreatedDate: "2021-07-27"},
			ErrStr:  "",
		},
		{
			Context: "id null",
			Task:    storages.Task{},
			ErrStr:  "pq: null value in column \"id\" violates not-null constraint",
		},
		{
			Context: "user_id null",
			Task:    storages.Task{ID: "124567"},
			ErrStr:  "pq: insert or update on table \"tasks\" violates foreign key constraint \"tasks_fk\"",
		},
		{
			Context: "user_id not exsist",
			Task:    storages.Task{ID: "124567", UserID: "0000"},
			ErrStr:  "pq: insert or update on table \"tasks\" violates foreign key constraint \"tasks_fk\"",
		},
	}

	for _, c := range cases {
		t.Run(c.Context, func(t *testing.T) {
			db := testfixture.SetupRepo(fixturePath)
			taskStore := task.NewTaskStorage(db)
			err := taskStore.CreateTask(&c.Task)
			if c.Task.ID == "11001000" {
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

func TestRetrieveTasks(t *testing.T) {
	cases := []struct {
		Context     string
		UserID      string
		CreatedDate string
		ErrStr      string
		Expected    []*storages.Task
	}{
		{
			Context:  "success",
			UserID:   "1000",
			Expected: []*storages.Task{&task1, &task1},
		},
		{
			Context:  "user_id null",
			Expected: []*storages.Task{&task1, &task1},
		},
		{
			Context:  "user_id not exsist",
			UserID:   "abcd",
			Expected: []*storages.Task{},
		},
		{
			Context:     "created_date not exists",
			UserID:      "1000",
			CreatedDate: "1998-03-19",
			Expected:    []*storages.Task{},
		},
	}

	for _, c := range cases {
		t.Run(c.Context, func(t *testing.T) {
			db := testfixture.SetupRepo(fixturePath)
			taskStore := task.NewTaskStorage(db)
			tasks, err := taskStore.RetrieveTasks(c.UserID, c.CreatedDate)
			assert.Equal(t, len(c.Expected), len(tasks))
			if err != nil {
				assert.Equal(t, c.ErrStr, err.Error())
			}
		})
	}
}
