package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/api/handler"
	"github.com/manabie-com/togo/api/route"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	gr := r.Group("/test")
	route.RegisterTodo(gr, nil)
	return r
}

func TestUserMiddleware(t *testing.T) {
	r := setupRouter()

	tc := []struct {
		header       string
		expectedCode int
	}{
		{"", 400},
		{"1", 500},
	}

	for _, v := range tc {
		t.Run("User validation", func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test"+route.Todo, nil)
			req.Header.Add(handler.UserIdHeader, v.header)
			r.ServeHTTP(w, req)

			if w.Code != v.expectedCode {
				t.Fatalf("invalid status  expected %d but got %d", v.expectedCode, w.Code)
			}
		})
	}

}
