package http

import (
	"encoding/json"
	"io"
	"net/http"
)

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func ListenAndServe(port string, handler http.Handler) error {
	return http.ListenAndServe(port, handler)
}
