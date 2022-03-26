// +build integration

package postgres_test

import (
	"testing"

	"github.com/laghodessa/togo/infra/postgres"
	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	err := postgres.Migrate(dbURL)
	assert.NoError(t, err)
}
