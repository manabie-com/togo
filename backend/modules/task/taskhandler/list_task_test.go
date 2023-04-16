package taskhandler

import (
	"context"
	"testing"

	"github.com/golang-module/carbon/v2"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/modules/task/taskmodel"
)

type mockListTaskRepo struct {
}

func (m *mockListTaskRepo) ListItem(ctx context.Context, filter *taskmodel.Filter, paging *sdkcm.Paging) ([]taskmodel.Task, error) {
	if filter.UserId == 2 {
		return []taskmodel.Task{}, nil
	}

	if filter.UserId == 3 {
		return nil, sdkcm.ErrCannotListEntity("tasks", nil)
	}

	return []taskmodel.Task{
		taskmodel.Task{
			SQLModel:    sdkcm.SQLModel{ID: 1},
			Content:     "content",
			UserId:      1,
			CreatedDate: carbon.Date{},
			IsDone:      false,
		},
	}, nil
}

func TestListTask(t *testing.T) {
	repo := &mockListTaskRepo{}

	dataTables := []struct {
		UserId     int
		Expect     error
		ExpectLeng int
	}{
		{1, nil, 1},
		{2, nil, 0},
		{3, sdkcm.ErrCannotListEntity("tasks", nil), 0},
	}

	for _, item := range dataTables {
		user := sdkcm.SimpleUser{SQLModel: sdkcm.SQLModel{ID: item.UserId}}
		hdl := NewListTaskHdl(repo, &user)

		filter := taskmodel.Filter{
			UserId:      0,
			IsDone:      nil,
			CreatedDate: nil,
		}
		paging := sdkcm.Paging{}
		paging.FullFill()
		tasks, actual := hdl.Response(context.Background(), &filter, &paging)

		if len(tasks) != item.ExpectLeng {
			t.Errorf("expect leng is %d but leng is %d", item.ExpectLeng, len(tasks))
		}

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
			t.Errorf("expect err is %s but err is %s", item.Expect, actual)
		}

	}
}
