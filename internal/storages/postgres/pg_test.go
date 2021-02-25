package postgres

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToConStr(t *testing.T) {
	test := assert.New(t)

	expected := "postgresql://test:123456@localhost:5432/test_db"

	config := &Config{
		Host: "localhost",
		Port: "5432",
		Usr:  "test",
		Pwd:  "123456",
		Db:   "test_db",
	}
	test.Equal(expected, config.toConnStr())
}
