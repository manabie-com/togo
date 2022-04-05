package util

import "regexp"

// Regex for username that only contains alphanumerics and underscores and dashes and whitespaces
func IsSupportedFullname(fullname string) bool {
	re := regexp.MustCompile(`^[\w\s-]+$`).MatchString(fullname)
	return re
}
