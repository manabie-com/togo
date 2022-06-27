package test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/huynhhuuloc129/todo/util"
	"github.com/stretchr/testify/assert"
)

// fucntion create a random task
func RandomTask() models.Task {
	task := models.Task{
		Id:      int(util.RandomId()),
		Content: util.RandomContent(),
		Status:  "pending",
		Time:    time.Now(),
		UserId:  int(util.RandomUserid()),
	}
	return task
}

//function create a random new task
func RandomNewTask() models.NewTask {
	task := models.NewTask{
		Content: util.RandomContent(),
		Status:  "pending",
		Time:    time.Now(),
		UserId:  int(util.RandomUserid()),
	}
	return task
}

// unit test for get all task
func TestGetAllTask(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db)

	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"})
	for i := 0; i < 10; i++ { // create 10 random new task and add row
		task := RandomTask()
		rows.AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, 1)
	}

	query := regexp.QuoteMeta(`SELECT * FROM tasks WHERE userid = $1`)
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	tasks, err := models.GetAllTasks(h.DB, 1)
	fmt.Println(tasks)
	assert.NotEmpty(t, tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 10)
}

// unit test for get task by id
func TestFindTaskById(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db)

	task := RandomTask()
	rows := sqlmock.NewRows([]string{"id", "content", "status", "time", "timedone", "userid"}).AddRow(task.Id, task.Content, task.Status, task.Time, task.TimeDone, task.UserId)

	query := regexp.QuoteMeta(`SELECT * FROM tasks WHERE id = $1 AND userid = $2`)
	mock.ExpectQuery(query).WithArgs(task.Id, task.UserId).WillReturnRows(rows)

	mock.ExpectQuery(query).WithArgs(123).WillReturnRows(rows)
	task, valid := models.FindTaskByID(h.DB, int(task.Id), task.UserId)

	assert.NotNil(t, task)
	assert.NotEqual(t, false, valid)
}

// unit test for delete task
func TestDeleteTask(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db)

	task := RandomTask()
	query := regexp.QuoteMeta(`DELETE FROM tasks WHERE id = $1 AND userid = $2;`)
	mock.ExpectExec(query).WithArgs(task.Id, task.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := models.DeleteTask(h.DB, task.Id, task.UserId)
	assert.NoError(t, err)
}

// unit test for insert task
func TestInsertTask(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db)

	query := regexp.QuoteMeta(`INSERT INTO tasks(content, status,time, timedone, userid) VALUES ($1, $2, $3, $4, $5);`)

	newTask := RandomNewTask()
	mock.ExpectExec(query).WithArgs(newTask.Content, newTask.Status, newTask.Time, newTask.TimeDone, newTask.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := models.InsertTask(h.DB, newTask)
	assert.NoError(t, err)
}

// unit test for update task
func TestUpdateTask(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db)

	query := regexp.QuoteMeta(`UPDATE tasks SET content =COALESCE($1, content), status = COALESCE($2, status), timedone = COALESCE($3, timedone) WHERE id = $4 AND userid = $5;`)

	task := RandomTask()
	mock.ExpectExec(query).WithArgs(task.Content, task.Status, task.TimeDone, task.Id, task.UserId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := models.UpdateTask(h.DB, task, task.Id, task.UserId)
	assert.NoError(t, err)
}
