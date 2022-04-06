package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	mux := routes()

	switch v := mux.(type) {
	case *chi.Mux:
		// true, do nothing
	default:
		t.Errorf("expect type *chi.Mux, not %T", v)
	}
}
