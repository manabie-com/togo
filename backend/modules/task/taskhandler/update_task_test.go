package taskhandler

import (
	"context"
	"errors"
	"testing"

	"github.com/golang-module/carbon/v2"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/modules/task/taskmodel"
)

type mockUpdateTaskRepo struct {
}

func (m *mockUpdateTaskRepo) GetTask(ctx context.Context, cond map[string]interface{}) (*taskmodel.Task, error) {
	if cond["id"] == 2 {
		return nil, sdkcm.ErrCannotGetEntity("task", nil)
	}
	date := carbon.Date{Carbon: carbon.Now()}

	if cond["id"] == 3 {
		task := taskmodel.Task{
			SQLModel:    sdkcm.SQLModel{ID: 3},
			Content:     "content",
			UserId:      2,
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

func (m *mockUpdateTaskRepo) UpdateTask(ctx context.Context, cond map[string]interface{}, dataUpdate *taskmodel.TaskUpdate) error {
	if cond["id"] == 3 {
		return sdkcm.ErrCannotUpdateEntity("task", nil)
	}

	return nil
}

func TestUpdateTask(t *testing.T) {
	repo := &mockUpdateTaskRepo{}

	dataTables := []struct {
		Id     int
		UserId int
		Expect error
	}{
		{1, 1, nil},
		{1, 2, sdkcm.ErrNoPermission(errors.New("no permission"))},
		{2, 2, sdkcm.ErrCannotGetEntity("task", nil)},
		{3, 2, sdkcm.ErrCannotUpdateEntity("task", nil)},
	}

	for _, item := range dataTables {
		user := sdkcm.SimpleUser{SQLModel: sdkcm.SQLModel{ID: item.UserId}}

		hdl := NewUpdateTaskHdl(repo, &user)
		isDone := true
		data := taskmodel.TaskUpdate{
			IsDone: &isDone,
		}
		actual := hdl.Response(context.Background(), item.Id, &data)

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
