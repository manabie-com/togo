package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConnectDatabase_Postgres(t *testing.T) {
	// set env for testing
	{
		os.Setenv("DB_USER", "togo_user")
		os.Setenv("DB_PASSWORD", "togo_password")
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_NAME", "togo_db")

		os.Setenv("DB_DRIVER", "postgres")
	}

	srv, db, err := run()

	require.Nil(t, err)
	require.Nil(t, db.SQL.Ping())
	require.NotNil(t, srv)

	// test is passed, unset env
	{
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASSWORD")
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("DB_NAME")
		os.Unsetenv("DB_DRIVER")
	}
}
