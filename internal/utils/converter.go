package utils

import (
	"fmt"
	"strconv"
)

func BytesToString(data []byte) string {
	return string(data[:])
}

func InterfaceFloat64ToInt(item interface{}) int {
	result, err := strconv.ParseFloat(fmt.Sprintf("%f", item), 10)
	if err != nil {
		fmt.Print("ERROR WHILE PARSING")
		fmt.Println(err)
	}
	return int(result)
}

func CheckIfNilBool(booleanValue *bool) bool {
	if booleanValue == nil {
		return false
	}
	return *booleanValue
}
