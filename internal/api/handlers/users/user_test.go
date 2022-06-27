package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
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

	{ //Fail case
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
}
