package test

import (
	"ansidev.xyz/pkg/log"
	"github.com/ansidev/togo/gingo"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

func init() {
	log.InitLogger("console")
}

type ControllerTestSuite struct {
	suite.Suite
	PostgresTestSuite
	RedisTestSuite
	Router *gin.Engine
}

func (s *ControllerTestSuite) SetupSuite() {
	s.PostgresTestSuite.SetT(s.T())
	s.RedisTestSuite.SetT(s.T())

	s.PostgresTestSuite.SetupSuite()
	s.PostgresTestSuite.SetMigrationDir("../../../migration")
	s.RedisTestSuite.SetupSuite()
	s.Router = gingo.DefaultRouter()
}

func (s *ControllerTestSuite) BeforeTest(suite string, method string) {
	s.PostgresTestSuite.BeforeTest(suite, method)
}

func (s *ControllerTestSuite) AfterTest(suite string, method string) {
	s.PostgresTestSuite.AfterTest(suite, method)
	s.RedisTestSuite.AfterTest(suite, method)
}

func (s *ControllerTestSuite) TearDownSuite() {
	s.PostgresTestSuite.TearDownSuite()
	s.RedisTestSuite.TearDownSuite()
}
