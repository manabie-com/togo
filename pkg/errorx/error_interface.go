package errorx

type ErrorInterface interface {
	// GetCode is of the valid err codes.
	GetCode() string

	// 	GetStatusCode int is of the status code.
	GetStatusCode() int

	GetTitle() string

	// Msg returns a human-readable, unstructured messages describing the err.
	Msg() string

	// WithMeta returns a copy of the Error with the given key-value pair attached
	// as metadata. If the key is already set, it is overwritten.
	WithMeta(key string, val string) ErrorInterface

	// GetMeta returns the stored value for the given key. If the key has no set
	// value, Meta returns an empty string. There is no way to distinguish between
	// an unset value and an explicit empty string.
	GetMeta(key string) string

	// MetaMap returns the complete key-value metadata map stored on the err.
	MetaMap() map[string]string

	// Error returns a string of the form "twirp err <Type>: <Msg>"
	Error() string

	// Location returns a string of the location infomation
	Location() string
}
