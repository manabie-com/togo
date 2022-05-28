package util

import (
	cr "crypto/rand"
	"fmt"
	"math/rand"
)

const aToZ = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	bytes := make([]byte, n)
	if _, err := cr.Read(bytes); err != nil {
		return ""
	}

	for i, b := range bytes {
		bytes[i] = aToZ[b%byte(len(aToZ))]
	}
	return string(bytes)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

func RandomEmail() string {
	return fmt.Sprintf("%v@gmail.com", RandomString(6))
}
