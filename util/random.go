package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnoprstuwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random interger between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string with length of n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomUsername generates a random username
func RandomUsername() string{
	return RandomString(15)
}

// RandomPassword generates a random password
func RandomPassword() string{
	return RandomString(17)
}

// RandomContent generates a random content of a task
func RandomContent() string{
	return RandomString(17)
}

// RandomUserid generates a random number of userid
func RandomUserid() int64{
	return RandomInt(0, 100)
}

// RandomId generates a random number of id
func RandomId() int64{
	return RandomInt(0, 100)
}

func RandomLimittask() int64 {
	return RandomInt(0, 100)
}