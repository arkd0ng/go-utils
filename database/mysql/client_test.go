package mysql

import (
	"testing"
	"time"
)

// TestNew tests the client creation
// TestNew는 클라이언트 생성을 테스트합니다
func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		opts    []Option
		wantErr bool
	}{
		{
			name: "valid DSN",
			opts: []Option{
				WithDSN("user:pass@tcp(localhost:3306)/test"),
			},
			wantErr: true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
		{
			name:    "no DSN",
			opts:    []Option{},
			wantErr: true, // Should fail validation / 검증 실패해야 함
		},
		{
			name: "with credential refresh",
			opts: []Option{
				WithCredentialRefresh(
					func() (string, error) {
						return "user:pass@tcp(localhost:3306)/test", nil
					},
					3,
					1*time.Hour,
				),
			},
			wantErr: true, // Will fail without actual DB / 실제 DB 없이는 실패
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestConfig tests the configuration
// TestConfig는 설정을 테스트합니다
func TestConfig(t *testing.T) {
	t.Run("default config", func(t *testing.T) {
		cfg := defaultConfig()
		if cfg.maxOpenConns <= 0 {
			t.Error("maxOpenConns should be positive")
		}
		if cfg.maxIdleConns < 0 {
			t.Error("maxIdleConns should be non-negative")
		}
	})

	t.Run("validate - no DSN", func(t *testing.T) {
		cfg := defaultConfig()
		cfg.dsn = ""
		cfg.credRefreshFunc = nil
		if err := cfg.validate(); err == nil {
			t.Error("validate should fail without DSN or credential function")
		}
	})

	t.Run("validate - invalid pool count", func(t *testing.T) {
		cfg := defaultConfig()
		cfg.dsn = ""
		cfg.credRefreshFunc = func() (string, error) { return "", nil }
		cfg.poolCount = 1 // Should be at least 2 for rotation / 순환을 위해 최소 2개
		if err := cfg.validate(); err == nil {
			t.Error("validate should fail with poolCount < 2 for credential rotation")
		}
	})
}

// TestOptions tests the option functions
// TestOptions는 옵션 함수를 테스트합니다
func TestOptions(t *testing.T) {
	t.Run("WithDSN", func(t *testing.T) {
		cfg := defaultConfig()
		opt := WithDSN("test-dsn")
		if err := opt(cfg); err != nil {
			t.Fatalf("WithDSN failed: %v", err)
		}
		if cfg.dsn != "test-dsn" {
			t.Errorf("dsn = %v, want %v", cfg.dsn, "test-dsn")
		}
	})

	t.Run("WithMaxOpenConns", func(t *testing.T) {
		cfg := defaultConfig()
		opt := WithMaxOpenConns(50)
		if err := opt(cfg); err != nil {
			t.Fatalf("WithMaxOpenConns failed: %v", err)
		}
		if cfg.maxOpenConns != 50 {
			t.Errorf("maxOpenConns = %v, want %v", cfg.maxOpenConns, 50)
		}
	})

	t.Run("WithCredentialRefresh", func(t *testing.T) {
		cfg := defaultConfig()
		fn := func() (string, error) { return "dsn", nil }
		opt := WithCredentialRefresh(fn, 3, 1*time.Hour)
		if err := opt(cfg); err != nil {
			t.Fatalf("WithCredentialRefresh failed: %v", err)
		}
		if cfg.credRefreshFunc == nil {
			t.Error("credRefreshFunc should be set")
		}
		if cfg.poolCount != 3 {
			t.Errorf("poolCount = %v, want %v", cfg.poolCount, 3)
		}
		if cfg.rotationInterval != 1*time.Hour {
			t.Errorf("rotationInterval = %v, want %v", cfg.rotationInterval, 1*time.Hour)
		}
	})
}
