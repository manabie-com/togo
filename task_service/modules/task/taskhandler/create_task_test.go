package taskhandler

import (
	"context"
	"testing"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/common"
	"togo/modules/task/taskmodel"
)

type mockCreateTaskRepo struct{}

func (m *mockCreateTaskRepo) CreateTask(ctx context.Context, data *taskmodel.TaskCreate) error {
	if data.UserId == 4 {
		return sdkcm.ErrCannotCreateEntity("task", nil)
	}
	return nil
}

func (m *mockCreateTaskRepo) CountTaskToday(ctx context.Context, userId int) (int, error) {
	if userId == 2 {
		return 5, nil
	}

	return 4, nil
}

func (m *mockCreateTaskRepo) IncrByNumberTaskToday(ctx context.Context, userId, number int) (int, error) {
	if userId == 2 {
		return 6, nil
	}

	if userId == 3 {
		return 0, sdkcm.ErrCannotCreateEntity("task", nil)
	}

	return 4, nil
}

type mockCreateTaskUserRepo struct{}

func (m *mockCreateTaskUserRepo) GetUserLimit(ctx context.Context, userId int) (int, error) {
	return 5, nil
}

func TestCreateTaskHdl(t *testing.T) {
	repo := &mockCreateTaskRepo{}
	userRepo := &mockCreateTaskUserRepo{}

	dataTable := []struct {
		Content string
		UserId  int
		Expect  error
	}{
		{"content", 1, nil},
		{"content", 2, common.ErrLimitTaskToday},
		{"content", 3, nil},
		{"content", 4, sdkcm.ErrCannotCreateEntity("task", nil)},
	}
	for _, item := range dataTable {
		user := sdkcm.SimpleUser{SQLModel: sdkcm.SQLModel{ID: item.UserId}}
		hdl := NewCreateTaskHdl(repo, userRepo, &user)

		data := taskmodel.TaskCreate{Content: item.Content}

		actual := hdl.Response(context.Background(), &data)

		if actual != nil {
			if actual.Error() != item.Expect.Error() {
				t.Errorf("expect err is %s but err is %s", item.Expect, actual.Error())

			}

			continue
		}

		if actual != item.Expect {
			t.Errorf("expect err is %s but err is %s", item.Expect, actual.Error())
		}
	}
}
