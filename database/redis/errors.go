package redis

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Predefined errors / 사전 정의된 에러
var (
	ErrConnectionFailed = errors.New("redis connection failed")
	ErrCommandFailed    = errors.New("redis command failed")
	ErrTimeout          = errors.New("operation timeout")
	ErrClosed           = errors.New("redis connection closed")
	ErrInvalidAddr      = errors.New("invalid redis address")
	ErrNil              = errors.New("redis: nil")
)

// RedisError represents a Redis operation error
// RedisError는 Redis 작업 에러를 나타냅니다
type RedisError struct {
	Op       string        // Operation name / 작업 이름
	Key      string        // Redis key / Redis 키
	Args     []interface{} // Command arguments / 명령어 인자
	Err      error         // Original error / 원본 에러
	Time     time.Time     // Error timestamp / 에러 타임스탬프
	Duration time.Duration // Operation duration / 작업 소요 시간
}

// Error implements the error interface
// Error는 error 인터페이스를 구현합니다
func (e *RedisError) Error() string {
	if e.Key != "" {
		return fmt.Sprintf("redis %s failed for key '%s': %v (took %v)",
			e.Op, e.Key, e.Err, e.Duration)
	}
	return fmt.Sprintf("redis %s failed: %v (took %v)",
		e.Op, e.Err, e.Duration)
}

// Unwrap returns the wrapped error
// Unwrap은 래핑된 에러를 반환합니다
func (e *RedisError) Unwrap() error {
	return e.Err
}

// isRetriableError checks if an error is retriable
// isRetriableError는 에러가 재시도 가능한지 확인합니다
func isRetriableError(err error) bool {
	if err == nil {
		return false
	}

	// Context errors are not retriable / Context 에러는 재시도 불가
	if errors.Is(err, ErrTimeout) {
		return false
	}

	// Network errors are retriable / 네트워크 에러는 재시도 가능
	retriableErrors := []string{
		"connection refused",
		"connection reset",
		"broken pipe",
		"i/o timeout",
		"EOF",
		"connection closed",
		"no route to host",
	}

	errStr := strings.ToLower(err.Error())
	for _, retryErr := range retriableErrors {
		if strings.Contains(errStr, retryErr) {
			return true
		}
	}

	return false
}
