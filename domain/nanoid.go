package domain

import gonanoid "github.com/matoous/go-nanoid/v2"

func NewID() string {
	return gonanoid.MustGenerate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 11)
}
