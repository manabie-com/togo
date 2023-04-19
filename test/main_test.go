package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"togo/pkg/db_client"
	"togo/pkg/models"
	"togo/pkg/services"
	"togo/pkg/utils"
)

var serverTest = SetupTestDatabase()

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestRegisterAccount(t *testing.T) {
	//serverTest := SetupTestDatabase()
	r := SetUpRouter()
	r.POST("/auth/register", serverTest.RegisterAccount)
	var jsonStr = []byte(`{"UserName":"admin", "passWord": "123456"}`)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonStr))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code == http.StatusConflict {
		t.Log("Register account by userName is exits.")
		return
	}
	if res.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, res.Code)
		return
	}
	bodyRes, _ := ioutil.ReadAll(res.Body)
	var registerResult models.Response[models.LoginResponse]
	err := json.Unmarshal(bodyRes, &registerResult)
	if err != nil {
		t.Errorf("Error parse body response register account.")
		return
	}
	t.Log("Register account success")
}

func TestLoginAccount(t *testing.T) {
	//serverTest := SetupTestDatabase()
	r := SetUpRouter()
	r.POST("/auth/login", serverTest.LoginAccount)
	var jsonStr = []byte(`{"UserName":"admin", "passWord": "123456"}`)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonStr))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code == http.StatusNotFound {
		t.Errorf("User not found.")
		return
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, res.Code)
		return
	}

	bodyRes, _ := ioutil.ReadAll(res.Body)
	var result models.LoginResponse
	err := json.Unmarshal(bodyRes, &result)

	if err != nil {
		t.Errorf("Error parse body response login")
		return
	}
	if result.Status != http.StatusOK {
		t.Errorf("Login fail, message = " + result.Message)
		return
	}
	if result.Token == "" {
		t.Errorf("Create token by login fail")
		return
	}
	t.Log("Token login success: " + result.Token)
}

func LoginGetToken() models.LoginResponse {
	//serverTest := SetupTestDatabase()
	r := SetUpRouter()
	r.POST("/auth/login", serverTest.LoginAccount)
	var jsonStr = []byte(`{"UserName":"admin", "passWord": "123456"}`)
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonStr))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	if res.Code == http.StatusNotFound {
		return models.LoginResponse{}
	}
	if res.Code != http.StatusOK {
		return models.LoginResponse{}
	}
	bodyRes, _ := ioutil.ReadAll(res.Body)
	var result models.LoginResponse
	json.Unmarshal(bodyRes, &result)
	return result
}

func TestCreateTask(t *testing.T) {
	token := LoginGetToken().Token
	//serverTest := SetupTestDatabase()
	r := SetUpRouter()

	r.Use(serverTest.Validate)
	r.POST("/task/create", serverTest.CreateTask)

	var jsonCreate = []byte(`{"content":"Task test danh 4"}`)
	req, _ := http.NewRequest("POST", "/task/create", bytes.NewBuffer(jsonCreate))
	req.Header.Set("Authorization", "Bearer "+token)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code == http.StatusUnauthorized {
		t.Errorf("Token is expire or invalid.")
		return
	}

	if res.Code == http.StatusBadRequest {
		t.Log("User limit task.")
		return
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, res.Code)
		return
	}

	bodyRes, _ := ioutil.ReadAll(res.Body)
	var createTaskResul models.Response[models.Task]
	err := json.Unmarshal(bodyRes, &createTaskResul)
	if err != nil {
		t.Errorf("Error parse body response create task.")
		return
	}

	if createTaskResul.Status != 200 {
		t.Errorf("Create task fail, message = " + createTaskResul.Message)
		return
	}
	t.Log("Create task success")
}

func TestGetListTask(t *testing.T) {
	token := LoginGetToken().Token
	//serverTest := SetupTestDatabase()
	r := SetUpRouter()

	r.Use(serverTest.Validate)
	r.GET("/task/list", serverTest.GetTaskByDate)
	var createDateNow = time.Now().Format("2006-01-02")
	req, _ := http.NewRequest("GET", "/task/list?createdDate="+createDateNow, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	if res.Code == http.StatusUnauthorized {
		t.Errorf("Token is expire or invalid.")
		return
	}

	if res.Code != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, res.Code)
		return
	}

	bodyRes, _ := ioutil.ReadAll(res.Body)
	var createTaskResul models.Response[models.Task]
	err := json.Unmarshal(bodyRes, &createTaskResul)
	if err != nil {
		t.Errorf("Error parse body response create task.")
		return
	}

	if createTaskResul.Status != 200 {
		t.Errorf("Get list data fail, message" + createTaskResul.Message)
		return
	}
	t.Log("Get list create success")
}

func SetupTestDatabase() services.Server {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
	}

	dbContainer, _ := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")

	// 3.2 Create db connection string and connect
	var url = fmt.Sprintf("postgres://postgres:postgres@%v:%v/testdb", host, port.Port())
	dbTest := db_client.Init(url)
	dbTest.DB.AutoMigrate(&models.User{})
	dbTest.DB.AutoMigrate(&models.Task{})
	jwtTest := utils.JwtWrapper{
		SecretKey:       "key-test",
		Issuer:          "manibie-todo-test",
		ExpirationHours: 24 * 60,
	}

	redisRequest := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	redisServer, _ := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: redisRequest,
		Started:          true,
	})

	hostRedis, _ := redisServer.Host(context.Background())

	redisTest := redis.NewClient(&redis.Options{Addr: hostRedis + ":6379", DB: 0})

	serverTest := services.Server{
		H:     dbTest,
		Jwt:   jwtTest,
		Redis: redisTest,
	}
	return serverTest
}
