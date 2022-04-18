package controller

import (
	"testing"
	"net/http"
	"bytes"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Unit Testing for the User Controller

func TestLoginAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockUserLoginJSON(ctx, "todo_test_user", "secret")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.LoginUser(ctx)
	
	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestLoginEmptyFieldsAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockUserLoginJSON(ctx, "", "")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.LoginUser(ctx)
	
	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestLoginFailedAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockUserLoginJSON(ctx, "todo_test_user", "secret")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, false, nil)

	userCont.LoginUser(ctx)
	
	assert.EqualValues(t, http.StatusUnauthorized, w.Code, "Output should be HTTP Code Unauthorized")
}

func TestLoginErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockUserLoginJSON(ctx, "todo_test_user", "secret")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, false, createMockErrorMessage("Test Message", "Test Error"))

	userCont.LoginUser(ctx)
	
	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestRegisterUserAPI (t *testing.T) {
	w, ctx := createRegularContext()
	
	jsonBytes, err := createMockUserRegisterJSON(ctx, "todo_test_user", "Todo Test User", "todotestuser@sample.com", "secret", "secret", 10)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.RegisterUser(ctx)
	
	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestRegisterUserPasswordNotMatchAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockUserRegisterJSON(ctx, "todo_test_user", "Todo Test User", "todotestuser@sample.com", "secret", "sikreto", 10)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.RegisterUser(ctx)
	
	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestRegisterUserEmptyFieldsAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockUserRegisterJSON(ctx, "", "", "", "", "", 0)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.RegisterUser(ctx)
	
	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestRegisterUserFailedAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockUserRegisterJSON(ctx, "todo_test_user", "Todo Test User", "todotestuser@sample.com", "secret", "secret", 10)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, false, nil)

	userCont.RegisterUser(ctx)
	
	assert.EqualValues(t, http.StatusNotAcceptable, w.Code, "Output should be HTTP Code Not Acceptable")
}

func TestRegisterUserErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

	jsonBytes, err := createMockUserRegisterJSON(ctx, "todo_test_user", "Todo Test User", "todotestuser@sample.com", "secret", "secret", 10)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, false, createMockErrorMessage("Test Message", "Test Error"))

	userCont.RegisterUser(ctx)
	
	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestFetchUserDetailsAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	userCont := getUserController(createMockUserDetails(false), false, nil)

	userCont.FetchUserDetails(ctx)
	
	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestFetchUserDetailsEmptyUsernameAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "",
		},
	}

	userCont := getUserController(createMockUserDetails(false), false, nil)

	userCont.FetchUserDetails(ctx)
	
	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestFetchUserDetailsEmptyOutputAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	userCont := getUserController(createMockUserDetails(true), false, nil)

	userCont.FetchUserDetails(ctx)
	
	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestFetchUserDetailsErrortAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	userCont := getUserController(createMockUserDetails(true), false, createMockErrorMessage("Test Message", "Test Error"))

	userCont.FetchUserDetails(ctx)
	
	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestUpdateUserDetailsAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	jsonBytes, err := createMockUserUpdateJSON(ctx, "todo_test_user", "Todo Test User", "todotestuser@sample.com", 10)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.UpdateUserDetails(ctx)
	
	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestUpdateUserDetailsEmptyUsernameAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "",
		},
	}

	jsonBytes, err := createMockUserUpdateJSON(ctx, "todo_test_user", "Todo Test User", "todotestuser@sample.com", 10)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.UpdateUserDetails(ctx)
	
	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUpdateUserDetailsEmptyDetailsAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	jsonBytes, err := createMockUserUpdateJSON(ctx, "", "", "", 0)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.UpdateUserDetails(ctx)
	
	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUpdateUserDetailsFailedAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	jsonBytes, err := createMockUserUpdateJSON(ctx, "todo_test_user", "Todo Test User", "todotestuser@sample.com", 10)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, false, nil)

	userCont.UpdateUserDetails(ctx)
	
	assert.EqualValues(t, http.StatusNotFound, w.Code, "Output should be HTTP Code Not Found")
}

func TestUpdateUserDetailsErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	jsonBytes, err := createMockUserUpdateJSON(ctx, "todo_test_user", "Todo Test User", "todotestuser@sample.com", 10)

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, false, createMockErrorMessage("Test Message", "Test Error"))

	userCont.UpdateUserDetails(ctx)
	
	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}

func TestUserPasswordChangeAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	jsonBytes, err := createMockUserPwdChangeJSON(ctx, "secret", "newsecret", "newsecret")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.UserPasswordChange(ctx)
	
	assert.EqualValues(t, http.StatusOK, w.Code, "Output should be HTTP Code OK")
}

func TestUserPasswordChangeEmptyUsernameAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "",
		},
	}

	jsonBytes, err := createMockUserPwdChangeJSON(ctx, "secret", "newsecret", "newsecret")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.UserPasswordChange(ctx)
	
	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUserPasswordChangeEmptyDetailsAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	jsonBytes, err := createMockUserPwdChangeJSON(ctx, "", "", "")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.UserPasswordChange(ctx)
	
	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUserPasswordChangeNewPasswordNotMatchAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	jsonBytes, err := createMockUserPwdChangeJSON(ctx, "secret", "newsecret", "newsikreto")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, true, nil)

	userCont.UserPasswordChange(ctx)
	
	assert.EqualValues(t, http.StatusBadRequest, w.Code, "Output should be HTTP Code Bad Request")
}

func TestUserPasswordChangeFailedAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	jsonBytes, err := createMockUserPwdChangeJSON(ctx, "secret", "newsecret", "newsecret")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, false, nil)

	userCont.UserPasswordChange(ctx)
	
	assert.EqualValues(t, http.StatusUnauthorized, w.Code, "Output should be HTTP Code Unauthorized")
}

func TestUserPasswordChangeErrorAPI (t *testing.T) {
	w, ctx := createRegularContext()

	ctx.Params = []gin.Param{
		{
			Key: "username",
			Value: "todo_test_user",
		},
	}

	jsonBytes, err := createMockUserPwdChangeJSON(ctx, "secret", "newsecret", "newsecret")

	if err != nil {
		t.Errorf("Error encountered when creating mock JSON: %v", err.Error())
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBytes))

	userCont := getUserController(nil, false, createMockErrorMessage("Test Message", "Test Error"))

	userCont.UserPasswordChange(ctx)
	
	assert.EqualValues(t, http.StatusInternalServerError, w.Code, "Output should be HTTP Code Internal Server Error")
}