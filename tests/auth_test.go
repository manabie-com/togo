package tests

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func Test_Register_InvalidInput(t *testing.T) {
	ctx := context.Background()
	err := sendRegister(ctx, &registerReq{
		Username: "",
		Password: "",
	})
	assert.Error(t, err)
}

func Test_Register_Successful(t *testing.T) {
	ctx := context.Background()
	err := sendRegister(ctx, &registerReq{
		Username: faker.Username(),
		Password: "a123456",
	})
	assert.NoError(t, err)
}

func Test_Login_InvalidInput(t *testing.T) {
	ctx := context.Background()
	res, err := sendLogin(ctx, &loginReq{
		Username: "",
		Password: "",
	})
	assert.Error(t, err)
	assert.Nil(t, res)
}

func Test_Login_Successful(t *testing.T) {
	ctx := context.Background()
	username := faker.Username()
	password := "a123456"
	err := sendRegister(ctx, &registerReq{
		Username: username,
		Password: password,
	})
	assert.NoError(t, err)
	res, err := sendLogin(ctx, &loginReq{
		Username: username,
		Password: password,
	})
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
