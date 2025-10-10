package mysql

import (
	"crypto/tls"
	"time"

	"github.com/arkd0ng/go-utils/logging"
)

// Option is a function that configures the MySQL client
// Option은 MySQL 클라이언트를 설정하는 함수입니다
type Option func(*config) error

// WithDSN sets the data source name (connection string)
// WithDSN은 데이터 소스 이름(연결 문자열)을 설정합니다
//
// Example / 예제:
//
//	mysql.WithDSN("user:password@tcp(localhost:3306)/dbname?parseTime=true")
func WithDSN(dsn string) Option {
	return func(c *config) error {
		c.dsn = dsn
		return nil
	}
}

// WithCredentialRefresh sets up dynamic credential rotation
// WithCredentialRefresh는 동적 자격 증명 순환을 설정합니다
//
// Parameters / 매개변수:
//   - fn: User function that returns new DSN / 새 DSN을 반환하는 사용자 함수
//   - poolCount: Number of connection pools to maintain / 유지할 연결 풀 개수
//   - interval: How often to rotate one connection / 하나의 연결을 순환할 주기
//
// Example / 예제:
//
//	func getDSN() (string, error) {
//	    // Fetch from Vault, file, env var, etc.
//	    return "user:pass@tcp(localhost:3306)/db", nil
//	}
//
//	mysql.WithCredentialRefresh(getDSN, 3, 1*time.Hour)
//	// Result: 3 pools, rotate 1 per hour
func WithCredentialRefresh(fn CredentialRefreshFunc, poolCount int, interval time.Duration) Option {
	return func(c *config) error {
		c.credRefreshFunc = fn
		c.poolCount = poolCount
		c.rotationInterval = interval
		return nil
	}
}

// WithMaxOpenConns sets the maximum number of open connections to the database
// WithMaxOpenConns는 데이터베이스에 대한 최대 오픈 연결 수를 설정합니다
//
// Example / 예제:
//
//	mysql.WithMaxOpenConns(50)
func WithMaxOpenConns(n int) Option {
	return func(c *config) error {
		c.maxOpenConns = n
		return nil
	}
}

// WithMaxIdleConns sets the maximum number of idle connections
// WithMaxIdleConns는 최대 유휴 연결 수를 설정합니다
//
// Example / 예제:
//
//	mysql.WithMaxIdleConns(10)
func WithMaxIdleConns(n int) Option {
	return func(c *config) error {
		c.maxIdleConns = n
		return nil
	}
}

// WithConnMaxLifetime sets the maximum lifetime of a connection
// WithConnMaxLifetime은 연결의 최대 수명을 설정합니다
//
// Example / 예제:
//
//	mysql.WithConnMaxLifetime(5 * time.Minute)
func WithConnMaxLifetime(d time.Duration) Option {
	return func(c *config) error {
		c.connMaxLifetime = d
		return nil
	}
}

// WithConnMaxIdleTime sets the maximum idle time of a connection
// WithConnMaxIdleTime은 연결의 최대 유휴 시간을 설정합니다
//
// Example / 예제:
//
//	mysql.WithConnMaxIdleTime(2 * time.Minute)
func WithConnMaxIdleTime(d time.Duration) Option {
	return func(c *config) error {
		c.connMaxIdleTime = d
		return nil
	}
}

// WithConnectTimeout sets the connection timeout
// WithConnectTimeout은 연결 타임아웃을 설정합니다
//
// Example / 예제:
//
//	mysql.WithConnectTimeout(10 * time.Second)
func WithConnectTimeout(d time.Duration) Option {
	return func(c *config) error {
		c.connectTimeout = d
		return nil
	}
}

// WithQueryTimeout sets the query execution timeout
// WithQueryTimeout은 쿼리 실행 타임아웃을 설정합니다
//
// Example / 예제:
//
//	mysql.WithQueryTimeout(30 * time.Second)
func WithQueryTimeout(d time.Duration) Option {
	return func(c *config) error {
		c.queryTimeout = d
		return nil
	}
}

// WithMaxRetries sets the maximum number of retry attempts
// WithMaxRetries는 최대 재시도 횟수를 설정합니다
//
// Example / 예제:
//
//	mysql.WithMaxRetries(5)
func WithMaxRetries(n int) Option {
	return func(c *config) error {
		c.maxRetries = n
		return nil
	}
}

// WithRetryDelay sets the delay between retry attempts
// WithRetryDelay는 재시도 간 지연 시간을 설정합니다
//
// Example / 예제:
//
//	mysql.WithRetryDelay(200 * time.Millisecond)
func WithRetryDelay(d time.Duration) Option {
	return func(c *config) error {
		c.retryDelay = d
		return nil
	}
}

// WithLogger sets the logger instance
// WithLogger는 로거 인스턴스를 설정합니다
//
// Example / 예제:
//
//	logger, _ := logging.New(logging.WithFilePath("./logs/mysql.log"))
//	mysql.WithLogger(logger)
func WithLogger(logger *logging.Logger) Option {
	return func(c *config) error {
		c.logger = logger
		return nil
	}
}

// WithQueryLogging enables or disables query logging
// WithQueryLogging은 쿼리 로깅을 활성화하거나 비활성화합니다
//
// Example / 예제:
//
//	mysql.WithQueryLogging(true)
func WithQueryLogging(enable bool) Option {
	return func(c *config) error {
		c.logQueries = enable
		return nil
	}
}

// WithSlowQueryLogging enables or disables slow query logging
// WithSlowQueryLogging은 느린 쿼리 로깅을 활성화하거나 비활성화합니다
//
// Example / 예제:
//
//	mysql.WithSlowQueryLogging(true)
func WithSlowQueryLogging(enable bool) Option {
	return func(c *config) error {
		c.logSlowQueries = enable
		return nil
	}
}

// WithSlowQueryThreshold sets the threshold for slow query logging
// WithSlowQueryThreshold는 느린 쿼리 로깅 임계값을 설정합니다
//
// Example / 예제:
//
//	mysql.WithSlowQueryThreshold(500 * time.Millisecond)
func WithSlowQueryThreshold(d time.Duration) Option {
	return func(c *config) error {
		c.slowQueryThreshold = d
		return nil
	}
}

// WithHealthCheck enables or disables automatic health checks
// WithHealthCheck는 자동 헬스 체크를 활성화하거나 비활성화합니다
//
// Example / 예제:
//
//	mysql.WithHealthCheck(true)
func WithHealthCheck(enable bool) Option {
	return func(c *config) error {
		c.enableHealthCheck = enable
		return nil
	}
}

// WithHealthCheckInterval sets the health check interval
// WithHealthCheckInterval은 헬스 체크 주기를 설정합니다
//
// Example / 예제:
//
//	mysql.WithHealthCheckInterval(60 * time.Second)
func WithHealthCheckInterval(d time.Duration) Option {
	return func(c *config) error {
		c.healthCheckInterval = d
		return nil
	}
}

// WithTLS enables TLS with the provided configuration
// WithTLS는 제공된 설정으로 TLS를 활성화합니다
//
// Example / 예제:
//
//	tlsConfig := &tls.Config{InsecureSkipVerify: false}
//	mysql.WithTLS(tlsConfig)
func WithTLS(tlsConfig *tls.Config) Option {
	return func(c *config) error {
		c.enableTLS = true
		c.tlsConfig = tlsConfig
		return nil
	}
}
