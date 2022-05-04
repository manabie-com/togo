package env_config

import "strings"

// tagOptions contains a slice of tag options
type tagOptions []string

// parseTag splits a struct field's tag into its name and a list of options
// which comes after a name. A tag is in the form of: "name,option1,option2".
// The name can be neglectected.
func parseTag(tag string) (string, string) {
	// tag is one of followings:
	// ""
	// "envName"
	// "envName,default"

	res := strings.Split(tag, ",")
	if len(res) == 0 {
		return "", ""
	}
	if len(res) == 1 {
		return res[0], ""
	}
	return res[0], res[1]
}