// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/tap/v3alpha/wrapper.proto

package envoy_data_tap_v3alpha

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
var _wrapper_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on TraceWrapper with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TraceWrapper) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Trace.(type) {

	case *TraceWrapper_HttpBufferedTrace:

		if v, ok := interface{}(m.GetHttpBufferedTrace()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TraceWrapperValidationError{
					field:  "HttpBufferedTrace",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TraceWrapper_HttpStreamedTraceSegment:

		if v, ok := interface{}(m.GetHttpStreamedTraceSegment()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TraceWrapperValidationError{
					field:  "HttpStreamedTraceSegment",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TraceWrapper_SocketBufferedTrace:

		if v, ok := interface{}(m.GetSocketBufferedTrace()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TraceWrapperValidationError{
					field:  "SocketBufferedTrace",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TraceWrapper_SocketStreamedTraceSegment:

		if v, ok := interface{}(m.GetSocketStreamedTraceSegment()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TraceWrapperValidationError{
					field:  "SocketStreamedTraceSegment",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return TraceWrapperValidationError{
			field:  "Trace",
			reason: "value is required",
		}

	}

	return nil
}

// TraceWrapperValidationError is the validation error returned by
// TraceWrapper.Validate if the designated constraints aren't met.
type TraceWrapperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TraceWrapperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TraceWrapperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TraceWrapperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TraceWrapperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TraceWrapperValidationError) ErrorName() string { return "TraceWrapperValidationError" }

// Error satisfies the builtin error interface
func (e TraceWrapperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTraceWrapper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TraceWrapperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TraceWrapperValidationError{}
