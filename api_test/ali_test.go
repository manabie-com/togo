package api_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/manabie/project/config"
	todoHttp "github.com/manabie/project/internal/http"
	todoRepository "github.com/manabie/project/internal/repository"
	todoRouter "github.com/manabie/project/internal/router"
	todoUsecase "github.com/manabie/project/internal/usecase"
	"github.com/manabie/project/model"
	"github.com/manabie/project/pkg/hash"
	"github.com/manabie/project/pkg/jwt"
	"github.com/manabie/project/pkg/postgres"
	"github.com/manabie/project/pkg/snowflake"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	conf, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Println("can't connect to config file ", err.Error())
	}

	jwt := jwt.NewTokenUser(conf)
	snowflake := snowflake.NewSnowflake()
	hash := hash.NewHashPassword()

	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Postgres.PostgresqlHost,conf.Postgres.PostgresqlUser,conf.Postgres.PostgresqlPassword,
		conf.Postgres.PostgresqlDbname,conf.Postgres.PostgresqlPort)

	postgresConn:= postgres.NewPostgresConn(dsn)
	if err := postgresConn.AutoMigrate(&model.User{},&model.Task{}); err != nil {
		fmt.Println("Can't create table in database ", err.Error())
	}

	todoRepository := todoRepository.NewRepository(postgresConn)
	todoUsecase := todoUsecase.NewUsecase(todoRepository, jwt, hash)
	todoHttp := todoHttp.NewHttp(todoUsecase)
	todoRouter := todoRouter.NewRouter(todoHttp, snowflake)

	router.POST("localhost:5000/api/login", todoRouter.Login)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}

func TestSignup(t *testing.T) {
	w := httptest.NewRecorder()
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	conf, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Println("can't connect to config file ", err.Error())
	}

	jwt := jwt.NewTokenUser(conf)
	snowflake := snowflake.NewSnowflake()
	hash := hash.NewHashPassword()

	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Postgres.PostgresqlHost,conf.Postgres.PostgresqlUser,conf.Postgres.PostgresqlPassword,
		conf.Postgres.PostgresqlDbname,conf.Postgres.PostgresqlPort)

	postgresConn:= postgres.NewPostgresConn(dsn)
	if err := postgresConn.AutoMigrate(&model.User{},&model.Task{}); err != nil {
		fmt.Println("Can't create table in database ", err.Error())
	}

	todoRepository := todoRepository.NewRepository(postgresConn)
	todoUsecase := todoUsecase.NewUsecase(todoRepository, jwt, hash)
	todoHttp := todoHttp.NewHttp(todoUsecase)
	todoRouter := todoRouter.NewRouter(todoHttp, snowflake)

	router.POST("localhost:5000/api/register", todoRouter.SignUp)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}

func TestCreateTask(t *testing.T) {
	w := httptest.NewRecorder()
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	conf, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Println("can't connect to config file ", err.Error())
	}

	jwt := jwt.NewTokenUser(conf)
	snowflake := snowflake.NewSnowflake()
	hash := hash.NewHashPassword()

	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Postgres.PostgresqlHost,conf.Postgres.PostgresqlUser,conf.Postgres.PostgresqlPassword,
		conf.Postgres.PostgresqlDbname,conf.Postgres.PostgresqlPort)

	postgresConn:= postgres.NewPostgresConn(dsn)
	if err := postgresConn.AutoMigrate(&model.User{},&model.Task{}); err != nil {
		fmt.Println("Can't create table in database ", err.Error())
	}

	todoRepository := todoRepository.NewRepository(postgresConn)
	todoUsecase := todoUsecase.NewUsecase(todoRepository, jwt, hash)
	todoHttp := todoHttp.NewHttp(todoUsecase)
	todoRouter := todoRouter.NewRouter(todoHttp, snowflake)

	router.POST("localhost:5000/api/task/1469320290763804672", todoRouter.CreateTask)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}

func TestUpdateTask(t *testing.T) {
	w := httptest.NewRecorder()
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	conf, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Println("can't connect to config file ", err.Error())
	}

	jwt := jwt.NewTokenUser(conf)
	snowflake := snowflake.NewSnowflake()
	hash := hash.NewHashPassword()

	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Postgres.PostgresqlHost,conf.Postgres.PostgresqlUser,conf.Postgres.PostgresqlPassword,
		conf.Postgres.PostgresqlDbname,conf.Postgres.PostgresqlPort)

	postgresConn:= postgres.NewPostgresConn(dsn)
	if err := postgresConn.AutoMigrate(&model.User{},&model.Task{}); err != nil {
		fmt.Println("Can't create table in database ", err.Error())
	}

	todoRepository := todoRepository.NewRepository(postgresConn)
	todoUsecase := todoUsecase.NewUsecase(todoRepository, jwt, hash)
	todoHttp := todoHttp.NewHttp(todoUsecase)
	todoRouter := todoRouter.NewRouter(todoHttp, snowflake)

	router.PUT("localhost:5000/api/task/1469525483627483136", todoRouter.UpdateTask)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}

func TestDeleteTask(t *testing.T) {
	w := httptest.NewRecorder()
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	conf, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Println("can't connect to config file ", err.Error())
	}

	jwt := jwt.NewTokenUser(conf)
	snowflake := snowflake.NewSnowflake()
	hash := hash.NewHashPassword()

	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Postgres.PostgresqlHost,conf.Postgres.PostgresqlUser,conf.Postgres.PostgresqlPassword,
		conf.Postgres.PostgresqlDbname,conf.Postgres.PostgresqlPort)

	postgresConn:= postgres.NewPostgresConn(dsn)
	if err := postgresConn.AutoMigrate(&model.User{},&model.Task{}); err != nil {
		fmt.Println("Can't create table in database ", err.Error())
	}

	todoRepository := todoRepository.NewRepository(postgresConn)
	todoUsecase := todoUsecase.NewUsecase(todoRepository, jwt, hash)
	todoHttp := todoHttp.NewHttp(todoUsecase)
	todoRouter := todoRouter.NewRouter(todoHttp, snowflake)

	router.DELETE("localhost:5000/api/task/1469647784243105792", todoRouter.DeleteTask)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}

func TestTaskAll(t *testing.T) {
	w := httptest.NewRecorder()
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	conf, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Println("can't connect to config file ", err.Error())
	}

	jwt := jwt.NewTokenUser(conf)
	snowflake := snowflake.NewSnowflake()
	hash := hash.NewHashPassword()

	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Postgres.PostgresqlHost,conf.Postgres.PostgresqlUser,conf.Postgres.PostgresqlPassword,
		conf.Postgres.PostgresqlDbname,conf.Postgres.PostgresqlPort)

	postgresConn:= postgres.NewPostgresConn(dsn)
	if err := postgresConn.AutoMigrate(&model.User{},&model.Task{}); err != nil {
		fmt.Println("Can't create table in database ", err.Error())
	}

	todoRepository := todoRepository.NewRepository(postgresConn)
	todoUsecase := todoUsecase.NewUsecase(todoRepository, jwt, hash)
	todoHttp := todoHttp.NewHttp(todoUsecase)
	todoRouter := todoRouter.NewRouter(todoHttp, snowflake)

	router.GET("localhost:5000/api/task", todoRouter.TaskAll)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}

func TestTaskById(t *testing.T) {
	w := httptest.NewRecorder()
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	conf, err := config.ReadConf("config/config-docker.yml")
	if err != nil {
		log.Println("can't connect to config file ", err.Error())
	}

	jwt := jwt.NewTokenUser(conf)
	snowflake := snowflake.NewSnowflake()
	hash := hash.NewHashPassword()

	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.Postgres.PostgresqlHost,conf.Postgres.PostgresqlUser,conf.Postgres.PostgresqlPassword,
		conf.Postgres.PostgresqlDbname,conf.Postgres.PostgresqlPort)

	postgresConn:= postgres.NewPostgresConn(dsn)
	if err := postgresConn.AutoMigrate(&model.User{},&model.Task{}); err != nil {
		fmt.Println("Can't create table in database ", err.Error())
	}

	todoRepository := todoRepository.NewRepository(postgresConn)
	todoUsecase := todoUsecase.NewUsecase(todoRepository, jwt, hash)
	todoHttp := todoHttp.NewHttp(todoUsecase)
	todoRouter := todoRouter.NewRouter(todoHttp, snowflake)

	router.GET("localhost:5000/api/task/1469647787426582528", todoRouter.TaskById)

	t.Run("get json data", func(t *testing.T) {
		assert.Equal(t, 200, w.Code)
	})
}