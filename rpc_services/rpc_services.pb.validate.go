// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_services.proto

package rpc_services

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _rpc_services_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on NewTaskReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *NewTaskReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TaskName

	return nil
}

// NewTaskReqValidationError is the validation error returned by
// NewTaskReq.Validate if the designated constraints aren't met.
type NewTaskReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewTaskReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewTaskReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewTaskReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewTaskReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewTaskReqValidationError) ErrorName() string { return "NewTaskReqValidationError" }

// Error satisfies the builtin error interface
func (e NewTaskReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewTaskReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewTaskReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewTaskReqValidationError{}

// Validate checks the field values on NewTaskResp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *NewTaskResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TaskName

	return nil
}

// NewTaskRespValidationError is the validation error returned by
// NewTaskResp.Validate if the designated constraints aren't met.
type NewTaskRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewTaskRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewTaskRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewTaskRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewTaskRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewTaskRespValidationError) ErrorName() string { return "NewTaskRespValidationError" }

// Error satisfies the builtin error interface
func (e NewTaskRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewTaskResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewTaskRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewTaskRespValidationError{}