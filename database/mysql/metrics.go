package mysql

import (
	"context"
	"database/sql"
	"time"
)

// PoolMetrics represents connection pool metrics
// PoolMetrics는 연결 풀 메트릭을 나타냅니다
type PoolMetrics struct {
	PoolCount        int                 // Number of connection pools / 연결 풀 수
	TotalConnections int                 // Total connections across all pools / 모든 풀의 총 연결 수
	PoolStats        []PoolStat          // Individual pool statistics / 개별 풀 통계
	Healthy          bool                // Overall health status / 전체 상태
	LastChecked      time.Time           // Last health check time / 마지막 헬스 체크 시간
}

// PoolStat represents statistics for a single connection pool
// PoolStat는 단일 연결 풀의 통계를 나타냅니다
type PoolStat struct {
	Index            int           // Pool index / 풀 인덱스
	MaxOpenConns     int           // Maximum open connections / 최대 열린 연결 수
	OpenConnections  int           // Currently open connections / 현재 열린 연결 수
	InUse            int           // Connections in use / 사용 중인 연결 수
	Idle             int           // Idle connections / 유휴 연결 수
	WaitCount        int64         // Total number of connections waited for / 대기한 총 연결 수
	WaitDuration     time.Duration // Total time blocked waiting for connections / 연결 대기로 차단된 총 시간
	MaxIdleClosed    int64         // Total connections closed due to max idle / 최대 유휴로 인해 닫힌 총 연결 수
	MaxIdleTimeClosed int64        // Total connections closed due to max idle time / 최대 유휴 시간으로 인해 닫힌 총 연결 수
	MaxLifetimeClosed int64        // Total connections closed due to max lifetime / 최대 수명으로 인해 닫힌 총 연결 수
}

// PoolHealthInfo represents health information for connection pools
// PoolHealthInfo는 연결 풀의 상태 정보를 나타냅니다
type PoolHealthInfo struct {
	Healthy       bool          // Whether all pools are healthy / 모든 풀이 정상인지
	UnhealthyPool []int         // Indices of unhealthy pools / 비정상 풀의 인덱스
	LastCheck     time.Time     // Last health check time / 마지막 헬스 체크 시간
	CheckDuration time.Duration // Duration of last health check / 마지막 헬스 체크 소요 시간
	Details       []PoolHealth  // Health details per pool / 풀별 상태 세부 정보
}

// PoolHealth represents health status for a single pool
// PoolHealth는 단일 풀의 상태를 나타냅니다
type PoolHealth struct {
	Index      int           // Pool index / 풀 인덱스
	Healthy    bool          // Health status / 상태
	PingTime   time.Duration // Ping response time / Ping 응답 시간
	Error      error         // Error if unhealthy / 비정상인 경우 에러
}

// GetPoolMetrics returns connection pool metrics
// GetPoolMetrics는 연결 풀 메트릭을 반환합니다
//
// Example / 예제:
//
//	metrics := client.GetPoolMetrics()
//	fmt.Printf("Total pools: %d\n", metrics.PoolCount)
//	fmt.Printf("Total connections: %d\n", metrics.TotalConnections)
//
//	for _, pool := range metrics.PoolStats {
//	    fmt.Printf("Pool %d: %d in use, %d idle\n",
//	        pool.Index, pool.InUse, pool.Idle)
//	}
func (c *Client) GetPoolMetrics() PoolMetrics {
	c.connectionsMu.RLock()
	defer c.connectionsMu.RUnlock()

	metrics := PoolMetrics{
		PoolCount:   len(c.connections),
		PoolStats:   make([]PoolStat, 0, len(c.connections)),
		Healthy:     true,
		LastChecked: time.Now(),
	}

	totalConns := 0
	for i, db := range c.connections {
		stats := db.Stats()

		poolStat := PoolStat{
			Index:             i,
			MaxOpenConns:      stats.MaxOpenConnections,
			OpenConnections:   stats.OpenConnections,
			InUse:             stats.InUse,
			Idle:              stats.Idle,
			WaitCount:         stats.WaitCount,
			WaitDuration:      stats.WaitDuration,
			MaxIdleClosed:     stats.MaxIdleClosed,
			MaxIdleTimeClosed: stats.MaxIdleTimeClosed,
			MaxLifetimeClosed: stats.MaxLifetimeClosed,
		}

		metrics.PoolStats = append(metrics.PoolStats, poolStat)
		totalConns += stats.OpenConnections
	}

	metrics.TotalConnections = totalConns

	return metrics
}

// GetPoolHealthInfo performs a health check on all connection pools
// GetPoolHealthInfo는 모든 연결 풀에 대한 헬스 체크를 수행합니다
//
// Example / 예제:
//
//	ctx := context.Background()
//	health := client.GetPoolHealthInfo(ctx)
//
//	if !health.Healthy {
//	    fmt.Printf("Unhealthy pools: %v\n", health.UnhealthyPool)
//	    for _, detail := range health.Details {
//	        if !detail.Healthy {
//	            fmt.Printf("Pool %d error: %v\n", detail.Index, detail.Error)
//	        }
//	    }
//	}
func (c *Client) GetPoolHealthInfo(ctx context.Context) PoolHealthInfo {
	c.connectionsMu.RLock()
	defer c.connectionsMu.RUnlock()

	start := time.Now()
	info := PoolHealthInfo{
		Healthy:       true,
		UnhealthyPool: make([]int, 0),
		LastCheck:     start,
		Details:       make([]PoolHealth, 0, len(c.connections)),
	}

	for i, db := range c.connections {
		health := PoolHealth{
			Index: i,
		}

		pingStart := time.Now()
		err := db.PingContext(ctx)
		health.PingTime = time.Since(pingStart)

		if err != nil {
			health.Healthy = false
			health.Error = err
			info.Healthy = false
			info.UnhealthyPool = append(info.UnhealthyPool, i)
		} else {
			health.Healthy = true
		}

		info.Details = append(info.Details, health)
	}

	info.CheckDuration = time.Since(start)

	return info
}

// GetPoolStats returns raw database statistics for all pools
// GetPoolStats는 모든 풀의 원시 데이터베이스 통계를 반환합니다
//
// This is an alias for the Stats() method for backward compatibility.
// 이것은 하위 호환성을 위한 Stats() 메서드의 별칭입니다.
//
// Example / 예제:
//
//	stats := client.GetPoolStats()
//	for i, stat := range stats {
//	    fmt.Printf("Pool %d:\n", i)
//	    fmt.Printf("  Open: %d, InUse: %d, Idle: %d\n",
//	        stat.OpenConnections, stat.InUse, stat.Idle)
//	}
func (c *Client) GetPoolStats() []sql.DBStats {
	return c.Stats()
}

// GetConnectionUtilization returns connection utilization percentage for each pool
// GetConnectionUtilization은 각 풀의 연결 사용률을 반환합니다
//
// Returns a slice of utilization percentages (0-100) for each pool.
// 각 풀의 사용률 백분율(0-100)을 포함하는 슬라이스를 반환합니다.
//
// Example / 예제:
//
//	utilization := client.GetConnectionUtilization()
//	for i, util := range utilization {
//	    fmt.Printf("Pool %d utilization: %.2f%%\n", i, util)
//	    if util > 80 {
//	        fmt.Printf("  WARNING: High utilization on pool %d\n", i)
//	    }
//	}
func (c *Client) GetConnectionUtilization() []float64 {
	c.connectionsMu.RLock()
	defer c.connectionsMu.RUnlock()

	utilization := make([]float64, len(c.connections))

	for i, db := range c.connections {
		stats := db.Stats()
		if stats.MaxOpenConnections > 0 {
			utilization[i] = float64(stats.InUse) / float64(stats.MaxOpenConnections) * 100
		}
	}

	return utilization
}

// GetWaitStatistics returns wait statistics for all pools
// GetWaitStatistics는 모든 풀의 대기 통계를 반환합니다
//
// Returns wait count and average wait duration for each pool.
// 각 풀의 대기 횟수와 평균 대기 시간을 반환합니다.
//
// Example / 예제:
//
//	type WaitStat struct {
//	    PoolIndex   int
//	    WaitCount   int64
//	    AvgWaitTime time.Duration
//	}
//
//	waitStats := client.GetWaitStatistics()
//	for _, ws := range waitStats {
//	    if ws.WaitCount > 0 {
//	        fmt.Printf("Pool %d: %d waits, avg %v\n",
//	            ws.PoolIndex, ws.WaitCount, ws.AvgWaitTime)
//	    }
//	}
func (c *Client) GetWaitStatistics() []WaitStatistic {
	c.connectionsMu.RLock()
	defer c.connectionsMu.RUnlock()

	waitStats := make([]WaitStatistic, len(c.connections))

	for i, db := range c.connections {
		stats := db.Stats()
		ws := WaitStatistic{
			PoolIndex: i,
			WaitCount: stats.WaitCount,
		}

		if stats.WaitCount > 0 {
			ws.AvgWaitTime = time.Duration(int64(stats.WaitDuration) / stats.WaitCount)
		}

		waitStats[i] = ws
	}

	return waitStats
}

// WaitStatistic represents wait statistics for a connection pool
// WaitStatistic는 연결 풀의 대기 통계를 나타냅니다
type WaitStatistic struct {
	PoolIndex   int           // Pool index / 풀 인덱스
	WaitCount   int64         // Number of waits / 대기 횟수
	AvgWaitTime time.Duration // Average wait time / 평균 대기 시간
}

// GetCurrentConnectionIndex returns the current connection index (for round-robin)
// GetCurrentConnectionIndex는 현재 연결 인덱스를 반환합니다(round-robin용)
//
// Example / 예제:
//
//	idx := client.GetCurrentConnectionIndex()
//	fmt.Printf("Next query will use pool %d\n", idx)
func (c *Client) GetCurrentConnectionIndex() int {
	c.connectionsMu.RLock()
	defer c.connectionsMu.RUnlock()

	return c.currentIdx
}

// GetRotationIndex returns the current rotation index
// GetRotationIndex는 현재 순환 인덱스를 반환합니다
//
// This is only relevant when credential rotation is enabled.
// 이것은 자격 증명 순환이 활성화된 경우에만 관련이 있습니다.
//
// Example / 예제:
//
//	rotIdx := client.GetRotationIndex()
//	fmt.Printf("Next rotation will replace pool %d\n", rotIdx)
func (c *Client) GetRotationIndex() int {
	c.connectionsMu.RLock()
	defer c.connectionsMu.RUnlock()

	return c.rotationIdx
}
