package service

import (
	"context"
	"log"
	"testing"

	"github.com/manabie-com/togo/internal/user/repository"

	"github.com/manabie-com/togo/pkg/errorx"

	repository2 "github.com/manabie-com/togo/internal/task/repository"

	"github.com/manabie-com/togo/core/config"
	"github.com/manabie-com/togo/pkg/database"
	"github.com/stretchr/testify/assert"
)

var taskServiceTest TaskService

const (
	testUser     = "testpostgres"
	testPassword = "testpostgres"
	testDbName   = "testpostgres"
	testHost     = "localhost"
	testPort     = 5430
)

func TestMain(m *testing.M) {
	db := database.New(config.DBConfig{Test_PostgresDB: &config.PostgresConfig{
		Host:     testHost,
		Port:     testPort,
		Username: testUser,
		Password: testPassword,
		Database: testDbName,
		SSLMode:  "disable",
	}})
	taskRepo := repository2.NewTaskRepository(db.TestManabieDB)
	userRepo := repository.NewUserRepository(db.TestManabieDB)
	taskServiceTest = NewTaskService(userRepo, taskRepo, db.TestManabieDB)
	m.Run()
	d, err := db.TestManabieDB.DB()
	if err != nil {
		log.Fatalf("Could not connect sql database: %s", err)
	}
	defer d.Close()
}

func Test_CreateTask(t *testing.T) {
	t.Run("Test_CreateTask_Success", func(t *testing.T) {
		err := taskServiceTest.CreateTask(context.Background(), &CreateTaskArgs{
			Content: "Quét nhà",
			UserID:  1,
		})
		assert.Nil(t, err)
	})
	t.Run("Test_CreateTask_Fail_Missing_Arguments", func(t *testing.T) {
		err := taskServiceTest.CreateTask(context.Background(), &CreateTaskArgs{})
		if assert.NotNil(t, err) {
			e := err.(errorx.ErrorInterface)
			assert.Equal(t, errorx.ErrInvalidParameter(err).GetTitle(), e.GetTitle())
		}
	})
	t.Run("Test_CreateTask_Fail_Out_Of_Scope_Limit", func(t *testing.T) {
		err := taskServiceTest.CreateTask(context.Background(), &CreateTaskArgs{
			Content: "Quét nhà",
			UserID:  2,
		})
		if assert.NotNil(t, err) {
			e := err.(errorx.ErrorInterface)
			assert.Equal(t, errorx.ErrInternal(err).GetTitle(), e.GetTitle())
		}
	})
}

func Test_DeleteTask(t *testing.T) {
	t.Run("Test_DeleteTask_Fail_Task_Not_Found", func(t *testing.T) {
		err := taskServiceTest.DeleteTask(context.Background(), &DeleteTaskArgs{
			UserID: 1,
			ID:     70000,
		})
		if assert.NotNil(t, err) {
			e := err.(errorx.ErrorInterface)
			assert.Equal(t, errorx.ErrTaskNotFound(err).GetTitle(), e.GetTitle())
		}
	})
}

func Test_UpdateTask(t *testing.T) {
	t.Run("Test_UpdateTask_Success", func(t *testing.T) {
		err := taskServiceTest.UpdateTask(context.Background(), &UpdateTaskArgs{
			Content: "Quét nhà updated",
			UserID:  1,
			TaskID:  2,
		})
		assert.Nil(t, err)
	})
	t.Run("Test_UpdateTask_Fail_Invalid_Arguments", func(t *testing.T) {
		err := taskServiceTest.UpdateTask(context.Background(), &UpdateTaskArgs{
			UserID: 1,
			TaskID: 10000,
		})
		if assert.NotNil(t, err) {
			e := err.(errorx.ErrorInterface)
			assert.Equal(t, errorx.ErrTaskNotFound(err).GetTitle(), e.GetTitle())
		}
	})
}
