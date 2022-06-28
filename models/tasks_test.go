package models

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huynhhuuloc129/todo/util"
	"github.com/stretchr/testify/assert"
)

// unit test for get all task
func TestGetAllTask(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	for i := 0; i < 10; i++ { // create 10 random new task and add row
		task := RandomTask()
		rows.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, 1)
	}

	query := regexp.QuoteMeta(`SELECT * FROM tasks WHERE userid = $1`)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	tasks, err := h.BaseCtrl.GetAllTasks(1)
	fmt.Println(tasks)
	assert.NotEmpty(t, tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 10)
}

// unit test for get task by id
func TestFindTaskById(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	task := RandomTask()
	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"}).AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, task.UserId)

	query := regexp.QuoteMeta(`SELECT * FROM tasks WHERE id = $1 AND userid = $2`)
	mock.ExpectQuery(query).WithArgs(task.Id, task.UserId).WillReturnRows(rows)
	newTask, valid := h.BaseCtrl.FindTaskByID(int(task.Id), task.UserId)

	assert.Equal(t, newTask.Content, task.Content)
	assert.Equal(t, newTask.UserId, task.UserId)
	assert.NotNil(t, task)
	assert.NotEqual(t, false, valid)
}

// unit test for delete task
func TestDeleteTask(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	task := RandomTask()
	query := regexp.QuoteMeta(`DELETE FROM tasks WHERE id = $1 AND userid = $2`)
	mock.ExpectExec(query).WithArgs(task.Id, task.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.DeleteTask(task.Id, task.UserId)
	assert.NoError(t, err)
}

// unit test for delete task
func TestDeleteAllTaskFromUser(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	userid := util.RandomInt(1, 100)
	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	for i := 0; i < 10; i++ { // create 10 random new task and add row
		task := RandomTask()
		rows.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, userid)
	}

	query := regexp.QuoteMeta(`DELETE FROM tasks WHERE userid = $1`)
	mock.ExpectExec(query).WithArgs(userid).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.DeleteAllTaskFromUser(int(userid))
	assert.NoError(t, err)
}

// unit test for insert task
func TestInsertTask(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	query := regexp.QuoteMeta(`INSERT INTO tasks(content, status,time, timedone, userid) VALUES ($1, $2, $3, $4, $5)`)

	newTask := RandomNewTask()
	mock.ExpectExec(query).WithArgs(newTask.Content, newTask.Status, newTask.Time, newTask.TimeDone, newTask.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.InsertTask(newTask)
	assert.NoError(t, err)
}

// unit test for update task
func TestUpdateTask(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	query := regexp.QuoteMeta(`UPDATE tasks SET content =COALESCE($1, content), status = COALESCE($2, status), timedone = COALESCE($3, timedone) WHERE id = $4 AND userid = $5`)

	task := RandomTask()
	mock.ExpectExec(query).WithArgs(task.Content, task.Status, task.TimeDone, task.Id, task.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.UpdateTask(task, task.Id, task.UserId)
	assert.NoError(t, err)
}

// unit test for check limit task user
func TestCheckLimitTaskUser(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	user := RandomUser()
	rows1 := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	rows1.AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	rows2 := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	for i := 0; i < user.LimitTask; i++ { // create user.limit random new task and add row
		task := RandomTask()
		rows2.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, user.Id)
	}

	queryUser := regexp.QuoteMeta(`SELECT * FROM users WHERE id = $1`)
	queryAllTask := regexp.QuoteMeta(`SELECT * FROM tasks WHERE userid = $1`)

	mock.ExpectQuery(queryUser).WithArgs(user.Id).WillReturnRows(rows1)
	mock.ExpectQuery(queryAllTask).WithArgs(user.Id).WillReturnRows(rows2)

	valid, err := h.BaseCtrl.CheckLimitTaskUser(int(user.Id))

	assert.NoError(t, err)
	assert.Equal(t, false, valid) // should return false because limittask >= user.limittask
}
