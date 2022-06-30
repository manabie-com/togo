package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/stretchr/testify/assert"
)

// create a mock database
func CreateMockingDB() (sqlmock.Sqlmock, *BaseHandler) {
	db, mock := models.NewMock()
	dbConn := models.NewdbConn(db)
	h := NewBaseHandler(dbConn)
	return mock, h
}

// test controller register
func TestRegister(t *testing.T) {
	mock, h := CreateMockingDB()

	newUser := models.RandomNewUser()
	newUserJSON, err := json.Marshal(newUser)
	if err != nil {
		t.Errorf("Can't marshal user, err: " + err.Error())
	}

	mock.ExpectExec(regexp.QuoteMeta(models.InsertUserText)).WithArgs(newUser.Username, newUser.Password, newUser.LimitTask).WillReturnResult(sqlmock.NewResult(1, 1))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "localhost:8000/register", bytes.NewReader(newUserJSON))
	h.Register(w, req)

	resp := w.Result()
	var user models.NewUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		t.Fatal("decode failed, err: " + err.Error())
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, user.Username, newUser.Username)
	assert.Equal(t, user.Password, newUser.Password)
	assert.Equal(t, user.LimitTask, newUser.LimitTask)
}

// test controller login
func TestLogin(t *testing.T) {
	mock, h := CreateMockingDB()

	user := models.RandomUser()
	passworhHashed, _ := models.Hash(user.Password)
	newUser := models.NewUser{
		Username:  user.Username,
		Password:  user.Password,
		LimitTask: user.LimitTask,
	}
	newUserJSON, err := json.Marshal(newUser)
	if err != nil {
		t.Errorf("Can't marshal user, err: " + err.Error())
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	rows.AddRow(user.Id, user.Username, passworhHashed, user.LimitTask)

	mock.ExpectQuery(regexp.QuoteMeta(models.QueryAllUsernameText)).WithArgs(newUser.Username).WillReturnRows(rows)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "localhost:8000/login", bytes.NewReader(newUserJSON))
	h.Login(w, req)

	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't read body response")
	}

	var resToken ResponseToken
	err = json.Unmarshal(respBody, &resToken)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, resToken.Message, "login success")
}
