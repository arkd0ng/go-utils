package mysql

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/arkd0ng/go-utils/logging"
)

// config holds all configuration options for the MySQL client
// config는 MySQL 클라이언트의 모든 설정 옵션을 보유합니다
type config struct {
	// Basic connection / 기본 연결
	dsn string

	// Connection pool settings / 연결 풀 설정
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration
	connMaxIdleTime time.Duration

	// Credential rotation (optional) / 자격 증명 순환 (선택)
	credRefreshFunc  CredentialRefreshFunc // User-provided function / 사용자 제공 함수
	poolCount        int                   // Number of connection pools / 연결 풀 개수
	rotationInterval time.Duration         // Rotation interval / 순환 주기

	// Timeout settings / 타임아웃 설정
	connectTimeout time.Duration
	queryTimeout   time.Duration

	// Retry settings / 재시도 설정
	maxRetries int
	retryDelay time.Duration

	// Logging settings / 로깅 설정
	logger             *logging.Logger
	logQueries         bool
	logSlowQueries     bool
	slowQueryThreshold time.Duration

	// Health check settings / 헬스 체크 설정
	enableHealthCheck   bool
	healthCheckInterval time.Duration

	// Query statistics settings / 쿼리 통계 설정
	enableStats bool

	// Security settings / 보안 설정
	enableTLS bool
	tlsConfig *tls.Config
}

// defaultConfig returns the default configuration
// defaultConfig는 기본 설정을 반환합니다
func defaultConfig() *config {
	return &config{
		// Connection pool defaults / 연결 풀 기본값
		maxOpenConns:    25,                // 25 connections / 25개 연결
		maxIdleConns:    10,                // 10 idle connections / 10개 유휴 연결
		connMaxLifetime: 5 * time.Minute,   // 5 minutes / 5분
		connMaxIdleTime: 2 * time.Minute,   // 2 minutes / 2분
		poolCount:       1,                 // Single pool by default / 기본적으로 단일 풀
		rotationInterval: 1 * time.Hour,    // 1 hour default rotation / 1시간 기본 순환

		// Timeout defaults / 타임아웃 기본값
		connectTimeout: 10 * time.Second, // 10 seconds / 10초
		queryTimeout:   30 * time.Second, // 30 seconds / 30초

		// Retry defaults / 재시도 기본값
		maxRetries: 3,                      // 3 retries / 3번 재시도
		retryDelay: 100 * time.Millisecond, // 100ms between retries / 재시도 간 100ms

		// Logging defaults / 로깅 기본값
		logQueries:         false,          // Don't log all queries by default / 기본적으로 모든 쿼리 로깅 안 함
		logSlowQueries:     true,           // Log slow queries / 느린 쿼리 로깅
		slowQueryThreshold: 1 * time.Second, // Queries > 1s are slow / 1초 이상은 느린 쿼리

		// Health check defaults / 헬스 체크 기본값
		enableHealthCheck:   true,           // Enable by default / 기본적으로 활성화
		healthCheckInterval: 30 * time.Second, // Every 30 seconds / 30초마다

		// Query statistics defaults / 쿼리 통계 기본값
		enableStats: false, // Disabled by default for performance / 성능을 위해 기본적으로 비활성화

		// Security defaults / 보안 기본값
		enableTLS: false,
		tlsConfig: nil,
	}
}

// validate validates the configuration
// validate는 설정을 검증합니다
func (c *config) validate() error {
	// Check DSN or credential refresh function
	// DSN 또는 자격 증명 갱신 함수 확인
	if c.dsn == "" && c.credRefreshFunc == nil {
		return fmt.Errorf("%w: either DSN or credential refresh function must be provided", ErrInvalidConfig)
	}

	// Validate connection pool settings
	// 연결 풀 설정 검증
	if c.maxOpenConns <= 0 {
		return fmt.Errorf("%w: maxOpenConns must be positive", ErrInvalidConfig)
	}
	if c.maxIdleConns < 0 {
		return fmt.Errorf("%w: maxIdleConns must be non-negative", ErrInvalidConfig)
	}
	if c.maxIdleConns > c.maxOpenConns {
		return fmt.Errorf("%w: maxIdleConns cannot exceed maxOpenConns", ErrInvalidConfig)
	}

	// Validate pool count for credential rotation
	// 자격 증명 순환을 위한 풀 개수 검증
	if c.credRefreshFunc != nil {
		if c.poolCount < 2 {
			return fmt.Errorf("%w: poolCount must be at least 2 for credential rotation", ErrInvalidConfig)
		}
		if c.rotationInterval <= 0 {
			return fmt.Errorf("%w: rotationInterval must be positive", ErrInvalidConfig)
		}
	}

	// Validate timeout settings
	// 타임아웃 설정 검증
	if c.connectTimeout <= 0 {
		return fmt.Errorf("%w: connectTimeout must be positive", ErrInvalidConfig)
	}
	if c.queryTimeout <= 0 {
		return fmt.Errorf("%w: queryTimeout must be positive", ErrInvalidConfig)
	}

	// Validate retry settings
	// 재시도 설정 검증
	if c.maxRetries < 0 {
		return fmt.Errorf("%w: maxRetries must be non-negative", ErrInvalidConfig)
	}
	if c.retryDelay < 0 {
		return fmt.Errorf("%w: retryDelay must be non-negative", ErrInvalidConfig)
	}

	// Validate health check settings
	// 헬스 체크 설정 검증
	if c.enableHealthCheck && c.healthCheckInterval <= 0 {
		return fmt.Errorf("%w: healthCheckInterval must be positive when health check is enabled", ErrInvalidConfig)
	}

	return nil
}
