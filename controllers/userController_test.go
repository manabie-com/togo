package controllers

import (
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
	var users []models.User

	db, mock := models.NewMock()
	h := NewBaseHandler(db)

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
	err = json.Unmarshal(respBody, &users)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, respBody)
	assert.Len(t, users, 10)
}

func TestResponseOneUser(t *testing.T) {
	var userfromdb models.User

	db, mock := models.NewMock()
	h := NewBaseHandler(db)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	user := models.RandomUser()
	rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/users/"+fmt.Sprintf("%v",user.Id), nil)
	context.Set(req, "id", user.Id)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = $1`)).WithArgs(user.Id).WillReturnRows(rows)

	h.ResponseOneUser(w, req)
	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response")
	}

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
	var userfromdb models.User

	db, mock := models.NewMock()
	h := NewBaseHandler(db)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	user := models.RandomUser()
	rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "localhost:8000/users/"+fmt.Sprintf("%v",user.Id), nil)
	context.Set(req, "id", user.Id)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = $1`)).WithArgs(user.Id).WillReturnRows(rows)

	h.ResponseOneUser(w, req)
	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response")
	}

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