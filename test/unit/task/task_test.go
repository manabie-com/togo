package task_test

import (
	"os"
	"testing"
	"togo/globals/database"
	"togo/migration"
	"togo/models"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestShouldNotCreateTask(t *testing.T){
	task := models.Task{UserID: 0, Detail: "test create task"}
	_, err := models.CreateTask(task)
	assert.Contains(t, err.Error(), "1452")
}

func TestShouldCreateTask(t *testing.T){
	database.SQL.Model(models.User{}).Create(&models.User{ID: 1})

	task := models.Task{UserID: 1, Detail: "test create task"}
	createdTask, err := models.CreateTask(task)

	assert.Equal(t, uint(1), createdTask.UserID)
	assert.Equal(t, nil, err)
}

func setup() {
	database.InitDBConnection()
	migration.Migrate(database.SQL)
	database.SQL.Model(models.User{}).Create(&models.User{ID: 1})
}

func teardown() {
	migration.Rollback(database.SQL)
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	setup()
	test := m.Run()
	teardown()
	os.Exit(test)
}