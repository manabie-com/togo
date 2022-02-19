package configs

import "time"

const (
	defaultEncryptSalt             = "salt"
	defaultJwtSigningKey           = "strongJwTKeY"
	defaultJwtAccessTokenDuration  = 5 * time.Minute
	defaultJwtRefreshTokenDuration = 24 * time.Hour
)
