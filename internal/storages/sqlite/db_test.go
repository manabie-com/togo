package sqllite

import (
	"bytes"
	"database/sql"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/utils/test_constants"
)

var (
	timeNow  string = time.Now().Format("2006-01-02")
	userId          = sql.NullString{String: test_constants.UserName, Valid: true}
	password        = sql.NullString{String: test_constants.Password, Valid: true}
)

const (
	findUserQuery      = "SELECT id FROM users WHERE id = \\? AND password = \\?"
	getMaxToDoQuery    = "SELECT max_todo FROM users WHERE id = \\?"
	retrieveTasksQuery = "SELECT id, content, user_id, created_date FROM tasks WHERE user_id = \\? AND created_date = \\?"
	createTaskQuery    = "INSERT INTO tasks \\(id, content, user_id, created_date\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
)

func TestValidateUserWithCorrectCredentialsReturnsTrue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	loginReq, loginErr := http.NewRequest(http.MethodGet, test_constants.LoginUrl, nil)
	if loginErr != nil {
		t.Fatal(loginErr)
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(test_constants.UserName)
	mock.ExpectQuery(findUserQuery).WithArgs(test_constants.UserName, test_constants.Password).WillReturnRows(rows)

	actual := liteDb.ValidateUser(loginReq.Context(), userId, password)
	expected := true

	if actual != expected {
		t.Errorf("returned wrong result: actual: %v expected: %v", actual, expected)
	}
}

func TestValidateUserWithIncorrectCredentialsReturnsFalse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	userId := sql.NullString{
		String: test_constants.WrongUserName,
		Valid:  true,
	}
	password := sql.NullString{
		String: test_constants.WrongPassword,
		Valid:  true,
	}

	loginReq, loginErr := http.NewRequest(http.MethodGet, "/login?user_id="+test_constants.WrongUserName+"&password="+test_constants.WrongPassword, nil)
	if loginErr != nil {
		t.Fatal(loginErr)
	}

	rows := sqlmock.NewRows([]string{"id"})
	mock.ExpectQuery(findUserQuery).WithArgs(test_constants.WrongUserName, test_constants.WrongPassword).WillReturnRows(rows)

	actual := liteDb.ValidateUser(loginReq.Context(), userId, password)
	expected := false

	if actual != expected {
		t.Errorf("returned wrong result: actual: %v expected: %v", actual, expected)
	}
}

func TestGetMaxToDoWithCorrectCredentials(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"max_todo"}).AddRow(test_constants.DefaultMaxToDo)
	mock.ExpectQuery(getMaxToDoQuery).WithArgs(test_constants.UserName).WillReturnRows(rows)

	actual := liteDb.GetMaxToDo(req.Context(), userId)
	expected := test_constants.DefaultMaxToDo

	if actual != expected {
		t.Errorf("returned wrong result: actual: %v expected: %v", actual, expected)
	}
}

func TestGetMaxToDoWithIncorrectCredentials(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	userId := sql.NullString{
		String: test_constants.WrongUserName,
		Valid:  true,
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"max_todo"})
	mock.ExpectQuery(getMaxToDoQuery).WithArgs(test_constants.WrongUserName).WillReturnRows(rows)

	actual := liteDb.GetMaxToDo(req.Context(), userId)
	expected := -1
	if actual != expected {
		t.Errorf("returned wrong result: actual: %v expected: %v", actual, expected)
	}
}

func TestRetrieveTasksCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	createdDate := sql.NullString{
		String: timeNow,
		Valid:  true,
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow("1", "first content", test_constants.UserName, timeNow).
		AddRow("2", "second content", test_constants.UserName, timeNow)

	mock.ExpectQuery(retrieveTasksQuery).WithArgs(test_constants.UserName, timeNow).WillReturnRows(rows)

	actual := liteDb.RetrieveTasksCount(req.Context(), userId, createdDate)
	expected := 2

	if actual != expected {
		t.Errorf("returned wrong result: actual: %v expected: %v", actual, expected)
	}
}

func TestRetrieveTasksCountWithIncorrectCredentials(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	userId := sql.NullString{
		String: test_constants.WrongUserName,
		Valid:  true,
	}

	createdDate := sql.NullString{
		String: timeNow,
		Valid:  true,
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"})
	mock.ExpectQuery(retrieveTasksQuery).WithArgs(test_constants.WrongUserName, timeNow).WillReturnRows(rows)

	actual := liteDb.RetrieveTasksCount(req.Context(), userId, createdDate)
	expected := 0

	if actual != expected {
		t.Errorf("returned wrong result: actual: %v expected: %v", actual, expected)
	}
}

func TestRetrieveTasksCountWithError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	createdDate := sql.NullString{
		String: timeNow,
		Valid:  true,
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow("1", "first content", test_constants.UserName, timeNow).
		AddRow("2", "second content", test_constants.UserName, timeNow).
		RowError(1, errors.New("Error retrieving tasks"))

	mock.ExpectQuery(retrieveTasksQuery).WithArgs(test_constants.UserName, timeNow).WillReturnRows(rows)

	actual := liteDb.RetrieveTasksCount(req.Context(), userId, createdDate)
	expected := -1

	if actual != expected {
		t.Errorf("returned wrong result: actual: %v expected: %v", actual, expected)
	}
}

func TestRetrieveTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	req, err := http.NewRequest(http.MethodGet, test_constants.GetTasksUrl, nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedTask := storages.Task{
		ID:          "001",
		Content:     test_constants.TaskContent,
		UserID:      test_constants.UserName,
		CreatedDate: test_constants.CreatedDate,
	}

	createdDate := sql.NullString{
		String: test_constants.CreatedDate,
		Valid:  true,
	}

	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(expectedTask.ID, expectedTask.Content, expectedTask.UserID, expectedTask.CreatedDate)
	mock.ExpectQuery(retrieveTasksQuery).WithArgs(test_constants.UserName, test_constants.CreatedDate).WillReturnRows(rows)

	actual, _ := liteDb.RetrieveTasks(req.Context(), userId, createdDate)
	expectedResult := []*storages.Task{
		&expectedTask,
	}

	for i, expected := range expectedResult {
		if actual[i].ID != expected.ID {
			t.Errorf("returned wrong result: actual: %v expected: %v", actual[i].ID, expected.ID)
		}
		if actual[i].Content != expected.Content {
			t.Errorf("returned wrong result: actual: %v expected: %v", actual[i].Content, expected.Content)
		}
		if actual[i].UserID != expected.UserID {
			t.Errorf("returned wrong result: actual: %v expected: %v", actual[i].UserID, expected.UserID)
		}
		if actual[i].CreatedDate != expected.CreatedDate {
			t.Errorf("returned wrong result: actual: %v expected: %v", actual[i].CreatedDate, expected.CreatedDate)
		}
	}
}

func TestRetrieveTasksWithError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	req, err := http.NewRequest(http.MethodGet, test_constants.GetTasksUrl, nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedTask := storages.Task{
		ID:          "001",
		Content:     test_constants.TaskContent,
		UserID:      test_constants.UserName,
		CreatedDate: test_constants.CreatedDate,
	}

	createdDate := sql.NullString{
		String: test_constants.CreatedDate,
		Valid:  true,
	}

	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"}).
		AddRow(expectedTask.ID, expectedTask.Content, expectedTask.UserID, expectedTask.CreatedDate).
		RowError(0, errors.New("Error retrieving tasks"))
	mock.ExpectQuery(retrieveTasksQuery).WithArgs(test_constants.UserName, test_constants.CreatedDate).WillReturnRows(rows)

	actual, err := liteDb.RetrieveTasks(req.Context(), userId, createdDate)

	if err == nil {
		t.Errorf("returned nil error, should return an error object")
	}

	if actual != nil {
		t.Errorf("returned wrong result: actual: %v expected: nil", actual)
	}
}

func TestRetrieveEmptyTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	req, err := http.NewRequest(http.MethodGet, test_constants.GetTasksUrl, nil)
	if err != nil {
		t.Fatal(err)
	}

	createdDate := sql.NullString{
		String: test_constants.CreatedDate,
		Valid:  true,
	}

	rows := sqlmock.NewRows([]string{"id", "content", "user_id", "created_date"})
	mock.ExpectQuery(retrieveTasksQuery).WithArgs(test_constants.UserName, test_constants.CreatedDate).WillReturnRows(rows)

	actual, _ := liteDb.RetrieveTasks(req.Context(), userId, createdDate)

	if len(actual) != 0 {
		t.Errorf("returned wrong result: actual: %v expected: []", actual)
	}
}

func TestCreateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	expectedTask := storages.Task{
		ID:          "001",
		Content:     test_constants.TaskContent,
		UserID:      test_constants.UserName,
		CreatedDate: timeNow,
	}

	mock.ExpectExec(createTaskQuery).WithArgs(expectedTask.ID, expectedTask.Content, expectedTask.UserID, expectedTask.CreatedDate).WillReturnResult(sqlmock.NewResult(0, 1))

	actual := liteDb.AddTask(req.Context(), &expectedTask)
	if actual != nil {
		t.Errorf("returned wrong result: actual: %v expected: nil", actual)
	}
}

func TestCreateTaskWithError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	liteDb := LiteDB{
		DB: db,
	}

	jsonStr := []byte(test_constants.CreateTaskReqBody)
	req, err := http.NewRequest(http.MethodPost, test_constants.PostTasksUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	expectedTask := storages.Task{
		ID:          "001",
		Content:     test_constants.TaskContent,
		UserID:      test_constants.UserName,
		CreatedDate: timeNow,
	}

	mock.ExpectExec(createTaskQuery).WithArgs(expectedTask.ID, expectedTask.Content, expectedTask.UserID, expectedTask.CreatedDate).WillReturnError(errors.New("some error"))

	actual := liteDb.AddTask(req.Context(), &expectedTask)
	if actual == nil {
		t.Errorf("returned nil error, should return an error object")
	}
}
