package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	l := []string{"some", "elements", "for", "testing", "purpose"}
	// a string, list []string
	testCases := []struct {
		name     string
		el       string
		list     []string
		expected bool
	}{
		{
			name:     "Exists",
			el:       "elements",
			list:     l,
			expected: true,
		},
		{
			name:     "Not exists",
			el:       "not-exits",
			list:     l,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := StringInSlice(tc.el, tc.list)
			assert.Equal(t, tc.expected, got)
		})
	}
}
