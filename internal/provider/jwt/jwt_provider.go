package jwt

import (
	"time"
	"togo/internal/domain"
	"togo/internal/provider"

	"github.com/dgrijalva/jwt-go"
)

// tokenProvider token provider
type tokenProvider struct {
	*JWTProvider
	jwtAccessTokenDuration  time.Duration
	jwtRefreshTokenDuration time.Duration
}

func addClaims(target jwt.MapClaims, data interface{}) {
	user := data.(*domain.User)
	target["user_id"] = user.ID
}

func extractClaims(data jwt.Claims) interface{} {
	mapClaims := data.(jwt.MapClaims)
	userID := mapClaims["user_id"].(float64)
	user := &domain.User{
		ID: uint(userID),
	}
	return user
}

// NewJWTProvider constructor
func NewJWTProvider(
	signingKey string,
	jwtAccessTokenDuration time.Duration,
	jwtRefreshTokenDuration time.Duration,
) provider.TokenProvider {
	return &tokenProvider{
		NewRootJWTProvider(signingKey, addClaims, extractClaims),
		jwtAccessTokenDuration,
		jwtRefreshTokenDuration,
	}
}

func (p tokenProvider) GenerateToken(data interface{}) (string, error) {
	return p.JWTProvider.GenerateToken(data, &SigningOption{Duration: p.jwtAccessTokenDuration})
}
