package gintask_test

import (
	"github.com/japananh/togo/component/tokenprovider"
	"github.com/japananh/togo/modules/task/tasktransport/gintask"
	"gorm.io/gorm"
	"testing"
)

type mockAppCtx struct{}

func (m mockAppCtx) GetMainDBConnection() *gorm.DB {
	return &gorm.DB{}
}

func (m mockAppCtx) GetTokenConfig() *tokenprovider.TokenConfig {
	if tokenConfig, err := tokenprovider.NewTokenConfig(86400, 604800); err != nil {
		return tokenConfig
	}
	return nil
}

func (m mockAppCtx) SecretKey() string {
	return "secret-key"
}

func TestGintask_CreateTask(t *testing.T) {
	gintask.CreateTask(mockAppCtx{})
}
