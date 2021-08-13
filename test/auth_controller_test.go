package test

import (
	"encoding/json"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/controller"
	"github.com/manabie-com/togo/internal/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthControllerSuccess(t *testing.T) {
	db, err := config.GetPostgersDB()

	if err != nil{
		t.Fatal(err)
	}

	auth :=controller.NewAuthController(db)
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("user_id", "firstUser")
	q.Add("password", "example")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(auth.GetAuthToken)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	jwt := utils.NewJwt()
	token, _ := jwt.CreateToken("firstUser")

	test := map[string]string{
		"data": token,
	}
	jsonString, _ := json.Marshal(test)
	print(jsonString)
	if !strings.Contains(rr.Body.String(), string(jsonString)) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(jsonString))
	}

}