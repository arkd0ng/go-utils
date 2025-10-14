package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Client represents a MySQL database client with auto-management features
// Client는 자동 관리 기능을 갖춘 MySQL 데이터베이스 클라이언트를 나타냅니다
type Client struct {
	// Connection pool / 연결 풀
	connections   []*sql.DB    // Multiple connections for rotation / 순환을 위한 여러 연결
	currentIdx    int          // Current connection index (round-robin) / 현재 연결 인덱스 (round-robin)
	rotationIdx   int          // Rotation index / 순환 인덱스
	connectionsMu sync.RWMutex // Connection array synchronization / 연결 배열 동기화

	// Configuration / 설정
	config *config

	// Query statistics / 쿼리 통계
	statsTracker *queryStatsTracker

	// Background tasks / 백그라운드 작업
	stopChan        chan struct{} // Stop signal / 종료 신호
	healthCheckStop chan struct{} // Health check stop / 헬스 체크 중지
	rotationStop    chan struct{} // Rotation stop / 순환 중지

	// Synchronization / 동기화
	mu     sync.RWMutex // General synchronization / 일반 동기화
	closed bool         // Closed flag / 종료 플래그
}

// New creates a new MySQL client with the given options
// New는 주어진 옵션으로 새 MySQL 클라이언트를 생성합니다
//
// Example (static credentials) / 예제 (정적 자격 증명):
//
//	db, err := mysql.New(
//	    mysql.WithDSN("user:pass@tcp(localhost:3306)/dbname"),
//	)
//
// Example (dynamic credentials) / 예제 (동적 자격 증명):
//
//	db, err := mysql.New(
//	    mysql.WithCredentialRefresh(
//	        func() (string, error) {
//	            return "user:pass@tcp(localhost:3306)/db", nil
//	        },
//	        3,            // 3 connection pools
//	        1*time.Hour,  // Rotate one per hour
//	    ),
//	    mysql.WithLogger(logger),
//	)
func New(opts ...Option) (*Client, error) {
	// Create default config / 기본 설정 생성
	cfg := defaultConfig()

	// Apply options / 옵션 적용
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	// Validate config / 설정 검증
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	// Create client / 클라이언트 생성
	client := &Client{
		config:          cfg,
		connections:     make([]*sql.DB, 0, cfg.poolCount),
		statsTracker:    newQueryStatsTracker(),
		stopChan:        make(chan struct{}),
		healthCheckStop: make(chan struct{}),
		rotationStop:    make(chan struct{}),
	}

	// Enable stats tracking if configured / 설정된 경우 통계 추적 활성화
	if cfg.enableStats {
		client.statsTracker.enabled = true
	}

	// Initialize connections / 연결 초기화
	if err := client.initializeConnections(); err != nil {
		return nil, err
	}

	// Start background tasks / 백그라운드 작업 시작
	if cfg.enableHealthCheck {
		client.startHealthCheck()
	}

	// Start credential rotation if configured / 설정된 경우 자격 증명 순환 시작
	if cfg.credRefreshFunc != nil {
		client.startConnectionRotation()
	}

	return client, nil
}

// initializeConnections creates initial connection pools
// initializeConnections는 초기 연결 풀을 생성합니다
func (c *Client) initializeConnections() error {
	// Get DSN / DSN 가져오기
	dsn := c.config.dsn
	if c.config.credRefreshFunc != nil {
		var err error
		dsn, err = c.config.credRefreshFunc()
		if err != nil {
			return fmt.Errorf("failed to get DSN from credential function: %w", err)
		}
	}

	// Create connection pools / 연결 풀 생성
	for i := 0; i < c.config.poolCount; i++ {
		db, err := c.createConnection(dsn)
		if err != nil {
			// Close any created connections / 생성된 연결 모두 닫기
			for _, conn := range c.connections {
				conn.Close()
			}
			return fmt.Errorf("failed to create connection pool %d: %w", i, err)
		}
		c.connections = append(c.connections, db)
	}

	return nil
}

// createConnection creates a new database connection
// createConnection은 새 데이터베이스 연결을 생성합니다
func (c *Client) createConnection(dsn string) (*sql.DB, error) {
	// Open connection / 연결 열기
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConnectionFailed, err)
	}

	// Configure connection pool / 연결 풀 설정
	db.SetMaxOpenConns(c.config.maxOpenConns)
	db.SetMaxIdleConns(c.config.maxIdleConns)
	db.SetConnMaxLifetime(c.config.connMaxLifetime)
	db.SetConnMaxIdleTime(c.config.connMaxIdleTime)

	// Test connection / 연결 테스트
	ctx, cancel := context.WithTimeout(context.Background(), c.config.connectTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("%w: ping failed: %v", ErrConnectionFailed, err)
	}

	return db, nil
}

// Close closes all database connections and stops background tasks
// Close는 모든 데이터베이스 연결을 닫고 백그라운드 작업을 중지합니다
func (c *Client) Close() error {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return nil
	}
	c.closed = true
	c.mu.Unlock()

	// Stop background tasks / 백그라운드 작업 중지
	close(c.stopChan)
	close(c.healthCheckStop)
	close(c.rotationStop)

	// Close all connections / 모든 연결 닫기
	c.connectionsMu.Lock()
	defer c.connectionsMu.Unlock()

	var lastErr error
	for i, db := range c.connections {
		if err := db.Close(); err != nil {
			lastErr = err
			if c.config.logger != nil {
				c.config.logger.Error("Failed to close connection",
					"index", i,
					"error", err)
			}
		}
	}

	return lastErr
}

// Ping verifies a connection to the database is still alive
// Ping은 데이터베이스 연결이 여전히 활성 상태인지 확인합니다
func (c *Client) Ping(ctx context.Context) error {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return ErrClosed
	}
	c.mu.RUnlock()

	db := c.getCurrentConnection()
	return db.PingContext(ctx)
}

// Stats returns database statistics
// Stats는 데이터베이스 통계를 반환합니다
func (c *Client) Stats() []sql.DBStats {
	c.connectionsMu.RLock()
	defer c.connectionsMu.RUnlock()

	stats := make([]sql.DBStats, len(c.connections))
	for i, db := range c.connections {
		stats[i] = db.Stats()
	}
	return stats
}

// getCurrentConnection returns the current connection using round-robin
// getCurrentConnection은 round-robin을 사용하여 현재 연결을 반환합니다
func (c *Client) getCurrentConnection() *sql.DB {
	c.connectionsMu.RLock()
	defer c.connectionsMu.RUnlock()

	// Round-robin selection / Round-robin 선택
	conn := c.connections[c.currentIdx]
	c.currentIdx = (c.currentIdx + 1) % len(c.connections)

	return conn
}

// Query executes a query that returns rows
// Query는 행을 반환하는 쿼리를 실행합니다
func (c *Client) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	start := time.Now()
	db := c.getCurrentConnection()

	rows, err := db.QueryContext(ctx, query, args...)
	duration := time.Since(start)

	// Record stats / 통계 기록
	c.statsTracker.recordQuery(query, args, duration, err)

	// Log query / 쿼리 로깅
	c.logQuery(query, args, duration, err)

	if err != nil {
		return nil, c.wrapError("Query", query, args, err, duration)
	}

	return rows, nil
}

// QueryRow executes a query that returns a single row
// QueryRow는 단일 행을 반환하는 쿼리를 실행합니다
func (c *Client) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil
	}
	c.mu.RUnlock()

	db := c.getCurrentConnection()
	return db.QueryRowContext(ctx, query, args...)
}

// Exec executes a query that doesn't return rows
// Exec는 행을 반환하지 않는 쿼리를 실행합니다
func (c *Client) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	start := time.Now()
	db := c.getCurrentConnection()

	result, err := db.ExecContext(ctx, query, args...)
	duration := time.Since(start)

	// Record stats / 통계 기록
	c.statsTracker.recordQuery(query, args, duration, err)

	// Log query / 쿼리 로깅
	c.logQuery(query, args, duration, err)

	if err != nil {
		return nil, c.wrapError("Exec", query, args, err, duration)
	}

	return result, nil
}

// Begin starts a new transaction
// Begin은 새 트랜잭션을 시작합니다
func (c *Client) Begin(ctx context.Context) (*Tx, error) {
	c.mu.RLock()
	if c.closed {
		c.mu.RUnlock()
		return nil, ErrClosed
	}
	c.mu.RUnlock()

	db := c.getCurrentConnection()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTransactionFailed, err)
	}

	return &Tx{
		tx:     tx,
		client: c,
	}, nil
}

// wrapError wraps an error with context information
// wrapError는 컨텍스트 정보로 에러를 래핑합니다
func (c *Client) wrapError(op, query string, args []interface{}, err error, duration time.Duration) error {
	return &DBError{
		Op:       op,
		Query:    query,
		Args:     args,
		Err:      err,
		Time:     time.Now(),
		Duration: duration,
	}
}

// logQuery logs a query if logging is enabled
// logQuery는 로깅이 활성화된 경우 쿼리를 로깅합니다
func (c *Client) logQuery(query string, args []interface{}, duration time.Duration, err error) {
	if c.config.logger == nil {
		return
	}

	// Log all queries if enabled / 활성화된 경우 모든 쿼리 로깅
	if c.config.logQueries {
		if err != nil {
			c.config.logger.Error("Query failed",
				"query", query,
				"args", args,
				"duration", duration,
				"error", err)
		} else {
			c.config.logger.Debug("Query executed",
				"query", query,
				"args", args,
				"duration", duration)
		}
	}

	// Log slow queries / 느린 쿼리 로깅
	if c.config.logSlowQueries && duration >= c.config.slowQueryThreshold {
		c.config.logger.Warn("Slow query detected",
			"query", query,
			"args", args,
			"duration", duration,
			"threshold", c.config.slowQueryThreshold)
	}
}
