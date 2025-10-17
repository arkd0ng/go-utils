package mysql

import (
	"sync"
	"time"
)

// QueryStats represents query execution statistics
// QueryStats는 쿼리 실행 통계를 나타냅니다
type QueryStats struct {
	TotalQueries   int64         // Total number of queries executed / 실행된 총 쿼리 수
	SuccessQueries int64         // Number of successful queries / 성공한 쿼리 수
	FailedQueries  int64         // Number of failed queries / 실패한 쿼리 수
	TotalDuration  time.Duration // Total execution time / 총 실행 시간
	AvgDuration    time.Duration // Average execution time / 평균 실행 시간
	SlowQueries    int64         // Number of slow queries / 느린 쿼리 수
}

// SlowQueryInfo represents information about a slow query
// SlowQueryInfo는 느린 쿼리에 대한 정보를 나타냅니다
type SlowQueryInfo struct {
	Query     string        // SQL query / SQL 쿼리
	Args      []interface{} // Query arguments / 쿼리 인자
	Duration  time.Duration // Execution time / 실행 시간
	Timestamp time.Time     // When the query was executed / 쿼리 실행 시간
}

// SlowQueryHandler is a callback function for handling slow queries
// SlowQueryHandler는 느린 쿼리를 처리하기 위한 콜백 함수입니다
type SlowQueryHandler func(info SlowQueryInfo)

// queryStatsTracker tracks query execution statistics
// queryStatsTracker는 쿼리 실행 통계를 추적합니다
type queryStatsTracker struct {
	mu                 sync.RWMutex        // Synchronization / 동기화
	totalQueries       int64               // Total queries / 총 쿼리
	successQueries     int64               // Success queries / 성공 쿼리
	failedQueries      int64               // Failed queries / 실패 쿼리
	totalDuration      time.Duration       // Total duration / 총 실행 시간
	slowQueries        int64               // Slow queries / 느린 쿼리
	slowQueryThreshold time.Duration       // Slow query threshold / 느린 쿼리 임계값
	slowQueryHandler   SlowQueryHandler    // Slow query handler / 느린 쿼리 핸들러
	slowQueryLog       []SlowQueryInfo     // Slow query log / 느린 쿼리 로그
	maxSlowQueryLog    int                 // Maximum slow query log size / 최대 느린 쿼리 로그 크기
	enabled            bool                // Whether stats tracking is enabled / 통계 추적 활성화 여부
}

// newQueryStatsTracker creates a new query stats tracker
// newQueryStatsTracker는 새 쿼리 통계 추적기를 생성합니다
func newQueryStatsTracker() *queryStatsTracker {
	return &queryStatsTracker{
		slowQueryLog:    make([]SlowQueryInfo, 0, 100),
		maxSlowQueryLog: 100, // Keep last 100 slow queries / 최근 100개의 느린 쿼리 유지
		enabled:         false,
	}
}

// recordQuery records a query execution
// recordQuery는 쿼리 실행을 기록합니다
func (t *queryStatsTracker) recordQuery(query string, args []interface{}, duration time.Duration, err error) {
	if !t.enabled {
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	t.totalQueries++
	t.totalDuration += duration

	if err != nil {
		t.failedQueries++
	} else {
		t.successQueries++
	}

	// Check if this is a slow query
	// 느린 쿼리인지 확인
	if t.slowQueryThreshold > 0 && duration >= t.slowQueryThreshold {
		t.slowQueries++

		info := SlowQueryInfo{
			Query:     query,
			Args:      args,
			Duration:  duration,
			Timestamp: time.Now(),
		}

		// Add to log
		// 로그에 추가
		if len(t.slowQueryLog) >= t.maxSlowQueryLog {
			// Remove oldest entry
			// 가장 오래된 항목 제거
			t.slowQueryLog = t.slowQueryLog[1:]
		}
		t.slowQueryLog = append(t.slowQueryLog, info)

		// Call handler if set
		// 핸들러가 설정된 경우 호출
		if t.slowQueryHandler != nil {
			// Call handler in goroutine to avoid blocking
			// 차단을 피하기 위해 고루틴에서 핸들러 호출
			go t.slowQueryHandler(info)
		}
	}
}

// getStats returns current statistics
// getStats는 현재 통계를 반환합니다
func (t *queryStatsTracker) getStats() QueryStats {
	t.mu.RLock()
	defer t.mu.RUnlock()

	stats := QueryStats{
		TotalQueries:   t.totalQueries,
		SuccessQueries: t.successQueries,
		FailedQueries:  t.failedQueries,
		TotalDuration:  t.totalDuration,
		SlowQueries:    t.slowQueries,
	}

	if t.totalQueries > 0 {
		stats.AvgDuration = time.Duration(int64(t.totalDuration) / t.totalQueries)
	}

	return stats
}

// reset resets all statistics
// reset은 모든 통계를 재설정합니다
func (t *queryStatsTracker) reset() {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.totalQueries = 0
	t.successQueries = 0
	t.failedQueries = 0
	t.totalDuration = 0
	t.slowQueries = 0
	t.slowQueryLog = make([]SlowQueryInfo, 0, t.maxSlowQueryLog)
}

// getSlowQueries returns recent slow queries
// getSlowQueries는 최근 느린 쿼리를 반환합니다
func (t *queryStatsTracker) getSlowQueries(limit int) []SlowQueryInfo {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if limit <= 0 || limit > len(t.slowQueryLog) {
		limit = len(t.slowQueryLog)
	}

	// Return the most recent entries
	// 가장 최근 항목 반환
	start := len(t.slowQueryLog) - limit
	result := make([]SlowQueryInfo, limit)
	copy(result, t.slowQueryLog[start:])

	return result
}

// enableSlowQueryLog enables slow query logging
// enableSlowQueryLog는 느린 쿼리 로깅을 활성화합니다
func (t *queryStatsTracker) enableSlowQueryLog(threshold time.Duration, handler SlowQueryHandler) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.slowQueryThreshold = threshold
	t.slowQueryHandler = handler
}

// GetQueryStats returns query execution statistics
// GetQueryStats는 쿼리 실행 통계를 반환합니다
//
// Example
// 예제:
//
//	stats := client.GetQueryStats()
//	fmt.Printf("Total queries: %d\n", stats.TotalQueries)
//	fmt.Printf("Success rate: %.2f%%\n",
//	    float64(stats.SuccessQueries)/float64(stats.TotalQueries)*100)
//	fmt.Printf("Average duration: %v\n", stats.AvgDuration)
func (c *Client) GetQueryStats() QueryStats {
	if c.statsTracker == nil {
		return QueryStats{}
	}
	return c.statsTracker.getStats()
}

// ResetQueryStats resets query execution statistics
// ResetQueryStats는 쿼리 실행 통계를 재설정합니다
//
// Example
// 예제:
//
// // Reset stats to start fresh
// 새로 시작하기 위해 통계 재설정
//	client.ResetQueryStats()
func (c *Client) ResetQueryStats() {
	if c.statsTracker != nil {
		c.statsTracker.reset()
	}
}

// EnableSlowQueryLog enables slow query logging with a threshold and optional handler
// EnableSlowQueryLog는 임계값과 선택적 핸들러로 느린 쿼리 로깅을 활성화합니다
//
// The handler function is called asynchronously for each slow query.
// 핸들러 함수는 각 느린 쿼리에 대해 비동기적으로 호출됩니다.
//
// Example
// 예제:
//
// // Log queries that take longer than 1 second
// 1초 이상 걸리는 쿼리 로깅
//	client.EnableSlowQueryLog(1*time.Second, func(info mysql.SlowQueryInfo) {
//	    log.Printf("Slow query detected: %s (took %v)", info.Query, info.Duration)
//	})
//
// Example with custom handling
// 커스텀 처리 예제:
//
// client.EnableSlowQueryLog(500*time.Millisecond, func(info mysql.SlowQueryInfo) {
// Send to monitoring system / 모니터링 시스템으로 전송
//	    metrics.RecordSlowQuery(info.Query, info.Duration)
//
// // Log with details
// 세부 정보와 함께 로깅
//	    logger.Warn("Slow query",
//	        "query", info.Query,
//	        "args", info.Args,
//	        "duration", info.Duration,
//	        "timestamp", info.Timestamp)
//	})
func (c *Client) EnableSlowQueryLog(threshold time.Duration, handler SlowQueryHandler) {
	if c.statsTracker != nil {
		c.statsTracker.enableSlowQueryLog(threshold, handler)
	}
}

// GetSlowQueries returns recent slow queries
// GetSlowQueries는 최근 느린 쿼리를 반환합니다
//
// The limit parameter specifies how many recent slow queries to return.
// limit 매개변수는 반환할 최근 느린 쿼리의 수를 지정합니다.
//
// Example
// 예제:
//
// // Get last 10 slow queries
// 최근 10개의 느린 쿼리 가져오기
//	slowQueries := client.GetSlowQueries(10)
//	for _, sq := range slowQueries {
//	    fmt.Printf("Query: %s\n", sq.Query)
//	    fmt.Printf("Duration: %v\n", sq.Duration)
//	    fmt.Printf("Timestamp: %v\n", sq.Timestamp)
//	    fmt.Println("---")
//	}
//
// Example with analysis
// 분석 예제:
//
//	slowQueries := client.GetSlowQueries(50)
//	if len(slowQueries) > 0 {
//	    var totalDuration time.Duration
//	    for _, sq := range slowQueries {
//	        totalDuration += sq.Duration
//	    }
//	    avgDuration := totalDuration / time.Duration(len(slowQueries))
//	    fmt.Printf("Average slow query duration: %v\n", avgDuration)
//	}
func (c *Client) GetSlowQueries(limit int) []SlowQueryInfo {
	if c.statsTracker == nil {
		return []SlowQueryInfo{}
	}
	return c.statsTracker.getSlowQueries(limit)
}

// EnableQueryStats enables query statistics tracking
// EnableQueryStats는 쿼리 통계 추적을 활성화합니다
//
// This must be called before any queries are executed to track statistics.
// 통계를 추적하려면 쿼리가 실행되기 전에 호출해야 합니다.
//
// Example
// 예제:
//
//	client, _ := mysql.New(mysql.WithDSN("..."))
//	client.EnableQueryStats()
//
// // Execute queries...
// 쿼리 실행...
//
// // Check statistics
// 통계 확인
//	stats := client.GetQueryStats()
//	fmt.Printf("Total queries: %d\n", stats.TotalQueries)
func (c *Client) EnableQueryStats() {
	if c.statsTracker != nil {
		c.statsTracker.mu.Lock()
		c.statsTracker.enabled = true
		c.statsTracker.mu.Unlock()
	}
}

// DisableQueryStats disables query statistics tracking
// DisableQueryStats는 쿼리 통계 추적을 비활성화합니다
//
// Existing statistics are preserved but no new queries will be tracked.
// 기존 통계는 보존되지만 새 쿼리는 추적되지 않습니다.
//
// Example
// 예제:
//
// // Temporarily disable stats tracking
// 일시적으로 통계 추적 비활성화
//	client.DisableQueryStats()
//	// ... perform operations ...
//	client.EnableQueryStats()
func (c *Client) DisableQueryStats() {
	if c.statsTracker != nil {
		c.statsTracker.mu.Lock()
		c.statsTracker.enabled = false
		c.statsTracker.mu.Unlock()
	}
}
