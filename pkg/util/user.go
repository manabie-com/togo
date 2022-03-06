package util

import (
	"encoding/base64"
)

func ExtractUsername(token string) string {
	parse, err := ParseToken(token)
	if err != nil {
		return ""
	}
	username, err := base64.StdEncoding.DecodeString(parse.Username)
	if err != nil {
		return ""
	}
	return string(username)
}

func ExtractMemberID(token string) int {
	parse, err := ParseToken(token)
	if err != nil {
		return 0
	}
	return parse.UserID
}
