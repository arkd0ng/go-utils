package mysql

import (
	"testing"
	"time"
)

func TestNewRequiresDSN(t *testing.T) {
	if _, err := New(); err == nil {
		t.Fatal("expected error when DSN is missing")
	}
}

func TestNewWithValidDSN(t *testing.T) {
	client := newTestClient(t)
	if client.config.dsn != testMySQLDSN {
		t.Fatalf("dsn mismatch: got %s", client.config.dsn)
	}
}

func TestConfigValidation(t *testing.T) {
	cfg := defaultConfig()
	cfg.dsn = ""
	cfg.credRefreshFunc = nil
	if err := cfg.validate(); err == nil {
		t.Fatal("expected validation failure without DSN")
	}

	cfg = defaultConfig()
	cfg.dsn = ""
	cfg.credRefreshFunc = func() (string, error) { return "", nil }
	cfg.poolCount = 1
	if err := cfg.validate(); err == nil {
		t.Fatal("expected validation failure when poolCount < 2 for rotation")
	}
}

func TestOptionsApply(t *testing.T) {
	cfg := defaultConfig()

	if err := WithDSN("mysql-dsn")(cfg); err != nil {
		t.Fatalf("WithDSN failed: %v", err)
	}
	if cfg.dsn != "mysql-dsn" {
		t.Fatalf("dsn not set by WithDSN: %s", cfg.dsn)
	}

	if err := WithMaxOpenConns(42)(cfg); err != nil {
		t.Fatalf("WithMaxOpenConns failed: %v", err)
	}
	if cfg.maxOpenConns != 42 {
		t.Fatalf("maxOpenConns not applied: %d", cfg.maxOpenConns)
	}

	fn := func() (string, error) { return "rotation-dsn", nil }
	if err := WithCredentialRefresh(fn, 3, 30*time.Minute)(cfg); err != nil {
		t.Fatalf("WithCredentialRefresh failed: %v", err)
	}
	if cfg.credRefreshFunc == nil || cfg.poolCount != 3 {
		t.Fatal("credential refresh options not applied")
	}
}

func TestClientClosePreventsQueries(t *testing.T) {
	client := newTestClient(t)

	if err := client.Close(); err != nil {
		t.Fatalf("Close failed: %v", err)
	}

	if _, err := client.SelectAll("users"); err == nil {
		t.Fatal("expected SelectAll after Close to fail")
	}
}
