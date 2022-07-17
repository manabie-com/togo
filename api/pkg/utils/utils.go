package utils

import (
	"math/rand"
	"time"
)

func RamdomID() int {
	min, max := 1000, 100000
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
