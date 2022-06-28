package repository

import (
	"database/sql"
	e "lntvan166/togo/internal/entities"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var u = &e.User{
	ID:       1,
	Username: "admin",
	Password: "test",
	Plan:     "free",
	MaxTodo:  int64(10),
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestAddUser(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`INSERT INTO users (username, password, plan, max_todo) VALUES ($1, $2, $3, $4)`)

	mock.ExpectBegin()
	// prep := mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(u.Username, u.Password, u.Plan, u.MaxTodo).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.AddUser(u)
	assert.NoError(t, err)
}

func TestGetAllUser(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT * FROM users`)

	mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "plan", "max_todo"}).AddRow(u.ID, u.Username, u.Password, u.Plan, u.MaxTodo))

	users, err := repo.GetAllUsers()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
}

func TestGetUserByName(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT * FROM users WHERE username = $1`)

	mock.ExpectQuery(query).WithArgs(u.Username).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "plan", "max_todo"}).AddRow(u.ID, u.Username, u.Password, u.Plan, u.MaxTodo))

	user, err := repo.GetUserByName(u.Username)
	assert.NotEmpty(t, user)
	assert.NoError(t, err)
	assert.Equal(t, u.ID, user.ID)
}

func TestGetUserByID(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT * FROM users WHERE id = $1`)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "plan", "max_todo"}).AddRow(u.ID, u.Username, u.Password, u.Plan, u.MaxTodo))

	user, err := repo.GetUserByID(u.ID)
	assert.NotEmpty(t, user)
	assert.NoError(t, err)
	assert.Equal(t, u.ID, user.ID)
}

func TestGetUserIDByUsername(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT id FROM users WHERE username = $1`)

	mock.ExpectQuery(query).WithArgs(u.Username).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(u.ID))

	id, err := repo.GetUserIDByUsername(u.Username)
	assert.NoError(t, err)
	assert.Equal(t, u.ID, id)
}

func TestGetPlanByID(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT plan FROM users WHERE id = $1`)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(sqlmock.NewRows([]string{"plan"}).AddRow(u.Plan))

	plan, err := repo.GetPlanByID(u.ID)
	assert.NoError(t, err)
	assert.Equal(t, u.Plan, plan)
}

func TestGetPlanByUsername(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`SELECT plan FROM users WHERE username = $1`)

	mock.ExpectQuery(query).WithArgs(u.Username).WillReturnRows(sqlmock.NewRows([]string{"plan"}).AddRow(u.Plan))

	plan, err := repo.GetPlanByUsername(u.Username)
	assert.NoError(t, err)
	assert.Equal(t, u.Plan, plan)
}

func TestUpdateUser(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`UPDATE users SET username = $1, password = $2, plan = $3, max_todo = $4 WHERE id = $5`)

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(u.Username, u.Password, u.Plan, u.MaxTodo, u.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.UpdateUser(u)
	assert.NoError(t, err)
}

func TestDeleteUserByID(t *testing.T) {
	db, mock := NewMock()
	repo := &userRepository{db}
	defer db.Close()

	query := regexp.QuoteMeta(`DELETE FROM users WHERE id = $1`)

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.DeleteUserByID(u.ID)
	assert.NoError(t, err)
}
