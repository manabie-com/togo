package model

import "time"

type Token struct {
	AccessToken string
	ExpiredAt   time.Time
}
