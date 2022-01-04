package services


import (
	"fmt"
	"github.com/namnhatdoan/togo/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testService = ToGoServiceImpl{}

// Test AddTask

func TestAddTaskWithExistConfig(t *testing.T) {
	t.Skip("TODO")
}

func TestAddTaskWithNonExistConfig(t *testing.T) {
	a := assert.New(t)
	email := fmt.Sprintf("dummy_add_task_%v@gmail.com", time.Now().Unix())
	taskDetail := "task 1"

	task, err := testService.AddNewTask(email, taskDetail)
	a.Nil(err)
	a.NotNil(task)
	a.NotEmpty(task.ID)
	a.Equal(task.Task, taskDetail)
	a.Equal(task.Email, email)
}

func TestAddTaskWithEmptyTask(t *testing.T) {
	t.Skip("TODO")
}

func TestAddTaskWithDuplicateTask(t *testing.T) {
	t.Skip("Not implement yet")
}

func TestAddTaskWithInvalidEmail(t *testing.T) {
	t.Skip("TODO")
}

func TestAddTaskWithEmptyEmail(t *testing.T) {
	t.Skip("TODO")
}

func TestAddTaskWithExceedLimit(t *testing.T) {
	t.Skip("TODO")
}

// Test SetConfig

func TestSetConfigCreateNewOne(t *testing.T) {
	a := assert.New(t)
	email := fmt.Sprintf("dummy_set_config_%v@gmail.com", time.Now().Unix())
	limit := int8(5)
	today := utils.GetCurrentDate()

	config, err := testService.SetUserConfig(email, limit, today)
	a.Nil(err)
	a.NotNil(config)
	a.Equal(config.Email, email)
	a.Equal(config.Limit, limit)
	a.Equal(config.Current, int8(0))
	a.True(config.Date.Equal(today))
}

func TestSetConfigUpdateOne(t *testing.T) {
	t.Skip("TODO")
}
