package bcrypt_lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBcrypt(t *testing.T) {
	pd, err := GenerateFromPassword("abc")
	assert.NotEqual(t, pd, "")
	assert.Equal(t, err, nil)
	ok := CompareHashAndPassword(pd, "abc")
	assert.Equal(t, ok, true)
}
