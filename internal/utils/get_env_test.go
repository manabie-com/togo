package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test case for get string env
var GetStringEnvTestCases = map[string]struct {
	EnvKey         string
	EnvVal         string
	DefaultVal     string
	ExpectedResult string
	ShouldSetEnv   bool
}{
	"Should return default value on empty environment variable": {
		EnvKey:         "test_env_key",
		EnvVal:         "test_val",
		DefaultVal:     "test_default_val",
		ExpectedResult: "test_default_val",
		ShouldSetEnv:   false,
	},
	"Should return env on non-empty environment variable": {
		EnvKey:         "test_env_key",
		EnvVal:         "test_val",
		DefaultVal:     "test_default_val",
		ExpectedResult: "test_val",
		ShouldSetEnv:   true,
	},
	"Should return default value on blank environment variable": {
		EnvKey:         "test_env_key",
		EnvVal:         "",
		DefaultVal:     "test_default_val",
		ExpectedResult: "test_default_val",
		ShouldSetEnv:   true,
	},
}

// Retrieve env key, return default value if value of env is empty
func TestGetStringEnv(t *testing.T) {
	t.Parallel()

	for caseName, tCase := range GetStringEnvTestCases {
		t.Run(caseName, func(t *testing.T) {
			if tCase.ShouldSetEnv {
				os.Setenv(tCase.EnvKey, tCase.EnvVal)
			}
			got := GetStringEnv(tCase.EnvKey, tCase.DefaultVal)
			assert.Equal(t, tCase.ExpectedResult, got)
		})
	}
}
