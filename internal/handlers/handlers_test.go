package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/driver"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repository/postgres"
	"github.com/manabie-com/togo/internal/usecases/user"

	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/suite"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type IntegrationTestSuite struct {
	suite.Suite
	port        int
	dbConn      *driver.DB
	dbMigration *migrate.Migrate
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, &IntegrationTestSuite{})
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.port = appPort
	dbConn, err := driver.ConnectDB(dsn)
	s.Require().NoError(err)
	s.dbConn = dbConn

	s.dbMigration, err = migrate.New("file://../../db/migrations", dsn)
	s.Require().NoError(err)
	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}

	os.Setenv("JWT_KEY", "eyJhbGciOiJIUzI1NiJ9")
	os.Setenv("JWT_TIMEOUT", "1000000")

	// Set up a test server
	repo := NewRepo(dbConn)
	SetRepoForHandlers(repo)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: routes(),
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Fatalf("error listening at port %d: %+s", s.port, err)
		}
	}()
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.dbConn.SQL.Close()

	os.Unsetenv("JWT_KEY")
	os.Unsetenv("JWT_TIMEOUT")
}

func (s *IntegrationTestSuite) SetupTest() {
	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}

	// seed user for testing
	stmt := "INSERT INTO users(id, username, password) VALUES (1, 'khxingn', 'Qq@1234567');"
	s.dbConn.SQL.ExecContext(context.Background(), stmt)
}

func (s *IntegrationTestSuite) TearDownTest() {
	s.Require().NoError(s.dbMigration.Down())
}

func (s *IntegrationTestSuite) TestIntegration_Login_Success() {
	reqBodyStr := `{"username": "khxingn", "password": "Qq@1234567"}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/login", s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()

	s.Require().Equal(response.StatusCode, http.StatusOK)
}

func (s *IntegrationTestSuite) TestIntegration_Login_Error() {
	reqBodyStr := `{"username": "some_username", "password": "Qq@1234567"}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/login", s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()

	s.Require().Equal(http.StatusUnauthorized, response.StatusCode)
}

func (s *IntegrationTestSuite) TestIntegration_RetrieveTasks_Success() {
	tx, err := s.dbConn.SQL.Begin()
	s.Require().NoError(err)

	// get token
	repository := postgres.NewPostgresRepository(s.dbConn.SQL)
	userUsecase := user.NewUserUsecase(repository)
	token, err := userUsecase.GenerateToken(uint(1), uint(5))
	s.Require().NoError(err)

	currentDate := time.Now().Format("2006-01-02")
	// prepare tasks for test
	{
		taskIds := []string{uuid.New().String(), uuid.New().String(), uuid.New().String()}
		tasks := []*models.Task{
			{
				ID:          taskIds[0],
				Detail:      "first task",
				UserID:      uint(1),
				CreatedDate: currentDate,
			},
			{
				ID:          taskIds[1],
				Detail:      "second task",
				UserID:      uint(1),
				CreatedDate: currentDate,
			},
			{
				ID:          taskIds[2],
				Detail:      "third task",
				UserID:      uint(1),
				CreatedDate: currentDate,
			},
		}

		for _, t := range tasks {
			stmt := `INSERT INTO tasks(id, detail, user_id, created_date) VALUES ($1, $2, $3, $4)`
			_, err := tx.ExecContext(context.Background(), stmt, &t.ID, &t.Detail, &t.UserID, &t.CreatedDate)
			if err != nil {
				tx.Rollback()
				s.Require().NoError(err)
			}
		}

		err = tx.Commit()
		s.Require().NoError(err)
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/tasks", s.port), nil)
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	q := req.URL.Query()
	q.Add("created_date", currentDate)
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()

	s.Require().Equal(response.StatusCode, http.StatusOK)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.Require().NoError(err)

	data := map[string]interface{}{}
	err = json.Unmarshal(byteBody, &data)
	s.Require().NoError(err)

	s.Require().Equal(len(data["data"].([]interface{})), 3)
}

func (s *IntegrationTestSuite) TestIntegration_AddTaskSuccess() {
	// get token
	repository := postgres.NewPostgresRepository(s.dbConn.SQL)
	userUsecase := user.NewUserUsecase(repository)
	token, err := userUsecase.GenerateToken(uint(1), uint(5))
	s.Require().NoError(err)

	reqBodyStr := `{"detail": "new task"}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/tasks", s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()
	s.Require().Equal(response.StatusCode, http.StatusOK)

	byteResBody, err := ioutil.ReadAll(response.Body)
	s.Require().NoError(err)

	data := map[string]interface{}{}
	err = json.Unmarshal(byteResBody, &data)
	s.Require().NoError(err)

	byteTask, err := json.Marshal(data["data"])
	s.Require().NoError(err)

	task := models.Task{}
	err = json.Unmarshal(byteTask, &task)
	s.Require().NoError(err)

	s.Require().Equal(task.Detail, "new task")

}

func (s *IntegrationTestSuite) TestIntegration_AddTaskFailDueToThreshHold() {
	tx, err := s.dbConn.SQL.Begin()
	s.Require().NoError(err)

	// get token
	resository := postgres.NewPostgresRepository(s.dbConn.SQL)
	userUsecase := user.NewUserUsecase(resository)
	token, err := userUsecase.GenerateToken(uint(1), uint(5))
	s.Require().NoError(err)

	currentDate := time.Now().Format("2006-01-02")
	{
		// prepare tasks for testing
		taskIds := []string{uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String(), uuid.New().String()}
		tasks := []*models.Task{
			{
				ID:          taskIds[0],
				Detail:      "first task",
				UserID:      uint(1),
				CreatedDate: currentDate,
			},
			{
				ID:          taskIds[1],
				Detail:      "second task",
				UserID:      uint(1),
				CreatedDate: currentDate,
			},
			{
				ID:          taskIds[2],
				Detail:      "third task",
				UserID:      uint(1),
				CreatedDate: currentDate,
			},
			{
				ID:          taskIds[3],
				Detail:      "fourth task",
				UserID:      uint(1),
				CreatedDate: currentDate,
			},
			{
				ID:          taskIds[4],
				Detail:      "fifth task",
				UserID:      uint(1),
				CreatedDate: currentDate,
			},
		}

		for _, t := range tasks {
			stmt := `INSERT INTO tasks(id, detail, user_id, created_date) VALUES ($1, $2, $3, $4)`
			_, err := tx.ExecContext(context.Background(), stmt, &t.ID, &t.Detail, &t.UserID, &t.CreatedDate)
			if err != nil {
				tx.Rollback()
				s.Require().NoError(err)
			}
		}

		err = tx.Commit()
		s.Require().NoError(err)
	}

	reqBodyStr := `{"detail": "new task"}`
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/tasks", s.port), bytes.NewBuffer([]byte(reqBodyStr)))
	s.Require().NoError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(req)
	s.Require().NoError(err)
	defer response.Body.Close()
	s.Require().Equal(response.StatusCode, http.StatusBadRequest)
}
