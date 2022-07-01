package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser_FailCase(t *testing.T) {
	t.Parallel()
	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	defer db.Close()
	r := SetUpRouter()
	service := handlers.HandleService(db)
	r.POST("/users", CreateUser(service))

	{ //Fail case - Not found
		input := &models.User{
			ID:            "2",
			Username:      "",
			Password:      "123456",
			MaxTaskPerDay: 5,
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1) ORDER BY "users"."id" ASC LIMIT 1`)).
			WithArgs(input.Username).
			WillReturnError(fmt.Errorf(""))

		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WithArgs(input.ID, input.Username, input.Password, input.MaxTaskPerDay)
		mock.ExpectCommit()

		jsonValue, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/example", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	}

	{ //Fail case - Username is empty
		input := &models.User{
			ID:            "2",
			Username:      "",
			Password:      "123456",
			MaxTaskPerDay: 5,
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1) ORDER BY "users"."id" ASC LIMIT 1`)).
			WithArgs(input.Username).
			WillReturnError(fmt.Errorf(""))

		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WithArgs(input.ID, input.Username, input.Password, input.MaxTaskPerDay)
		mock.ExpectCommit()

		jsonValue, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}

	{ //Fail case - Password is empty
		input := &models.User{
			ID:            "3",
			Username:      "manabie-fail-1",
			Password:      "",
			MaxTaskPerDay: 5,
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1) ORDER BY "users"."id" ASC LIMIT 1`)).
			WithArgs(input.Username).
			WillReturnError(fmt.Errorf(""))

		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WithArgs(input.ID, input.Username, input.Password, input.MaxTaskPerDay)
		mock.ExpectCommit()

		jsonValue, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}

	{ //Fail case - Usernamd Password are empty
		input := &models.User{
			ID:            "3",
			Username:      "manabie-fail-1",
			Password:      "",
			MaxTaskPerDay: 5,
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1) ORDER BY "users"."id" ASC LIMIT 1`)).
			WithArgs(input.Username).
			WillReturnError(fmt.Errorf("invalid password"))

		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
			WithArgs(input.ID, input.Username, input.Password, input.MaxTaskPerDay)
		mock.ExpectCommit()

		jsonValue, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}

	{ //Fail case - conflic data
		input := &models.User{
			ID:            "user-1",
			Username:      "new-manabie-1",
			Password:      "123456",
			MaxTaskPerDay: 5,
		}
		rows := sqlmock.NewRows([]string{"id", "username", "password", "max_task_per_day"}).
			AddRow("user-3", input.Username, input.Password, 5)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1) ORDER BY "users"."id" ASC LIMIT 1`)).
			WithArgs(input.Username).
			WillReturnRows(rows)

		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("id","username","password","max_task_per_day") VALUES ($1,$2,$3,$4) RETURNING "users"."id"`)).
			WithArgs(input.ID, input.Username, input.Password, input.MaxTaskPerDay).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		jsonValue, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}
}

func TestCreateUser_SuccessCase(t *testing.T) {
	t.Parallel()
	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	defer db.Close()
	r := SetUpRouter()
	service := handlers.HandleService(db)
	r.POST("/users", CreateUser(service))

	{ //Success case
		input := &models.User{
			ID:            "user-3",
			Username:      "new-manabie-3",
			Password:      "123456",
			MaxTaskPerDay: 5,
		}

		// Create mock data for test
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("id","username","password","max_task_per_day") VALUES ($1,$2,$3,$4) RETURNING "users"."id"`)).
			WithArgs(input.ID, input.Username, input.Password, input.MaxTaskPerDay).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		jsonValue, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}

}
