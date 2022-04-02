package jwt

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// AddClaimsFunc claim adding handle type
type AddClaimsFunc = func(target jwt.MapClaims, data interface{})

// ExtractClaimsFunc claim extracting handle type
type ExtractClaimsFunc = func(data jwt.Claims) interface{}

// SigningOption options for sign JWT token
type SigningOption struct {
	Duration time.Duration
}

// Add add option method
func (o *SigningOption) Add(options ...*SigningOption) {
	for _, option := range options {
		if option.Duration != 0 {
			o.Duration = option.Duration
		}
	}
}

// JWTProvider token provider
type JWTProvider struct {
	// Signing part
	signingKey    string
	addClaims     AddClaimsFunc
	extractClaims ExtractClaimsFunc
}

// NewJWTProvider constructor
func NewRootJWTProvider(
	signingKey string,

	makeClaims AddClaimsFunc,
	extractClaims ExtractClaimsFunc,
) *JWTProvider {
	return &JWTProvider{
		signingKey,
		makeClaims,
		extractClaims,
	}
}

// GenerateToken sign function
func (p *JWTProvider) GenerateToken(data interface{}, options ...*SigningOption) (string, error) {
	option := &SigningOption{}
	option.Add(options...)
	claims := jwt.MapClaims{}
	p.addClaims(claims, data)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(option.Duration).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(p.signingKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// VerifyToken verify token function
func (p *JWTProvider) VerifyToken(token string) (interface{}, error) {
	getKey := func(token *jwt.Token) (interface{}, error) {
		return []byte(p.signingKey), nil
	}
	jwtToken, err := jwt.Parse(token, getKey)
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, errors.New("TOKEN_INVALID")
	}
	data := p.extractClaims(jwtToken.Claims)
	return data, nil
}