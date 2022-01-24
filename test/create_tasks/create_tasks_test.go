package create_tasks

import (
	"context"
	"errors"
	"testing"
	model2 "togo/module/task/model"
	"togo/module/task/repo"
	"togo/module/userconfig/model"
	model3 "togo/module/usertask/model"
)

type userCfgStore struct {}
func (u *userCfgStore) Get(ctx context.Context, cond map[string]interface{}) (*model.UserConfig, error) {
	userId := cond["user_id"].(uint)
	for _, usr := range UserConfigs {
		if usr.UserId == userId {
			return &usr, nil
		}
	}

	return nil, errors.New("DataNotFound")
}

type createTaskStore struct {}
func (t *createTaskStore) CreateTasks(ctx context.Context, data []model2.CreateTask) error {
	return nil
}

type createUserTaskStore struct {}
func (t *createUserTaskStore) CreateUserTasks(ctx context.Context, data []model3.CreateUserTask) error {
	return nil
}

type testCase struct {
	Title string
	UserId uint
	Expect string
}

func TestCreateTasks(t *testing.T)  {

	tcs := []testCase{
		{
			Title: "Validate Maximum Task of User's ID is 2",
			UserId: 2,
			Expect: "Length of task is bigger than user's limit task",
		},
		{
			Title: "Validate Maximum Task of User's ID is 1",
			UserId: 1,
			Expect: "",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Title, func(t *testing.T) {
			ctx := context.Background()
			usrCfgStr := &userCfgStore{}
			taskStr := &createTaskStore{}
			usrTaskStr := &createUserTaskStore{}
			taskRepo := repo.NewCreateTaskRepo(usrCfgStr, taskStr, usrTaskStr)

			tasks := CreateTasks()
			if err := taskRepo.CreateTasks(ctx, tc.UserId, tasks); err != nil {
				if err.Error() != tc.Expect {
					t.Fail()
				}
			}
		})
	}
}