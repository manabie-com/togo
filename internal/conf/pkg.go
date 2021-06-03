package conf

import "time"

const (
	TokenTimeOutEnvKey = "JWT_TOKEN_EXPIRE"
	TokenIssuerEnvKey  = "JWT_TOKEN_ISSUER"
	JwtSecretEnvKey    = "JWT_SECRET_KEY"
)
const (
	DefaultTokenTimeOut = time.Minute * 10
	DefaultTokenIssuer  = "manabie"
	DefaultJwtSecretKey = "test_data"
)
