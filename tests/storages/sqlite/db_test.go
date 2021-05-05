package sqlite

import (
	"context"
	"database/sql"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/tests/utils"
)

func TestAddTask(t *testing.T) {
	ctx := context.Background()
	httpServer, serv, _ := utils.Install()
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		httpServer.Close()
		serv.Store.DB.Close()
	}(httpServer, serv)
	task := &storages.Task{}

	now := time.Now()

	task.ID = uuid.New().String()
	task.UserID = "1234"
	task.CreatedDate = now.Format("2006-01-02")

	err := serv.Store.AddTask(ctx, task)

	if nil != err {
		t.Error(err.Error())
	}

	stmt := `SELECT * FROM tasks WHERE id = $1`
	row := serv.Store.DB.QueryRowContext(ctx, stmt, task.ID)
	result := &storages.Task{}
	err = row.Scan(&result.ID, &result.Content, &result.UserID, &result.CreatedDate)

	if nil != err {
		t.Error(err.Error())
	}

	if task.ID != result.ID {
		t.Error("ID not equal")
	}

	if task.UserID != result.UserID {
		t.Error("UserID not equal")
	}

	if task.CreatedDate != result.CreatedDate {
		t.Error("CreatedDate not equal")
	}
}

func TestRetrieveTasks(t *testing.T) {
	ctx := context.Background()
	httpServer, serv, _ := utils.Install()
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		httpServer.Close()
		serv.Store.DB.Close()
	}(httpServer, serv)
	task := &storages.Task{}
	for i := 0; i < 5; i++ {
		now := time.Now()

		task.ID = uuid.New().String()
		task.UserID = "1234"
		task.CreatedDate = now.Format("2006-01-02")

		err := serv.Store.AddTask(ctx, task)

		if nil != err {
			t.Error(err.Error())
		}
	}

	tasks, err := serv.Store.RetrieveTasks(ctx,
		sql.NullString{String: task.UserID, Valid: true},
		sql.NullString{String: task.CreatedDate, Valid: true})

	if nil != err {
		t.Error(err.Error())
	}

	if 5 != len(tasks) {
		t.Error("Task is must be 5")
	}
}

func TestValidateUser(t *testing.T) {
	ctx := context.Background()
	httpServer, serv, _ := utils.Install()
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		httpServer.Close()
		serv.Store.DB.Close()
	}(httpServer, serv)

	val := serv.Store.ValidateUser(ctx,
		sql.NullString{String: utils.USERID, Valid: true},
		sql.NullString{String: utils.PASSWORD, Valid: true})

	if !val {
		t.Error("ValidateUser can't validate user")
	}

	val = serv.Store.ValidateUser(ctx,
		sql.NullString{String: utils.USERID + "no-existed", Valid: true},
		sql.NullString{String: utils.PASSWORD + "no-existed", Valid: true})

	if val {
		t.Error("ValidateUser is validate failed")
	}

}

func TestValidateTask(t *testing.T) {
	ctx := context.Background()
	httpServer, serv, _ := utils.Install()
	defer func(httpServer *httptest.Server, serv *services.ToDoService) {
		httpServer.Close()
		serv.Store.DB.Close()
	}(httpServer, serv)
	task := &storages.Task{}
	for i := 0; i < 5; i++ {
		now := time.Now()

		task.ID = uuid.New().String()
		task.UserID = "1234"
		task.CreatedDate = now.Format("2006-01-02")

		err := serv.Store.AddTask(ctx, task)

		if nil != err {
			t.Error(err.Error())
		}
	}

	countTask := serv.Store.ValidateTask(ctx, time.Now())

	if 5 != countTask {
		t.Error("Task is must be 5")
	}
}
