package test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/dbinterface"
	"github.com/manabie-com/togo/internal/storages/pg"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"strconv"
	"testing"
	"time"
)

func initDBForTest(driverName string, driverInfo string) dbinterface.DBInterface {
	db, _ := dbinterface.NewDB(driverName, driverInfo)
	switch driverName {
	case config.DBType.Postgres:
		return &pg.PostgresDB{DB: db}
	case config.DBType.Sqlite:
		return &sqllite.LiteDB{DB: db}
	}
	return nil
}

//var DBTest = initDBForTest(config.DBType.Sqlite,"../../../data.db?_foreign_keys=on")
var DBTest = initDBForTest(config.DBType.Postgres, config.GetPostgresDBConfig().ToString())

func TestRunner(t *testing.T) {
	t.Run("Scenario 1: Test validate user with valid user_id and password", func(t *testing.T) {
		testValidateUserWithValidUserIdAndPassword(t, sql.NullString{
			String: "firstUser",
			Valid:  true,
		}, sql.NullString{String: "example", Valid: true})
	})

	t.Run("Scenario 2: Test validate user with valid user_id and invalid password", func(t *testing.T) {
		testValidateUserWithInvalidInfo(t, sql.NullString{
			String: "firstUser",
			Valid:  true,
		}, sql.NullString{String: "exampleInvalid", Valid: true})
	})

	t.Run("Scenario 3: Test validate user with valid user_id and invalid password", func(t *testing.T) {
		testValidateUserWithInvalidInfo(t, sql.NullString{
			String: "firstUser",
			Valid:  true,
		}, sql.NullString{String: "exampleInvalid", Valid: true})
	})

	t.Run("Scenario 4: Test validate user with invalid user_id and invalid password", func(t *testing.T) {
		testValidateUserWithInvalidInfo(t, sql.NullString{
			String: "firstUserABCS",
			Valid:  true,
		}, sql.NullString{String: "exampleInvalid", Valid: true})
	})

	t.Run("Scenario 5: Test retrieve task with user has task at specific created_date", func(t *testing.T) {
		testRetrieveTasksWithUserHasTaskAtSpecificCreatedDate(t, sql.NullString{
			String: "firstUser",
			Valid:  true,
		}, sql.NullString{
			String: "2020-06-29",
			Valid:  true,
		})
	})

	t.Run("Scenario 6: Test retrieve task with user don't have task at specific created_date", func(t *testing.T) {
		testRetrieveTaskWithUserDontHaveTaskAtSpecificCreatedDate(t, sql.NullString{
			String: "firstUser",
			Valid:  true,
		}, sql.NullString{
			String: "2001-06-29",
			Valid:  true,
		})
	})

	t.Run("Scenario 7: Test retrieve task with user is not exists in user table", func(t *testing.T) {
		testRetrieveTaskWithUserIsNotExistInUserTable(t, sql.NullString{
			String: "firstUser",
			Valid:  true,
		}, sql.NullString{
			String: "2001-06-29",
			Valid:  true,
		})
	})

	t.Run("Scenario 8: Test add task with valid info", func(t *testing.T) {
		testAddTaskWithValidUserAndCreatedDate(t, &storages.Task{
			ID:          uuid.New().String(),
			Content:     "Content in db test - scenario 8",
			UserID:      "firstUser",
			CreatedDate: time.Now().Format("2006-01-02"),
		})
	})

	t.Run("Scenario 9: Test add task with user is not exists in table users", func(t *testing.T) {
		testAddTaskWithUserNotExistsInUserTable(t, &storages.Task{
			ID:          uuid.New().String(),
			Content:     "Content in db test - scenario 9",
			UserID:      "invalidUser",
			CreatedDate: time.Now().Format("2020-01-02"),
		})
	})

	t.Run("Scenario 10: Test add task with user has reached to limit of task in day", func(t *testing.T) {
		arrTask := make([]*storages.Task, 0, 6)
		now := time.Now().Format("2020-01-02")
		for i := 1; i <= 6; i++ {
			arrTask = append(arrTask, &storages.Task{
				ID:          uuid.New().String(),
				Content:     "Do homework " + strconv.Itoa(i),
				UserID:      "secondUser",
				CreatedDate: now,
			})
		}
		testAddTaskWithUserHasReachedToLimitTaskOfDay(t, arrTask)
	})
}

func testRetrieveTasksWithUserHasTaskAtSpecificCreatedDate(t *testing.T, userId sql.NullString, createdDate sql.NullString) {
	data, err := DBTest.RetrieveTasks(context.Background(), userId, createdDate)
	if err != nil {
		t.Errorf("Need error is nil but the fact is %v", err)
		return
	}
	if data == nil {
		t.Errorf("Need receive task data but the fact is nil")
		return
	}
}

func testRetrieveTaskWithUserDontHaveTaskAtSpecificCreatedDate(t *testing.T, userId sql.NullString, createdDate sql.NullString) {
	data, err := DBTest.RetrieveTasks(context.Background(), userId, createdDate)
	if err != nil {
		t.Errorf("Need error is nil but the fact is %v", err)
		return
	}
	if data != nil {
		t.Errorf("Need receive task data is nil but the fact is different. Detail: %v", data)
		return
	}
}

func testRetrieveTaskWithUserIsNotExistInUserTable(t *testing.T, userId sql.NullString, createdDate sql.NullString) {
	data, err := DBTest.RetrieveTasks(context.Background(), userId, createdDate)
	if err != nil {
		t.Errorf("Need error is nil but the fact is %v", err)
		return
	}
	if data != nil {
		t.Errorf("Need receive task data is nil but the fact is different. Detail: %v", data)
		return
	}
}

func testAddTaskWithValidUserAndCreatedDate(t *testing.T, input *storages.Task) error {
	err := DBTest.AddTask(context.Background(), input)
	if err != nil {
		t.Errorf("Need error is nil but the fact is %v", err)
		return errors.New(fmt.Sprintf("Need error is nil but the fact is %v", err))
	}
	return nil
}

func testAddTaskWithUserNotExistsInUserTable(t *testing.T, input *storages.Task) {
	err := DBTest.AddTask(context.Background(), input)
	if err == nil {
		t.Errorf("Need add task error but the fact error is nil")
	}
}

func testAddTaskWithUserHasReachedToLimitTaskOfDay(t *testing.T, input []*storages.Task) {
	n := len(input)
	for i := 0; i < n; i++ {
		if i < 5 {
			err := testAddTaskWithValidUserAndCreatedDate(t, input[i])
			if err != nil {
				t.Errorf(err.Error())
				return
			}
		} else {
			err := DBTest.AddTask(context.Background(), input[i])
			if err == nil {
				t.Errorf("Need add task error but the fact error is nil")
			}
		}
	}
}

func testValidateUserWithValidUserIdAndPassword(t *testing.T, userId sql.NullString, pwd sql.NullString) {
	valid := DBTest.ValidateUser(context.Background(), userId, pwd)
	if !valid {
		t.Errorf("But return user info is valid but the fact is different")
	}
}

func testValidateUserWithInvalidInfo(t *testing.T, userId sql.NullString, pwd sql.NullString) {
	valid := DBTest.ValidateUser(context.Background(), userId, pwd)
	if valid {
		t.Errorf("But return user info is invalid but the fact is different")
	}
}
