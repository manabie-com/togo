package services

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/manabie-com/togo/internal/config"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CacheTestSuite struct {
	suite.Suite
	redisPool *redis.Pool
	cfg *config.Config
	cache ICache
}

func (s *CacheTestSuite) SetupSuite() {
	m, err := miniredis.Run()
	s.NoError(err)

	s.cfg = &config.Config{
		Redis:       &config.Redis{
			Address:                m.Addr(),
			Password:               "",
			DatabaseNum:            0,
			MaxIdle:                3,
			MaxActive:              0,
			MaxIdleTimeout:         300,
			Wait:                   false,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
		},
	}
	s.redisPool = newRedisPool(s.cfg)
	s.cache = &RedisCache{redisPool: s.redisPool}
}

func (s *CacheTestSuite) TearDownSuite() {
	s.redisPool.Close()
}

func (s *CacheTestSuite) Test_1_GetNumberOfTasks_KeyDoesNotExist() {
	result, err := s.cache.GetNumberOfTasks("firstUser", "01-01-2021")
	s.NoError(err)
	s.Equal(int32(-1), result)
}

func (s *CacheTestSuite) Test_2_IncTask() {
	userId := "firstUser"
	createdDate := "01-01-2021"
	err := s.cache.IncTask(userId, createdDate)
	s.NoError(err)
	result, err := s.cache.GetNumberOfTasks(userId, createdDate)
	s.NoError(err)
	s.Equal(int32(1), result)
}

func (s *CacheTestSuite) Test_3_SetNumberOfTask() {
	userId := "firstUser"
	createdDate := "01-01-2021"
	err := s.cache.SetNumberOfTasks(userId, createdDate, 4)
	s.NoError(err)
	result, err := s.cache.GetNumberOfTasks(userId, createdDate)
	s.NoError(err)
	s.Equal(int32(4), result)
}

func (s *CacheTestSuite) Test_4_GetMaxToDo_keyDoesNotExist() {
	userId := "firstUser"
	maxTodo, err := s.cache.GetMaxTodo(userId)
	s.NoError(err)
	s.Equal(int32(-1), maxTodo)
}

func (s *CacheTestSuite) Test_5_SetMaxToDo() {
	userId := "firstUser"
	err := s.cache.SetMaxTodo(userId, 5)
	s.NoError(err)
	maxTodo, err := s.cache.GetMaxTodo(userId)
	s.NoError(err)
	s.Equal(int32(5), maxTodo)
}

func TestNewRedisPool(t *testing.T) {
	m, err := miniredis.Run()
	require.NoError(t, err)

	cfg := &config.Config{
		Redis:       &config.Redis{
			Address:                m.Addr(),
			Password:               "",
			DatabaseNum:            0,
			MaxIdle:                3,
			MaxActive:              0,
			MaxIdleTimeout:         300,
			Wait:                   false,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
		},
	}
	pool := newRedisPool(cfg)
	conn := pool.Get()
	data, err := conn.Do("PING")
	require.NoError(t, err)
	require.Equal(t, "PONG", data.(string))
	pool.Close()
}

func TestNewRedisPool_WithPassword(t *testing.T) {
	m, err := miniredis.Run()
	require.NoError(t, err)
	m.RequireAuth("12345")
	cfg := &config.Config{
		Redis:       &config.Redis{
			Address:                m.Addr(),
			Password:               "12345",
			DatabaseNum:            0,
			MaxIdle:                3,
			MaxActive:              0,
			MaxIdleTimeout:         300,
			Wait:                   false,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
		},
	}
	pool := newRedisPool(cfg)
	conn := pool.Get()
	data, err := conn.Do("PING")
	require.NoError(t, err)
	require.Equal(t, "PONG", data.(string))
	pool.Close()
}

func TestGetNumberFromRedis(t *testing.T) {
	result, err := getNumberFromRedis([]byte("1"))
	require.NoError(t, err)
	require.Equal(t, int32(1), result)
}

func TestGetNumberFromRedis_Nil(t *testing.T) {
	result, err := getNumberFromRedis(nil)
	require.NoError(t, err)
	require.Equal(t, int32(-1), result)
}

func TestGetNumberFromRedis_CannotCast(t *testing.T) {
	result, err := getNumberFromRedis(1)
	require.Equal(t, ErrCastValue, err)
	require.Equal(t, int32(-1), result)
}

func TestGetNumberFromRedis_EmptyString(t *testing.T) {
	result, err := getNumberFromRedis([]byte(""))
	require.NoError(t, err)
	require.Equal(t, int32(-1), result)
}

func TestGetNumberFromRedis_ConvertToIntError(t *testing.T) {
	result, err := getNumberFromRedis([]byte("abc"))
	require.Error(t, err)
	require.Equal(t, int32(-1), result)
}

func TestRedisCache(t *testing.T) {
	suite.Run(t, new(CacheTestSuite))
}


