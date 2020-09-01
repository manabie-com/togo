package integration_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

var toDoService http.Handler

func init() {
	createServices()
}

func createServices() {
	db, err := sql.Open("sqlite3", "../../data.db")
	if err != nil {
		log.Fatalf("error opening db %v", err)
	}
	toDoService = &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}
}

func getToken(method string, validLoginUrl string) (string, error) {
	r := httptest.NewRequest(method, validLoginUrl, nil)
	w := httptest.NewRecorder()
	toDoService.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		return "", fmt.Errorf("wrong status code want %v but get %v", http.StatusOK, w.Code)
	}
	dataJson := make(map[string]interface{})
	err := json.Unmarshal(w.Body.Bytes(), &dataJson)
	if err != nil {
		return "", fmt.Errorf("can not parse response data %v", err)
	}
	token, ok := dataJson["data"].(string)
	if !ok {
		return "", fmt.Errorf("data is invalid, get %v", token)
	}

	return token, nil
}
