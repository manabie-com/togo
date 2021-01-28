package services

import (
	"database/sql"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/usecase"
)

var TestTodoService = &ToDoService{}

type clientRequest struct {
	method string
	path   string
	body   io.Reader
	token  string
}

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSourceTest)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries := &postgres.PostgresDB{
		DB: testDB,
	}

	TestTodoService = &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		TaskUsecase: &usecase.TaskUsecase{
			Store: testQueries,
		},
		UserUsecase: &usecase.UserUsecase{
			Store: testQueries,
		},
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
