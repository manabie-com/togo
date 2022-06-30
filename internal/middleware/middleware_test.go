package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/context"
	"github.com/stretchr/testify/assert"
)

const ADMIN_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VybmFtZSI6ImFkbWluIn0.ei4kWxPWuJyiIQBok-ojPpwY8CA6NcFw-APrjOuI_rk"

type testHandler struct {
	w *httptest.ResponseRecorder
	r *http.Request
}

func (t testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func TestAuthorization(t *testing.T) {
	//test middleware
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost:8080/user", nil)
	r.Header.Set("Authorization", "Bearer "+"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VybmFtZSI6ImFkbWluIn0.ei4kWxPWuJyiIQBok-ojPpwY8CA6NcFw-APrjOuI_rk")

	testHandler := testHandler{w: w, r: r}

	authorization := Authorization(testHandler)
	authorization.ServeHTTP(w, r)

	assert.Equal(t, w.Code, http.StatusOK)

	//test invalid token
	assert.Equal(t, context.Get(r, "username"), "admin")

}
