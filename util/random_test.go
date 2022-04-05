package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	randomInt := RandomInt(0, 100)
	require.True(t, randomInt >= 0 && randomInt <= 100)
}

func TestRandomString(t *testing.T) {
	randomString := RandomString(10)
	require.True(t, len(randomString) == 10)
}

func TestRandomName(t *testing.T) {
	randomName := RandomName()
	require.True(t, len(randomName) == 6)
}

func TestRandomCost(t *testing.T) {
	randomCost := RandomCost()
	require.True(t, randomCost >= 0 && randomCost <= 1000)
}

func TestRandomQuantity(t *testing.T) {
	randomQuantity := RandomQuantity()
	require.True(t, randomQuantity >= 2 && randomQuantity <= 10)
}

func TestRandomEmail(t *testing.T) {
	randomEmail := RandomEmail()
	require.True(t, strings.HasSuffix(randomEmail, "@email.com"))
}

func TestRandomPassword(t *testing.T) {
	randomPassword := RandomPassword()
	require.True(t, len(randomPassword) == 12)
}
