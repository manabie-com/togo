package common

import "math/rand"

var letters = []rune("abcdefghiklmnopqrstuvwxyzABCDEFGHIKLMNOPQRSTVXYZ")

func randSequence(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(99999)%len(letters)]
	}
	return string(b)
}

func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}
	return randSequence(length)
}
