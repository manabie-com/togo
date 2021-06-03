package test

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/handlers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/stretchr/testify/assert"
)

var LoginTestCases = map[string]struct {
	userID           string
	password         string
	mockFn           func(sqlmock.Sqlmock)
	expectedRespData map[string]string
	expectedRespCode int
}{
	"Should login successfully": {
		userID:   "test_user_id",
		password: "test_password",
		mockFn: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(`SELECT (.+) FROM "manabie"."users" (.+)`).
				WithArgs("test_user_id", "test_password").
				WillReturnRows(sqlmock.NewRows([]string{
					models.UserColumns.ID,
					models.UserColumns.Password,
				}).AddRow("test_user_id", "test_password"))
		},
		expectedRespData: nil,
		expectedRespCode: http.StatusOK,
	},
	"Should return unauthorized error no now data found": {
		userID:   "test_user_id",
		password: "test_password",
		mockFn: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(`SELECT (.+) FROM "manabie"."users" (.+)`).
				WithArgs("test_user_id", "test_password").
				WillReturnRows(sqlmock.NewRows([]string{}).RowError(0, sql.ErrNoRows))
		},
		expectedRespData: nil,
		expectedRespCode: http.StatusBadRequest,
	},
	"Should return unauthorized error on invalid user id": {
		userID:   "",
		password: "test_password",
		mockFn: func(mock sqlmock.Sqlmock) {
			//don't expect
		},
		expectedRespData: nil,
		expectedRespCode: http.StatusBadRequest,
	},
	"Should return unauthorized error on invalid password": {
		userID:   "test_user_id",
		password: "",
		mockFn: func(mock sqlmock.Sqlmock) {
			//don't expect
		},
		expectedRespData: nil,
		expectedRespCode: http.StatusBadRequest,
	},
	"Should return internal server error": {
		userID:   "test_user_id",
		password: "test_password",
		mockFn: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(`SELECT (.+) FROM "manabie"."users" (.+)`).
				WithArgs("test_user_id", "test_password").
				WillReturnError(nil)
		},
		expectedRespData: nil,
		expectedRespCode: http.StatusInternalServerError,
	},
}

func TestLogin(t *testing.T) {
	defaultJwt := "test"
	for caseName, tCase := range LoginTestCases {
		t.Run(caseName, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			request, _ := http.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/login?user_id=%s&password=%s", tCase.userID, tCase.password),
				nil,
			)
			tCase.mockFn(mock)
			response := httptest.NewRecorder()
			handlers := &handlers.Handlers{JWTSecret: defaultJwt, DB: db}
			handlers.LoadHandlers()
			handlers.ServeHTTP(response, request)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
			assert.Equal(t, tCase.expectedRespCode, response.Code)
		})

	}
}
