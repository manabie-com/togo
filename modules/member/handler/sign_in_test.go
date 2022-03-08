package handler

import (
	"testing"

	"github.com/khoale193/togo/models/dbcon"
	"github.com/khoale193/togo/models/migration"
)

func init() {
	dbcon.SetupTest()
	migration.Migrate()
}

func TestIsCorrectUser(t *testing.T) {
	if is, _ := isCorrectUser("test1", "123456"); is == false {
		t.Logf("Expected: Username: test1 and Password: 123456 was right.")
		t.Fail()
	}
	if is, _ := isCorrectUser("test2", "123456"); is == false {
		t.Logf("Expected: Username: test2 and Password: 123456 was right.")
		t.Fail()
	}
}
