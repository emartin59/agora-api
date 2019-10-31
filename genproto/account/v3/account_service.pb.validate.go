// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: account/v3/account_service.proto

package account

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
var _account_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAccountRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAccountId() == nil {
		return CreateAccountRequestValidationError{
			field:  "AccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateAccountRequestValidationError{
				field:  "AccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateAccountRequestValidationError is the validation error returned by
// CreateAccountRequest.Validate if the designated constraints aren't met.
type CreateAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAccountRequestValidationError) ErrorName() string {
	return "CreateAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAccountRequestValidationError{}

// Validate checks the field values on CreateAccountResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateAccountResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	return nil
}

// CreateAccountResponseValidationError is the validation error returned by
// CreateAccountResponse.Validate if the designated constraints aren't met.
type CreateAccountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAccountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAccountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAccountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAccountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAccountResponseValidationError) ErrorName() string {
	return "CreateAccountResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateAccountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAccountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAccountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAccountResponseValidationError{}
