package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/context"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/stretchr/testify/assert"
)

// test controller response all user
func TestResponseAllUser(t *testing.T) {
	mock, h := CreateMockingDB()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	for i := 0; i < 10; i++ {
		user := models.RandomUser()
		rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)
	}

	mock.ExpectQuery(regexp.QuoteMeta(models.QueryAllUserText)).WillReturnRows(rows)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/users", nil)
	h.ResponseAllUser(w, req)

	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response")
	}

	var users []models.User
	err = json.Unmarshal(respBody, &users)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, respBody)
	assert.Len(t, users, 10)
}

// test controller response one user
func TestResponseOneUser(t *testing.T) {
	mock, h := CreateMockingDB()

	user := models.RandomUser()
	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	mock.ExpectQuery(regexp.QuoteMeta(models.FindUserByIDText)).WithArgs(user.Id).WillReturnRows(rows)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/users/"+fmt.Sprintf("%v", user.Id), nil)
	context.Set(req, "id", user.Id)
	h.ResponseOneUser(w, req)

	var userfromdb models.User
	resp := w.Result()
	err := json.NewDecoder(resp.Body).Decode(&userfromdb)
	if err != nil {
		t.Fatal("decode failed, err: " +err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, userfromdb)
	assert.Equal(t, userfromdb.Username, user.Username)
	assert.Equal(t, userfromdb.Password, user.Password)
	assert.Equal(t, userfromdb.LimitTask, user.LimitTask)
}

// test controller create user
func TestCreateUser(t *testing.T) {
	mock, h := CreateMockingDB()

	user := models.RandomUser()
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Can't marshal user, err: " + err.Error())
	}

	//exec
	mock.ExpectExec(regexp.QuoteMeta(models.InsertUserText)).WithArgs(user.Username, user.Password, 10).WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder() // set custom writer and response
	req := httptest.NewRequest("POST", "localhost:8000/users", bytes.NewReader(userJSON))
	h.CreateUser(w, req)
	
	var userfromdb models.User
	resp := w.Result()
	err = json.NewDecoder(resp.Body).Decode(&userfromdb)
	if err != nil {
		t.Fatal("decode failed, err: " +err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, userfromdb)
	assert.Equal(t, userfromdb.Username, user.Username)
}

// test controller delete user
func TestDeleteFromUser(t *testing.T) {
	mock, h := CreateMockingDB()

	user := models.RandomUser()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	mock.ExpectQuery(regexp.QuoteMeta(models.FindUserByIDText)).WithArgs(user.Id).WillReturnRows(rows)
	mock.ExpectExec(regexp.QuoteMeta(models.DeleteAllTaskText)).WithArgs(user.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta(models.DeleteUserText)).WithArgs(user.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("localhost:8000/%v", user.Id), nil)
	context.Set(req, "id", user.Id)
	h.DeleteFromUser(w, req)

	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	respBodyString := string(respBody)
	if err != nil {
		t.Errorf("Can't read body response")
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, respBodyString, "message: delete success")
}

// test controller update to user
func TestUpdateToUser(t *testing.T) {
	mock, h := CreateMockingDB()

	user := models.RandomUser()
	newUser := models.RandomNewUser()
	newUserJSON, err := json.Marshal(newUser)
	if err != nil {
		t.Errorf("Can't marshal user, err: " + err.Error())
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	mock.ExpectQuery(regexp.QuoteMeta(models.FindUserByIDText)).WithArgs(user.Id).WillReturnRows(rows)
	mock.ExpectExec(regexp.QuoteMeta(models.UpdateUserText)).WithArgs(newUser.Username, newUser.Password, newUser.LimitTask, user.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", fmt.Sprintf("localhost:8000/%v", user.Id), bytes.NewReader(newUserJSON))
	context.Set(req, "id", user.Id)
	h.UpdateToUser(w, req)

	var userfromdb models.User
	resp := w.Result()
	err = json.NewDecoder(resp.Body).Decode(&userfromdb)
	if err != nil {
		t.Fatal("decode failed, err: " +err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, userfromdb.Username, newUser.Username)
	assert.Equal(t, userfromdb.Password, newUser.Password)
	assert.Equal(t, userfromdb.LimitTask, newUser.LimitTask)
}
