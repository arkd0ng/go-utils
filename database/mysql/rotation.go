package mysql

import (
	"fmt"
	"time"
)

// startConnectionRotation starts the connection rotation goroutine
// startConnectionRotation은 연결 순환 goroutine을 시작합니다
func (c *Client) startConnectionRotation() {
	// Only start if credential refresh function is provided
	// 자격 증명 갱신 함수가 제공된 경우에만 시작
	if c.config.credRefreshFunc == nil {
		return
	}

	go func() {
		ticker := time.NewTicker(c.config.rotationInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := c.rotateOneConnection(); err != nil {
					if c.config.logger != nil {
						c.config.logger.Error("Connection rotation failed",
							"error", err)
					}
				}
			case <-c.rotationStop:
				return
			}
		}
	}()
}

// rotateOneConnection rotates one connection pool
// rotateOneConnection은 하나의 연결 풀을 순환합니다
func (c *Client) rotateOneConnection() error {
	// 1. Get new DSN from user function
	// 사용자 함수에서 새 DSN 가져오기
	newDSN, err := c.config.credRefreshFunc()
	if err != nil {
		return fmt.Errorf("credential refresh function failed: %w", err)
	}

	// 2. Create new connection
	// 새 연결 생성
	newDB, err := c.createConnection(newDSN)
	if err != nil {
		return fmt.Errorf("failed to create new connection: %w", err)
	}

	// 3. Find the connection to replace (round-robin)
	// 교체할 연결 찾기 (round-robin)
	c.connectionsMu.Lock()
	oldestIdx := c.rotationIdx % len(c.connections)
	c.rotationIdx++
	oldDB := c.connections[oldestIdx]

	// 4. Replace the connection
	// 연결 교체
	c.connections[oldestIdx] = newDB
	c.connectionsMu.Unlock()

	// 5. Close old connection gracefully (after 30 seconds)
	// 5. 오래된 연결을 우아하게 닫기 (30초 후)
	go func() {
		time.Sleep(30 * time.Second) // Wait for active queries to complete / 활성 쿼리 완료 대기
		oldDB.Close()
	}()

	if c.config.logger != nil {
		c.config.logger.Info("Connection rotated successfully",
			"pool_index", oldestIdx,
			"total_pools", len(c.connections))
	}

	return nil
}

// RotateNow immediately rotates one connection (for testing or manual rotation)
// RotateNow는 즉시 하나의 연결을 순환합니다 (테스트 또는 수동 순환용)
func (c *Client) RotateNow() error {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return ErrClosed
	}
	c.mu.RUnlock()

	if c.config.credRefreshFunc == nil {
		return fmt.Errorf("credential refresh function not configured")
	}

	return c.rotateOneConnection()
}
