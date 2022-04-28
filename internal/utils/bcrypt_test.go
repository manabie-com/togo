package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "123456"
	_, err := HashPassword(password)
	if err != nil {
		panic(err)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "123456"
	hash, _ := HashPassword(password)
	if CheckPasswordHash(password, hash) {
		t.Log("Password is correct")
	} else {
		t.Error("Password is incorrect")
	}
}
