package test

import (
	"ansidev.xyz/pkg/log"
	"ansidev.xyz/pkg/rds"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func init() {
	log.InitLogger("console")
}

type RedisTestSuite struct {
	suite.Suite
	RedisClient *redis.Client
	Rdb         *rds.RedisDB
}

func (s *RedisTestSuite) initDbConnection(dbConfig rds.RedisConfig) {
	s.RedisClient = rds.NewRedisClient(dbConfig)
	s.Rdb = rds.NewRedisDB(context.Background(), s.RedisClient)
}

func (s *RedisTestSuite) SetupSuite() {
	dbConfig, err := GetTestRedisDbConfig()
	require.NoError(s.T(), err)

	s.initDbConnection(dbConfig)
}

func (s *RedisTestSuite) AfterTest(suite string, method string) {
	log.Info(fmt.Sprintf("Suite: %s, After running %s", suite, method))
	s.RedisClient.FlushDB(context.Background())
}

func (s *RedisTestSuite) TearDownSuite() {
	err := s.RedisClient.Close()
	require.NoError(s.T(), err)
}
