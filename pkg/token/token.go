package token

import "time"

// Token is an interface of managing token
type Token interface {
	// CreateToken create new token for username with duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken check if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
