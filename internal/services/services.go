package services

// ServiceCommonError defines general error
// which should be handle be developer
type ServiceCommonError string

func (e ServiceCommonError) Error() string {
	return string(e)
}

var (
	// ErrServiceUnhandledException addresses the error that is not being well-handled by developer
	// the error doesn't belong to the scope of the service package, but somehow
	// the service face this issue
	ErrServiceUnhandledException = ServiceCommonError("Service can't handle the exception")
)
