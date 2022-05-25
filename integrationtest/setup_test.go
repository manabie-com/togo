package integrationtest

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mmysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component/tokenprovider"
	"github.com/japananh/togo/server"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"
)

type e2eTestSuite struct {
	port            int
	dbConnectionStr string
	dbConn          *gorm.DB
	dbMigration     *migrate.Migrate
	suite.Suite
}

func (s *e2eTestSuite) getClientURL(path string) string {
	if path == "" {
		return fmt.Sprintf("http://localhost:%d/api/v1", s.port)
	}
	return fmt.Sprintf("http://localhost:%d/api/v1/%s", s.port, path)
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	// Load config from `.env` file
	config := common.NewConfig()
	s.Require().NoError(config.Load("../"))

	gormDB, err := gorm.Open(mysql.Open(config.DBConnectionURLTest()), &gorm.Config{})
	s.Require().NoError(err)

	sqlDB, err := sql.Open("mysql", config.DBConnectionURLTest())
	s.Require().NoError(err)
	driver, _ := mmysql.WithInstance(sqlDB, &mmysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations",
		"mysql",
		driver,
	)
	s.Require().NoError(err)

	s.port = config.AppPort()
	s.dbConnectionStr = config.DBConnectionURLTest()
	s.dbConn = gormDB
	s.dbMigration = m

	// Run migration up
	if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
		s.Require().NoError(err)
	}

	// create token configs
	tokenConfig, err := tokenprovider.NewTokenConfig(config.AtExpiry(), config.RtExpiry())
	s.Require().NoError(err)

	apiServer := server.Server{
		Port:        config.AppPort(),
		SecretKey:   config.SecretKey(),
		DBConn:      s.dbConn,
		TokenConfig: tokenConfig,
		ServerReady: make(chan bool),
	}

	go apiServer.Start()
	<-apiServer.ServerReady
	close(apiServer.ServerReady)

	// TODO: figure out create user at setup suite is best practice or not
	createUser(s)
}

func createUser(s *e2eTestSuite) string {
	req, err := http.NewRequest(
		http.MethodPost,
		s.getClientURL("register"),
		strings.NewReader(`{"email":"login@gmail.com", "password": "login@123"}`),
	)
	s.NoError(err)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	s.NoError(err)
	s.Equal(http.StatusOK, res.StatusCode)

	registerBody, err := ioutil.ReadAll(res.Body)
	s.NoError(err)
	s.NotNil(strings.Trim(string(registerBody), ""))

	return string(registerBody)
}

func (s *e2eTestSuite) TearDownSuite() {
	s.NoError(s.dbMigration.Down())
	p, _ := os.FindProcess(syscall.Getpid())
	s.Require().NoError(p.Signal(syscall.SIGINT))
}

func (s *e2eTestSuite) SetupTest() {
	//log.Println("---------- setup test -----------")
}

func (s *e2eTestSuite) TearDownTest() {
	//s.NoError(s.dbMigration.Down())
	//log.Println("________down test -----")
}
