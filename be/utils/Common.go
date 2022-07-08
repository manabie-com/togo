package utils

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/base64"
	"io"
	"math/rand"
	"time"
)

const (
	idChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	lenChar = 62
)

func IsError(err error) bool {
	return err != nil
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		num := rand.Intn(lenChar)
		bytes[i] = idChars[num]
	}
	return string(bytes)
}

func RandomNumber(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

func EncryptAES(key string, plainText []byte) (string, bool) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "1", false
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(crand.Reader, iv); err != nil {
		return "2", false
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	encmess := base64.URLEncoding.EncodeToString(cipherText)
	return encmess, true
}

func DecryptAES(key string, ct string) ([]byte, bool) {
	cipherText, err := base64.URLEncoding.DecodeString(ct)
	if err != nil {
		return cipherText, false
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return cipherText, false
	}

	if len(cipherText) < aes.BlockSize {
		return cipherText, false
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return cipherText, true
}
