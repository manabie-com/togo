package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/manabie-com/togo/internal/pkg/db/postgres"
	"github.com/manabie-com/togo/internal/todo/handler"
	pgr "github.com/manabie-com/togo/internal/todo/repository/postgres"
)

var appHandler http.Handler
var testDB *TestDB

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	os.Setenv("DB_NAME", os.Getenv("DB_NAME")+"_test")
	sqlxConn := postgres.NewSQLXConn()
	defer sqlxConn.Close()

	testDB = NewTestDB(sqlxConn)
	baseRepo := pgr.PGRepository{DBConn: sqlxConn}
	appHandler = handler.NewTodoHandler(handler.TodoRepositoryList{
		UserRepo: &pgr.PGUserRepository{PGRepository: baseRepo},
		TaskRepo: &pgr.PGTaskRepository{PGRepository: baseRepo},
	})

	return m.Run()
}
