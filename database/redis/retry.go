package redis

import (
	"context"
	"fmt"
	"time"
)

// executeWithRetry executes a Redis command with retry logic
// executeWithRetry는 재시도 로직으로 Redis 명령을 실행합니다
func (c *Client) executeWithRetry(ctx context.Context, fn func() error) error {
	var lastErr error

	for i := 0; i < c.config.MaxRetries; i++ {
		// Check context cancellation / Context 취소 확인
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Execute the function / 함수 실행
		startTime := time.Now()
		err := fn()
		duration := time.Since(startTime)

		// Success / 성공
		if err == nil {
			return nil
		}

		// Check if error is retriable / 에러가 재시도 가능한지 확인
		if !isRetriableError(err) {
			return &RedisError{
				Op:       "execute",
				Err:      err,
				Time:     startTime,
				Duration: duration,
			}
		}

		lastErr = err

		// Don't sleep on last retry / 마지막 재시도에서는 sleep하지 않음
		if i < c.config.MaxRetries-1 {
			// Exponential backoff / 지수 백오프
			backoff := time.Duration(i+1) * c.config.RetryInterval
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(backoff):
				continue
			}
		}
	}

	return &RedisError{
		Op:  "execute",
		Err: fmt.Errorf("max retries exceeded: %w", lastErr),
	}
}
