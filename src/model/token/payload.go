package token

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

// Payload structure data token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// Valid valid payload
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// InitPayload  Init Payload
func InitPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.New()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
