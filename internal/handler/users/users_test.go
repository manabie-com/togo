package users

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/manabie-com/togo/pkg/database"
	"github.com/manabie-com/togo/pkg/seeding"
	"github.com/stretchr/testify/assert"
)

const assignTasksPath = "/users/:id/tasks"

func init() {
	// load env
	err := godotenv.Load("../../../test.env")
	if err != nil {
		panic("load env error")
	}
	database.Init()
}

func TestAssignTasks(t *testing.T) {
	seeding.Truncate()
	seeding.SeedUsers(3)
	seeding.SeedTasks(1, sql.NullInt16{
		Int16: 1,
		Valid: true,
	})
	seeding.SeedTasks(5, sql.NullInt16{
		Valid: false,
	})
	assert := assert.New(t)
	gin.SetMode(gin.TestMode)
	uh := UserHandler{}

	testCases := []testAssignTasks{
		{
			name:   "id parse error",
			userID: "a",
			body:   "{}",
			wantFunc: func(code int) {
				assert.Equal(http.StatusBadRequest, code)
			},
		},
		{
			name:   "body parse error",
			userID: "1",
			body:   "",
			wantFunc: func(code int) {
				assert.Equal(http.StatusBadRequest, code)
			},
		},
		{
			name:   "validate error",
			userID: "1",
			body:   "{}",
			wantFunc: func(code int) {
				assert.Equal(http.StatusBadRequest, code)
			},
		},
		{
			name:   "AssignUserTasks return error",
			userID: "1",
			body:   `{"task_ids": [1]}`,
			wantFunc: func(code int) {
				assert.Equal(http.StatusBadRequest, code)
			},
		},
		{
			name:   "success",
			userID: "2",
			body:   `{"task_ids": [2]}`,
			wantFunc: func(code int) {
				assert.Equal(http.StatusOK, code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)
			r.POST(assignTasksPath, uh.AssignTasks)

			path := strings.Replace(assignTasksPath, ":id", testCase.userID, 1)
			c.Request, _ = http.NewRequest(http.MethodPost, path, bytes.NewBuffer([]byte(testCase.body)))
			r.ServeHTTP(w, c.Request)
			testCase.wantFunc(w.Code)
		})
	}
}

type testAssignTasks struct {
	name     string
	userID   string
	body     string
	wantFunc func(int)
}
