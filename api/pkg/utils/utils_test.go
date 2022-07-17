package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRamdomID(t *testing.T) {
	out := RamdomID()
	require.True(t, out >= 1000)
	require.True(t, out <= 100000)
}
