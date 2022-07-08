package utils

import (
	"errors"
	"testing"
)

func TestConstValues(t *testing.T) {
	if idChars != "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" {
		t.Errorf("Output expect '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz' instead of %v", idChars)
	}
	if lenChar != len(idChars) {
		t.Errorf("Output expect %d instead of %d", len(idChars), lenChar)
	}
}

func TestIsError(t *testing.T) {
	result := IsError(nil)
	if result {
		t.Errorf("Output expect 'false' instead of 'true'")
	}
	result = IsError(errors.New("Error string"))
	if !result {
		t.Errorf("Output expect 'true' instead of 'false'")
	}
}

func TestRandomString(t *testing.T) {
	length := 10
	result := RandomString(length)
	if length != len(result) {
		t.Errorf("Output expect length=%d instead of %d", length, len(result))
	}
}

func TestRandomNumber(t *testing.T) {
	var min int64 = 10
	var max int64 = 15
	for i := 1; i <= 25; i++ {
		result := RandomNumber(min, max)
		if result < min || max < result {
			t.Errorf("Output expect in range [%d, %d] instead of %d", min, max, result)
			return
		}
	}
}

func TestAES(t *testing.T) {
	key := "kdfjka475ushs7483y9e8f8f84ihg834"
	if len(key) != 32 {
		t.Errorf("key expect length is 32 instead of %d", len(key))
	}
	text := "text to encrypt"
	encText, encSuccess := EncryptAES(key, []byte(text))
	if !encSuccess {
		t.Errorf("Output expect true instead of false")
	}
	decBytes, decSuccess := DecryptAES(key, encText)
	if !decSuccess {
		t.Errorf("Output expect true instead of false")
	}
	decText := string(decBytes)
	if decText != text {
		t.Errorf("Output expect %v instead of %v", text, decText)
	}
}
