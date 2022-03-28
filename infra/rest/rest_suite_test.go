// +build integration

package rest_test

import (
	"testing"

	"github.com/gofiber/fiber/v2"
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

var cleanup func()
var server *fiber.App

var _ = BeforeSuite(func() {
	db, dbURL, stop := pgtest.NewContainerDB()
	cleanup = stop

	Expect(postgres.Migrate(dbURL)).To(Succeed())

	server = rest.NewFiber(db)
})

var _ = AfterSuite(func() {
	cleanup()
})
