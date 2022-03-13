package integration_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
	"github.com/triet-truong/todo/config"
	"github.com/triet-truong/todo/servers"
	"github.com/triet-truong/todo/todo/repository"
)

type EndToEndTestSuite struct {
	suite.Suite
	server *servers.Server
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndTestSuite))
}

func (s *EndToEndTestSuite) SetupSuite() {
	//Load config from env vars
	config.Load()

	// Setup repositories
	repo := repository.NewTodoMysqlRepository(config.DatabaseDSN())
	cacheStore := repository.NewTodoRedisRepository(redis.Options{
		Addr:     config.CacheConnectioURL(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//Start local server
	s.server = servers.NewServer(repo, cacheStore)
	s.server.Run()
}

func (s *EndToEndTestSuite) TearDownSuite() {
	s.server.Stop()
}

func (s *EndToEndTestSuite) TestWhenAddingTodoItem_ThenInsertedToDb() {
	s.HTTPSuccess(func(w http.ResponseWriter, r *http.Request) {

	}, http.MethodPost, fmt.Sprintf("localhost:%v/user/todo", config.Port()), nil, `{"content":"Hello world","user_id":1}`)
}
