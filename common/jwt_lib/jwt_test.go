package jwt_lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type wrapper struct {
	Abc string `json:"abc"`
	Def string `json:"def"`
}

func Test2Ways(t *testing.T) {
	value := wrapper{"abc", "def"}
	token, err := Encrypt(value)
	assert.Equal(t, err, nil)

	var data map[string]interface{}
	data, err = Decrypt(token)
	assert.Equal(t, err, nil)
	assert.Equal(t, data["abc"], "abc")
	assert.Equal(t, data["def"], "def")
}
