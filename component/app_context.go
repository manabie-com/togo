package component

import (
	"github.com/japananh/togo/component/tokenprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	GetTokenConfig() *tokenprovider.TokenConfig
}

type appCtx struct {
	db          *gorm.DB
	secretKey   string
	tokenConfig *tokenprovider.TokenConfig
}

func NewAppContext(
	db *gorm.DB,
	secretKey string,
	tokenConfig *tokenprovider.TokenConfig,
) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, tokenConfig: tokenConfig}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetTokenConfig() *tokenprovider.TokenConfig {
	return ctx.tokenConfig
}
