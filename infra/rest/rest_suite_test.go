// +build integration

package rest_test

import (
	"database/sql"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/laghodessa/togo/domain/todo"
	"github.com/laghodessa/togo/infra/postgres"
	"github.com/laghodessa/togo/infra/rest"
	"github.com/laghodessa/togo/test/pgtest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rest Suite")
}

var db *sql.DB
var cleanup func()
var server *fiber.App
var userRepo todo.UserRepo

var _ = BeforeSuite(func() {
	var dbURL string
	db, dbURL, cleanup = pgtest.NewContainerDB()

	Expect(postgres.Migrate(dbURL)).To(Succeed())

	server = rest.NewFiber(db)
	userRepo = postgres.NewTodoUserRepo(db)
})

var _ = AfterSuite(func() {
	cleanup()
})

var _ = AfterEach(func() {
	pgtest.ClearDB(db)
})
