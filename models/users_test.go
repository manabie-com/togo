package models

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// unit test for get all user
func TestGetAllUser(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"})
	for i := 0; i < 10; i++ {
		user := RandomUser()
		rows.AddRow(user.Id, user.Username, user.Password, user.LimitTask)
	}

	query := regexp.QuoteMeta(`SELECT * FROM users`)
	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := h.BaseCtrl.GetAllUser()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 10)
}

// unit test for get user by id
func TestFindUserById(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	user := RandomUser()
	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"}).AddRow(user.Id, user.Username, user.Password, user.LimitTask)

	query := regexp.QuoteMeta(`SELECT * FROM users WHERE id = $1`)

	mock.ExpectQuery(query).WithArgs(user.Id).WillReturnRows(rows)
	newUser, valid := h.BaseCtrl.FindUserByID(int(user.Id))

	assert.Equal(t, newUser.Username, user.Username)
	assert.Equal(t, newUser.Password, user.Password)
	assert.NotNil(t, user)
	assert.NotEqual(t, false, valid)
}

// unit test check username exist
func TestCheckUserNameExist(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	user := RandomUser()
	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"}).AddRow(user.Id, user.Username, user.Password, user.LimitTask)
	query := regexp.QuoteMeta(`SELECT * FROM users WHERE username = $1`)
	mock.ExpectQuery(query).WithArgs(user.Username).WillReturnRows(rows)

	user, valid := h.BaseCtrl.CheckUserNameExist(user.Username)
	assert.NotNil(t, user)
	assert.NotEqual(t, false, valid)
}

// unit test for delete user
func TestDeleteUser(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	user := RandomUser()
	query := regexp.QuoteMeta(`DELETE FROM users WHERE id = $1`)
	mock.ExpectExec(query).WithArgs(user.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.DeleteUser(int(user.Id))
	assert.NoError(t, err)
}

// unit test for insert user
func TestInsertUser(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	query := regexp.QuoteMeta(`INSERT INTO users(username, password, limittask) VALUES ($1, $2, $3)`)

	newUser := RandomNewUser()
	mock.ExpectExec(query).WithArgs(newUser.Username, newUser.Password, newUser.LimitTask).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.InsertUser(newUser)
	assert.NoError(t, err)
}

// unit test for delete user
func TestUpdateUser(t *testing.T) {
	db, mock := NewMock()
	DbConn := NewdbConn(db)
	h := NewBaseHandler(DbConn)

	query := regexp.QuoteMeta(`UPDATE users SET username = $1, password = $2, limittask = $3 WHERE id = $4`)

	user := RandomUser()
	newUser := RandomNewUser()
	mock.ExpectExec(query).WithArgs(newUser.Username, newUser.Password, newUser.LimitTask, user.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := h.BaseCtrl.UpdateUser(newUser, int(user.Id))
	assert.NoError(t, err)
}
