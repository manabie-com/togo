package request

import (
	"net/http"
)

// Should be improved later, allow parse into
// a given struct object with reflect keys
func QueryParam(req *http.Request, p string) string {
	return req.FormValue(p)
}
