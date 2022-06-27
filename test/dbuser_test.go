package test

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huynhhuuloc129/todo/controllers"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/huynhhuuloc129/todo/util"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// create random user
func RandomUser() models.User {
	user := models.User{
		Id:        util.RandomId(),
		Username:  util.RandomUsername(),
		Password:  util.RandomPassword(),
		LimitTask: int(util.RandomLimittask()),
	}
	return user
}

// create random new user
func RandomNewUser() models.NewUser {
	newUser := models.NewUser{
		Username:  util.RandomUsername(),
		Password:  util.RandomPassword(),
		LimitTask: int(util.RandomLimittask()),
	}
	return newUser
}

// unit test for get all user
func TestGetAllUser(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db) 

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	for i := 0; i < 10; i++ {
		user := RandomUser()
		rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)
	}

	query := regexp.QuoteMeta(`SELECT * FROM users`)
	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := models.GetAllUser(h.DB)
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 10)
}

// unit test for get user by id
func TestFindUserById(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db) 

	user := RandomUser()
	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"}).AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	query := regexp.QuoteMeta(`SELECT * FROM users WHERE id = $1`)

	mock.ExpectQuery(query).WithArgs(user.Id).WillReturnRows(rows)
	user, valid := models.FindUserByID(h.DB, int(user.Id))

	assert.NotNil(t, user)
	assert.NotEqual(t, false, valid)
}

// unit test check username exist
func TestCheckUserNameExist(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db) 

	user := RandomUser()
	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"}).AddRow(user.Id, user.Username, user.Password, user.LimitTask)
	query := regexp.QuoteMeta(`SELECT * FROM users WHERE username = $1`)
	mock.ExpectQuery(query).WithArgs(user.Username).WillReturnRows(rows)

	user, valid := models.CheckUserNameExist(h.DB, user.Username)
	assert.NotNil(t, user)
	assert.NotEqual(t, false, valid)
}

// unit test for delete user
func TestDeleteUser(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db) 

	user := RandomUser()
	query := regexp.QuoteMeta(`DELETE FROM users WHERE id = $1`)
	mock.ExpectExec(query).WithArgs(user.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := models.DeleteUser(h.DB, int(user.Id))
	assert.NoError(t, err)
}

// unit test for insert user
func TestInsertUser(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db) 

	query := regexp.QuoteMeta(`INSERT INTO users(username, password, limittask) VALUES ($1, $2, $3);`)

	newUser := RandomNewUser()
	mock.ExpectExec(query).WithArgs(newUser.Username, newUser.Password, newUser.LimitTask).WillReturnResult(sqlmock.NewResult(0, 1))

	err := models.InsertUser(h.DB, newUser)
	assert.NoError(t, err)
}

// unit test for delete user
func TestUpdateUser(t *testing.T) {
	db, mock := NewMock()
	h := controllers.NewBaseHandler(db) 

	query := regexp.QuoteMeta(`UPDATE users SET username = $1, password = $2, limittask = $3 WHERE id = $4`)

	user := RandomUser()
	newUser := RandomNewUser()
	mock.ExpectExec(query).WithArgs(newUser.Username, newUser.Password, newUser.LimitTask, user.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := models.UpdateUser(h.DB, newUser, int(user.Id))
	assert.NoError(t, err)
}
