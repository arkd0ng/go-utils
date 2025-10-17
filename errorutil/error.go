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
// Parameters / 매개변수:
//   - message: The error message / 에러 메시지
//
// Returns / 반환:
//   - error: A new error instance / 새로운 에러 인스턴스
//
// Example / 예제:
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
// Parameters / 매개변수:
//   - format: The format string / 포맷 문자열
//   - args: The format arguments / 포맷 인자들
//
// Returns / 반환:
//   - error: A new error instance with formatted message / 포맷된 메시지를 가진 새로운 에러 인스턴스
//
// Example / 예제:
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
// Parameters / 매개변수:
//   - code: The error code (e.g., "ERR001", "VALIDATION_ERROR") / 에러 코드 (예: "ERR001", "VALIDATION_ERROR")
//   - message: The error message / 에러 메시지
//
// Returns / 반환:
//   - error: A new coded error instance / 새로운 코드가 있는 에러 인스턴스
//
// The returned error implements the Coder interface and can be checked with HasCode.
// 반환된 에러는 Coder 인터페이스를 구현하며 HasCode로 확인할 수 있습니다.
//
// Example / 예제:
//
//	err := errorutil.WithCode("ERR001", "invalid input")
//	fmt.Println(err.Error()) // Output: [ERR001] invalid input
//	if errorutil.HasCode(err, "ERR001") {
//	    // handle specific error / 특정 에러 처리
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
// Parameters / 매개변수:
//   - code: The error code / 에러 코드
//   - format: The format string / 포맷 문자열
//   - args: The format arguments / 포맷 인자들
//
// Returns / 반환:
//   - error: A new coded error instance with formatted message / 포맷된 메시지를 가진 새로운 코드 에러 인스턴스
//
// Example / 예제:
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
// Parameters / 매개변수:
//   - code: The numeric error code (e.g., 404, 500) / 숫자 에러 코드 (예: 404, 500)
//   - message: The error message / 에러 메시지
//
// Returns / 반환:
//   - error: A new numeric coded error instance / 새로운 숫자 코드 에러 인스턴스
//
// The returned error implements the NumericCoder interface and can be checked with HasNumericCode.
// 반환된 에러는 NumericCoder 인터페이스를 구현하며 HasNumericCode로 확인할 수 있습니다.
//
// Example / 예제:
//
//	err := errorutil.WithNumericCode(404, "user not found")
//	fmt.Println(err.Error()) // Output: [404] user not found
//	if errorutil.HasNumericCode(err, 404) {
//	    // handle 404 error / 404 에러 처리
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
// Parameters / 매개변수:
//   - code: The numeric error code / 숫자 에러 코드
//   - format: The format string / 포맷 문자열
//   - args: The format arguments / 포맷 인자들
//
// Returns / 반환:
//   - error: A new numeric coded error instance with formatted message / 포맷된 메시지를 가진 새로운 숫자 코드 에러 인스턴스
//
// Example / 예제:
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
