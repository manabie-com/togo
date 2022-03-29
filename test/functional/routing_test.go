package functional

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"togo/globals/database"
	"togo/migration"
	"togo/models"
	"togo/routes"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	routers := routes.InitRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/to-do", nil)
	routers.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "To-do API")
}

func TestCreateTodoRoute(t *testing.T) {
	routers := routes.InitRoutes()
	t.Run("Should fail validation", func(t *testing.T) {
		t.Run("Missing User ID and Task Detail", func(t *testing.T) {
			w, req := createReqRes("POST", "/api/to-do", nil)
			routers.ServeHTTP(w, req)

			assert.Equal(t, 400, w.Code)
			assert.Equal(t, "{\"message\":[{\"field\":\"Form.UserID\",\"validate\":\"required\"},{\"field\":\"Form.TaskDetail\",\"validate\":\"required\"}]}", w.Body.String())
		})

		t.Run("Missing User ID", func(t *testing.T) {
			form := url.Values{}
			form.Add("task_detail", "check the test")
			b := bytes.NewBufferString(form.Encode())

			w, req := createReqRes("POST", "/api/to-do", b)
			routers.ServeHTTP(w, req)

			assert.Equal(t, 400, w.Code)
			assert.Equal(t, "{\"message\":[{\"field\":\"Form.UserID\",\"validate\":\"required\"}]}", w.Body.String())
		})

		t.Run("Missing Task Detail", func(t *testing.T) {
			form := url.Values{}
			form.Add("user_id", "1")
			b := bytes.NewBufferString(form.Encode())

			w, req := createReqRes("POST", "/api/to-do", b)
			routers.ServeHTTP(w, req)

			assert.Equal(t, 400, w.Code)
			assert.Equal(t, "{\"message\":[{\"field\":\"Form.TaskDetail\",\"validate\":\"required\"}]}", w.Body.String())
		})
	})

	t.Run("Should create Todo", func(t *testing.T) {
		t.Run("Should create new User with Todo", func(t *testing.T) {
			userBefore := models.User{}
			database.SQL.First(&userBefore, 10)
			assert.Empty(t, userBefore)

			form := url.Values{}
			form.Add("user_id", "10")
			form.Add("task_detail", "test create user id = 10")
			b := bytes.NewBufferString(form.Encode())

			w, req := createReqRes("POST", "/api/to-do", b)

			routers.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)

			userAfter := models.User{}
			database.SQL.First(&userAfter, 10)
			assert.NotEmpty(t, userAfter)
		})

		t.Run("Should create new Todo", func(t *testing.T) {
			var (
				taskCountBefore int64
				taskCountAfter int64
			)
			database.SQL.Model(&models.Task{}).Where("user_id = ?", 10).Count(&taskCountBefore)
			assert.Equal(t, int64(1), taskCountBefore)

			form := url.Values{}
			form.Add("user_id", "10")
			form.Add("task_detail", "test create todo")
			b := bytes.NewBufferString(form.Encode())

			w, req := createReqRes("POST", "/api/to-do", b)

			routers.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)

			database.SQL.Model(&models.Task{}).Where("user_id = ?", 10).Count(&taskCountAfter)
			assert.Equal(t, int64(2), taskCountAfter)
		})
	})

	t.Run("Should reach daily limit", func(t *testing.T) {
		form := url.Values{}
		form.Add("user_id", "5")
		form.Add("task_detail", "test create todo")

		for i := 0; i < 10; i++ {
			b := bytes.NewBufferString(form.Encode())
			w, req := createReqRes("POST", "/api/to-do", b)
			routers.ServeHTTP(w, req)

			if i < 8 {
				assert.Equal(t, 200, w.Code)
				assert.Contains(t, w.Body.String(), "Create ToDo successful")
			}
			if i >= 8 {
				assert.Equal(t, 400, w.Code)
				assert.Contains(t, w.Body.String(), "{\"message\":\"Daily limit reached\"}")
			}
		}
	})
}

func createReqRes(method string, url string, b io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, b)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return w, req
}

func setup() {
	database.InitDBConnection()
	migration.Migrate(database.SQL)
}

func teardown() {
	migration.Rollback(database.SQL)
}

func TestMain(m *testing.M){
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	setup()
	test := m.Run()
	teardown()
	os.Exit(test)
}