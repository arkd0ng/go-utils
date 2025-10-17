// Package errorutil provides comprehensive error handling utilities for Go applications.
// errorutil 패키지는 Go 애플리케이션을 위한 포괄적인 에러 처리 유틸리티를 제공합니다.
//
// This package offers enhanced error creation, wrapping, inspection, and formatting
// capabilities beyond the standard library, while maintaining full compatibility.
//
// 이 패키지는 표준 라이브러리를 넘어서는 향상된 에러 생성, 래핑, 검사 및 포매팅 기능을
// 제공하며, 완전한 호환성을 유지합니다.
//
// # Key Features
// 주요 기능
//
// - Error wrapping with context preservation
// 컨텍스트 보존과 함께 에러 래핑
// - Error codes (string and numeric)
// 에러 코드 (문자열 및 숫자)
// - Stack trace capture and display
// 스택 트레이스 캡처 및 표시
// - Contextual errors with key-value data
// 키-값 데이터를 가진 컨텍스트 에러
// - Error classification and inspection
// 에러 분류 및 검사
// - Advanced error formatting
// 고급 에러 포매팅
//
// # Example Usage
// 사용 예제
//
//	import "github.com/arkd0ng/go-utils/errorutil"
//
// // Create error with code
// 코드와 함께 에러 생성
//
//	err := errorutil.WithCode("ERR001", "invalid input")
//
// // Wrap error with context
// 컨텍스트와 함께 에러 래핑
//
//	err = errorutil.Wrap(err, "failed to process user data")
//
// // Check error code
// 에러 코드 확인
// if errorutil.HasCode(err, "ERR001") {
// handle specific error
// 특정 에러 처리
//
//	}
package errorutil

import (
	"fmt"

	"github.com/arkd0ng/go-utils/internal/version"
)

// Version is the current version of the errorutil package.
// Version은 errorutil 패키지의 현재 버전입니다.
var Version = version.Get()

// Unwrapper is the interface for errors that wrap other errors.
// This is compatible with the standard library errors.Unwrap function.
// Unwrapper는 다른 에러를 래핑하는 에러를 위한 인터페이스입니다.
// 표준 라이브러리 errors.Unwrap 함수와 호환됩니다.
type Unwrapper interface {
	error
	Unwrap() error
}

// Coder is the interface for errors that have an associated string code.
// Error codes are useful for API responses and error categorization.
// Coder는 연결된 문자열 코드를 가진 에러를 위한 인터페이스입니다.
// 에러 코드는 API 응답 및 에러 분류에 유용합니다.
type Coder interface {
	error
	Code() string
}

// NumericCoder is the interface for errors that have an associated numeric code.
// Numeric codes are useful for HTTP status codes and error numbers.
// NumericCoder는 연결된 숫자 코드를 가진 에러를 위한 인터페이스입니다.
// 숫자 코드는 HTTP 상태 코드 및 에러 번호에 유용합니다.
type NumericCoder interface {
	error
	Code() int
}

// StackTracer is the interface for errors that capture stack traces.
// Stack traces help with debugging by showing where errors originated.
// StackTracer는 스택 트레이스를 캡처하는 에러를 위한 인터페이스입니다.
// 스택 트레이스는 에러가 어디서 발생했는지 보여줌으로써 디버깅에 도움을 줍니다.
type StackTracer interface {
	error
	StackTrace() []Frame
}

// Contexter is the interface for errors that carry structured contextual data.
// Context data provides additional information about the error condition.
// Contexter는 구조화된 컨텍스트 데이터를 전달하는 에러를 위한 인터페이스입니다.
// 컨텍스트 데이터는 에러 조건에 대한 추가 정보를 제공합니다.
type Contexter interface {
	error
	Context() map[string]interface{}
}

// Frame represents a single stack frame with file, line, and function information.
// Frame은 파일, 라인 및 함수 정보를 가진 단일 스택 프레임을 나타냅니다.
type Frame struct {
	// File is the full path to the source file
	// File은 소스 파일의 전체 경로입니다
	File string
	// Line is the line number in the source file
	// Line은 소스 파일의 라인 번호입니다
	Line int
	// Function is the fully qualified function name
	// Function은 완전한 함수명입니다
	Function string
}

// String returns a formatted string representation of the frame.
// String은 프레임의 포맷된 문자열 표현을 반환합니다.
func (f Frame) String() string {
	return fmt.Sprintf("%s:%d %s", f.File, f.Line, f.Function)
}

// wrappedError is a basic error that wraps another error.
// This is the foundation for all error wrapping functionality.
// wrappedError는 다른 에러를 래핑하는 기본 에러입니다.
// 이것은 모든 에러 래핑 기능의 기초입니다.
type wrappedError struct {
	// msg is the error message
	// msg는 에러 메시지입니다
	msg string
	// cause is the wrapped error
	// cause는 래핑된 에러입니다
	cause error
}

// Error returns the error message.
// If there is a wrapped error, it includes that error's message.
// Error는 에러 메시지를 반환합니다.
// 래핑된 에러가 있으면 해당 에러의 메시지를 포함합니다.
func (e *wrappedError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.cause)
	}
	return e.msg
}

// Unwrap returns the wrapped error.
// This allows errors.Is and errors.As to work correctly.
// Unwrap은 래핑된 에러를 반환합니다.
// 이것은 errors.Is 및 errors.As가 올바르게 작동하도록 합니다.
func (e *wrappedError) Unwrap() error {
	return e.cause
}

// codedError is an error with an associated string code.
// This is useful for API error responses and error categorization.
// codedError는 연결된 문자열 코드를 가진 에러입니다.
// 이것은 API 에러 응답 및 에러 분류에 유용합니다.
type codedError struct {
	// msg is the error message
	// msg는 에러 메시지입니다
	msg string
	// code is the error code
	// code는 에러 코드입니다
	code string
	// cause is the wrapped error
	// cause는 래핑된 에러입니다
	cause error
}

// Error returns the error message with code prefix.
// Error는 코드 접두사와 함께 에러 메시지를 반환합니다.
func (e *codedError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.code, e.msg, e.cause)
	}
	return fmt.Sprintf("[%s] %s", e.code, e.msg)
}

// Code returns the error code.
// Code는 에러 코드를 반환합니다.
func (e *codedError) Code() string {
	return e.code
}

// Unwrap returns the wrapped error.
// Unwrap은 래핑된 에러를 반환합니다.
func (e *codedError) Unwrap() error {
	return e.cause
}

// numericCodedError is an error with an associated numeric code.
// This is useful for HTTP status codes and numeric error codes.
// numericCodedError는 연결된 숫자 코드를 가진 에러입니다.
// 이것은 HTTP 상태 코드 및 숫자 에러 코드에 유용합니다.
type numericCodedError struct {
	// msg is the error message
	// msg는 에러 메시지입니다
	msg string
	// code is the numeric error code
	// code는 숫자 에러 코드입니다
	code int
	// cause is the wrapped error
	// cause는 래핑된 에러입니다
	cause error
}

// Error returns the error message with code prefix.
// Error는 코드 접두사와 함께 에러 메시지를 반환합니다.
func (e *numericCodedError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("[%d] %s: %v", e.code, e.msg, e.cause)
	}
	return fmt.Sprintf("[%d] %s", e.code, e.msg)
}

// Code returns the numeric error code.
// Code는 숫자 에러 코드를 반환합니다.
func (e *numericCodedError) Code() int {
	return e.code
}

// Unwrap returns the wrapped error.
// Unwrap은 래핑된 에러를 반환합니다.
func (e *numericCodedError) Unwrap() error {
	return e.cause
}

// stackError is an error that captures the stack trace at creation time.
// This helps with debugging by showing where the error originated.
// stackError는 생성 시점의 스택 트레이스를 캡처하는 에러입니다.
// 이것은 에러가 어디서 발생했는지 보여줌으로써 디버깅에 도움을 줍니다.
type stackError struct {
	// msg is the error message
	// msg는 에러 메시지입니다
	msg string
	// stack is the captured stack trace
	// stack은 캡처된 스택 트레이스입니다
	stack []Frame
	// cause is the wrapped error
	// cause는 래핑된 에러입니다
	cause error
}

// Error returns the error message.
// Error는 에러 메시지를 반환합니다.
func (e *stackError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.cause)
	}
	return e.msg
}

// StackTrace returns the captured stack frames.
// StackTrace는 캡처된 스택 프레임을 반환합니다.
func (e *stackError) StackTrace() []Frame {
	return e.stack
}

// Unwrap returns the wrapped error.
// Unwrap은 래핑된 에러를 반환합니다.
func (e *stackError) Unwrap() error {
	return e.cause
}

// contextError is an error that carries structured key-value contextual data.
// This provides additional information about the error condition.
// contextError는 구조화된 키-값 컨텍스트 데이터를 전달하는 에러입니다.
// 이것은 에러 조건에 대한 추가 정보를 제공합니다.
type contextError struct {
	// msg is the error message
	// msg는 에러 메시지입니다
	msg string
	// ctx is the contextual data
	// ctx는 컨텍스트 데이터입니다
	ctx map[string]interface{}
	// cause is the wrapped error
	// cause는 래핑된 에러입니다
	cause error
}

// Error returns the error message.
// Error는 에러 메시지를 반환합니다.
func (e *contextError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.cause)
	}
	return e.msg
}

// Context returns the contextual data.
// Context는 컨텍스트 데이터를 반환합니다.
func (e *contextError) Context() map[string]interface{} {
	// Return a copy to prevent external modification
	// 외부 수정을 방지하기 위해 복사본을 반환합니다
	ctx := make(map[string]interface{}, len(e.ctx))
	for k, v := range e.ctx {
		ctx[k] = v
	}
	return ctx
}

// Unwrap returns the wrapped error.
// Unwrap은 래핑된 에러를 반환합니다.
func (e *contextError) Unwrap() error {
	return e.cause
}

// compositeError combines multiple error features (code, stack, context).
// This is the most feature-rich error type in the package.
// compositeError는 여러 에러 기능(코드, 스택, 컨텍스트)을 결합합니다.
// 이것은 패키지에서 가장 기능이 풍부한 에러 타입입니다.
type compositeError struct {
	// msg is the error message
	// msg는 에러 메시지입니다
	msg string
	// code is the optional error code
	// code는 선택적 에러 코드입니다
	code string
	// numCode is the optional numeric error code
	// numCode는 선택적 숫자 에러 코드입니다
	numCode int
	// stack is the optional stack trace
	// stack은 선택적 스택 트레이스입니다
	stack []Frame
	// ctx is the optional contextual data
	// ctx는 선택적 컨텍스트 데이터입니다
	ctx map[string]interface{}
	// cause is the wrapped error
	// cause는 래핑된 에러입니다
	cause error
}

// Error returns the error message with code if present.
// Error는 코드가 있으면 코드와 함께 에러 메시지를 반환합니다.
func (e *compositeError) Error() string {
	prefix := ""
	if e.code != "" {
		prefix = fmt.Sprintf("[%s] ", e.code)
	} else if e.numCode != 0 {
		prefix = fmt.Sprintf("[%d] ", e.numCode)
	}

	if e.cause != nil {
		return fmt.Sprintf("%s%s: %v", prefix, e.msg, e.cause)
	}
	return fmt.Sprintf("%s%s", prefix, e.msg)
}

// Code returns the string error code if set.
// Code는 설정된 경우 문자열 에러 코드를 반환합니다.
func (e *compositeError) Code() string {
	return e.code
}

// NumericCode returns the numeric error code if set.
// NumericCode는 설정된 경우 숫자 에러 코드를 반환합니다.
func (e *compositeError) NumericCode() int {
	return e.numCode
}

// StackTrace returns the stack trace if captured.
// StackTrace는 캡처된 경우 스택 트레이스를 반환합니다.
func (e *compositeError) StackTrace() []Frame {
	return e.stack
}

// Context returns the contextual data if set.
// Context는 설정된 경우 컨텍스트 데이터를 반환합니다.
func (e *compositeError) Context() map[string]interface{} {
	if e.ctx == nil {
		return nil
	}
	// Return a copy to prevent external modification
	// 외부 수정을 방지하기 위해 복사본을 반환합니다
	ctx := make(map[string]interface{}, len(e.ctx))
	for k, v := range e.ctx {
		ctx[k] = v
	}
	return ctx
}

// Unwrap returns the wrapped error.
// Unwrap은 래핑된 에러를 반환합니다.
func (e *compositeError) Unwrap() error {
	return e.cause
}
