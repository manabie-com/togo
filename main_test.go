package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/entities"
	dbPostgres "github.com/manabie-com/togo/internal/storages/postgres"
	dbRedis "github.com/manabie-com/togo/internal/storages/redis"
)

var service *services.ToDoService

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Load .env file failed. Use default config")
	}

	storeManager := dbPostgres.GetStorageManager(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
	)

	maxRequestPeHour, err := strconv.Atoi(os.Getenv("MAX_REQ_PER_HOUR"))
	if err != nil || maxRequestPeHour <= 0 {
		log.Fatalln("Invalid rate limit config")
	}
	rateLimiter := dbRedis.GetRateLimiter(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		maxRequestPeHour,
	)

	service = services.GetTodoService(os.Getenv("JWT_KEY"), storeManager, rateLimiter)

	m.Run()
}

func TestLogin(t *testing.T) {
	cleanDB()

	type args struct {
		username string
		password string
	}
	type wantResp struct {
		Code int
	}

	testCases := []struct {
		name     string
		args     args
		wantResp wantResp
	}{
		{
			name: "User is already existed, login success",
			args: args{
				username: "test_user_1",
				password: "p@ssworD!",
			},
			wantResp: wantResp{
				Code: http.StatusOK,
			},
		},
		{
			name: "User not found, login failed",
			args: args{
				username: "not_found_user",
				password: "dummy",
			},
			wantResp: wantResp{
				Code: http.StatusUnauthorized,
			},
		},
		{
			name: "User exist, wrong password, login failed",
			args: args{
				username: "test_user_1",
				password: "wr0n9p@ss",
			},
			wantResp: wantResp{
				Code: http.StatusUnauthorized,
			},
		},
	}

	// Setup DB
	err := service.Store.AddUser(context.Background(), "test_user_1", "p@ssworD!")
	if err != nil {
		t.Fatalf("Init database failed, err %s", err)
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/login?user_id=%s&password=%s", tt.args.username, tt.args.password), nil)
			response := httptest.NewRecorder()

			service.ServeHTTP(response, request)

			gotCode := response.Code
			if gotCode != tt.wantResp.Code {
				t.Errorf("Login failed. Expected http code %d, but got %d\n", tt.wantResp.Code, gotCode)
			}
		})
	}
}

func TestListTask(t *testing.T) {
	cleanDB()

	type args struct {
		username string
		password string
		date     string
	}
	type wantResp struct {
		Code  int
		Tasks []*entities.Task
	}

	testCases := []struct {
		name     string
		args     args
		wantResp wantResp
	}{
		{
			name: "valid token then get 2 task at date 2021-04-15",
			args: args{
				username: "test_user_1",
				password: "p@ssworD!",
				date:     "2021-04-15",
			},
			wantResp: wantResp{
				Code: http.StatusOK,
				Tasks: []*entities.Task{
					{
						ID:          "1",
						Content:     "task 1",
						UserID:      "test_user_1",
						CreatedDate: "2021-04-15",
					},
					{
						ID:          "2",
						Content:     "task 2",
						UserID:      "test_user_1",
						CreatedDate: "2021-04-15",
					},
				},
			},
		},
		{
			name: "Valid token but no task found",
			args: args{
				username: "test_user_1",
				password: "p@ssworD!",
				date:     "1111-12-12",
			},
			wantResp: wantResp{
				Code:  http.StatusOK,
				Tasks: []*entities.Task{},
			},
		},
	}

	// Setup DB
	err := service.Store.AddUser(context.Background(), "test_user_1", "p@ssworD!")
	if err != nil {
		t.Fatalf("Init database failed, err %s", err)
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			loginReq, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/login?user_id=%s&password=%s", tt.args.username, tt.args.password), nil)
			loginResp := httptest.NewRecorder()

			service.ServeHTTP(loginResp, loginReq)
			resp := make(map[string]interface{})
			if err := json.Unmarshal(loginResp.Body.Bytes(), &resp); err != nil {
				t.Errorf("Login failed, err %s", err)
			}

			token, ok := resp["data"].(string)
			if !ok {
				t.Errorf("Get token failed, got %v", resp)
			}

			for _, task := range tt.wantResp.Tasks {
				if err := service.Store.AddTask(context.Background(), task); err != nil {
					t.Errorf("insert task failed, err %s", err)
				}
			}

			taskReq, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks?created_date=%s", tt.args.date), nil)
			taskReq.Header.Set("Authorization", token)
			taskResp := httptest.NewRecorder()
			service.ServeHTTP(taskResp, taskReq)

			gotCode := taskResp.Code
			if gotCode != tt.wantResp.Code {
				t.Errorf("Get task failed. Expected http code %d, but got %d\n", tt.wantResp.Code, gotCode)
			}

			gotResp := map[string][]*entities.Task{}
			json.Unmarshal(taskResp.Body.Bytes(), &gotResp)
			if !reflect.DeepEqual(gotResp["data"], tt.wantResp.Tasks) {
				t.Errorf("Get task failed. Expected get %v, but got %v\n", tt.wantResp.Tasks, gotResp["data"])
			}
		})
	}
}

func TestCreateTask(t *testing.T) {
	cleanDB()

	type args struct {
		username string
		password string
		reqs     []string
	}
	type wantResp struct {
		Code  int
		Tasks []*entities.Task
	}

	testCases := []struct {
		name     string
		args     args
		wantResp wantResp
	}{
		{
			name: "create 3 task successful",
			args: args{
				username: "test_user_1",
				password: "p@ssworD!",
				reqs:     []string{"description 1", "description 2", "description 3"},
			},
			wantResp: wantResp{
				Code: http.StatusOK,
				Tasks: []*entities.Task{
					{
						UserID:      "test_user_1",
						Content:     "description 1",
						CreatedDate: time.Now().Format("2006-01-02"),
					},
					{
						UserID:      "test_user_1",
						Content:     "description 2",
						CreatedDate: time.Now().Format("2006-01-02"),
					},
					{
						UserID:      "test_user_1",
						Content:     "description 3",
						CreatedDate: time.Now().Format("2006-01-02"),
					},
				},
			},
		},
		{
			name: "create 6 task but got rate limited error",
			args: args{
				username: "test_user_1",
				password: "p@ssworD!",
				reqs: []string{
					"description 1", "description 2", "description 3",
					"description 4", "description 5", "description 6",
				},
			},
			wantResp: wantResp{
				Code: http.StatusTooManyRequests,
				Tasks: []*entities.Task{
					{
						UserID:      "test_user_1",
						Content:     "description 1",
						CreatedDate: time.Now().Format("2006-01-02"),
					},
					{
						UserID:      "test_user_1",
						Content:     "description 2",
						CreatedDate: time.Now().Format("2006-01-02"),
					},
					{
						UserID:      "test_user_1",
						Content:     "description 3",
						CreatedDate: time.Now().Format("2006-01-02"),
					},
					{
						UserID:      "test_user_1",
						Content:     "description 4",
						CreatedDate: time.Now().Format("2006-01-02"),
					},
					{
						UserID:      "test_user_1",
						Content:     "description 5",
						CreatedDate: time.Now().Format("2006-01-02"),
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			cleanDB()
			cleanCache()

			err := service.Store.AddUser(context.Background(), "test_user_1", "p@ssworD!")
			if err != nil {
				t.Fatalf("Init database failed, err %s", err)
			}

			loginReq, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/login?user_id=%s&password=%s", tt.args.username, tt.args.password), nil)
			loginResp := httptest.NewRecorder()

			service.ServeHTTP(loginResp, loginReq)
			resp := make(map[string]interface{})
			if err := json.Unmarshal(loginResp.Body.Bytes(), &resp); err != nil {
				t.Errorf("Login failed, err %s", err)
			}

			token, ok := resp["data"].(string)
			if !ok {
				t.Errorf("Get token failed, got %v", resp)
			}

			for _, task := range tt.args.reqs {
				taskReq, _ := http.NewRequest(
					http.MethodPost,
					"/tasks",
					strings.NewReader(fmt.Sprintf(`{"content":"%s"}`, task)),
				)
				taskReq.Header.Set("Authorization", token)
				taskResp := httptest.NewRecorder()
				service.ServeHTTP(taskResp, taskReq)

				gotCode := taskResp.Code
				if gotCode != http.StatusOK && gotCode != tt.wantResp.Code {
					t.Errorf("Get task failed. Expected http code %d, but got %d\n", tt.wantResp.Code, gotCode)
				}
			}

			gotTask, err := service.Store.RetrieveTasks(context.Background(),
				sql.NullString{
					String: "test_user_1",
					Valid:  true,
				}, sql.NullString{
					String: time.Now().Format("2006-01-02"),
					Valid:  true,
				},
			)

			if err != nil {
				t.Errorf("get tasks failed, err %s", err)
			}

			// Discard random id
			for i := range gotTask {
				gotTask[i].ID = ""
			}

			if !reflect.DeepEqual(gotTask, tt.wantResp.Tasks) {
				t.Errorf("get tasks failed, want %v but got %v", tt.wantResp.Tasks, gotTask)
			}
		})
	}
}

func cleanDB() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"))

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalln("Init DB connection failed", err)
	}
	db.Exec("DELETE FROM tasks WHERE 1=1")
	db.Exec("DELETE FROM users WHERE 1=1")
}

func cleanCache() {
	rdb := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		},
	)

	rdb.FlushDB(context.Background())
}
