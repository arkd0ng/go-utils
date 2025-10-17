package redis

import (
	"context"
	"fmt"
	"time"
)

// healthCheck runs periodic health checks on the Redis connection
// healthCheck는 Redis 연결에 대한 정기적인 헬스 체크를 실행합니다
func (c *Client) healthCheck() {
	ticker := time.NewTicker(c.config.HealthCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			err := c.rdb.Ping(ctx).Err()
			cancel()

			if err != nil {
				// Log error but don't panic
				// 에러 로깅하지만 패닉하지 않음
				fmt.Printf("Redis health check failed: %v\n", err)
			}

		case <-c.done:
			return
		}
	}
}

// Ping checks the connection to Redis
// Ping은 Redis 연결을 확인합니다
func (c *Client) Ping(ctx context.Context) error {
	return c.executeWithRetry(ctx, func() error {
		return c.rdb.Ping(ctx).Err()
	})
}

// Close closes the Redis connection
// Close는 Redis 연결을 닫습니다
func (c *Client) Close() error {
	close(c.done)
	return c.rdb.Close()
}
