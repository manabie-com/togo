package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := Load()
	require.Equal(t, "T", cfg.Environment)
	require.Equal(t, "wqGyEBBfPK9w3Lxw", cfg.JWTKey)
}
