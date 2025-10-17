package mysql

import (
	"context"
	"database/sql"
	"time"
)

// startHealthCheck starts the health check goroutine
// startHealthCheck는 헬스 체크 goroutine을 시작합니다
func (c *Client) startHealthCheck() {
	go func() {
		ticker := time.NewTicker(c.config.healthCheckInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				c.performHealthCheck()
			case <-c.healthCheckStop:
				return
			}
		}
	}()
}

// performHealthCheck performs a health check on all connections
// performHealthCheck는 모든 연결에 대해 헬스 체크를 수행합니다
func (c *Client) performHealthCheck() {
	c.connectionsMu.RLock()
	connections := make([]*sql.DB, len(c.connections))
	copy(connections, c.connections)
	c.connectionsMu.RUnlock()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i, db := range connections {
		if err := db.PingContext(ctx); err != nil {
			if c.config.logger != nil {
				c.config.logger.Warn("Health check failed for connection",
					"index", i,
					"error", err)
			}
		}
	}
}

// ensureConnected ensures that the connection is healthy
// ensureConnected는 연결이 정상인지 확인합니다
func (c *Client) ensureConnected() error {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return ErrClosed
	}
	c.mu.RUnlock()

	// Try to ping
	// Ping 시도
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := c.getCurrentConnection()
	if err := db.PingContext(ctx); err != nil {
		if c.config.logger != nil {
			c.config.logger.Warn("Connection unhealthy",
				"error", err)
		}
		return err
	}

	return nil
}

// IsHealthy checks if the database connection is healthy
// IsHealthy는 데이터베이스 연결이 정상인지 확인합니다
func (c *Client) IsHealthy(ctx context.Context) bool {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return false
	}
	c.mu.RUnlock()

	db := c.getCurrentConnection()
	return db.PingContext(ctx) == nil
}
