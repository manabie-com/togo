package services

import (
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"github.com/techschool/simplebank/util"

	"database/sql"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var TestTodoService = &ToDoService{}

type clientRequest struct {
	method string
	path   string
	body   io.Reader
	token  string
}

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries := &sqllite.LiteDB{
		DB: testDB,
	}

	TestTodoService = &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  testQueries,
	}

	os.Exit(m.Run())
}

func apiClient(req clientRequest) *httptest.ResponseRecorder {
	request := httptest.NewRequest(req.method, req.path, req.body)
	responseRecorder := httptest.NewRecorder()
	if req.token != "" {
		request.Header.Set("Authorization", req.token)
	}
	TestTodoService.ServeHTTP(responseRecorder, request)

	return responseRecorder
}
