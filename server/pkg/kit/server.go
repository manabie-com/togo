package kit

import (
	"context"
	"net/http"
)

// DecodeRequestFunc extracts a user-domain request object from an HTTP
// request object
type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

// EncodeResponseFunc encodes the passed response object to the HTTP response
// writer
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error

// RequestFunc may take information from an HTTP request and put it into a
// request context
type RequestFunc func(context.Context, *http.Request) context.Context

// ServerResponseFunc may take information from a request context and use it to
// manipulate a ResponseWriter
type ServerResponseFunc func(context.Context, http.ResponseWriter) context.Context

// ErrorEncoder is responsible for encoding an error to the ResponseWriter.
type ErrorEncoder func(ctx context.Context, err error, w http.ResponseWriter)

// Server wraps an endpoint and implements http.Handler.
type Server struct {
	endpoint            Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	before       []RequestFunc
	after        []ServerResponseFunc
	errorEncoder ErrorEncoder
}

// NewServer constructs a new server, which implements http.Handler and wraps
// the provided endpoint.
func NewServer(
	e Endpoint,
	dec DecodeRequestFunc,
	options ...ServerOption,
) *Server {
	s := &Server{
		endpoint:            e,
		dec:          dec,
	}

	s.enc =          defaultEncodeResponse
	s.errorEncoder = defaultEncodeError

	for _, option := range options {
		option(s)
	}
	return s
}

// ServerOption sets an optional parameter for servers.
type ServerOption func(*Server)

// ServerBefore functions are executed on the HTTP request object before the
// request is decoded.
func ServerBefore(before ...RequestFunc) ServerOption {
	return func(s *Server) { s.before = append(s.before, before...) }
}

// ServerAfter functions are executed on the HTTP response writer after the
// endpoint is invoked, but before anything is written to the client.
func ServerAfter(after ...ServerResponseFunc) ServerOption {
	return func(s *Server) { s.after = append(s.after, after...) }
}

// ServerErrorEncoder is used to encode errors to the http.ResponseWriter
// whenever they're encountered in the processing of a request.
func ServerErrorEncoder(ee ErrorEncoder) ServerOption {
	return func(s *Server) { s.errorEncoder = ee }
}

// ServerResponseEncoder is used to encode response to the http.ResponseWriter
// whenever they're encountered in the processing of a request.
func ServerResponseEncoder(erf EncodeResponseFunc) ServerOption {
	return func(s *Server) { s.enc = erf }
}

// ServeHTTP implements http.Handler.
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	for _, f := range s.before {
		ctx = f(ctx, r)
	}

	request, err := s.dec(ctx, r)
	if err != nil {
		s.errorEncoder(ctx, err, w)
		return
	}

	response, err := s.endpoint(ctx, request)
	if err != nil {
		s.errorEncoder(ctx, err, w)
		return
	}

	for _, f := range s.after {
		ctx = f(ctx, w)
	}

	if err := s.enc(ctx, w, response); err != nil {
		s.errorEncoder(ctx, err, w)
		return
	}
}
