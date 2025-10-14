package mysql

import (
	"fmt"
	"testing"
	"time"
)

// TestQueryStatsTracker tests the query stats tracker
// TestQueryStatsTracker는 쿼리 통계 추적기를 테스트합니다
func TestQueryStatsTracker(t *testing.T) {
	tracker := newQueryStatsTracker()

	if tracker == nil {
		t.Fatal("newQueryStatsTracker() returned nil")
	}

	if tracker.enabled {
		t.Error("tracker should be disabled by default")
	}

	if len(tracker.slowQueryLog) != 0 {
		t.Errorf("slowQueryLog should be empty, got %d entries", len(tracker.slowQueryLog))
	}
}

// TestRecordQuery tests recording query execution
// TestRecordQuery는 쿼리 실행 기록을 테스트합니다
func TestRecordQuery(t *testing.T) {
	tracker := newQueryStatsTracker()
	tracker.enabled = true

	// Record successful query / 성공한 쿼리 기록
	tracker.recordQuery("SELECT * FROM users", nil, 100*time.Millisecond, nil)

	stats := tracker.getStats()
	if stats.TotalQueries != 1 {
		t.Errorf("TotalQueries = %d, want 1", stats.TotalQueries)
	}
	if stats.SuccessQueries != 1 {
		t.Errorf("SuccessQueries = %d, want 1", stats.SuccessQueries)
	}
	if stats.FailedQueries != 0 {
		t.Errorf("FailedQueries = %d, want 0", stats.FailedQueries)
	}

	// Record failed query / 실패한 쿼리 기록
	tracker.recordQuery("SELECT * FROM nonexistent", nil, 50*time.Millisecond, errTestError)

	stats = tracker.getStats()
	if stats.TotalQueries != 2 {
		t.Errorf("TotalQueries = %d, want 2", stats.TotalQueries)
	}
	if stats.FailedQueries != 1 {
		t.Errorf("FailedQueries = %d, want 1", stats.FailedQueries)
	}
}

// errTestError is a test error for testing purposes
// errTestError는 테스트 목적의 테스트 에러입니다
var errTestError = fmt.Errorf("test error")

// TestSlowQueryDetection tests slow query detection
// TestSlowQueryDetection는 느린 쿼리 감지를 테스트합니다
func TestSlowQueryDetection(t *testing.T) {
	tracker := newQueryStatsTracker()
	tracker.enabled = true
	tracker.slowQueryThreshold = 100 * time.Millisecond

	// Record fast query / 빠른 쿼리 기록
	tracker.recordQuery("SELECT * FROM users WHERE id = ?", []interface{}{1}, 50*time.Millisecond, nil)

	stats := tracker.getStats()
	if stats.SlowQueries != 0 {
		t.Errorf("SlowQueries = %d, want 0", stats.SlowQueries)
	}

	// Record slow query / 느린 쿼리 기록
	tracker.recordQuery("SELECT * FROM large_table", nil, 200*time.Millisecond, nil)

	stats = tracker.getStats()
	if stats.SlowQueries != 1 {
		t.Errorf("SlowQueries = %d, want 1", stats.SlowQueries)
	}

	slowQueries := tracker.getSlowQueries(10)
	if len(slowQueries) != 1 {
		t.Fatalf("getSlowQueries() returned %d queries, want 1", len(slowQueries))
	}

	if slowQueries[0].Duration != 200*time.Millisecond {
		t.Errorf("slow query duration = %v, want 200ms", slowQueries[0].Duration)
	}
}

// TestSlowQueryLog tests slow query log management
// TestSlowQueryLog는 느린 쿼리 로그 관리를 테스트합니다
func TestSlowQueryLog(t *testing.T) {
	tracker := newQueryStatsTracker()
	tracker.enabled = true
	tracker.slowQueryThreshold = 10 * time.Millisecond
	tracker.maxSlowQueryLog = 5

	// Record more slow queries than the log can hold
	// 로그가 보관할 수 있는 것보다 많은 느린 쿼리 기록
	for i := 0; i < 10; i++ {
		tracker.recordQuery("SELECT * FROM users", nil, 20*time.Millisecond, nil)
	}

	slowQueries := tracker.getSlowQueries(100)
	if len(slowQueries) != 5 {
		t.Errorf("slowQueryLog should contain max 5 entries, got %d", len(slowQueries))
	}
}

// TestSlowQueryHandler tests slow query handler callback
// TestSlowQueryHandler는 느린 쿼리 핸들러 콜백을 테스트합니다
func TestSlowQueryHandler(t *testing.T) {
	tracker := newQueryStatsTracker()
	tracker.enabled = true
	tracker.slowQueryThreshold = 50 * time.Millisecond

	called := false
	var receivedInfo SlowQueryInfo

	handler := func(info SlowQueryInfo) {
		called = true
		receivedInfo = info
	}

	tracker.enableSlowQueryLog(50*time.Millisecond, handler)

	// Record a slow query / 느린 쿼리 기록
	query := "SELECT * FROM large_table"
	args := []interface{}{1, 2, 3}
	duration := 100 * time.Millisecond

	tracker.recordQuery(query, args, duration, nil)

	// Give handler goroutine time to execute / 핸들러 고루틴 실행 시간 제공
	time.Sleep(50 * time.Millisecond)

	if !called {
		t.Error("slow query handler was not called")
	}

	if receivedInfo.Query != query {
		t.Errorf("handler received query = %v, want %v", receivedInfo.Query, query)
	}

	if receivedInfo.Duration != duration {
		t.Errorf("handler received duration = %v, want %v", receivedInfo.Duration, duration)
	}
}

// TestGetStats tests statistics retrieval
// TestGetStats는 통계 검색을 테스트합니다
func TestGetStats(t *testing.T) {
	tracker := newQueryStatsTracker()
	tracker.enabled = true

	// Record some queries with different durations
	// 다른 지속 시간으로 일부 쿼리 기록
	tracker.recordQuery("query1", nil, 100*time.Millisecond, nil)
	tracker.recordQuery("query2", nil, 200*time.Millisecond, nil)
	tracker.recordQuery("query3", nil, 300*time.Millisecond, errTestError)

	stats := tracker.getStats()

	if stats.TotalQueries != 3 {
		t.Errorf("TotalQueries = %d, want 3", stats.TotalQueries)
	}

	if stats.SuccessQueries != 2 {
		t.Errorf("SuccessQueries = %d, want 2", stats.SuccessQueries)
	}

	if stats.FailedQueries != 1 {
		t.Errorf("FailedQueries = %d, want 1", stats.FailedQueries)
	}

	expectedTotalDuration := 600 * time.Millisecond
	if stats.TotalDuration != expectedTotalDuration {
		t.Errorf("TotalDuration = %v, want %v", stats.TotalDuration, expectedTotalDuration)
	}

	expectedAvgDuration := 200 * time.Millisecond
	if stats.AvgDuration != expectedAvgDuration {
		t.Errorf("AvgDuration = %v, want %v", stats.AvgDuration, expectedAvgDuration)
	}
}

// TestResetStats tests statistics reset
// TestResetStats는 통계 재설정을 테스트합니다
func TestResetStats(t *testing.T) {
	tracker := newQueryStatsTracker()
	tracker.enabled = true
	tracker.slowQueryThreshold = 50 * time.Millisecond

	// Record some queries / 일부 쿼리 기록
	tracker.recordQuery("query1", nil, 100*time.Millisecond, nil)
	tracker.recordQuery("query2", nil, 200*time.Millisecond, nil)

	// Reset stats / 통계 재설정
	tracker.reset()

	stats := tracker.getStats()
	if stats.TotalQueries != 0 {
		t.Errorf("TotalQueries after reset = %d, want 0", stats.TotalQueries)
	}
	if stats.SuccessQueries != 0 {
		t.Errorf("SuccessQueries after reset = %d, want 0", stats.SuccessQueries)
	}
	if stats.TotalDuration != 0 {
		t.Errorf("TotalDuration after reset = %v, want 0", stats.TotalDuration)
	}
	if stats.SlowQueries != 0 {
		t.Errorf("SlowQueries after reset = %d, want 0", stats.SlowQueries)
	}

	slowQueries := tracker.getSlowQueries(100)
	if len(slowQueries) != 0 {
		t.Errorf("slowQueryLog after reset should be empty, got %d entries", len(slowQueries))
	}
}

// TestGetSlowQueries tests slow query retrieval with limit
// TestGetSlowQueries는 제한이 있는 느린 쿼리 검색을 테스트합니다
func TestGetSlowQueries(t *testing.T) {
	tracker := newQueryStatsTracker()
	tracker.enabled = true
	tracker.slowQueryThreshold = 10 * time.Millisecond

	// Record multiple slow queries / 여러 느린 쿼리 기록
	for i := 0; i < 10; i++ {
		tracker.recordQuery("SELECT * FROM users", nil, 50*time.Millisecond, nil)
	}

	tests := []struct {
		name      string
		limit     int
		wantCount int
	}{
		{
			name:      "get all queries",
			limit:     0,
			wantCount: 10,
		},
		{
			name:      "get last 5 queries",
			limit:     5,
			wantCount: 5,
		},
		{
			name:      "get more than available",
			limit:     20,
			wantCount: 10,
		},
		{
			name:      "negative limit",
			limit:     -1,
			wantCount: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queries := tracker.getSlowQueries(tt.limit)
			if len(queries) != tt.wantCount {
				t.Errorf("getSlowQueries(%d) returned %d queries, want %d",
					tt.limit, len(queries), tt.wantCount)
			}
		})
	}
}

// TestClientQueryStats tests client-level query stats methods
// TestClientQueryStats는 클라이언트 수준 쿼리 통계 메서드를 테스트합니다
func TestClientQueryStats(t *testing.T) {
	client := &Client{
		config:       &config{dsn: "invalid-dsn"},
		statsTracker: newQueryStatsTracker(),
	}

	// Enable stats tracking / 통계 추적 활성화
	client.EnableQueryStats()

	// Get stats / 통계 가져오기
	stats := client.GetQueryStats()
	if stats.TotalQueries != 0 {
		t.Errorf("initial TotalQueries = %d, want 0", stats.TotalQueries)
	}

	// Reset stats / 통계 재설정
	client.ResetQueryStats()

	// Disable stats / 통계 비활성화
	client.DisableQueryStats()

	// Get slow queries / 느린 쿼리 가져오기
	slowQueries := client.GetSlowQueries(10)
	if len(slowQueries) != 0 {
		t.Errorf("initial slow queries count = %d, want 0", len(slowQueries))
	}
}

// TestClientWithNilStatsTracker tests client methods with nil stats tracker
// TestClientWithNilStatsTracker는 nil 통계 추적기로 클라이언트 메서드를 테스트합니다
func TestClientWithNilStatsTracker(t *testing.T) {
	client := &Client{
		config:       &config{dsn: "invalid-dsn"},
		statsTracker: nil, // No stats tracker / 통계 추적기 없음
	}

	// These should not panic / 이것들은 패닉하지 않아야 함
	stats := client.GetQueryStats()
	if stats.TotalQueries != 0 {
		t.Error("GetQueryStats with nil tracker should return zero stats")
	}

	client.ResetQueryStats()
	client.EnableQueryStats()
	client.DisableQueryStats()

	slowQueries := client.GetSlowQueries(10)
	if len(slowQueries) != 0 {
		t.Error("GetSlowQueries with nil tracker should return empty slice")
	}

	client.EnableSlowQueryLog(100*time.Millisecond, nil)
}

// TestStatsTrackerDisabled tests that disabled tracker doesn't record
// TestStatsTrackerDisabled는 비활성화된 추적기가 기록하지 않는지 테스트합니다
func TestStatsTrackerDisabled(t *testing.T) {
	tracker := newQueryStatsTracker()
	// Don't enable tracker / 추적기를 활성화하지 않음

	// Record queries / 쿼리 기록
	tracker.recordQuery("query1", nil, 100*time.Millisecond, nil)
	tracker.recordQuery("query2", nil, 200*time.Millisecond, nil)

	stats := tracker.getStats()
	if stats.TotalQueries != 0 {
		t.Errorf("disabled tracker recorded TotalQueries = %d, want 0", stats.TotalQueries)
	}
}

// BenchmarkRecordQuery benchmarks query recording performance
// BenchmarkRecordQuery는 쿼리 기록 성능을 벤치마크합니다
func BenchmarkRecordQuery(b *testing.B) {
	tracker := newQueryStatsTracker()
	tracker.enabled = true

	query := "SELECT * FROM users WHERE id = ?"
	args := []interface{}{1}
	duration := 100 * time.Millisecond

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tracker.recordQuery(query, args, duration, nil)
	}
}

// BenchmarkGetStats benchmarks stats retrieval performance
// BenchmarkGetStats는 통계 검색 성능을 벤치마크합니다
func BenchmarkGetStats(b *testing.B) {
	tracker := newQueryStatsTracker()
	tracker.enabled = true

	// Pre-populate with some queries / 일부 쿼리로 미리 채우기
	for i := 0; i < 1000; i++ {
		tracker.recordQuery("query", nil, 100*time.Millisecond, nil)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = tracker.getStats()
	}
}
