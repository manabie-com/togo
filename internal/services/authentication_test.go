package services

import (
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/storages"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:5050/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("user_id", "firstUser")
	q.Add("password", "example")
	req.URL.RawQuery = q.Encode()
	res := httptest.NewRecorder()

	db, err := storages.GetConnection("postgres", "localhost", 5432, "postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	service := &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &storages.LiteDB{
			DB: db,
		},
	}

	handler := http.HandlerFunc(service.GetAuthToken)
	handler.ServeHTTP(res, req)

	if statusCode := res.Code; statusCode != http.StatusOK {
		t.Errorf("GetAuthToken() returned wrong status code: got %d - expect %d", statusCode, http.StatusOK)
	}
}

//func getConnectionPostgres(host string, port int, user string, password string, dbname string) string {
//	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
//}