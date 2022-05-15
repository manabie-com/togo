package tokenprovider

import (
	"errors"
	"github.com/japananh/togo/common"
	"strconv"
	"time"
)

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)
	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrInvalidToken",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"` // milliseconds
}

type TokenPayload struct {
	UserId int `json:"user_id"`
}

type TokenConfig struct {
	AccessTokenExpiry  int
	RefreshTokenExpiry int
}

func NewTokenConfig(ate, rte string) (*TokenConfig, error) {
	AccessTokenExpiry, err := strconv.Atoi(ate)
	if err != nil {
		return nil, err
	}

	RefreshTokenExpiry, err := strconv.Atoi(rte)
	if err != nil {
		return nil, err
	}

	return &TokenConfig{
		AccessTokenExpiry:  AccessTokenExpiry,
		RefreshTokenExpiry: RefreshTokenExpiry,
	}, nil
}
