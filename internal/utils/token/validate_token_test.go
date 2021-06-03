package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ValidateTokenTestCases = map[string]struct {
	input string
}{}

func TestValidateToken(t *testing.T) {
	t.Parallel()

	for caseName, tCase := range ValidateTokenTestCases {
		t.Run(caseName, func(t *testing.T) {
			assert.NotNil(t, tCase.input)
			//test
		})
	}
}
