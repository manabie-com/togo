package helper

import (
	"net/http"
	"time"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func QueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

type RequestParam string

func (p RequestParam) String() string {
	return string(p)
}

func (p RequestParam) Time() (time.Time, error) {
	return time.Parse("2006-01-02", p.String())
}
