// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v4alpha/backoff.proto

package envoy_config_core_v4alpha

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
var _backoff_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on BackoffStrategy with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *BackoffStrategy) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetBaseInterval() == nil {
		return BackoffStrategyValidationError{
			field:  "BaseInterval",
			reason: "value is required",
		}
	}

	if d := m.GetBaseInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return BackoffStrategyValidationError{
				field:  "BaseInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

		if dur < gte {
			return BackoffStrategyValidationError{
				field:  "BaseInterval",
				reason: "value must be greater than or equal to 1ms",
			}
		}

	}

	if d := m.GetMaxInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return BackoffStrategyValidationError{
				field:  "MaxInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return BackoffStrategyValidationError{
				field:  "MaxInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	return nil
}

// BackoffStrategyValidationError is the validation error returned by
// BackoffStrategy.Validate if the designated constraints aren't met.
type BackoffStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BackoffStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BackoffStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BackoffStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BackoffStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BackoffStrategyValidationError) ErrorName() string { return "BackoffStrategyValidationError" }

// Error satisfies the builtin error interface
func (e BackoffStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBackoffStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BackoffStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BackoffStrategyValidationError{}
