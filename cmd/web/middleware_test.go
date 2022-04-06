package main

import (
	"net/http"
	"testing"
)

func Test_DefaultMiddleware(t *testing.T) {
	var myH myHandler
	h := DefaultMiddleWare(&myH)

	switch v := h.(type) {
	case http.Handler:
		// true, do nothing
	default:
		t.Errorf("expect type http.Handler, not %T", v)
	}
}
