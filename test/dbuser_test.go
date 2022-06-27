package test

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huynhhuuloc129/todo/models"
	"github.com/stretchr/testify/assert"
)

var admin = models.User{
	Id:        1,
	Username:  "admin",
	Password:  "admin",
	LimitTask: 0,
}
var newUser = models.NewUser{
	Username:  "newAccountTest",
	Password:  "newAccountTest",
	LimitTask: 10,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetAllUser(t *testing.T) {
	db, mock := NewMock()
	repo := &models.Repository{DB: db}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"}).AddRow(admin.Id, admin.Username, admin.Password, admin.LimitTask)

	query := regexp.QuoteMeta(`SELECT * FROM users`)
	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := repo.GetAllUser()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestFindUserById(t *testing.T) {
	db, mock := NewMock()
	repo := &models.Repository{DB: db}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"}).AddRow(admin.Id, admin.Username, admin.Password, admin.LimitTask)

	query := regexp.QuoteMeta(`SELECT * FROM users WHERE id = $1`)

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)


	mock.ExpectQuery(query).WithArgs(admin.Id).WillReturnRows(rows)

	user, valid := repo.FindUserByID(int(admin.Id))
	
	assert.NotNil(t, user)
	assert.NotEqual(t, false, valid)
}

func TestCheckUserNameExist(t *testing.T) {
	db, mock := NewMock()
	repo := models.Repository{DB: db}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "limittask"}).AddRow(admin.Id, admin.Username, admin.Password, admin.LimitTask)

	query := regexp.QuoteMeta(`SELECT * FROM users WHERE username = $1`)
	mock.ExpectQuery(query).WithArgs(admin.Username).WillReturnRows(rows)

	user, valid := repo.CheckUserNameExist(admin.Username)
	assert.NotNil(t, user)
	assert.NotEqual(t, false ,valid)
}
func TestDeleteUser(t *testing.T) {
	db, mock := NewMock()
	repo := models.Repository{DB: db}

	query := regexp.QuoteMeta(`DELETE FROM users WHERE id = $1`)
	mock.ExpectExec(query).WithArgs(admin.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteUser(int(admin.Id))
	assert.NoError(t, err)
}

func TestInsertUser(t *testing.T) {
	db, mock := NewMock()
	repo := &models.Repository{DB: db}

	defer repo.Close()

	query := regexp.QuoteMeta(`INSERT INTO users(username, password, limittask) VALUES ($1, $2, $3);`)

	mock.ExpectExec(query).WithArgs(newUser.Username, newUser.Password, newUser.LimitTask).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.InsertUser(newUser)
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	db, mock := NewMock()
	repo := models.Repository{DB: db}

	query := regexp.QuoteMeta(`UPDATE users SET username = $1, password = $2, limittask = $3 WHERE id = $4`)

	mock.ExpectExec(query).WithArgs(newUser.Username, newUser.Password, newUser.LimitTask, admin.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateUser(newUser, 1)
	assert.NoError(t, err)
}
