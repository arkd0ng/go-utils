package httputil

import (
	"fmt"
	"net/http"
)

// HTTPError represents an HTTP error with status code and response details.
// HTTPError는 상태 코드 및 응답 세부 정보를 포함하는 HTTP 에러를 나타냅니다.
type HTTPError struct {
	// HTTP status code
	// HTTP 상태 코드
	StatusCode int
	// HTTP status text
	// HTTP 상태 텍스트
	Status string
	// Response body
	// 응답 본문
	Body string
	// Request URL
	// 요청 URL
	URL string
	// HTTP method
	// HTTP 메서드
	Method string
}

// Error implements the error interface.
// Error는 error 인터페이스를 구현합니다.
func (e *HTTPError) Error() string {
	if e.Body != "" {
		return fmt.Sprintf("HTTP %d %s: %s (URL: %s %s, Body: %s)",
			e.StatusCode, e.Status, http.StatusText(e.StatusCode),
			e.Method, e.URL, e.Body)
	}
	return fmt.Sprintf("HTTP %d %s: %s (URL: %s %s)",
		e.StatusCode, e.Status, http.StatusText(e.StatusCode),
		e.Method, e.URL)
}

// IsHTTPError checks if an error is an HTTPError.
// IsHTTPError는 에러가 HTTPError인지 확인합니다.
func IsHTTPError(err error) bool {
	_, ok := err.(*HTTPError)
	return ok
}

// GetStatusCode extracts the status code from an HTTPError.
// Returns 0 if the error is not an HTTPError.
// GetStatusCode는 HTTPError에서 상태 코드를 추출합니다.
// 에러가 HTTPError가 아닌 경우 0을 반환합니다.
func GetStatusCode(err error) int {
	if httpErr, ok := err.(*HTTPError); ok {
		return httpErr.StatusCode
	}
	return 0
}

// RetryError represents an error after all retry attempts have failed.
// RetryError는 모든 재시도 시도가 실패한 후의 에러를 나타냅니다.
type RetryError struct {
	// Number of retry attempts
	// 재시도 시도 횟수
	Attempts int
	// Last error encountered
	// 마지막으로 발생한 에러
	LastErr error
	// Request URL
	// 요청 URL
	URL string
	// HTTP method
	// HTTP 메서드
	Method string
}

// Error implements the error interface.
// Error는 error 인터페이스를 구현합니다.
func (e *RetryError) Error() string {
	return fmt.Sprintf("request failed after %d attempts (URL: %s %s): %v",
		e.Attempts, e.Method, e.URL, e.LastErr)
}

// Unwrap returns the last error for error unwrapping.
// Unwrap은 에러 언래핑을 위해 마지막 에러를 반환합니다.
func (e *RetryError) Unwrap() error {
	return e.LastErr
}

// IsRetryError checks if an error is a RetryError.
// IsRetryError는 에러가 RetryError인지 확인합니다.
func IsRetryError(err error) bool {
	_, ok := err.(*RetryError)
	return ok
}

// TimeoutError represents a timeout error.
// TimeoutError는 타임아웃 에러를 나타냅니다.
type TimeoutError struct {
	// Request URL
	// 요청 URL
	URL string
	// HTTP method
	// HTTP 메서드
	Method string
}

// Error implements the error interface.
// Error는 error 인터페이스를 구현합니다.
func (e *TimeoutError) Error() string {
	return fmt.Sprintf("request timeout (URL: %s %s)", e.Method, e.URL)
}

// Timeout returns true to indicate this is a timeout error.
// Timeout은 이것이 타임아웃 에러임을 나타내기 위해 true를 반환합니다.
func (e *TimeoutError) Timeout() bool {
	return true
}

// IsTimeoutError checks if an error is a TimeoutError.
// IsTimeoutError는 에러가 TimeoutError인지 확인합니다.
func IsTimeoutError(err error) bool {
	_, ok := err.(*TimeoutError)
	return ok
}
