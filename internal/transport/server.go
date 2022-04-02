package transport

import "context"

// Server transport interface
type Server interface {
	Load(context.Context) error
	Serve(host string, port int)
}
