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

func TestResponseAllUser(t *testing.T) {
	db, mock := models.NewMock()
	dbConn := models.NewdbConn(db)
	h := NewBaseHandler(dbConn)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	for i := 0; i < 10; i++ {
		user := models.RandomUser()
		rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/users", nil)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).WillReturnRows(rows)

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

func TestResponseOneUser(t *testing.T) {
	db, mock := models.NewMock()
	dbConn := models.NewdbConn(db)
	h := NewBaseHandler(dbConn)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	user := models.RandomUser()
	rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/users/"+fmt.Sprintf("%v", user.Id), nil)
	context.Set(req, "id", user.Id)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = $1`)).WithArgs(user.Id).WillReturnRows(rows)

	h.ResponseOneUser(w, req)
	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response")
	}

	var userfromdb models.User
	err = json.Unmarshal(respBody, &userfromdb)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, userfromdb)
	assert.Equal(t, userfromdb.Username, user.Username)
	assert.Equal(t, userfromdb.Password, user.Password)
	assert.Equal(t, userfromdb.LimitTask, user.LimitTask)
}

func TestCreateUser(t *testing.T) {
	db, mock := models.NewMock()
	dbConn := models.NewdbConn(db)
	h := NewBaseHandler(dbConn)

	user := models.RandomUser()
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Can't marshal user, err: " + err.Error())
	}

	w := httptest.NewRecorder() // set custom writer and response
	req := httptest.NewRequest("POST", "localhost:8000/users", bytes.NewReader(userJSON))

	//exec
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users(username, password, limittask) VALUES ($1, $2, $3)`)).WithArgs(user.Username, user.Password, 10).WillReturnResult(sqlmock.NewResult(1, 1))

	h.CreateUser(w, req)
	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response")
	}

	var userfromdb models.NewUser
	err = json.Unmarshal(respBody, &userfromdb)
	if err != nil {
		fmt.Println(userfromdb, string(respBody))
		t.Errorf(err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, userfromdb)
	assert.Equal(t, userfromdb.Username, user.Username)
}
