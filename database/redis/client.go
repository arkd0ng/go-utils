package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// New creates a new Redis client with the given options
// New는 주어진 옵션으로 새로운 Redis 클라이언트를 생성합니다
func New(opts ...Option) (*Client, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	// Validate configuration / 설정 검증
	if cfg.Addr == "" {
		return nil, ErrInvalidAddr
	}

	// Create Redis client / Redis 클라이언트 생성
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	// Test connection / 연결 테스트
	ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	client := &Client{
		rdb:    rdb,
		config: cfg,
		done:   make(chan struct{}),
	}

	// Start health check goroutine / 헬스 체크 고루틴 시작
	if cfg.EnableHealthCheck {
		go client.healthCheck()
	}

	return client, nil
}
