package helper

import (
	"crypto/rand"
	"encoding/base64"
)

// GenKeyAEStoBase64 secret returns 32 bytes AES key view base64.
func GenKeyAEStoBase64() (string, error) {
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	keyBase64 := base64.StdEncoding.EncodeToString(key)
	return keyBase64, nil
}
