package errs

func Message(err error) string {
	type messager interface {
		Message() string
	}

	if err != nil {
		e, ok := err.(messager)
		if !ok {
			return ""
		}
		return e.Message()
	}

	return ""
}

func ErrorCode(err error) int {
	type errorCoder interface {
		ErrorCode() int
	}

	if err != nil {
		e, ok := err.(errorCoder)
		if !ok {
			return 0
		}
		return e.ErrorCode()
	}

	return 0
}
