package common

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/utils"
	"github.com/stretchr/testify/require"
)

func Test_recordStats(t *testing.T) {
	db, mock, err := setupMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	require.Nil(t, err)
	require.NotNil(t, mock)
	require.NotNil(t, db)

	defer db.Close()
	utils.LoadEnv("../../../../.env")
	{
		user := &models.User{
			Username: "manabie-test-2",
			Password: "123456",
		}

		// Create mock data for test
		mock.ExpectExec("INSERT INTO users").
			WithArgs(1, user.Username, user.Password, 5).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Create Data
		err = recordStats(db, user.Username, user.Password)
		require.Nil(t, err)
	}

	{ //Fail case
		user := &models.User{
			Username: "",
			Password: "",
		}

		// Create mock data for test
		mock.ExpectExec("INSERT INTO users").
			WithArgs(1, user.Username, user.Password, 5).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// Create Data
		err = recordStats(db, user.Username, user.Password)
		require.Nil(t, err)
	}
}
