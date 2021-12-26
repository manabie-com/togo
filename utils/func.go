package utils

import "fmt"

func GetKey(userID, date string) string {
	return fmt.Sprintf("%s||%s", userID, date)
}
