package repository

import (
	"context"
	"log"
	"testing"

	"github.com/manabie-com/togo/core/config"
	"github.com/manabie-com/togo/model"
	"github.com/manabie-com/togo/pkg/database"
	"github.com/manabie-com/togo/pkg/errorx"
	"github.com/stretchr/testify/assert"
)

var (
	db = &database.Database{}
)
var taskRepo TaskRepository

const (
	testUser     = "testpostgres"
	testPassword = "testpostgres"
	testDbName   = "testpostgres"
	testHost     = "localhost"
	testPort     = 5430
)

func TestMain(m *testing.M) {
	db = database.New(config.DBConfig{Test_PostgresDB: &config.PostgresConfig{
		Host:     testHost,
		Port:     testPort,
		Username: testUser,
		Password: testPassword,
		Database: testDbName,
		SSLMode:  "disable",
	}})
	taskRepo = NewTaskRepository(db.TestManabieDB)
	m.Run()
	d, err := db.TestManabieDB.DB()
	if err != nil {
		log.Fatalf("Could not connect sql database: %s", err)
	}
	defer d.Close()
}

func Test_GetTask(t *testing.T) {
	t.Run("Test_GetTask", func(t *testing.T) {
		task, err := taskRepo.GetTask(context.Background(), &model.Task{
			ID: 1,
		})
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, 1, task.ID)
		assert.Equal(t, "Quét nhà 1", task.Content)
		assert.Equal(t, 1, task.UserID)
	})

	t.Run("Test_GetTask_Not_Found_Error", func(t *testing.T) {
		_, err := taskRepo.GetTask(context.Background(), &model.Task{
			ID: 10000,
		})
		if assert.NotNil(t, err) {
			e := err.(errorx.ErrorInterface)
			assert.Equal(t, errorx.ErrTaskNotFound(err).GetTitle(), e.GetTitle())
		}
	})
}

func Test_SaveTask(t *testing.T) {
	t.Run("Test_SaveTask", func(t *testing.T) {
		err := taskRepo.SaveTask(db.TestManabieDB, &model.Task{
			Content: "Quét nhà saved",
			UserID:  1,
		})
		assert.Nil(t, err)
	})
}
