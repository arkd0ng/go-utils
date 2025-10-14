package redis

import (
	"context"
	"testing"
	"time"
)

func TestNewClientWithCustomOptions(t *testing.T) {
	client := newTestClient(t,
		WithPassword(""),
		WithDB(0),
		WithPoolSize(16),
		WithMinIdleConns(4),
		WithMaxRetries(5),
		WithRetryInterval(50*time.Millisecond),
		WithHealthCheck(true),
		WithHealthCheckInterval(2*time.Second),
		WithDialTimeout(2*time.Second),
		WithReadTimeout(1500*time.Millisecond),
		WithWriteTimeout(1500*time.Millisecond),
	)

	if got, want := client.config.PoolSize, 16; got != want {
		t.Fatalf("PoolSize mismatch: got %d, want %d", got, want)
	}
	if got, want := client.config.MinIdleConns, 4; got != want {
		t.Fatalf("MinIdleConns mismatch: got %d, want %d", got, want)
	}
	if got, want := client.config.MaxRetries, 5; got != want {
		t.Fatalf("MaxRetries mismatch: got %d, want %d", got, want)
	}
	if got, want := client.config.RetryInterval, 50*time.Millisecond; got != want {
		t.Fatalf("RetryInterval mismatch: got %v, want %v", got, want)
	}
	if got, want := client.config.HealthCheckInterval, 2*time.Second; got != want {
		t.Fatalf("HealthCheckInterval mismatch: got %v, want %v", got, want)
	}
	if got, want := client.config.DialTimeout, 2*time.Second; got != want {
		t.Fatalf("DialTimeout mismatch: got %v, want %v", got, want)
	}
	if got, want := client.config.ReadTimeout, 1500*time.Millisecond; got != want {
		t.Fatalf("ReadTimeout mismatch: got %v, want %v", got, want)
	}
	if got, want := client.config.WriteTimeout, 1500*time.Millisecond; got != want {
		t.Fatalf("WriteTimeout mismatch: got %v, want %v", got, want)
	}
}

func TestPing(t *testing.T) {
	client := newTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx); err != nil {
		t.Fatalf("expected successful ping, got error: %v", err)
	}
}

func TestClosePreventsFurtherCommands(t *testing.T) {
	ensureRedisRunning(t)

	client, err := New(WithAddr(testRedisAddr))
	if err != nil {
		t.Fatalf("failed to create redis client: %v", err)
	}

	ctx := context.Background()
	if err := client.rdb.FlushDB(ctx).Err(); err != nil {
		t.Fatalf("failed to flush db: %v", err)
	}

	if err := client.Close(); err != nil {
		t.Fatalf("close should not error, got %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := client.Ping(ctx); err == nil {
		t.Fatal("expected ping after close to fail, but it succeeded")
	}
}
