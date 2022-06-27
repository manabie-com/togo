package task

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/constants"
	"github.com/manabie-com/togo/internal/api/handlers"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories/task"
	"github.com/manabie-com/togo/internal/usecases/authorization"
	"github.com/manabie-com/togo/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddTask(t *testing.T) {
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
	r.POST("/tasks", AddTask(service))

	//Generate token
	repositories := handlers.NewRepositories(db)
	authUsecase := authorization.NewAuthUseCase(repositories.Auth)

	{ //Success case
		// Create mock data User for test
		user := models.User{
			Username: "manabie-test-2",
			Password: "123456",
		}
		mock.ExpectExec("INSERT INTO users").
			WithArgs(1, user.Username, user.Password, 5).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Create Data
		err := recordStatsUser(db, user.Username, user.Password)
		require.Nil(t, err)

		input := task.NewTask(models.Task{
			Content:    "Test Interview",
			CreateDate: time.Now().Format("2006-01-02"),
			UserID:     "1",
		})

		rows := sqlmock.NewRows([]string{"id", "content", "create_date", "user_id"}).
			AddRow("1", input.Content, input.CreateDate, 1)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" WHERE (user_id = $1 and create_date = $2)`)).
			WithArgs(input.UserID, input.CreateDate).
			WillReturnRows(rows)

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("id","content","create_date","user_id") VALUES ($1,$2,$3,$4) RETURNING "tasks"."id"`)).
			WithArgs(input.ID, input.Content, input.CreateDate, input.UserID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		//Generate token
		token, err := authUsecase.GenerateToken(input.UserID, "10")
		require.Nil(t, err)

		jsonValue, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(jsonValue))
		cookie := &http.Cookie{
			Name:   constants.CookieTokenKey,
			Value:  utils.SafeString(token),
			MaxAge: 300,
		}
		req.AddCookie(cookie)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}
}
