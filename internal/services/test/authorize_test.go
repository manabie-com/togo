package test

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/iservices"
	mockRepo "github.com/manabie-com/togo/internal/mocks/storages"
	mockTool "github.com/manabie-com/togo/internal/mocks/tools"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	t.Run("Login success", func(t *testing.T) {
		req := iservices.LoginRequest{UserId: "1", Password: "123"}
		resExpect := &iservices.LoginResponse{
			Data: "token valid",
		}
		authorizeRepo := mockRepo.IAuthorizeRepo{}
		authorizeRepo.On("ValidateUser", context.TODO(),
			sql.NullString{String: req.UserId, Valid: true},
			sql.NullString{String: req.Password, Valid: true}).Return(true)
		contextTool := mockTool.IContextTool{}
		tokenTool := mockTool.ITokenTool{}
		tokenTool.On("CreateToken", req.UserId, "1234").Return(resExpect.Data, nil)
		authorizeService := services.NewAuthorizeService(&authorizeRepo, "1234", &tokenTool, &contextTool)
		res, err := authorizeService.Login(context.TODO(), req)
		require.Nil(t, err)
		require.NotNil(t, res)
		assert.Equal(t, resExpect, res)
	})
	t.Run("Login Fail not found user id", func(t *testing.T) {
		req := iservices.LoginRequest{UserId: "1", Password: "123"}
		errExpect := tools.NewTodoError(http.StatusUnauthorized, "incorrect user_id/pwd")
		authorizeRepo := mockRepo.IAuthorizeRepo{}
		authorizeRepo.On("ValidateUser", context.TODO(),
			sql.NullString{String: req.UserId, Valid: true},
			sql.NullString{String: req.Password, Valid: true}).Return(false)
		contextTool := mockTool.IContextTool{}
		tokenTool := mockTool.ITokenTool{}
		authorizeService := services.NewAuthorizeService(&authorizeRepo, "1234", &tokenTool, &contextTool)
		res, err := authorizeService.Login(context.TODO(), req)
		require.NotNil(t, err)
		require.Nil(t, res)
		require.Equal(t, errExpect, err)
	})
	t.Run("Login Fail by create token error", func(t *testing.T) {
		req := iservices.LoginRequest{UserId: "1", Password: "123"}
		errExpect := tools.NewTodoError(http.StatusInternalServerError, "fail when create token")
		authorizeRepo := mockRepo.IAuthorizeRepo{}
		authorizeRepo.On("ValidateUser", context.TODO(),
			sql.NullString{String: req.UserId, Valid: true},
			sql.NullString{String: req.Password, Valid: true}).Return(true)
		contextTool := mockTool.IContextTool{}
		tokenTool := mockTool.ITokenTool{}
		tokenTool.On("CreateToken", req.UserId, "1234").Return("", errExpect)
		authorizeService := services.NewAuthorizeService(&authorizeRepo, "1234", &tokenTool, &contextTool)
		res, err := authorizeService.Login(context.TODO(), req)
		require.NotNil(t, err)
		require.Nil(t, res)
		require.Equal(t, errExpect, err)
	})
}

func TestValidate(t *testing.T) {
	t.Run("Validate success", func(t *testing.T) {
		req := &http.Request{}
		req.WithContext(context.TODO())
		contextExpect := context.WithValue(context.TODO(), tools.UserAuthKey(0), "1234")
		authorizeRepo := mockRepo.IAuthorizeRepo{}
		contextTool := mockTool.IContextTool{}
		contextTool.On("WriteUserIDToContext", context.TODO(), "1234").Return(contextExpect)
		tokenTool := mockTool.ITokenTool{}
		tokenTool.On("GetToken", req).Return("1234")
		tokenTool.On("ClaimToken", "1234", "1234").Return("1234", nil)
		authorizeService := services.NewAuthorizeService(&authorizeRepo, "1234", &tokenTool, &contextTool)
		ctx, err := authorizeService.Validate(req)
		require.Nil(t, err)
		require.NotNil(t, ctx)
		assert.Equal(t, contextExpect, ctx)
	})
	t.Run("Validate fail", func(t *testing.T) {
		req := &http.Request{}
		req.WithContext(context.TODO())
		errExpect := tools.NewTodoError(http.StatusUnauthorized, "Your request is unauthorized")
		authorizeRepo := mockRepo.IAuthorizeRepo{}
		contextTool := mockTool.IContextTool{}
		tokenTool := mockTool.ITokenTool{}
		tokenTool.On("GetToken", req).Return("1234")
		tokenTool.On("ClaimToken", "1234", "1234").Return("", errExpect)
		authorizeService := services.NewAuthorizeService(&authorizeRepo, "1234", &tokenTool, &contextTool)
		ctx, err := authorizeService.Validate(req)
		require.Nil(t, ctx)
		require.NotNil(t, err)
		assert.Equal(t, errExpect, err)
	})
}
