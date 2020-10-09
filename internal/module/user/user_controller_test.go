package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/util"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON   = `{"email":"test@mail.com", "password":"12345678"}`
	testConfig = config.Config{
		Env:            "local",
		Port:           "5050",
		DbHost:         "localhost",
		DbUser:         "postgres",
		DbPassword:     "password",
		DbName:         "postgres",
		DbPort:         "5432",
		JwtKey:         "ThisIsJwtKey",
		JwtExp:         1440,
		MaxTodoDefault: 5,
	}
	db, _ = util.CreateConnectionDB(testConfig)
)

func TestRegister(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userRepo, _ := NewUserRepository(db)
	userService, _ := NewUserService(userRepo)
	userController, _ := NewUserController(userService)

	if assert.NoError(t, userController.Register(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
