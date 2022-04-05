package integration__tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	handlers "togo/internal/pkg/deliveries/http"
	"togo/internal/pkg/domain/dtos"
	"togo/internal/pkg/domain/entities"
	"togo/internal/pkg/repositories"
	"togo/internal/pkg/usecases"
	"togo/pkg/auth"
	"togo/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {

	user := entities.User{
		ID:        1,
		Email:     "test@gmail.com",
		Password:  "jZae727K08KaOmKSgOaGzww/XVqGr/PKEgIMkjrcbJI=",
		LimitTodo: 1,
		CreatedAt: time.Now(),
	}

	// init db
	database()
	Migrate()
	gormDB.Create(&user)

	ur := repositories.NewUserRepository(gormDB)
	tr := repositories.NewToDoRepository(gormDB)
	tu := usecases.NewTodoUsecase(tr)
	th := handlers.NewTodoHandler(tu)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.AuthUser(ur))
	router.POST("/api/todo/create", th.Create)

	t.Run("success", func(t *testing.T) {
		todoReq := dtos.CreateTodoRequest{
			Task: "task1",
		}

		j, err := json.Marshal(todoReq)
		assert.NoError(t, err)

		token, err := auth.GenerateJWT(user.ID)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/api/todo/create", strings.NewReader(string(j)))
		assert.NoError(t, err)

		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Set("Content-type", "application/x-www-form-urlencoded")

		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)
		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rec.Body.Bytes(), &responseMap)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, http.StatusOK, rec.Code)

		todo := entities.Todo{}
		gormDB.First(&todo)
		assert.Equal(t, todo.Task, todoReq.Task)
	})

	t.Run("limited task fail", func(t *testing.T) {
		todoReq := dtos.CreateTodoRequest{
			Task: "task2",
		}
		j, err := json.Marshal(todoReq)
		assert.NoError(t, err)

		token, err := auth.GenerateJWT(user.ID)
		assert.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/api/todo/create", strings.NewReader(string(j)))
		assert.NoError(t, err)

		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Set("Content-type", "application/x-www-form-urlencoded")

		rec := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)

		router.ServeHTTP(rec, req)
		response := dtos.BaseResponse{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		if err != nil {
			t.Errorf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, response.Error.ErrorMessage, "LIMIT_DAILY_TASK")
	})
}
