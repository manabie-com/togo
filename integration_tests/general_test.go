package integration_tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/internal/app/task/model"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/router"
	"github.com/manabie-com/togo/internal/util"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open(util.DriverPostgres, config.Postgres{
		Host:     getEnvVarDefault("TOGO_DB_POSTGRES_HOST", "localhost"),
		Port:     getEnvVarDefault("TOGO_DB_POSTGRES_PORT", "5432"),
		Name:     getEnvVarDefault("TOGO_DB_POSTGRES_TEST_DB_NAME", "integration_tests"),
		UserName: getEnvVarDefault("TOGO_DB_POSTGRES_USERNAME", "root"),
		Password: getEnvVarDefault("TOGO_DB_POSTGRES_PASSWORD", "password"),
		SSLMode:  "disable",
	}.ToConnectionString())
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile("seed.sql")
	if err != nil {
		return nil, err
	}
	stmts := strings.Split(string(b), ";")
	for _, stmt := range stmts {
		_, err = db.Exec(stmt)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func getEnvVarDefault(name, defaultValue string) string {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}
	return v
}

func prepareServer(db *sql.DB) *httptest.Server {
	r := router.NewRouter(db, util.DriverPostgres, getEnvVarDefault("TOGO_JWT_KEY", "wqGyEBBfPK9w3Lxw"))
	ts := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		r.ServeHTTP(writer, request)
	}))
	return ts
}

func generateLoginRequest(endpoint, userID, password string) (*http.Request, error) {
	r, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	q := r.URL.Query()
	q.Add("user_id", userID)
	q.Add("password", password)
	r.URL.RawQuery = q.Encode()
	return r, nil
}

func login(endpoint, userID, password string) (string, error) {
	r, err := generateLoginRequest(endpoint, userID, password)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var data struct {
		Data string `json:"data"`
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return "", err
	}
	return data.Data, nil
}

func generateAddTaskRequest(endpoint, token string) (*http.Request, error) {
	var b []byte
	body := bytes.NewBuffer(b)
	d, err := json.Marshal(map[string]string{
		"content": "something like this",
	})
	if err != nil {
		return nil, err
	}
	body.Write(d)
	r, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Authorization", token)
	return r, nil
}

func generateRetrieveTasksRequest(endpoint, token string) (*http.Request, error) {
	r, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	q := r.URL.Query()
	q.Add("created_date", time.Now().Format("2006-01-02"))
	r.URL.RawQuery = q.Encode()
	r.Header.Set("Authorization", token)
	return r, nil
}

func retrieveTasks(endpoint, token string) ([]model.Task, error) {
	r, err := generateRetrieveTasksRequest(endpoint, token)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data struct {
		Data []model.Task `json:"data"`
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil
}
