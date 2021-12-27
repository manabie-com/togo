package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfigFailed(t *testing.T) {
	nonExistConfig, err := LoadConfig("")
	require.Error(t, err)
	require.Empty(t, nonExistConfig)
}

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, config)
}
