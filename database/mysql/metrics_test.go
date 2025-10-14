package mysql

import "testing"

func TestPoolMetrics(t *testing.T) {
	client := newTestClient(t, WithQueryStats(true))
	ctx := testContext()

	metrics := client.GetPoolMetrics()
	if metrics.PoolCount == 0 {
		t.Fatal("expected PoolCount > 0")
	}

	health := client.GetPoolHealthInfo(ctx)
	if len(health.Details) == 0 {
		t.Fatal("expected health details for pools")
	}

	stats := client.GetPoolStats()
	if len(stats) == 0 {
		t.Fatal("expected pool stats slice")
	}

	util := client.GetConnectionUtilization()
	if len(util) == 0 {
		t.Fatal("expected utilization entries")
	}
}
