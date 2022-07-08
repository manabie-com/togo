package db

import (
	"testing"
)

func TestInitDb(t *testing.T) {
	successDb := InitDb()
	if !successDb {
		t.Errorf("Output expect true instead of false")
	}
}
