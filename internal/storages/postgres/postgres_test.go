package postgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/storages"
	"reflect"
	"testing"
	"time"
)

func validateCorrectUserPwd(_t *testing.T, pg *Postgres) {
	ok := pg.ValidateUser(
		context.Background(),
		sql.NullString{
			String: "firstUser",
			Valid:  true,
		},
		sql.NullString{
			String: "example",
			Valid:  true,
		},
	)
	if !ok {
		_t.FailNow()
	}
}

func TestForPostgres(t *testing.T) {
	connStr := config.GetConfig().GetConnString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	pg := &Postgres{DB: db}

	t.Run("validate correct user/pwd", func(_t *testing.T) {
		validateCorrectUserPwd(_t, pg)
	})

	t.Run("validate incorrect user/pwd", func(_t *testing.T) {
		ok := pg.ValidateUser(
			context.Background(),
			sql.NullString{
				String: "firstUser",
				Valid:  true,
			},
			sql.NullString{
				String: "super_hero",
				Valid:  true,
			},
		)
		if ok {
			_t.FailNow()
		}
	})

	t.Run("get existing user info", func(_t *testing.T) {
		id := "firstUser"
		maxTodo := 5
		user := pg.GetUserInfo(
			context.Background(),
			sql.NullString{
				String: id,
				Valid:  true,
			},
		)
		if user == nil {
			_t.FailNow()
		}
		if user.ID != id {
			_t.FailNow()
		}
		if user.MaxTodo != maxTodo {
			_t.FailNow()
		}
	})

	t.Run("get invalid user info", func(_t *testing.T) {
		id := "superman"
		user := pg.GetUserInfo(
			context.Background(),
			sql.NullString{
				String: id,
				Valid:  true,
			},
		)
		if user != nil {
			_t.FailNow()
		}
	})

	t.Run("count tasks of user (1)", func(_t *testing.T) {
		id := "firstUser"
		createdDate := "2020-06-29"
		expectedTotal := 3
		total, err := pg.CountTasks(
			context.Background(),
			sql.NullString{
				String: id,
				Valid:  true,
			},
			sql.NullString{
				String: createdDate,
				Valid:  true,
			},
		)
		if err != nil {
			_t.Error(err.Error())
			_t.FailNow()
		}
		if total != expectedTotal {
			_t.FailNow()
		}
	})

	t.Run("count tasks of user (2)", func(_t *testing.T) {
		id := "firstUser"
		createdDate := "2020-02-30"
		expectedTotal := 0
		total, err := pg.CountTasks(
			context.Background(),
			sql.NullString{
				String: id,
				Valid:  true,
			},
			sql.NullString{
				String: createdDate,
				Valid:  true,
			},
		)
		if err != nil {
			_t.Error(err.Error())
			_t.FailNow()
		}
		if total != expectedTotal {
			_t.FailNow()
		}
	})

	t.Run("count tasks of user (3)", func(_t *testing.T) {
		id := "superman"
		createdDate := "2020-06-29"
		expectedTotal := 0
		total, err := pg.CountTasks(
			context.Background(),
			sql.NullString{
				String: id,
				Valid:  true,
			},
			sql.NullString{
				String: createdDate,
				Valid:  true,
			},
		)
		if err != nil {
			_t.Error(err.Error())
			_t.FailNow()
		}
		if total != expectedTotal {
			_t.FailNow()
		}
	})

	t.Run("retrieve tasks of user (1)", func(_t *testing.T) {
		id := "firstUser"
		createdDate := "2020-06-29"
		expectedTotal := 3
		tasks, err := pg.RetrieveTasks(
			context.Background(),
			sql.NullString{
				String: id,
				Valid:  true,
			},
			sql.NullString{
				String: createdDate,
				Valid:  true,
			},
		)
		if err != nil {
			_t.Error(err.Error())
			_t.FailNow()
		}
		if tasks == nil || len(tasks) != expectedTotal {
			_t.FailNow()
		}
	})

	t.Run("retrieve tasks of user (2)", func(_t *testing.T) {
		id := "firstUser"
		createdDate := "2020-02-30"
		expectedTotal := 0
		tasks, err := pg.RetrieveTasks(
			context.Background(),
			sql.NullString{
				String: id,
				Valid:  true,
			},
			sql.NullString{
				String: createdDate,
				Valid:  true,
			},
		)
		if err != nil {
			_t.Error(err.Error())
			_t.FailNow()
		}
		if tasks != nil && len(tasks) != expectedTotal {
			_t.FailNow()
		}
	})

	t.Run("add task", func(_t *testing.T) {
		now := time.Now()
		id := "firstUser"
		createdDate := now.Format("2006-01-02-150405")
		task := &storages.Task{
			ID:          uuid.New().String(),
			Content:     "this is content",
			UserID:      id,
			CreatedDate: createdDate,
		}
		err := pg.AddTask(context.Background(), task)
		if err != nil {
			_t.Error(err.Error())
			_t.FailNow()
		}

		// Recheck data after insert
		tasks, err := pg.RetrieveTasks(
			context.Background(),
			sql.NullString{
				String: id,
				Valid:  true,
			},
			sql.NullString{
				String: createdDate,
				Valid:  true,
			},
		)
		if err != nil {
			_t.Error(err.Error())
			_t.FailNow()
		}
		if tasks == nil || len(tasks) != 1 {
			_t.FailNow()
		}
		if !reflect.DeepEqual(tasks[0], task) {
			_t.FailNow()
		}
	})
}
