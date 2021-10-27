package psql

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/stretchr/testify/require"
)

const (
	jwtKey   = "wqGyEBBfPK9w3Lxw"
	email    = "me@here.com"
	password = "password"
	//hashpassword for mock
	hashPassword = "$2a$12$orZppdmhH.KRrxcZcjx0NeLPtIDpaf2GNUben4Rz7w53e5dSQJgdq"
	userID       = 1
)

func TestHandler(t *testing.T) {
	//for Login API tests
	// t.Run("ValidateUser", testLoginHandlerIncorrectHTTPMethod)
}

func TestValidateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	testCases := map[string]struct {
		email    string
		password string
		isValid  bool
	}{
		"valid": {
			email:    "me@here.com",
			password: password,
			isValid:  true,
		},
		"invalid": {
			email:    "test@gmail.com",
			password: "123456",
			isValid:  false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			u := &storages.User{
				Email:    tc.email,
				Password: tc.password,
			}

			rows := sqlmock.NewRows([]string{"password", "email"}).AddRow(hashPassword, email)
			mock.ExpectQuery(regexp.QuoteMeta(sqlValidateUser)).WithArgs(u.Email).WillReturnRows(rows)
			liteDB := NewModels(db)
			isValid := liteDB.ValidateUser(
				sql.NullString{
					String: u.Email,
					Valid:  true,
				},
				sql.NullString{
					String: u.Password,
					Valid:  true,
				},
			)
			require.NoError(t, err)
			require.Equal(t, tc.isValid, isValid)
		})
	}
}

func TestGetUserFromEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	testCases := map[string]struct {
		id       int
		max_todo int
		email    string
		isValid  bool
	}{
		"UserOne": {
			id:       1,
			max_todo: 5,
			email:    "test2@here.com",
		},
		"UserTwo": {
			id:       2,
			max_todo: 10,
			email:    "test@here.com",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			u := &storages.User{
				ID:      tc.id,
				MaxTodo: tc.max_todo,
				Email:   tc.email,
			}

			rows := sqlmock.NewRows([]string{"id", "max_todo", "email"}).AddRow(u.ID, u.MaxTodo, u.Email)
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, max_todo, email FROM users WHERE email = ?`)).WithArgs(u.Email).WillReturnRows(rows)
			liteDB := NewModels(db)
			user, err := liteDB.GetUserFromEmail(
				tc.email,
			)
			require.NoError(t, err)
			require.Equal(t, user.MaxTodo, u.MaxTodo)
			require.Equal(t, user.ID, u.ID)
		})
	}
}
