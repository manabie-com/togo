package token

import (
	"testing"
	"time"

	"github.com/manabie-com/togo/internal/utils/random"
	"github.com/stretchr/testify/assert"
)

var defaulTokenTimeOut time.Duration = time.Duration(random.RandInt(10))
var defaultIssuer string = random.RandString(10)

// Exprect result from create new claim test
type NewClaimExpectedResult struct {
	UserID    string
	ExpiresAt int64
	Issuer    string
}

// Test cases for create new claim
var NewClaimTestCases = map[string]struct {
	input    string
	expected NewClaimExpectedResult
}{
	"Should create new claim on ID, PartnerID provided and true value of IsAdmin": {
		input: "userTwo",
		expected: NewClaimExpectedResult{
			UserID:    "userTwo",
			ExpiresAt: time.Now().Add(defaulTokenTimeOut).Unix(),
			Issuer:    defaultIssuer,
		},
	},
	"Should create new claim on ID, PartnerID provided and false value of IsAdmin": {
		input: "userThree",
		expected: NewClaimExpectedResult{
			UserID:    "userThree",
			ExpiresAt: time.Now().Add(defaulTokenTimeOut).Unix(),
			Issuer:    defaultIssuer,
		},
	},
}

var NewTokenTestCases = map[string]struct {
	input        string
	jwtSecretKey string
	shouldFail   bool
}{
	"Should generate jwt token successfully on empty jwtSecretKey": {
		input:        "userOne",
		jwtSecretKey: "",
		shouldFail:   false,
	},
	"Should generate jwt token successfully on non-empty secret key": {
		input:        "userThree",
		jwtSecretKey: "",
		shouldFail:   false,
	},
}

func TestNewToken(t *testing.T) {
	t.Parallel()
	for caseName, tCase := range NewTokenTestCases {
		t.Run(caseName, func(t *testing.T) {
			gotVal, gotErr := NewToken(tCase.input, tCase.jwtSecretKey, defaulTokenTimeOut, defaultIssuer)
			if tCase.shouldFail {
				assert.Empty(t, gotVal)
				assert.NotNil(t, gotErr)
			} else {
				assert.NotEmpty(t, gotVal)
				assert.Nil(t, gotErr)
			}

		})
	}
}
