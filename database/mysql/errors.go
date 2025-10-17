package mysql

import (
	"errors"
	"fmt"
	"time"
)

// Predefined errors
// 사전 정의된 에러
var (
	ErrConnectionFailed  = errors.New("database connection failed")
	ErrQueryFailed       = errors.New("query execution failed")
	ErrTransactionFailed = errors.New("transaction failed")
	ErrTimeout           = errors.New("operation timeout")
	ErrClosed            = errors.New("database connection closed")
	ErrInvalidDSN        = errors.New("invalid DSN format")
	ErrNoRows            = errors.New("no rows in result set")
	ErrInvalidConfig     = errors.New("invalid configuration")
)

// DBError represents a database operation error
// DBError는 데이터베이스 작업 에러를 나타냅니다
type DBError struct {
	// Operation name
	// 작업 이름
	Op string
	// SQL query
	// SQL 쿼리
	Query string
	// Query arguments
	// 쿼리 인자
	Args []interface{}
	// Original error
	// 원본 에러
	Err error
	// Error timestamp
	// 에러 타임스탬프
	Time time.Time
	// Operation duration
	// 작업 소요 시간
	Duration time.Duration
}

// Error implements error interface
// Error는 error 인터페이스를 구현합니다
func (e *DBError) Error() string {
	if e.Query != "" {
		return fmt.Sprintf("%s failed: %v (query: %s, duration: %v)",
			e.Op, e.Err, e.Query, e.Duration)
	}
	return fmt.Sprintf("%s failed: %v (duration: %v)",
		e.Op, e.Err, e.Duration)
}

// Unwrap returns the underlying error
// Unwrap은 기본 에러를 반환합니다
func (e *DBError) Unwrap() error {
	return e.Err
}

// isRetryableError checks if an error is retryable
// isRetryableError는 에러가 재시도 가능한지 확인합니다
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// MySQL error codes that are retryable
	// 재시도 가능한 MySQL 에러 코드
	errStr := err.Error()

	// Connection errors
	// 연결 에러
	retryableErrors := []string{
		"connection refused",
		"connection reset",
		"broken pipe",
		"no such host",
		"network is unreachable",
		"i/o timeout",
		"EOF",
		"driver: bad connection",
		"invalid connection",
		"MySQL server has gone away",
		"Error 1213", // Deadlock
		"Error 1205", // Lock wait timeout
	}

	for _, retryable := range retryableErrors {
		if contains(errStr, retryable) {
			return true
		}
	}

	return false
}

// isConnectionError checks if an error is connection-related
// isConnectionError는 에러가 연결 관련인지 확인합니다
func isConnectionError(err error) bool {
	if err == nil {
		return false
	}

	errStr := err.Error()
	connectionErrors := []string{
		"connection refused",
		"connection reset",
		"broken pipe",
		"EOF",
		"driver: bad connection",
		"invalid connection",
		"MySQL server has gone away",
	}

	for _, connErr := range connectionErrors {
		if contains(errStr, connErr) {
			return true
		}
	}

	return false
}

// contains checks if a string contains a substring (case-insensitive)
// contains는 문자열이 부분 문자열을 포함하는지 확인합니다
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr || len(s) > len(substr) &&
			containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
