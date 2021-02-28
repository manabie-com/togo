package services

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	cmsqlmock "github.com/manabie-com/togo/pkg/common/cmsql/mock"
	"testing"
)

type IntegrationTest struct {
	db     *sql.DB
	mock   sqlmock.Sqlmock
	hander *ToDoService
}

func NewIntegrationTest() *IntegrationTest {
	db, mock := cmsqlmock.SetupMock()
	return &IntegrationTest{
		db:     db,
		mock:   mock,
		hander: NewToDoService(db, "secret", 5),
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var integrationTest *IntegrationTest

func TestMain(m *testing.M) {
	integrationTest = NewIntegrationTest()
	defer integrationTest.db.Close()
	m.Run()
}
