package handlers

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/manabie-com/togo/internal/driver"

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
