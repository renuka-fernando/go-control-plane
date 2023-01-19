// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/tap/v3/common.proto

package tapv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Body with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Body) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Body with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BodyMultiError, or nil if none found.
func (m *Body) ValidateAll() error {
	return m.validate(true)
}

func (m *Body) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Truncated

	switch v := m.BodyType.(type) {
	case *Body_AsBytes:
		if v == nil {
			err := BodyValidationError{
				field:  "BodyType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for AsBytes
	case *Body_AsString:
		if v == nil {
			err := BodyValidationError{
				field:  "BodyType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for AsString
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return BodyMultiError(errors)
	}

	return nil
}

// BodyMultiError is an error wrapping multiple validation errors returned by
// Body.ValidateAll() if the designated constraints aren't met.
type BodyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BodyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BodyMultiError) AllErrors() []error { return m }

// BodyValidationError is the validation error returned by Body.Validate if the
// designated constraints aren't met.
type BodyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BodyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BodyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BodyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BodyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BodyValidationError) ErrorName() string { return "BodyValidationError" }

// Error satisfies the builtin error interface
func (e BodyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBody.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BodyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BodyValidationError{}
