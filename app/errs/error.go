package errs

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
)

type ReadableError struct {
	error

	// Machine-readable error code.
	errorCode int

	// Human-readable message.
	message string

	// Logical operation and nested error.
	op string

	err error
}

func (e *ReadableError) ErrorCode() int {
	return e.errorCode
}

func (e *ReadableError) Message() string {
	return e.message
}

func (e *ReadableError) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.op != "" {
		fmt.Fprintf(&buf, "%s: ", e.op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise, print the error code & message.
	if e.err != nil {
		buf.WriteString(e.err.Error())
	} else {
		if e.errorCode != 0 {
			fmt.Fprintf(&buf, "<%d> ", e.errorCode)
		}
		buf.WriteString(e.message)
	}

	return buf.String()
}

// ReadableError builder
type readableErrorBuilder struct {
	errorCode int
	message   string
	op        string
	err       error
}

func New(message string) *readableErrorBuilder {
	return &readableErrorBuilder{message: message}
}

func (b *readableErrorBuilder) WithCode(errorCode int) *readableErrorBuilder {
	b.errorCode = errorCode
	return b
}

func (b *readableErrorBuilder) WithOp(op string) *readableErrorBuilder {
	b.op = op
	return b
}

func (b *readableErrorBuilder) WithErr(err error) *readableErrorBuilder {
	b.err = errors.Wrap(err, b.message)
	return b
}

func (b *readableErrorBuilder) Build() *ReadableError {
	return &ReadableError{
		errorCode: b.errorCode,
		message:   b.message,
		op:        b.op,
		err:       b.err,
	}
}
