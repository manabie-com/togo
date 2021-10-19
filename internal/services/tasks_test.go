package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/usecases/user"
	"github.com/stretchr/testify/suite"
)

var (
	dns = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"togo_user",
		"togo_password",
		"localhost",
		"5432",
		"togo_db_test",
	)
)

type IntegrationTestTaskSuite struct {
	suite.Suite
	port        int
	dbConn      *sql.DB
	dbMigration *migrate.Migrate
}

func TestIntegrationTestTaskSuite(t *testing.T) {
	suite.Run(t, &IntegrationTestTaskSuite{})
}

func (s *IntegrationTestTaskSuite) SetupSuite() {
	s.port = 6969
	dbConn, err := sql.Open("postgres", dns)
	s.Require().NoError(err)
	s.dbConn = dbConn

	// setup migration
	s.dbMigration, err = migrate.New("file://../../db/migrations", dns)
	s.Require().NoError(err)
	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}

	// setup JWT environment variable
	os.Setenv("JWT_KEY", "ddwqGyEBBdqfPK9w3Lxw")
	os.Setenv("JWT_TIMEOUT", "9999999")

	// Run http server
	go func() {
		if err = http.ListenAndServe(fmt.Sprintf(":%d", s.port), NewToDoService(dbConn)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()
}

func (s *IntegrationTestTaskSuite) TearDownSuite() {
	// p, _ := os.FindProcess(syscall.Getpid())
	// p.Signal(syscall.SIGINT)
}

func (s *IntegrationTestTaskSuite) SetupTest() {
	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}
}

func (s *IntegrationTestTaskSuite) TearDownTest() {
	s.NoError(s.dbMigration.Down())
}

func (s *IntegrationTestTaskSuite) Test_IntegrationTestTask_LoginSuccess() {
	// seed username and password
	stmt := "INSERT INTO users(id, username, password) VALUES (1, 'nohattee', '1qaz@WSX');"
	s.dbConn.ExecContext(context.Background(), stmt)

	reqStr := `{"username": "nohattee", "password": "1qaz@WSX"}`
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/login", s.port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	response.Body.Close()
}

func (s *IntegrationTestTaskSuite) Test_IntegrationTestTask_LoginError() {
	// seed username and password
	stmt := "INSERT INTO users(id, username, password) VALUES (1, 'nohattee', '1qaz@WSX');"
	s.dbConn.ExecContext(context.Background(), stmt)

	reqStr := `{"username": "wrong_username", "password": "1qaz@WSX"}`
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/login", s.port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusUnauthorized, response.StatusCode)

	response.Body.Close()
}

func (s *IntegrationTestTaskSuite) Test_IntegrationTestTask_ListTasksSuccess() {
	tx, err := s.dbConn.Begin()
	s.Require().NoError(err)

	// seed username and password
	stmt := "INSERT INTO users(id, username, password) VALUES (1, 'nohattee', '1qaz@WSX');"
	s.dbConn.ExecContext(context.Background(), stmt)

	// get token
	storeRepo := postgres.NewPostgresRepository(s.dbConn)
	userUseCase := user.NewUserUseCase(storeRepo)
	token, err := userUseCase.GenerateToken(1, 5)
	s.Require().NoError(err)

	// seed tasks
	layout := "2006-01-02"
	currentDate := time.Now().Format(layout)
	listTaskIds := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}
	tasks := []*storages.Task{
		{
			ID:          listTaskIds[0],
			Content:     "This is just the first test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{

			ID:          listTaskIds[1],
			Content:     "This is just the second test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{
			ID:          listTaskIds[2],
			Content:     "This is just the third test",
			UserID:      1,
			CreatedDate: currentDate,
		},
	}

	for _, t := range tasks {
		stmt := `INSERT INTO tasks(id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
		_, err := tx.ExecContext(context.TODO(), stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			tx.Rollback()
			s.Require().NoError(err)
		}
	}

	tx.Commit()

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/tasks", s.port), nil)
	s.NoError(err)

	q := req.URL.Query()
	q.Add("created_date", currentDate)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	data := map[string]interface{}{}
	err = json.Unmarshal(byteBody, &data)
	s.NoError(err)

	s.Equal(len(data["data"].([]interface{})), 5)
	response.Body.Close()
}

func (s *IntegrationTestTaskSuite) Test_IntegrationTestTask_AddTasksWhenNotReachedLimit() {
	tx, err := s.dbConn.Begin()
	s.Require().NoError(err)

	// seed username and password
	stmt := "INSERT INTO users(id, username, password) VALUES (1, 'nohattee', '1qaz@WSX');"
	s.dbConn.ExecContext(context.Background(), stmt)

	// get token
	storeRepo := postgres.NewPostgresRepository(s.dbConn)
	userUseCase := user.NewUserUseCase(storeRepo)
	token, err := userUseCase.GenerateToken(1, 5)
	s.Require().NoError(err)

	// seed tasks
	layout := "2006-01-02"
	currentDate := time.Now().Format(layout)
	listTaskIds := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}
	tasks := []*storages.Task{
		{
			ID:          listTaskIds[0],
			Content:     "This is just the first test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{

			ID:          listTaskIds[1],
			Content:     "This is just the second test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{
			ID:          listTaskIds[2],
			Content:     "This is just the third test",
			UserID:      1,
			CreatedDate: currentDate,
		},
	}

	for _, t := range tasks {
		stmt := `INSERT INTO tasks(id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
		_, err := tx.ExecContext(context.TODO(), stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			tx.Rollback()
			s.Require().NoError(err)
		}
	}

	tx.Commit()

	reqStr := `{"content": "This is just for the test"}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/tasks", s.port), strings.NewReader(reqStr))
	s.NoError(err)

	q := req.URL.Query()
	q.Add("created_date", currentDate)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	data := map[string]interface{}{}
	err = json.Unmarshal(byteBody, &data)
	s.NoError(err)
	byteResp, err := json.Marshal(data["data"])
	s.NoError(err)

	taskResp := storages.Task{}
	err = json.Unmarshal(byteResp, &taskResp)
	s.NoError(err)

	s.Equal(taskResp.Content, "This is just for the test")
	response.Body.Close()
}

func (s *IntegrationTestTaskSuite) Test_IntegrationTestTask_AddTasksWhenReachedLimited() {
	tx, err := s.dbConn.Begin()
	s.Require().NoError(err)

	// seed username and password
	stmt := "INSERT INTO users(id, username, password) VALUES (1, 'nohattee', '1qaz@WSX');"
	s.dbConn.ExecContext(context.Background(), stmt)

	// get token
	storeRepo := postgres.NewPostgresRepository(s.dbConn)
	userUseCase := user.NewUserUseCase(storeRepo)
	token, err := userUseCase.GenerateToken(1, 5)
	s.Require().NoError(err)

	// seed tasks
	layout := "2006-01-02"
	currentDate := time.Now().Format(layout)
	listTaskIds := []string{
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
	}

	tasks := []*storages.Task{
		{
			ID:          listTaskIds[0],
			Content:     "This is just the first test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{

			ID:          listTaskIds[1],
			Content:     "This is just the second test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{
			ID:          listTaskIds[2],
			Content:     "This is just the third test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{
			ID:          listTaskIds[3],
			Content:     "This is just the four test",
			UserID:      1,
			CreatedDate: currentDate,
		},
		{
			ID:          listTaskIds[4],
			Content:     "This is just the five test",
			UserID:      1,
			CreatedDate: currentDate,
		},
	}

	for _, t := range tasks {
		stmt := `INSERT INTO tasks(id, content, user_id, created_date) VALUES ($1, $2, $3, $4)`
		_, err := tx.ExecContext(context.TODO(), stmt, &t.ID, &t.Content, &t.UserID, &t.CreatedDate)
		if err != nil {
			tx.Rollback()
			s.Require().NoError(err)
		}
	}

	tx.Commit()

	reqStr := `{"content": "This is just for the test"}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/tasks", s.port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusBadRequest, response.StatusCode)

	response.Body.Close()
}
