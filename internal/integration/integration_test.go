package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lntvan166/togo/internal/config"
	"lntvan166/togo/internal/delivery"
	"lntvan166/togo/internal/middleware"
	"lntvan166/togo/internal/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const ADMIN_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VybmFtZSI6ImFkbWluIn0.ei4kWxPWuJyiIQBok-ojPpwY8CA6NcFw-APrjOuI_rk"
const ADMIN = "admin"

const (
	// Config is the global config
	DATABASE_URL string = ""
	PORT         string = ""
	FREE_LIMIT   int    = 0
	VIP_LIMIT    int    = 0
	LOCALHOST    string = "http://localhost:8080"
)

var Handler *delivery.Handler

func setupIntegrationTest() *delivery.Handler {

	db := repository.Connect()

	return delivery.NewHandler(db)
}

func teardownIntegrationTest() {

}

func TestRegisterIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("POST", LOCALHOST+"/auth/register", strings.NewReader(`
	{
		"username": "test_integration",
		"password":"admin"
	}`))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	route := mux.NewRouter()
	route.HandleFunc("/auth/register", handler.Register)
	route.ServeHTTP(w, r)

	assert.Equal(t, http.StatusCreated, w.Code)

	userID, _ := handler.UserDelivery.UserUsecase.GetUserIDByUsername("test_integration")
	handler.UserDelivery.UserUsecase.DeleteUserByID(userID)
}

func TestLoginIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("POST", LOCALHOST+"/auth/login", strings.NewReader(`
	{
		"username": "admin",
		"password":"admin"
	}`))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	route := mux.NewRouter()
	route.HandleFunc("/auth/login", handler.Login)
	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	// get token from body
	var element map[string]interface{}
	json.Unmarshal([]byte(body), &element)

	token := element["token"].(string)

	assert.NotEmpty(t, token)
}

func TestGetAllUsers(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	handler := setupIntegrationTest()
	defer teardownIntegrationTest()

	r, err := http.NewRequest("GET", LOCALHOST+"/user", nil)
	assert.NoError(t, err)

	r.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	config.ADMIN = ADMIN

	w := httptest.NewRecorder()

	route := mux.NewRouter()
	route.HandleFunc("/user", handler.GetAllUsers)
	route.Use(middleware.Authorization)

	fmt.Println(r.Header.Get("Authorization"))

	route.ServeHTTP(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.NotEmpty(t, body)
}
