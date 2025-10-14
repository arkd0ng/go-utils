package mysql

import (
	"context"
	"database/sql"
	"sync"
	"testing"
)

// TestGetPoolMetrics tests pool metrics retrieval
// TestGetPoolMetrics는 풀 메트릭 검색을 테스트합니다
func TestGetPoolMetrics(t *testing.T) {
	// Create a client with mock setup / 모의 설정으로 클라이언트 생성
	client := &Client{
		config:        &config{dsn: "invalid-dsn"},
		connections:   []*sql.DB{},
		connectionsMu: sync.RWMutex{},
	}

	metrics := client.GetPoolMetrics()

	if metrics.PoolCount != 0 {
		t.Errorf("PoolCount = %d, want 0", metrics.PoolCount)
	}

	if metrics.TotalConnections != 0 {
		t.Errorf("TotalConnections = %d, want 0", metrics.TotalConnections)
	}

	if !metrics.Healthy {
		t.Error("Healthy should be true for empty pool")
	}

	if len(metrics.PoolStats) != 0 {
		t.Errorf("PoolStats length = %d, want 0", len(metrics.PoolStats))
	}
}

// TestGetPoolHealthInfo tests health info retrieval
// TestGetPoolHealthInfo는 상태 정보 검색을 테스트합니다
func TestGetPoolHealthInfo(t *testing.T) {
	client := &Client{
		config:        &config{dsn: "invalid-dsn"},
		connections:   []*sql.DB{},
		connectionsMu: sync.RWMutex{},
	}

	ctx := context.Background()
	health := client.GetPoolHealthInfo(ctx)

	if !health.Healthy {
		t.Error("Healthy should be true for empty pool")
	}

	if len(health.UnhealthyPool) != 0 {
		t.Errorf("UnhealthyPool length = %d, want 0", len(health.UnhealthyPool))
	}

	if len(health.Details) != 0 {
		t.Errorf("Details length = %d, want 0", len(health.Details))
	}

	if health.CheckDuration < 0 {
		t.Error("CheckDuration should not be negative")
	}
}

// TestGetPoolStats tests raw pool statistics
// TestGetPoolStats는 원시 풀 통계를 테스트합니다
func TestGetPoolStats(t *testing.T) {
	client := &Client{
		config:        &config{dsn: "invalid-dsn"},
		connections:   []*sql.DB{},
		connectionsMu: sync.RWMutex{},
	}

	stats := client.GetPoolStats()

	if len(stats) != 0 {
		t.Errorf("GetPoolStats() length = %d, want 0", len(stats))
	}
}

// TestGetConnectionUtilization tests connection utilization calculation
// TestGetConnectionUtilization는 연결 사용률 계산을 테스트합니다
func TestGetConnectionUtilization(t *testing.T) {
	client := &Client{
		config:        &config{dsn: "invalid-dsn"},
		connections:   []*sql.DB{},
		connectionsMu: sync.RWMutex{},
	}

	utilization := client.GetConnectionUtilization()

	if len(utilization) != 0 {
		t.Errorf("GetConnectionUtilization() length = %d, want 0", len(utilization))
	}
}

// TestGetWaitStatistics tests wait statistics retrieval
// TestGetWaitStatistics는 대기 통계 검색을 테스트합니다
func TestGetWaitStatistics(t *testing.T) {
	client := &Client{
		config:        &config{dsn: "invalid-dsn"},
		connections:   []*sql.DB{},
		connectionsMu: sync.RWMutex{},
	}

	waitStats := client.GetWaitStatistics()

	if len(waitStats) != 0 {
		t.Errorf("GetWaitStatistics() length = %d, want 0", len(waitStats))
	}
}

// TestGetCurrentConnectionIndex tests current connection index
// TestGetCurrentConnectionIndex는 현재 연결 인덱스를 테스트합니다
func TestGetCurrentConnectionIndex(t *testing.T) {
	client := &Client{
		config:        &config{dsn: "invalid-dsn"},
		connections:   []*sql.DB{},
		connectionsMu: sync.RWMutex{},
		currentIdx:    0,
	}

	idx := client.GetCurrentConnectionIndex()

	if idx != 0 {
		t.Errorf("GetCurrentConnectionIndex() = %d, want 0", idx)
	}
}

// TestGetRotationIndex tests rotation index
// TestGetRotationIndex는 순환 인덱스를 테스트합니다
func TestGetRotationIndex(t *testing.T) {
	client := &Client{
		config:        &config{dsn: "invalid-dsn"},
		connections:   []*sql.DB{},
		connectionsMu: sync.RWMutex{},
		rotationIdx:   0,
	}

	idx := client.GetRotationIndex()

	if idx != 0 {
		t.Errorf("GetRotationIndex() = %d, want 0", idx)
	}
}

// TestPoolMetricsWithRealDB tests pool metrics with actual database
// TestPoolMetricsWithRealDB는 실제 데이터베이스로 풀 메트릭을 테스트합니다
func TestPoolMetricsWithRealDB(t *testing.T) {
	// This test would require a real database connection
	// 이 테스트는 실제 데이터베이스 연결이 필요합니다
	t.Skip("Requires actual database connection")
}

// TestPoolHealthWithRealDB tests pool health with actual database
// TestPoolHealthWithRealDB는 실제 데이터베이스로 풀 상태를 테스트합니다
func TestPoolHealthWithRealDB(t *testing.T) {
	// This test would require a real database connection
	// 이 테스트는 실제 데이터베이스 연결이 필요합니다
	t.Skip("Requires actual database connection")
}

// TestWaitStatisticsCalculation tests wait statistics calculation
// TestWaitStatisticsCalculation는 대기 통계 계산을 테스트합니다
func TestWaitStatisticsCalculation(t *testing.T) {
	// WaitStatistic structure test / WaitStatistic 구조체 테스트
	ws := WaitStatistic{
		PoolIndex:   0,
		WaitCount:   10,
		AvgWaitTime: 100,
	}

	if ws.PoolIndex != 0 {
		t.Errorf("PoolIndex = %d, want 0", ws.PoolIndex)
	}

	if ws.WaitCount != 10 {
		t.Errorf("WaitCount = %d, want 10", ws.WaitCount)
	}
}

// TestPoolMetricsStructure tests PoolMetrics structure
// TestPoolMetricsStructure는 PoolMetrics 구조체를 테스트합니다
func TestPoolMetricsStructure(t *testing.T) {
	metrics := PoolMetrics{
		PoolCount:        2,
		TotalConnections: 20,
		Healthy:          true,
		PoolStats:        []PoolStat{},
	}

	if metrics.PoolCount != 2 {
		t.Errorf("PoolCount = %d, want 2", metrics.PoolCount)
	}

	if metrics.TotalConnections != 20 {
		t.Errorf("TotalConnections = %d, want 20", metrics.TotalConnections)
	}

	if !metrics.Healthy {
		t.Error("Healthy should be true")
	}
}

// TestPoolStatStructure tests PoolStat structure
// TestPoolStatStructure는 PoolStat 구조체를 테스트합니다
func TestPoolStatStructure(t *testing.T) {
	stat := PoolStat{
		Index:            0,
		MaxOpenConns:     100,
		OpenConnections:  50,
		InUse:            25,
		Idle:             25,
		WaitCount:        10,
		WaitDuration:     1000,
		MaxIdleClosed:    5,
		MaxIdleTimeClosed: 3,
		MaxLifetimeClosed: 2,
	}

	if stat.Index != 0 {
		t.Errorf("Index = %d, want 0", stat.Index)
	}

	if stat.InUse+stat.Idle != stat.OpenConnections {
		t.Error("InUse + Idle should equal OpenConnections")
	}
}

// TestPoolHealthInfoStructure tests PoolHealthInfo structure
// TestPoolHealthInfoStructure는 PoolHealthInfo 구조체를 테스트합니다
func TestPoolHealthInfoStructure(t *testing.T) {
	health := PoolHealthInfo{
		Healthy:       true,
		UnhealthyPool: []int{},
		Details:       []PoolHealth{},
	}

	if !health.Healthy {
		t.Error("Healthy should be true")
	}

	if len(health.UnhealthyPool) != 0 {
		t.Errorf("UnhealthyPool length = %d, want 0", len(health.UnhealthyPool))
	}
}

// TestPoolHealthStructure tests PoolHealth structure
// TestPoolHealthStructure는 PoolHealth 구조체를 테스트합니다
func TestPoolHealthStructure(t *testing.T) {
	health := PoolHealth{
		Index:    0,
		Healthy:  true,
		PingTime: 10,
		Error:    nil,
	}

	if health.Index != 0 {
		t.Errorf("Index = %d, want 0", health.Index)
	}

	if !health.Healthy {
		t.Error("Healthy should be true")
	}

	if health.Error != nil {
		t.Error("Error should be nil")
	}
}

// BenchmarkGetPoolMetrics benchmarks pool metrics retrieval
// BenchmarkGetPoolMetrics는 풀 메트릭 검색을 벤치마크합니다
func BenchmarkGetPoolMetrics(b *testing.B) {
	client := &Client{
		config:        &config{dsn: "invalid-dsn"},
		connections:   []*sql.DB{},
		connectionsMu: sync.RWMutex{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.GetPoolMetrics()
	}
}

// BenchmarkGetConnectionUtilization benchmarks utilization calculation
// BenchmarkGetConnectionUtilization는 사용률 계산을 벤치마크합니다
func BenchmarkGetConnectionUtilization(b *testing.B) {
	client := &Client{
		config:        &config{dsn: "invalid-dsn"},
		connections:   []*sql.DB{},
		connectionsMu: sync.RWMutex{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.GetConnectionUtilization()
	}
}
