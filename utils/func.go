package utils

import (
	"fmt"
	"math/rand"
)

func GetKey(userID, date string) string {
	return fmt.Sprintf("%s||%s", userID, date)
}

func GetNumRandomTask() int {
	return rand.Intn(5) + 1
}
