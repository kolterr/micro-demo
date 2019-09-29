// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/core/config_source.proto

package core

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on ApiConfigSource with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ApiConfigSource) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := ApiConfigSource_ApiType_name[int32(m.GetApiType())]; !ok {
		return ApiConfigSourceValidationError{
			field:  "ApiType",
			reason: "value must be one of the defined enum values",
		}
	}

	for idx, item := range m.GetGrpcServices() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ApiConfigSourceValidationError{
						field:  fmt.Sprintf("GrpcServices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetRefreshDelay()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ApiConfigSourceValidationError{
					field:  "RefreshDelay",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if d := m.GetRequestTimeout(); d != nil {
		dur := *d

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return ApiConfigSourceValidationError{
				field:  "RequestTimeout",
				reason: "value must be greater than 0s",
			}
		}

	}

	{
		tmp := m.GetRateLimitSettings()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ApiConfigSourceValidationError{
					field:  "RateLimitSettings",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ApiConfigSourceValidationError is the validation error returned by
// ApiConfigSource.Validate if the designated constraints aren't met.
type ApiConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApiConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApiConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApiConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApiConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApiConfigSourceValidationError) ErrorName() string { return "ApiConfigSourceValidationError" }

// Error satisfies the builtin error interface
func (e ApiConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApiConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApiConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApiConfigSourceValidationError{}

// Validate checks the field values on AggregatedConfigSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AggregatedConfigSource) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// AggregatedConfigSourceValidationError is the validation error returned by
// AggregatedConfigSource.Validate if the designated constraints aren't met.
type AggregatedConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AggregatedConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AggregatedConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AggregatedConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AggregatedConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AggregatedConfigSourceValidationError) ErrorName() string {
	return "AggregatedConfigSourceValidationError"
}

// Error satisfies the builtin error interface
func (e AggregatedConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAggregatedConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AggregatedConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AggregatedConfigSourceValidationError{}

// Validate checks the field values on RateLimitSettings with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RateLimitSettings) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetMaxTokens()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RateLimitSettingsValidationError{
					field:  "MaxTokens",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if wrapper := m.GetFillRate(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			return RateLimitSettingsValidationError{
				field:  "FillRate",
				reason: "value must be greater than 0",
			}
		}

	}

	return nil
}

// RateLimitSettingsValidationError is the validation error returned by
// RateLimitSettings.Validate if the designated constraints aren't met.
type RateLimitSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitSettingsValidationError) ErrorName() string {
	return "RateLimitSettingsValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitSettingsValidationError{}

// Validate checks the field values on ConfigSource with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConfigSource) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetInitialFetchTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ConfigSourceValidationError{
					field:  "InitialFetchTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	switch m.ConfigSourceSpecifier.(type) {

	case *ConfigSource_Path:
		// no validation rules for Path

	case *ConfigSource_ApiConfigSource:

		{
			tmp := m.GetApiConfigSource()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ConfigSourceValidationError{
						field:  "ApiConfigSource",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *ConfigSource_Ads:

		{
			tmp := m.GetAds()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ConfigSourceValidationError{
						field:  "Ads",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	default:
		return ConfigSourceValidationError{
			field:  "ConfigSourceSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// ConfigSourceValidationError is the validation error returned by
// ConfigSource.Validate if the designated constraints aren't met.
type ConfigSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigSourceValidationError) ErrorName() string { return "ConfigSourceValidationError" }

// Error satisfies the builtin error interface
func (e ConfigSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigSourceValidationError{}
