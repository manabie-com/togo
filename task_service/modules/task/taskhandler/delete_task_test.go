package taskhandler

import (
	"context"
	"errors"
	"testing"

	"github.com/golang-module/carbon/v2"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/modules/task/taskmodel"
)

type mockDeleteTaskRepo struct {
}

func (m *mockDeleteTaskRepo) GetTask(ctx context.Context, cond map[string]interface{}) (*taskmodel.Task, error) {
	date := carbon.Date{Carbon: carbon.Now()}

	if cond["id"] == 3 {
		task := taskmodel.Task{
			SQLModel:    sdkcm.SQLModel{ID: 3},
			Content:     "content",
			UserId:      3,
			CreatedDate: date,
			IsDone:      false,
		}

		return &task, nil
	}
	task := taskmodel.Task{
		SQLModel:    sdkcm.SQLModel{ID: 1},
		Content:     "content",
		UserId:      1,
		CreatedDate: date,
		IsDone:      false,
	}

	return &task, nil
}

func (m *mockDeleteTaskRepo) DeleteTask(ctx context.Context, cond map[string]interface{}) error {
	if cond["id"] == 3 {
		return sdkcm.ErrCannotDeleteEntity("task", nil)
	}

	return nil
}

func (m *mockDeleteTaskRepo) IncrByNumberTaskToday(ctx context.Context, userId, number int) (int, error) {
	return 4, nil
}

func TestDeleteTask(t *testing.T) {
	repo := &mockDeleteTaskRepo{}

	dataTables := []struct {
		Id     int
		UserId int
		Expect error
	}{
		{1, 1, nil},
		{1, 2, sdkcm.ErrNoPermission(errors.New("no permission"))},
		{3, 3, sdkcm.ErrCannotDeleteEntity("task", nil)},
	}

	for _, item := range dataTables {
		user := sdkcm.SimpleUser{SQLModel: sdkcm.SQLModel{ID: item.UserId}}
		hdl := NewDeleteTaskHdl(repo, &user)
		actual := hdl.Response(context.Background(), item.Id)

		if actual != nil {
			if item.Expect != nil {
				if actual.Error() != item.Expect.Error() {
					t.Errorf("expect err is %s but err is %s", item.Expect, actual.Error())

				}
			} else {
				t.Errorf("expect err is %s but err is %s", item.Expect, actual.Error())
			}

			continue
		}

		if actual != item.Expect {
			t.Errorf("expect err is %s but err is %s", item.Expect, actual.Error())
		}
	}

}
