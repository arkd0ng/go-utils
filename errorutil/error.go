// Package errorutil provides comprehensive error handling utilities for Go applications.
// errorutil 패키지는 Go 애플리케이션을 위한 포괄적인 에러 처리 유틸리티를 제공합니다.
package errorutil

import (
	"fmt"
)

// New creates a new error with the given message.
// This is similar to errors.New but returns an error that supports unwrapping.
//
// New는 주어진 메시지로 새로운 에러를 생성합니다.
// errors.New와 유사하지만 unwrapping을 지원하는 에러를 반환합니다.
//
// Parameters
// 매개변수:
// - message: The error message
// 에러 메시지
//
// Returns
// 반환:
// - error: A new error instance
// 새로운 에러 인스턴스
//
// Example
// 예제:
//
//	err := errorutil.New("something went wrong")
//	fmt.Println(err.Error()) // Output: something went wrong
func New(message string) error {
	return &wrappedError{
		msg:   message,
		cause: nil,
	}
}

// Newf creates a new error with a formatted message.
// This is similar to fmt.Errorf but returns an error that supports unwrapping.
//
// Newf는 포맷된 메시지로 새로운 에러를 생성합니다.
// fmt.Errorf와 유사하지만 unwrapping을 지원하는 에러를 반환합니다.
//
// Parameters
// 매개변수:
// - format: The format string
// 포맷 문자열
// - args: The format arguments
// 포맷 인자들
//
// Returns
// 반환:
// - error: A new error instance with formatted message
// 포맷된 메시지를 가진 새로운 에러 인스턴스
//
// Example
// 예제:
//
//	err := errorutil.Newf("failed to process user %d", 123)
//	fmt.Println(err.Error()) // Output: failed to process user 123
func Newf(format string, args ...interface{}) error {
	return &wrappedError{
		msg:   fmt.Sprintf(format, args...),
		cause: nil,
	}
}

// WithCode creates a new error with a string error code.
// Error codes are useful for API responses and error categorization.
//
// WithCode는 문자열 에러 코드와 함께 새로운 에러를 생성합니다.
// 에러 코드는 API 응답 및 에러 분류에 유용합니다.
//
// Parameters
// 매개변수:
//   - code: The error code (e.g., "ERR001", "VALIDATION_ERROR") / 에러 코드 (예: "ERR001", "VALIDATION_ERROR")
// - message: The error message
// 에러 메시지
//
// Returns
// 반환:
// - error: A new coded error instance
// 새로운 코드가 있는 에러 인스턴스
//
// The returned error implements the Coder interface and can be checked with HasCode.
// 반환된 에러는 Coder 인터페이스를 구현하며 HasCode로 확인할 수 있습니다.
//
// Example
// 예제:
//
//	err := errorutil.WithCode("ERR001", "invalid input")
//	fmt.Println(err.Error()) // Output: [ERR001] invalid input
// if errorutil.HasCode(err, "ERR001") { / handle specific error / 특정 에러 처리
//	}
func WithCode(code, message string) error {
	return &codedError{
		msg:   message,
		code:  code,
		cause: nil,
	}
}

// WithCodef creates a new error with a string error code and formatted message.
// This combines the functionality of WithCode and Newf.
//
// WithCodef는 문자열 에러 코드와 포맷된 메시지로 새로운 에러를 생성합니다.
// WithCode와 Newf의 기능을 결합합니다.
//
// Parameters
// 매개변수:
// - code: The error code
// 에러 코드
// - format: The format string
// 포맷 문자열
// - args: The format arguments
// 포맷 인자들
//
// Returns
// 반환:
// - error: A new coded error instance with formatted message
// 포맷된 메시지를 가진 새로운 코드 에러 인스턴스
//
// Example
// 예제:
//
//	err := errorutil.WithCodef("ERR001", "invalid user: %d", 123)
//	fmt.Println(err.Error()) // Output: [ERR001] invalid user: 123
func WithCodef(code, format string, args ...interface{}) error {
	return &codedError{
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		cause: nil,
	}
}

// WithNumericCode creates a new error with a numeric error code.
// Numeric codes are useful for HTTP status codes and error numbers.
//
// WithNumericCode는 숫자 에러 코드와 함께 새로운 에러를 생성합니다.
// 숫자 코드는 HTTP 상태 코드 및 에러 번호에 유용합니다.
//
// Parameters
// 매개변수:
// - code: The numeric error code (e.g., 404, 500)
// 숫자 에러 코드 (예: 404, 500)
// - message: The error message
// 에러 메시지
//
// Returns
// 반환:
// - error: A new numeric coded error instance
// 새로운 숫자 코드 에러 인스턴스
//
// The returned error implements the NumericCoder interface and can be checked with HasNumericCode.
// 반환된 에러는 NumericCoder 인터페이스를 구현하며 HasNumericCode로 확인할 수 있습니다.
//
// Example
// 예제:
//
//	err := errorutil.WithNumericCode(404, "user not found")
//	fmt.Println(err.Error()) // Output: [404] user not found
// if errorutil.HasNumericCode(err, 404) {
// handle 404 error / 404 에러 처리
//	}
func WithNumericCode(code int, message string) error {
	return &numericCodedError{
		msg:   message,
		code:  code,
		cause: nil,
	}
}

// WithNumericCodef creates a new error with a numeric error code and formatted message.
// This combines the functionality of WithNumericCode and Newf.
//
// WithNumericCodef는 숫자 에러 코드와 포맷된 메시지로 새로운 에러를 생성합니다.
// WithNumericCode와 Newf의 기능을 결합합니다.
//
// Parameters
// 매개변수:
// - code: The numeric error code
// 숫자 에러 코드
// - format: The format string
// 포맷 문자열
// - args: The format arguments
// 포맷 인자들
//
// Returns
// 반환:
// - error: A new numeric coded error instance with formatted message
// 포맷된 메시지를 가진 새로운 숫자 코드 에러 인스턴스
//
// Example
// 예제:
//
//	err := errorutil.WithNumericCodef(500, "database error: %s", "connection timeout")
//	fmt.Println(err.Error()) // Output: [500] database error: connection timeout
func WithNumericCodef(code int, format string, args ...interface{}) error {
	return &numericCodedError{
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		cause: nil,
	}
}

// Wrap wraps an existing error with a new message.
// This is the most common way to add context to errors as they bubble up.
//
// Wrap은 기존 에러를 새로운 메시지로 래핑합니다.
// 에러가 위로 전파될 때 컨텍스트를 추가하는 가장 일반적인 방법입니다.
//
// Parameters
// 매개변수:
// - cause: The error to wrap
// 래핑할 에러
// - message: The wrapping message
// 래핑 메시지
//
// Returns
// 반환:
// - error: A new error wrapping the cause
// cause를 래핑하는 새로운 에러
//
// The returned error implements Unwrapper and works with errors.Is and errors.As.
// 반환된 에러는 Unwrapper를 구현하며 errors.Is 및 errors.As와 함께 작동합니다.
//
// Example
// 예제:
//
//	original := errors.New("connection failed")
//	wrapped := errorutil.Wrap(original, "failed to connect to database")
//	fmt.Println(wrapped.Error()) // Output: failed to connect to database: connection failed
//
// // Check the underlying error
// 기저 에러 확인
// if errors.Is(wrapped, original) {
// handle specific error / 특정 에러 처리
//	}
func Wrap(cause error, message string) error {
	if cause == nil {
		return nil
	}
	return &wrappedError{
		msg:   message,
		cause: cause,
	}
}

// Wrapf wraps an existing error with a formatted message.
// This combines Wrap with format string support.
//
// Wrapf는 기존 에러를 포맷된 메시지로 래핑합니다.
// Wrap과 포맷 문자열 지원을 결합합니다.
//
// Parameters
// 매개변수:
// - cause: The error to wrap
// 래핑할 에러
// - format: The format string
// 포맷 문자열
// - args: The format arguments
// 포맷 인자들
//
// Returns
// 반환:
// - error: A new error wrapping the cause with formatted message
// 포맷된 메시지로 cause를 래핑하는 새로운 에러
//
// Example
// 예제:
//
//	original := errors.New("not found")
//	wrapped := errorutil.Wrapf(original, "failed to find user %d", 123)
//	fmt.Println(wrapped.Error()) // Output: failed to find user 123: not found
func Wrapf(cause error, format string, args ...interface{}) error {
	if cause == nil {
		return nil
	}
	return &wrappedError{
		msg:   fmt.Sprintf(format, args...),
		cause: cause,
	}
}

// WrapWithCode wraps an existing error with a message and string code.
// This is useful for categorizing errors as they bubble up.
//
// WrapWithCode는 기존 에러를 메시지와 문자열 코드로 래핑합니다.
// 에러가 위로 전파될 때 분류하는 데 유용합니다.
//
// Parameters
// 매개변수:
// - cause: The error to wrap
// 래핑할 에러
// - code: The error code
// 에러 코드
// - message: The wrapping message
// 래핑 메시지
//
// Returns
// 반환:
// - error: A new coded error wrapping the cause
// cause를 래핑하는 새로운 코드 에러
//
// The returned error implements both Coder and Unwrapper interfaces.
// 반환된 에러는 Coder와 Unwrapper 인터페이스를 모두 구현합니다.
//
// Example
// 예제:
//
//	original := errors.New("validation failed")
//	wrapped := errorutil.WrapWithCode(original, "ERR001", "invalid input")
//	fmt.Println(wrapped.Error()) // Output: [ERR001] invalid input: validation failed
//
// // Check error code
// 에러 코드 확인
// if errorutil.HasCode(wrapped, "ERR001") { / handle validation error / 검증 에러 처리
//	}
func WrapWithCode(cause error, code, message string) error {
	if cause == nil {
		return nil
	}
	return &codedError{
		msg:   message,
		code:  code,
		cause: cause,
	}
}

// WrapWithCodef wraps an existing error with a formatted message and string code.
// This combines WrapWithCode with format string support.
//
// WrapWithCodef는 기존 에러를 포맷된 메시지와 문자열 코드로 래핑합니다.
// WrapWithCode와 포맷 문자열 지원을 결합합니다.
//
// Parameters
// 매개변수:
// - cause: The error to wrap
// 래핑할 에러
// - code: The error code
// 에러 코드
// - format: The format string
// 포맷 문자열
// - args: The format arguments
// 포맷 인자들
//
// Returns
// 반환:
// - error: A new coded error wrapping the cause
// cause를 래핑하는 새로운 코드 에러
//
// Example
// 예제:
//
//	original := errors.New("not found")
//	wrapped := errorutil.WrapWithCodef(original, "ERR404", "user %d not found", 123)
//	fmt.Println(wrapped.Error()) // Output: [ERR404] user 123 not found: not found
func WrapWithCodef(cause error, code, format string, args ...interface{}) error {
	if cause == nil {
		return nil
	}
	return &codedError{
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		cause: cause,
	}
}

// WrapWithNumericCode wraps an existing error with a message and numeric code.
// This is useful for HTTP errors and numeric error codes.
//
// WrapWithNumericCode는 기존 에러를 메시지와 숫자 코드로 래핑합니다.
// HTTP 에러 및 숫자 에러 코드에 유용합니다.
//
// Parameters
// 매개변수:
// - cause: The error to wrap
// 래핑할 에러
// - code: The numeric error code
// 숫자 에러 코드
// - message: The wrapping message
// 래핑 메시지
//
// Returns
// 반환:
// - error: A new numeric coded error wrapping the cause
// cause를 래핑하는 새로운 숫자 코드 에러
//
// The returned error implements both NumericCoder and Unwrapper interfaces.
// 반환된 에러는 NumericCoder와 Unwrapper 인터페이스를 모두 구현합니다.
//
// Example
// 예제:
//
//	original := errors.New("database error")
//	wrapped := errorutil.WrapWithNumericCode(original, 500, "internal server error")
//	fmt.Println(wrapped.Error()) // Output: [500] internal server error: database error
//
// // Check error code
// 에러 코드 확인
// if errorutil.HasNumericCode(wrapped, 500) {
// handle 500 error / 500 에러 처리
//	}
func WrapWithNumericCode(cause error, code int, message string) error {
	if cause == nil {
		return nil
	}
	return &numericCodedError{
		msg:   message,
		code:  code,
		cause: cause,
	}
}

// WrapWithNumericCodef wraps an existing error with a formatted message and numeric code.
// This combines WrapWithNumericCode with format string support.
//
// WrapWithNumericCodef는 기존 에러를 포맷된 메시지와 숫자 코드로 래핑합니다.
// WrapWithNumericCode와 포맷 문자열 지원을 결합합니다.
//
// Parameters
// 매개변수:
// - cause: The error to wrap
// 래핑할 에러
// - code: The numeric error code
// 숫자 에러 코드
// - format: The format string
// 포맷 문자열
// - args: The format arguments
// 포맷 인자들
//
// Returns
// 반환:
// - error: A new numeric coded error wrapping the cause
// cause를 래핑하는 새로운 숫자 코드 에러
//
// Example
// 예제:
//
//	original := errors.New("timeout")
//	wrapped := errorutil.WrapWithNumericCodef(original, 408, "request timeout after %d seconds", 30)
//	fmt.Println(wrapped.Error()) // Output: [408] request timeout after 30 seconds: timeout
func WrapWithNumericCodef(cause error, code int, format string, args ...interface{}) error {
	if cause == nil {
		return nil
	}
	return &numericCodedError{
		msg:   fmt.Sprintf(format, args...),
		code:  code,
		cause: cause,
	}
}
