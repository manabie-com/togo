// go:build ,!e2e_test

package tasks

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"manabie/todo/models"
	"manabie/todo/pkg/db"
	"manabie/todo/pkg/utils"
	"manabie/todo/repository/setting"
	"manabie/todo/repository/user"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {

	if !flag.Parsed() {
		flag.Parse()
	}

	if err := db.Setup(); err != nil {
		panic(err)
	}

	exit := m.Run()
	if err := db.Teardown(); err != nil {
		panic(err)
	}
	os.Exit(exit)

}

func Test_E2E_Integration(t *testing.T) {
	ur := user.NewUserRespository()
	sr := setting.NewSettingRespository()

	// Init data test
	u := &models.User{
		ID:    utils.RamdomID(),
		Email: "integration_test@test.com",
		Name:  "integration_test",
	}
	st := &models.Setting{
		MemberID:  u.ID,
		LimitTask: 5,
	}

	require.Nil(t, db.Transaction(context.Background(), nil, func(ctx context.Context, tx *sql.Tx) error {
		if err := ur.Create(ctx, tx, u); err != nil {
			return err
		}
		return sr.Create(ctx, tx, st)
	}))

	// Create
	{
		tk := &models.TaskCreateRequest{
			Content: "test_main 1", TargetDate: "2022-07-17",
		}

		data, _ := json.Marshal(tk)
		res, err := http.Post(fmt.Sprintf("http://localhost:8080/users/%d/tasks", u.ID), "application/json", bytes.NewBuffer(data))

		require.Nil(t, err)
		require.NotNil(t, res)

		assert.Equal(t, res.StatusCode, 200)
	}
	// Index
	var taskID int
	{
		r, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%d/tasks?date=%s", u.ID, "2022-07-17"))

		require.Nil(t, err)
		require.NotNil(t, r)

		defer r.Body.Close()

		assert.Equal(t, r.StatusCode, 200)

		body, err := ioutil.ReadAll(r.Body)
		require.Nil(t, err)

		var res models.TaskIndexResponse

		if assert.NoError(t, json.Unmarshal(body, &res)) {
			assert.Len(t, res.Tasks, 1)
			taskID = res.Tasks[0].ID
		}
	}
	// Show
	var tk *models.Task
	{
		r, err := http.Get(fmt.Sprintf("http://localhost:8080/tasks/%d", taskID))

		require.Nil(t, err)
		require.NotNil(t, r)

		defer r.Body.Close()

		assert.Equal(t, r.StatusCode, 200)

		body, err := ioutil.ReadAll(r.Body)
		require.Nil(t, err)

		assert.NoError(t, json.Unmarshal(body, &tk))
	}
	// Update
	{
		data, _ := json.Marshal(tk)

		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/tasks/%d", tk.ID), bytes.NewBuffer(data))
		require.Nil(t, err)

		// set the request header Content-Type for json
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		res, err := client.Do(req)

		require.Nil(t, err)
		require.NotNil(t, res)

		assert.Equal(t, res.StatusCode, 200)
	}
	// Delete
	{
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/tasks/%d", tk.ID), nil)
		require.Nil(t, err)

		// set the request header Content-Type for json
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		res, err := client.Do(req)

		require.Nil(t, err)
		require.NotNil(t, res)

		assert.Equal(t, res.StatusCode, 200)
	}
}

func Test_E2E_Create_Concurrency(t *testing.T) {

	ur := user.NewUserRespository()
	sr := setting.NewSettingRespository()

	// Init data test
	u := &models.User{
		ID:    utils.RamdomID(),
		Email: "concurrency_test@test.com",
		Name:  "concurrency_test",
	}
	st := &models.Setting{
		MemberID:  u.ID,
		LimitTask: 1,
	}

	require.Nil(t, db.Transaction(context.Background(), nil, func(ctx context.Context, tx *sql.Tx) error {
		if err := ur.Create(ctx, tx, u); err != nil {
			return err
		}
		return sr.Create(ctx, tx, st)
	}))

	vas := []models.TaskCreateRequest{
		{Content: "test_main 1", TargetDate: "2022-07-17"},
		{Content: "test_main 2", TargetDate: "2022-07-17"},
	}

	rltRes := make(chan *http.Response)
	rltErr := make(chan error)

	for i := range vas {
		go func(tk models.TaskCreateRequest, memeberID int) {
			data, _ := json.Marshal(tk)

			res, err := http.Post(fmt.Sprintf("http://localhost:8080/users/%d/tasks", memeberID), "application/json", bytes.NewBuffer(data))

			rltErr <- err
			rltRes <- res

		}(vas[i], u.ID)
	}

	var code200, code400 int

	for i := 0; i <= 1; i++ {
		err := <-rltErr
		require.Nil(t, err)

		res := <-rltRes
		if res.StatusCode == http.StatusOK {
			code200 = res.StatusCode
		} else {
			code400 = res.StatusCode
		}
	}

	assert.Equal(t, code200, 200)
	assert.Equal(t, code400, 400)
}
