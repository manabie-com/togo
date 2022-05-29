package token

import (
	"errors"
	"time"

	"github.com/dinhquockhanh/togo/internal/pkg/uuid"
)

var (
	ErrExpiredToken = errors.New("expired token")
	ErrInvalidToken = errors.New("invalid token")
)

type Payload struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	TierID    int       `json:"tier_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, tierID int, duration time.Duration) (*Payload, error) {
	return &Payload{
		ID:        uuid.New(),
		Username:  username,
		TierID:    tierID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func IsExpired(err error) bool {
	return errors.Is(err, ErrExpiredToken)
}
