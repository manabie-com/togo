package integration

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/trangmaiq/togo/internal/config"
	"github.com/trangmaiq/togo/internal/registry"
	"github.com/trangmaiq/togo/internal/server"
	"github.com/trangmaiq/togo/internal/server/handler"
)

var cfg *config.ToGo

func TestMain(m *testing.M) {
	cfg = config.Load()
	cfg.ServicePort = 9091
	err := registry.Init(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	go server.Start(cfg)
	time.Sleep(100 * time.Millisecond)

	os.Exit(m.Run())
}

func TestHandler_CreateTask(t *testing.T) {
	body :=
		`{
			"user_id": "c7cd294c-627f-452a-8c46-33b5dbfca47f",
			"title": "1st task",
			"note": "should have a better doc"
		}`

	res, err := http.Post(fmt.Sprintf("http://0.0.0.0:%d/tasks", cfg.ServicePort), echo.MIMEApplicationJSON, strings.NewReader(body))
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, http.StatusOK, res.StatusCode)

	var resBody handler.CreateTaskResponse
	err = json.NewDecoder(res.Body).Decode(&resBody)
	require.NoError(t, err)
	require.NotEmpty(t, resBody.ID)
}
