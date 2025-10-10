package mysql

import (
	"context"
	"time"
)

// executeWithRetry executes a function with exponential backoff retry logic
// executeWithRetry는 지수 백오프 재시도 로직으로 함수를 실행합니다
func (c *Client) executeWithRetry(ctx context.Context, fn func() error) error {
	var lastErr error

	for attempt := 0; attempt <= c.config.maxRetries; attempt++ {
		// Execute function / 함수 실행
		err := fn()
		if err == nil {
			return nil // Success / 성공
		}

		lastErr = err

		// Check if error is retryable / 에러가 재시도 가능한지 확인
		if !isRetryableError(err) {
			if c.config.logger != nil {
				c.config.logger.Debug("Error is not retryable, failing immediately",
					"error", err,
					"attempt", attempt)
			}
			return err
		}

		// Don't retry on last attempt / 마지막 시도에서는 재시도 안 함
		if attempt == c.config.maxRetries {
			break
		}

		// Check context / Context 확인
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Calculate backoff delay (exponential) / 백오프 지연 계산 (지수)
		delay := c.config.retryDelay * time.Duration(1<<uint(attempt))
		if delay > 5*time.Second {
			delay = 5 * time.Second // Cap at 5 seconds / 최대 5초로 제한
		}

		if c.config.logger != nil {
			c.config.logger.Debug("Retrying after error",
				"error", err,
				"attempt", attempt+1,
				"max_retries", c.config.maxRetries,
				"delay", delay)
		}

		// Wait before retry / 재시도 전 대기
		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	if c.config.logger != nil {
		c.config.logger.Error("All retry attempts exhausted",
			"error", lastErr,
			"attempts", c.config.maxRetries+1)
	}

	return lastErr
}

// retryableQuery executes a query with retry logic
// retryableQuery는 재시도 로직으로 쿼리를 실행합니다
func (c *Client) retryableQuery(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	var result interface{}
	var execErr error

	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		rows, err := db.QueryContext(ctx, query, args...)
		if err != nil {
			execErr = err
			return err
		}
		result = rows
		execErr = nil
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, execErr
}

// retryableExec executes a non-query statement with retry logic
// retryableExec는 재시도 로직으로 쿼리가 아닌 문을 실행합니다
func (c *Client) retryableExec(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	var result interface{}
	var execErr error

	err := c.executeWithRetry(ctx, func() error {
		db := c.getCurrentConnection()
		res, err := db.ExecContext(ctx, query, args...)
		if err != nil {
			execErr = err
			return err
		}
		result = res
		execErr = nil
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, execErr
}
