package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSupportedFullname(t *testing.T) {
	re := IsSupportedFullname("asd123")
	require.True(t, re)
}
