package common

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	t.Parallel()
	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	defer db.Close()
	utils.LoadEnv("../../../../.env")
	r := SetUpRouter()
	service := handlers.HandleService(db)
	r.POST("/login", Login(service))

	{ //Success case
		user := models.User{
			Username: "manabie-test-2",
			Password: "123456",
		}

		// Create mock data for test
		mock.ExpectExec("INSERT INTO users").
			WithArgs(1, user.Username, user.Password, 5).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Create Data
		err := recordStats(db, user.Username, user.Password)
		require.Nil(t, err)

		rows := sqlmock.NewRows([]string{"id", "username", "password", "max_task_per_day"}).
			AddRow("1", user.Username, user.Password, 5)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (username = $1 and password = $2) LIMIT 1`)).
			WithArgs(user.Username, user.Password).
			WillReturnRows(rows)

		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	}

	{ //Fail case
		user := models.User{
			Username: "manabie-test",
			Password: "123456",
		}

		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/example", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	}

	{ //Fail case
		user := models.User{
			Username: "manabie-test",
			Password: "123456",
		}

		jsonValue, _ := json.Marshal(user)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
}
