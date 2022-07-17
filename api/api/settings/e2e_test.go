// go:build ,!e2e_test

package settings

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

	// Init data test
	u := &models.User{
		ID:    utils.RamdomID(),
		Email: "integration_test@test.com",
		Name:  "integration_test",
	}
	require.Nil(t, db.Transaction(context.Background(), nil, func(ctx context.Context, tx *sql.Tx) error {
		return ur.Create(ctx, tx, u)
	}))

	// Create
	{
		got := &models.SettingCreateRequest{
			LimitTask: 10,
		}

		data, _ := json.Marshal(got)
		res, err := http.Post(fmt.Sprintf("http://localhost:8080/users/%d/settings", u.ID), "application/json", bytes.NewBuffer(data))

		require.Nil(t, err)
		require.NotNil(t, res)

		assert.Equal(t, res.StatusCode, 200)
	}
	// Show
	var st *models.Setting
	{
		r, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%d/settings", u.ID))

		require.Nil(t, err)
		require.NotNil(t, r)

		defer r.Body.Close()

		assert.Equal(t, r.StatusCode, 200)

		body, err := ioutil.ReadAll(r.Body)
		require.Nil(t, err)

		assert.NoError(t, json.Unmarshal(body, &st))
	}
	// Update
	{
		got := &models.SettingUpdateRequest{
			LimitTask: 1,
		}

		data, _ := json.Marshal(got)

		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/settings/%d", st.ID), bytes.NewBuffer(data))
		require.Nil(t, err)

		// set the request header Content-Type for json
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		res, err := client.Do(req)

		require.Nil(t, err)
		require.NotNil(t, res)

		assert.Equal(t, res.StatusCode, 200)
	}
}
