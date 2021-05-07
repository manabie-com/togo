package util

import (
	"encoding/json"
	"strings"
)

func NullOrBlankString(s *string) bool {
	if s == nil {
		return true
	}
	return len(strings.TrimSpace(*s)) == 0
}
func EmptyOrBlankString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func Object2JsonString(o interface{}) string {
	// struct object to json string
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return string(b)
}

//TrimSpaceToLower trim spaces and lower string
func TrimSpaceToLower(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

func TrimSpaceToUpper(str string) string {
	return strings.TrimSpace(strings.ToUpper(str))
}
