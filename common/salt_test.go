package common_test

import (
	"github.com/japananh/togo/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserCreate_GenSalt(t *testing.T) {
	var tcs = []struct {
		arg      int
		expected string
	}{
		{0, ""},
		{-1, "BMOcrdlEltpGCZxZkmVyBqyDwxrDXkxPLZMOFDSXNxGqrwKoxt"},
		{50, "BMOcrdlEltpGCZxZkmVyBqyDwxrDXkxPLZMOFDSXNxGqrwKoxt"},
		{20, "tNCdTDEnAVLkqXKcyOEp"},
	}

	for _, tc := range tcs {
		output := common.GenSalt(tc.arg)
		assert.Equal(t, len(output), len(tc.expected), "they should be equal")
	}
}
