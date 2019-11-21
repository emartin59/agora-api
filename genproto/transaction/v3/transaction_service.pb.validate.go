// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: transaction/v3/transaction_service.proto

package transaction

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
var _transaction_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetHistoryRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetHistoryRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetHistoryRequestValidationError is the validation error returned by
// GetHistoryRequest.Validate if the designated constraints aren't met.
type GetHistoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHistoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHistoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHistoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHistoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHistoryRequestValidationError) ErrorName() string {
	return "GetHistoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetHistoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHistoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHistoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHistoryRequestValidationError{}

// Validate checks the field values on GetHistoryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetHistoryResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetHistoryResponseValidationError is the validation error returned by
// GetHistoryResponse.Validate if the designated constraints aren't met.
type GetHistoryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHistoryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHistoryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHistoryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHistoryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHistoryResponseValidationError) ErrorName() string {
	return "GetHistoryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetHistoryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHistoryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHistoryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHistoryResponseValidationError{}

// Validate checks the field values on SubmitSendRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *SubmitSendRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetTransactionXdr()); l < 1 || l > 10240 {
		return SubmitSendRequestValidationError{
			field:  "TransactionXdr",
			reason: "value length must be between 1 and 10240 bytes, inclusive",
		}
	}

	return nil
}

// SubmitSendRequestValidationError is the validation error returned by
// SubmitSendRequest.Validate if the designated constraints aren't met.
type SubmitSendRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitSendRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitSendRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitSendRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitSendRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitSendRequestValidationError) ErrorName() string {
	return "SubmitSendRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitSendRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitSendRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitSendRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitSendRequestValidationError{}

// Validate checks the field values on SubmitSendResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubmitSendResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if m.GetHash() == nil {
		return SubmitSendResponseValidationError{
			field:  "Hash",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetHash()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubmitSendResponseValidationError{
				field:  "Hash",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Ledger

	if l := len(m.GetResultXdr()); l < 1 || l > 10240 {
		return SubmitSendResponseValidationError{
			field:  "ResultXdr",
			reason: "value length must be between 1 and 10240 bytes, inclusive",
		}
	}

	return nil
}

// SubmitSendResponseValidationError is the validation error returned by
// SubmitSendResponse.Validate if the designated constraints aren't met.
type SubmitSendResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitSendResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitSendResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitSendResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitSendResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitSendResponseValidationError) ErrorName() string {
	return "SubmitSendResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitSendResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitSendResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitSendResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitSendResponseValidationError{}

// Validate checks the field values on GetTransactionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetTransactionRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTransactionHash() == nil {
		return GetTransactionRequestValidationError{
			field:  "TransactionHash",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetTransactionHash()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTransactionRequestValidationError{
				field:  "TransactionHash",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetTransactionRequestValidationError is the validation error returned by
// GetTransactionRequest.Validate if the designated constraints aren't met.
type GetTransactionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransactionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransactionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransactionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransactionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransactionRequestValidationError) ErrorName() string {
	return "GetTransactionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransactionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransactionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransactionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransactionRequestValidationError{}

// Validate checks the field values on GetTransactionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetTransactionResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for State

	// no validation rules for Ledger

	if len(m.GetResultXdr()) > 10240 {
		return GetTransactionResponseValidationError{
			field:  "ResultXdr",
			reason: "value length must be at most 10240 bytes",
		}
	}

	return nil
}

// GetTransactionResponseValidationError is the validation error returned by
// GetTransactionResponse.Validate if the designated constraints aren't met.
type GetTransactionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransactionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransactionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransactionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransactionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransactionResponseValidationError) ErrorName() string {
	return "GetTransactionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransactionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransactionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransactionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransactionResponseValidationError{}
