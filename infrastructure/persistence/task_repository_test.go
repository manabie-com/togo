package persistence

import (
	"testing"

	"github.com/jfzam/togo/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestSaveTask_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var task = entity.Task{}
	task.Title = "task title"
	task.Description = "task description"
	task.UserID = 1

	var taskLimit int64 = 2

	repo := NewTaskRepository(conn)

	f, saveErr := repo.SaveTask(&task, taskLimit)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, f.Title, task.Title)
	assert.EqualValues(t, f.Description, task.Description)
	assert.EqualValues(t, f.UserID, task.UserID)
}

//Failure can be due to duplicate email, etc
//Here, we will attempt saving a Task that is already saved
func TestSaveTask_Failure(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the Task
	_, err = seedTask(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var task = entity.Task{}
	task.Title = "Task title"
	task.Description = "Task desc"
	task.UserID = 1

	var taskLimit int64 = 2

	repo := NewTaskRepository(conn)
	f, saveErr := repo.SaveTask(&task, taskLimit)

	dbMsg := map[string]string{
		"unique_title": "task title already taken",
	}
	assert.Nil(t, f)
	assert.EqualValues(t, dbMsg, saveErr)
}

//Failure due to reached task limit
//Here, we will attempt saving a Task that is already saved
func TestSaveTaskLimit_Failure(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the Task
	_, err = seedTask(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var task = entity.Task{}
	task.Title = "Task title"
	task.Description = "Task desc"
	task.UserID = 1

	var taskLimit int64 = 1

	repo := NewTaskRepository(conn)
	f, saveErr := repo.SaveTask(&task, taskLimit)

	dbMsg := map[string]string{
		"user_task_limit": "user reached its task limit for today",
	}
	assert.Nil(t, f)
	assert.EqualValues(t, dbMsg, saveErr)
}

func TestGetTask_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	Task, err := seedTask(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewTaskRepository(conn)

	f, saveErr := repo.GetTask(Task.ID)

	assert.Nil(t, saveErr)
	assert.EqualValues(t, f.Title, Task.Title)
	assert.EqualValues(t, f.Description, Task.Description)
	assert.EqualValues(t, f.UserID, Task.UserID)
}

func TestGetAllTask_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	_, err = seedTasks(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewTaskRepository(conn)
	foods, getErr := repo.GetAllTask()

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(foods), 2)
}
