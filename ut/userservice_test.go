package ut

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/repo"
	"github.com/manabie-com/togo/utils"
)

func InitUserService() *services.UserService {
	if err := utils.InitEnv(); err != nil {
		log.Fatal("error loading env", err)
	}

	db, err := utils.InitDB()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	return &services.UserService{
		Common: &services.TransportService{
			JWTKey: "wqGyEBBfPK9w3Lxw",
		},
		UserStore: &repo.UserStore{
			DB: db,
		},
	}
}

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestCreateUser(t *testing.T) {
	u := InitUserService()
	var jsonStr = []byte(`{"user_id":"` + RandStringRunes(32) + `","password":"example"}`)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(u.CreateUser)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("CreateUser returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}
}

func TestGetAuthToken(t *testing.T) {
	u := InitUserService()
	var jsonStr = []byte(`{"user_id":"firstUser","password":"foobar"}`)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(u.GetAuthToken)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("GetAuthToken returned wrong status code: got [%v] want [%v]",
			rr.Code, http.StatusOK)
	}
	/*expected := `{"data":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
}
