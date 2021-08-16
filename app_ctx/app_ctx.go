package appctx

import (
	tokenprovider "github.com/manabie-com/togo/token_provider"
	"gorm.io/gorm"
)

type IAppCtx interface {
	GetDbConn() *gorm.DB
	GetTokenProvider() tokenprovider.Provider
}

type AppCtx struct {
	db          *gorm.DB
	tokProvider tokenprovider.Provider
}

func NewAppContext(db *gorm.DB, tokProvider tokenprovider.Provider) *AppCtx {
	return &AppCtx{
		db:          db,
		tokProvider: tokProvider,
	}
}

func (a *AppCtx) GetDbConn() *gorm.DB {
	return a.db
}

func (a *AppCtx) GetTokenProvider() tokenprovider.Provider {
	return a.tokProvider
}
