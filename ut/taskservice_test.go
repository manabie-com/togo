package ut

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/repo"
	"github.com/manabie-com/togo/utils"
)

func createToken(id string) (string, error) {
	jwtKey := os.Getenv("APP_JWTKEY")

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 120).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func InitTaskService() *services.TaskService {
	if err := utils.InitEnv(); err != nil {
		log.Fatal("error loading env", err)
	}

	db, err := utils.InitDB()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	return &services.TaskService{
		TaskStore: &repo.TaskStore{
			DB: db,
		},
		UserStore: &repo.UserStore{
			DB: db,
		},
	}
}

func TestListTasks(t *testing.T) {
	c := InitTaskService()
	u := InitUserService()
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("created_date", "2021-08-29")
	req.URL.RawQuery = q.Encode()

	token, _ := createToken("testUser")
	req.Header.Set(echo.HeaderAuthorization, token)

	req, _ = u.IsValidToken(req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.ListTasks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got [%v] want [%v]",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"data":[{"id":"5edd0c84-5b22-4076-a243-10c8fc13d84c","content":"some tasks","user_id":"testUser","created_date":"2021-08-29"}]}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("ListTasks returned unexpected body: got [%v] want [%v]",
			rr.Body.String(), expected)
	}
}

func TestCreateTask(t *testing.T) {
	c := InitTaskService()
	u := InitUserService()
	var jsonStr = []byte(`{"content":"test adding task"}`)

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	token, _ := createToken("firstUser")
	req.Header.Set(echo.HeaderAuthorization, token)

	req, _ = u.IsValidToken(req)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.CreateTask)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("CreateUser returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}
