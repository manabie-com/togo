package configs

import "time"

const (
	defaultEncryptSalt             = "salt"
	defaultJwtSigningKey           = "strongJwTKeY"
	defaultJwtAccessTokenDuration  = 5 * time.Hour
	defaultJwtRefreshTokenDuration = 24 * time.Hour
)
