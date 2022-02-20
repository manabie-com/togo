package jwt

import (
	"time"
	"togo/internal/domain"
	"togo/internal/provider"

	"github.com/dgrijalva/jwt-go"
	"github.com/levanthanh-ptit/go-ez/ez_provider"
)

// tokenProvider token provider
type tokenProvider struct {
	*ez_provider.JWTProvider
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
		ez_provider.NewJWTProvider(signingKey, addClaims, extractClaims),
		jwtAccessTokenDuration,
		jwtRefreshTokenDuration,
	}
}

func (p tokenProvider) GenerateToken(data interface{}) (string, error) {
	return p.JWTProvider.GenerateToken(data, &ez_provider.SigningOption{Duration: p.jwtAccessTokenDuration})
}
