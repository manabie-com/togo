// go:build ,!e2e_test

package users

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
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

	// Index
	{
		r, err := http.Get("http://localhost:8080/users")

		require.Nil(t, err)
		require.NotNil(t, r)

		defer r.Body.Close()

		assert.Equal(t, r.StatusCode, 200)

		body, err := ioutil.ReadAll(r.Body)
		require.Nil(t, err)

		var res models.UserIndexResponse

		if assert.NoError(t, json.Unmarshal(body, &res)) {
			assert.GreaterOrEqual(t, len(res.Users), 1)
		}
	}
}
