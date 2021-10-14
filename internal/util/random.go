package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// responsible for generating a random number fo each generation
func init() {
	rand.Seed(time.Now().UnixNano())
}

// generates a random number given a min and max number
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1) 
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

