package postgres

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/utils"
	"github.com/stretchr/testify/assert"
)

func Test_User_CreateFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	req := &storages.User{
		ID: faker.Username(),
	}
	rErr := fmt.Errorf("sql error")

	mock.ExpectExec(regexp.QuoteMeta(CreateUserStmt)).WillReturnError(rErr)

	userStore := NewUserStore(db)
	err = userStore.Create(context.Background(), req)

	assert.Equal(t, rErr, err)
}

func Test_User_CreateSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	req := &storages.User{
		ID: faker.Username(),
	}

	mock.ExpectExec(regexp.QuoteMeta(CreateUserStmt)).WillReturnResult(sqlmock.NewResult(1, 1))

	userStore := NewUserStore(db)
	err = userStore.Create(context.Background(), req)

	assert.NoError(t, err)
}

func Test_User_FindUserFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	req := &storages.User{
		ID: faker.Username(),
	}
	rErr := fmt.Errorf("sql error")

	mock.ExpectQuery(regexp.QuoteMeta(FindUserStmt)).WillReturnError(rErr)

	userStore := NewUserStore(db)
	result, err := userStore.FindUser(context.Background(), req.ID)

	assert.Equal(t, rErr, err)
	assert.Nil(t, result)
}

func Test_User_FindUserSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	pwd, err := utils.HashPassword("somePassword")
	assert.NoError(t, err)

	req := &storages.User{
		ID:       faker.Username(),
		Password: pwd,
		MaxTodo:  5,
	}

	rows := sqlmock.NewRows([]string{"id", "password", "max_todo"}).AddRow(req.ID, req.Password, req.MaxTodo)

	mock.ExpectQuery(regexp.QuoteMeta(FindUserStmt)).WillReturnRows(rows)

	userStore := NewUserStore(db)
	u, err := userStore.FindUser(context.Background(), req.ID)

	assert.NoError(t, err)
	assert.Equal(t, req, u)
}
