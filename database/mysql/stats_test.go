package mysql

import (
	"testing"
	"time"
)

func TestQueryStatsAndSlowQueries(t *testing.T) {
	client := newTestClient(t, WithQueryStats(true))
	client.EnableQueryStats()
	client.EnableSlowQueryLog(1*time.Nanosecond, nil)

	ctx := testContext()
	resetTable(t, "users")

	if _, err := client.Exec(ctx, "INSERT INTO users (name, email, age, city) VALUES (?,?,?,?)",
		"Stats User", uniqueEmail("stats"), 30, "Seoul"); err != nil {
		t.Fatalf("Exec insert failed: %v", err)
	}

	rows, err := client.Query(ctx, "SELECT * FROM users")
	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}
	rows.Close()

	stats := client.GetQueryStats()
	if stats.TotalQueries == 0 {
		t.Fatal("expected query stats to record queries")
	}

	slow := client.GetSlowQueries(10)
	if len(slow) == 0 {
		t.Fatal("expected slow query log entries when threshold is 0")
	}

	client.ResetQueryStats()
	stats = client.GetQueryStats()
	if stats.TotalQueries != 0 {
		t.Fatal("expected stats reset to clear counters")
	}
}

func TestQueryStatsToggle(t *testing.T) {
	client := newTestClient(t, WithQueryStats(true))
	client.EnableQueryStats()

	client.DisableQueryStats()

	if stats := client.GetQueryStats(); stats.TotalQueries != 0 {
		t.Fatalf("expected no stats recorded when disabled, got %d", stats.TotalQueries)
	}

	client.EnableQueryStats()
	ctx := testContext()
	resetTable(t, "users")

	rows, err := client.Query(ctx, "SELECT 1")
	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}
	rows.Close()

	if stats := client.GetQueryStats(); stats.TotalQueries == 0 {
		t.Fatal("expected stats recorded after re-enabling")
	}
}
