package postgres

import (
	"os"
	"testing"

	"github.com/manabie-com/togo/internal/logs"
	"github.com/manabie-com/togo/internal/util"
)

var testPostgres *Postgres

func TestMain(m *testing.M) {
	logger := logs.WithPrefix("Test Postgres")
	err := util.LoadConfig("../../../configs")
	if err != nil {
		logger.Error("error loading config", "process", err.Error())
	}
	testPostgres = NewPostgres()

	os.Exit(m.Run())
}
