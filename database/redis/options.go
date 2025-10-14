package redis

import (
	"time"
)

// Option is a functional option for configuring the Redis client
// Option은 Redis 클라이언트 설정을 위한 함수형 옵션입니다
type Option func(*Config)

// WithAddr sets the Redis server address
// WithAddr은 Redis 서버 주소를 설정합니다
func WithAddr(addr string) Option {
	return func(c *Config) {
		c.Addr = addr
	}
}

// WithPassword sets the Redis password
// WithPassword는 Redis 비밀번호를 설정합니다
func WithPassword(password string) Option {
	return func(c *Config) {
		c.Password = password
	}
}

// WithDB sets the Redis database number
// WithDB는 Redis 데이터베이스 번호를 설정합니다
func WithDB(db int) Option {
	return func(c *Config) {
		c.DB = db
	}
}

// WithPoolSize sets the connection pool size
// WithPoolSize는 연결 풀 크기를 설정합니다
func WithPoolSize(size int) Option {
	return func(c *Config) {
		c.PoolSize = size
	}
}

// WithMinIdleConns sets the minimum number of idle connections
// WithMinIdleConns는 최소 유휴 연결 수를 설정합니다
func WithMinIdleConns(conns int) Option {
	return func(c *Config) {
		c.MinIdleConns = conns
	}
}

// WithMaxRetries sets the maximum number of retries
// WithMaxRetries는 최대 재시도 횟수를 설정합니다
func WithMaxRetries(retries int) Option {
	return func(c *Config) {
		c.MaxRetries = retries
	}
}

// WithRetryInterval sets the retry interval
// WithRetryInterval은 재시도 간격을 설정합니다
func WithRetryInterval(interval time.Duration) Option {
	return func(c *Config) {
		c.RetryInterval = interval
	}
}

// WithHealthCheck enables or disables background health checking
// WithHealthCheck는 백그라운드 헬스 체크를 활성화하거나 비활성화합니다
func WithHealthCheck(enabled bool) Option {
	return func(c *Config) {
		c.EnableHealthCheck = enabled
	}
}

// WithHealthCheckInterval sets the health check interval
// WithHealthCheckInterval은 헬스 체크 간격을 설정합니다
func WithHealthCheckInterval(interval time.Duration) Option {
	return func(c *Config) {
		c.HealthCheckInterval = interval
	}
}

// WithDialTimeout sets the dial timeout
// WithDialTimeout은 다이얼 타임아웃을 설정합니다
func WithDialTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.DialTimeout = timeout
	}
}

// WithReadTimeout sets the read timeout
// WithReadTimeout은 읽기 타임아웃을 설정합니다
func WithReadTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.ReadTimeout = timeout
	}
}

// WithWriteTimeout sets the write timeout
// WithWriteTimeout은 쓰기 타임아웃을 설정합니다
func WithWriteTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.WriteTimeout = timeout
	}
}
