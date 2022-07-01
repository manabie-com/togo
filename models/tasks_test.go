package models

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huynhhuuloc129/todo/util"
	"github.com/stretchr/testify/assert"
)

// unit test for get all task
func TestGetAllTask(t *testing.T) {
	mock, h := CreateMockingDB()

	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	for i := 0; i < 10; i++ { // create 10 random new task and add row
		task := RandomTask()
		rows.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, 1)
	}

	query := regexp.QuoteMeta(QueryAllTaskText)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	tasks, err := h.BaseCtrl.GetAllTasks(1)
	assert.NotEmpty(t, tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 10)
}

// unit test for get task by id
func TestFindTaskById(t *testing.T) {
	mock, h := CreateMockingDB()

	task := RandomTask()
	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"}).AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, task.UserId)

	query := regexp.QuoteMeta(FindTaskByIDText)
	mock.ExpectQuery(query).WithArgs(task.Id, task.UserId).WillReturnRows(rows)
	newTask, valid := h.BaseCtrl.FindTaskByID(int(task.Id), task.UserId)

	assert.Equal(t, newTask.Content, task.Content)
	assert.Equal(t, newTask.UserId, task.UserId)
	assert.NotNil(t, task)
	assert.NotEqual(t, false, valid)
}

// unit test for delete task
func TestDeleteTask(t *testing.T) {
	mock, h := CreateMockingDB()
	task := RandomTask()
	query := regexp.QuoteMeta(DeleteTaskText)
	mock.ExpectExec(query).WithArgs(task.Id, task.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.DeleteTask(task.Id, task.UserId)
	assert.NoError(t, err)
}

// unit test for delete task
func TestDeleteAllTaskFromUser(t *testing.T) {
	mock, h := CreateMockingDB()

	userid := util.RandomInt(1, 100)
	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	for i := 0; i < 10; i++ { // create 10 random new task and add row
		task := RandomTask()
		rows.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, userid)
	}

	query := regexp.QuoteMeta(DeleteAllTaskText)
	mock.ExpectExec(query).WithArgs(userid).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.DeleteAllTaskFromUser(int(userid))
	assert.NoError(t, err)
}

// unit test for insert task
func TestInsertTask(t *testing.T) {
	mock, h := CreateMockingDB()

	newTask := RandomNewTask()
	mock.ExpectExec(regexp.QuoteMeta(InsertTaskText)).WithArgs(newTask.Content, newTask.Status, newTask.Time, newTask.TimeDone, newTask.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.InsertTask(newTask)
	assert.NoError(t, err)
}

// unit test for update task
func TestUpdateTask(t *testing.T) {
	mock, h := CreateMockingDB()

	query := regexp.QuoteMeta(UpdateTaskText)

	task := RandomTask()
	mock.ExpectExec(query).WithArgs(task.Content, task.Status, task.TimeDone, task.Id, task.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.UpdateTask(task, task.Id, task.UserId)
	assert.NoError(t, err)
}

// unit test for check limit task user
func TestCheckLimitTaskUser(t *testing.T) {
	mock, h := CreateMockingDB()

	user := RandomUser()
	rows1 := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	rows1.AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	rows2 := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	for i := 0; i < user.LimitTask; i++ { // create user.limit random new task and add row
		task := RandomTask()
		rows2.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, user.Id)
	}

	mock.ExpectQuery(regexp.QuoteMeta(FindUserByIDText)).WithArgs(user.Id).WillReturnRows(rows1)
	mock.ExpectQuery(regexp.QuoteMeta(QueryAllTaskText)).WithArgs(user.Id).WillReturnRows(rows2)

	valid, err := h.BaseCtrl.CheckLimitTaskUser(int(user.Id))

	assert.NoError(t, err)
	assert.Equal(t, false, valid) // should return false because limittask >= user.limittask
}
