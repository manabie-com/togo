package services

import (
	"testing"

	"github.com/manabie-com/togo/internal/configurations"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/stretchr/testify/require"
)

func mockServiceController(t *testing.T, config configurations.Config, store postgres.Store) *ServiceController {
	sc := &ServiceController{Config: config, Store: store}
	require.NotNil(t, sc)

	return sc
}
