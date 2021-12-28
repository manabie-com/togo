package token

import "time"

// Managing tokens
type Maker interface {
	// Create and sign a new token for a specific username with a specific duration
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
