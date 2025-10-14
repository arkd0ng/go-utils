package redis

import (
	"time"
)

// Config holds the Redis client configuration
// Config는 Redis 클라이언트 설정을 보유합니다
type Config struct {
	// Addr is the Redis server address (host:port)
	// Addr은 Redis 서버 주소입니다 (host:port)
	Addr string

	// Password for Redis authentication
	// Password는 Redis 인증 비밀번호입니다
	Password string

	// DB is the Redis database number (0-15)
	// DB는 Redis 데이터베이스 번호입니다 (0-15)
	DB int

	// PoolSize is the maximum number of socket connections
	// PoolSize는 최대 소켓 연결 수입니다
	PoolSize int

	// MinIdleConns is the minimum number of idle connections
	// MinIdleConns는 최소 유휴 연결 수입니다
	MinIdleConns int

	// MaxRetries is the maximum number of retries before giving up
	// MaxRetries는 포기하기 전 최대 재시도 횟수입니다
	MaxRetries int

	// RetryInterval is the interval between retries
	// RetryInterval은 재시도 사이 간격입니다
	RetryInterval time.Duration

	// EnableHealthCheck enables background health checking
	// EnableHealthCheck는 백그라운드 헬스 체크를 활성화합니다
	EnableHealthCheck bool

	// HealthCheckInterval is the interval between health checks
	// HealthCheckInterval은 헬스 체크 사이 간격입니다
	HealthCheckInterval time.Duration

	// DialTimeout is the timeout for establishing new connections
	// DialTimeout은 새 연결 설정 타임아웃입니다
	DialTimeout time.Duration

	// ReadTimeout is the timeout for socket reads
	// ReadTimeout은 소켓 읽기 타임아웃입니다
	ReadTimeout time.Duration

	// WriteTimeout is the timeout for socket writes
	// WriteTimeout은 소켓 쓰기 타임아웃입니다
	WriteTimeout time.Duration
}

// defaultConfig returns the default configuration
// defaultConfig는 기본 설정을 반환합니다
func defaultConfig() *Config {
	return &Config{
		Addr:                "localhost:6379",
		Password:            "",
		DB:                  0,
		PoolSize:            10,
		MinIdleConns:        5,
		MaxRetries:          3,
		RetryInterval:       100 * time.Millisecond,
		EnableHealthCheck:   true,
		HealthCheckInterval: 30 * time.Second,
		DialTimeout:         5 * time.Second,
		ReadTimeout:         3 * time.Second,
		WriteTimeout:        3 * time.Second,
	}
}
