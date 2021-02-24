package sqlstore

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var tasks = []*model.Task{
	{
		ID:          uuid.New().String(),
		Content:     "new task 1",
		UserID:      "00001",
		CreatedDate: "2021-02-22",
	},
	{
		ID:          uuid.New().String(),
		Content:     "new task 2",
		UserID:      "00001",
		CreatedDate: "2021-02-22",
	},
	{
		ID:          uuid.New().String(),
		Content:     "new task 3",
		UserID:      "00001",
		CreatedDate: "2021-02-23",
	},
}

func TestAddTask(t *testing.T) {
	Convey("Add task", t, func() {
		store := setup()

		task := tasks[0]

		err := store.AddTask(context.Background(), task)
		So(err, ShouldBeNil)

		teardown(store)
	})
}

func TestRetrieveTasks(t *testing.T) {
	Convey("Retrieve tasks", t, func() {
		store := setup()

		for _, task := range tasks {
			err := store.AddTask(context.Background(), task)
			So(err, ShouldBeNil)
		}

		tasksResult, err := store.RetrieveTasks(context.Background(),
			sql.NullString{String: "00001", Valid: true},
			sql.NullString{String: "2021-02-22", Valid: true},
		)
		So(err, ShouldBeNil)
		So(len(tasksResult), ShouldEqual, 2)

		teardown(store)
	})
}
