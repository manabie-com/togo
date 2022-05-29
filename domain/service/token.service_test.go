package service_test

import (
	"context"
	"testing"
	"togo/domain/model"
	"togo/domain/service"
)

var secret = "aaaaaasdsds"

func TestTokenServiceImpl_CreateToken(t *testing.T) {
	tokenService := service.NewTokenService(secret)
	u := model.User{
		Username: "admin",
		Password: "admin",
		Token:    "",
		Limit:    0,
	}
	token, err := tokenService.CreateToken(context.Background(), u)
	if err != nil {
		t.Errorf("Error creating token: %v", err)
		return
	}
	t.Logf("Token created: %s", token)
}

func TestTokenServiceImpl_ValidateToken(t *testing.T) {
	signedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTM3OTg3ODM4NTYsIlVzZXJuYW1lIjoiYWRtaW4iLCJMaW1pdCI6MH0.ruMWKfxueOy5LIWkgyZsSpLsUbKNtt0QfLOG_UtUTnA"
	tokenService := service.NewTokenService(secret)
	userClaim, err := tokenService.ValidateToken(context.Background(), signedToken)
	if err != nil {
		t.Errorf("Error validte token: %#v", err)
		return
	}
	t.Logf("UserClaim: %#v", userClaim)
}
