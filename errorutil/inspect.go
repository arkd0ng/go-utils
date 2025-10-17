// Package errorutil provides comprehensive error handling utilities for Go applications.
// errorutil 패키지는 Go 애플리케이션을 위한 포괄적인 에러 처리 유틸리티를 제공합니다.
package errorutil

import (
	"errors"
)

// HasCode checks if an error or any error in its chain has the specified string code.
// This function walks the error chain using errors.As to find a Coder.
//
// HasCode는 에러 또는 에러 체인의 어떤 에러가 지정된 문자열 코드를 가지고 있는지 확인합니다.
// 이 함수는 errors.As를 사용하여 에러 체인을 탐색하여 Coder를 찾습니다.
//
// Parameters
// 매개변수:
// - err: The error to check
// 확인할 에러
// - code: The code to look for
// 찾을 코드
//
// Returns
// 반환:
// - bool: true if the code is found, false otherwise
// 코드를 찾으면 true, 아니면 false
//
// Example
// 예제:
//
//	err := errorutil.WithCode("ERR001", "validation failed")
// if errorutil.HasCode(err, "ERR001") { / handle validation error / 검증 에러 처리
//	}
//
//	wrapped := errorutil.Wrap(err, "failed to process")
// if errorutil.HasCode(wrapped, "ERR001") { / still found through the chain / 체인을 통해 여전히 찾을 수 있음
//	}
func HasCode(err error, code string) bool {
	if err == nil {
		return false
	}

	// Check the current error first
	// 현재 에러를 먼저 확인
	if coder, ok := err.(Coder); ok {
		if coder.Code() == code {
			return true
		}
	}

	// Walk the error chain
	// 에러 체인을 탐색
	var coder Coder
	if errors.As(err, &coder) {
		return coder.Code() == code
	}

	return false
}

// HasNumericCode checks if an error or any error in its chain has the specified numeric code.
// This function walks the error chain using errors.As to find a NumericCoder.
//
// HasNumericCode는 에러 또는 에러 체인의 어떤 에러가 지정된 숫자 코드를 가지고 있는지 확인합니다.
// 이 함수는 errors.As를 사용하여 에러 체인을 탐색하여 NumericCoder를 찾습니다.
//
// Parameters
// 매개변수:
// - err: The error to check
// 확인할 에러
// - code: The code to look for
// 찾을 코드
//
// Returns
// 반환:
// - bool: true if the code is found, false otherwise
// 코드를 찾으면 true, 아니면 false
//
// Example
// 예제:
//
//	err := errorutil.WithNumericCode(404, "not found")
// if errorutil.HasNumericCode(err, 404) {
// handle 404 error / 404 에러 처리
//	}
//
//	wrapped := errorutil.Wrap(err, "failed to fetch user")
// if errorutil.HasNumericCode(wrapped, 404) {
// still found through the chain / 체인을 통해 여전히 찾을 수 있음
//	}
func HasNumericCode(err error, code int) bool {
	if err == nil {
		return false
	}

	// Check the current error first
	// 현재 에러를 먼저 확인
	if coder, ok := err.(NumericCoder); ok {
		if coder.Code() == code {
			return true
		}
	}

	// Walk the error chain
	// 에러 체인을 탐색
	var coder NumericCoder
	if errors.As(err, &coder) {
		return coder.Code() == code
	}

	return false
}

// GetCode extracts the string code from an error or any error in its chain.
// If no code is found, returns an empty string and false.
//
// GetCode는 에러 또는 에러 체인의 어떤 에러에서 문자열 코드를 추출합니다.
// 코드를 찾지 못하면 빈 문자열과 false를 반환합니다.
//
// Parameters
// 매개변수:
// - err: The error to extract code from
// 코드를 추출할 에러
//
// Returns
// 반환:
// - string: The error code, or empty string if not found
// 에러 코드, 또는 찾지 못하면 빈 문자열
// - bool: true if code was found, false otherwise
// 코드를 찾으면 true, 아니면 false
//
// Example
// 예제:
//
//	err := errorutil.WithCode("ERR001", "validation failed")
//	if code, ok := errorutil.GetCode(err); ok {
//	    fmt.Printf("Error code: %s\n", code) // Output: Error code: ERR001
//	}
//
//	wrapped := errorutil.Wrap(err, "failed to process")
//	if code, ok := errorutil.GetCode(wrapped); ok {
//	    fmt.Printf("Error code: %s\n", code) // Output: Error code: ERR001
//	}
func GetCode(err error) (string, bool) {
	if err == nil {
		return "", false
	}

	// Check the current error first
	// 현재 에러를 먼저 확인
	if coder, ok := err.(Coder); ok {
		return coder.Code(), true
	}

	// Walk the error chain
	// 에러 체인을 탐색
	var coder Coder
	if errors.As(err, &coder) {
		return coder.Code(), true
	}

	return "", false
}

// GetNumericCode extracts the numeric code from an error or any error in its chain.
// If no code is found, returns 0 and false.
//
// GetNumericCode는 에러 또는 에러 체인의 어떤 에러에서 숫자 코드를 추출합니다.
// 코드를 찾지 못하면 0과 false를 반환합니다.
//
// Parameters
// 매개변수:
// - err: The error to extract code from
// 코드를 추출할 에러
//
// Returns
// 반환:
// - int: The error code, or 0 if not found
// 에러 코드, 또는 찾지 못하면 0
// - bool: true if code was found, false otherwise
// 코드를 찾으면 true, 아니면 false
//
// Example
// 예제:
//
//	err := errorutil.WithNumericCode(404, "not found")
//	if code, ok := errorutil.GetNumericCode(err); ok {
//	    fmt.Printf("HTTP status: %d\n", code) // Output: HTTP status: 404
//	}
//
//	wrapped := errorutil.Wrap(err, "failed to fetch user")
//	if code, ok := errorutil.GetNumericCode(wrapped); ok {
//	    fmt.Printf("HTTP status: %d\n", code) // Output: HTTP status: 404
//	}
func GetNumericCode(err error) (int, bool) {
	if err == nil {
		return 0, false
	}

	// Check the current error first
	// 현재 에러를 먼저 확인
	if coder, ok := err.(NumericCoder); ok {
		return coder.Code(), true
	}

	// Walk the error chain
	// 에러 체인을 탐색
	var coder NumericCoder
	if errors.As(err, &coder) {
		return coder.Code(), true
	}

	return 0, false
}

// GetStackTrace extracts the stack trace from an error or any error in its chain.
// If no stack trace is found, returns nil and false.
//
// GetStackTrace는 에러 또는 에러 체인의 어떤 에러에서 스택 트레이스를 추출합니다.
// 스택 트레이스를 찾지 못하면 nil과 false를 반환합니다.
//
// Parameters
// 매개변수:
// - err: The error to extract stack trace from
// 스택 트레이스를 추출할 에러
//
// Returns
// 반환:
// - []Frame: The stack trace, or nil if not found
// 스택 트레이스, 또는 찾지 못하면 nil
// - bool: true if stack trace was found, false otherwise
// 스택 트레이스를 찾으면 true, 아니면 false
//
// Example
// 예제:
//
//	err := errorutil.NewWithStack("something went wrong")
//	if stack, ok := errorutil.GetStackTrace(err); ok {
//	    for _, frame := range stack {
//	        fmt.Println(frame.String())
//	    }
//	}
func GetStackTrace(err error) ([]Frame, bool) {
	if err == nil {
		return nil, false
	}

	// Check the current error first
	// 현재 에러를 먼저 확인
	if tracer, ok := err.(StackTracer); ok {
		return tracer.StackTrace(), true
	}

	// Walk the error chain
	// 에러 체인을 탐색
	var tracer StackTracer
	if errors.As(err, &tracer) {
		return tracer.StackTrace(), true
	}

	return nil, false
}

// GetContext extracts the context data from an error or any error in its chain.
// If no context is found, returns nil and false.
//
// GetContext는 에러 또는 에러 체인의 어떤 에러에서 컨텍스트 데이터를 추출합니다.
// 컨텍스트를 찾지 못하면 nil과 false를 반환합니다.
//
// Parameters
// 매개변수:
// - err: The error to extract context from
// 컨텍스트를 추출할 에러
//
// Returns
// 반환:
// - map[string]interface{}: The context data, or nil if not found
// 컨텍스트 데이터, 또는 찾지 못하면 nil
// - bool: true if context was found, false otherwise
// 컨텍스트를 찾으면 true, 아니면 false
//
// Example
// 예제:
//
//	err := errorutil.WithContext("failed to process", map[string]interface{}{
//	    "user_id": 123,
//	    "action": "login",
//	})
//	if ctx, ok := errorutil.GetContext(err); ok {
//	    fmt.Printf("User ID: %v\n", ctx["user_id"])
//	}
func GetContext(err error) (map[string]interface{}, bool) {
	if err == nil {
		return nil, false
	}

	// Check the current error first
	// 현재 에러를 먼저 확인
	if contexter, ok := err.(Contexter); ok {
		return contexter.Context(), true
	}

	// Walk the error chain
	// 에러 체인을 탐색
	var contexter Contexter
	if errors.As(err, &contexter) {
		return contexter.Context(), true
	}

	return nil, false
}
