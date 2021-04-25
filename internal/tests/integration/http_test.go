package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"net/url"
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/storages/psql"
	"github.com/stretchr/testify/assert"

	"github.com/kelseyhightower/envconfig"
	"github.com/manabie-com/togo/internal/transport/http"
	"github.com/manabie-com/togo/internal/usecase"

	"github.com/labstack/echo/v4"
)

var (
	conf struct {
		Transport http.APIConf
		AuthUC    usecase.AuthUCConf
		Storage   psql.Config
		Port      int
	}
	fakeUser = domain.User{
		ID:             "admin",
		Password:       "adminadminadmin",
		MaxTasksPerDay: 10,
	}
)

func init() {
	panicIfErr(envconfig.Process("integration", &conf))
}

func TestHttpTransport(t *testing.T) {
	e := echo.New()
	storage, err := psql.NewStorage(conf.Storage)
	panicIfErr(err)
	defer func() {
		err := storage.CleanupDB()
		panicIfErr(err)
	}()
	taskUc := usecase.NewTaskUseCase(storage, storage)
	authUc, err := usecase.NewAuthUseCase(conf.AuthUC, storage)
	assert.NoError(t, err)
	http.BindAPI(conf.Transport, e, taskUc, authUc)
	go func() {
		panicIfErr(e.Start(fmt.Sprintf(":%d", conf.Port)))
	}()
	serverAddr := fmt.Sprintf("http://localhost:%d", conf.Port)
	jwtToken := ""
	t.Run("create user", func(t *testing.T) {
		assert.NoError(t, authUc.CreateUser(fakeUser.ID, fakeUser.Password))
	})
	t.Run("login", func(t *testing.T) {
		postReq, _ := json.Marshal(map[string]string{
			"user_id":  fakeUser.ID,
			"password": fakeUser.Password,
		})
		resp, err := nethttp.Post(fmt.Sprintf("%s/login", serverAddr), "application/json", bytes.NewBuffer(postReq))
		assert.NoError(t, err)
		defer resp.Body.Close()
		container := make(map[string]string)
		dec := json.NewDecoder(resp.Body)
		assert.NoError(t, dec.Decode(&container))
		token, exist := container["data"]
		assert.True(t, exist)
		jwtToken = token
		assert.NotEmpty(t, jwtToken)
	})
	t.Run("tasks", func(t *testing.T) {
		client := &nethttp.Client{}
		postReq, _ := json.Marshal(map[string]string{
			"content": "hello world",
		})
		req, err := nethttp.NewRequest("POST", fmt.Sprintf("%s/tasks", serverAddr), bytes.NewBuffer(postReq))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
		todate := time.Now().Format(domain.DateFormat)
		isMidnightTest := false
		taskCreated := map[string]bool{}
		addTaskSuccess := func(t *testing.T) {
			for try := 0; try < fakeUser.MaxTasksPerDay; try++ {
				res, err := client.Do(req)
				assert.NoError(t, err)
				defer res.Body.Close()

				container := struct {
					Data struct {
						ID          string
						UserID      string
						Content     string
						CreatedDate string
					}
				}{}

				dec := json.NewDecoder(res.Body)
				assert.NoError(t, dec.Decode(&container))
				assert.Equal(t, "hello world", container.Data.Content)
				assert.Equal(t, fakeUser.ID, container.Data.UserID)
				taskCreated[container.Data.ID] = true
				//midnight test
				if container.Data.CreatedDate != todate {
					if isMidnightTest {
						assert.FailNowf(t, "", "task created has different creation date before request: %s and %s", todate, container.Data.CreatedDate)
					}
					if time.Now().Format(domain.DateFormat) != todate {
						isMidnightTest = true
						return
					}
					assert.FailNowf(t, "", "task created has different creation date before request: %s and %s", todate, container.Data.CreatedDate)
				}
			}
		}
		t.Run("add task success", addTaskSuccess)
		queryDate := todate
		if isMidnightTest {
			t.Run("add task success", addTaskSuccess)
			//list tasks on new date
			newDate := time.Now().Format(domain.DateFormat)
			queryDate = newDate

		}
		t.Run("list created tasks", func(t *testing.T) {
			client := &nethttp.Client{}
			url := url.URL{
				Scheme: "http",
				Host:   fmt.Sprintf("localhost:%d", conf.Port),
				Path:   "/tasks",
				RawQuery: url.Values{
					"created_date": []string{queryDate},
				}.Encode(),
			}
			req, err := nethttp.NewRequest("GET", url.String(), nil)
			assert.NoError(t, err)
			req.Header.Set("Content-type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
			res, err := client.Do(req)
			assert.NoError(t, err)
			defer res.Body.Close()

			type TaskContainer struct {
				ID          string
				UserID      string
				Content     string
				CreatedDate string
			}
			var container struct {
				Data []TaskContainer
			}
			container.Data = []TaskContainer{}
			dec := json.NewDecoder(res.Body)
			assert.NoError(t, dec.Decode(&container))
			assert.Equal(t, fakeUser.MaxTasksPerDay, len(container.Data))
			for _, item := range container.Data {
				_, exist := taskCreated[item.ID]
				assert.True(t, exist)
				assert.Equal(t, "hello world", item.Content)
			}
		})
		t.Run("add task fail", func(t *testing.T) {
			res, err := client.Do(req)
			assert.NoError(t, err)
			defer res.Body.Close()
			//midnight test,rerun test
			if time.Now().Format(domain.DateFormat) != todate {
				t.Run("add task success", addTaskSuccess)
				res, err = client.Do(req)
				assert.NoError(t, err)
				defer res.Body.Close()
			}
			assert.Equal(t, res.StatusCode, nethttp.StatusConflict)
		})
	})
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
