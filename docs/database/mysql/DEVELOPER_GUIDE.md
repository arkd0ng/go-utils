# MySQL Package - Developer Guide / 개발자 가이드

**Package**: `github.com/arkd0ng/go-utils/database/mysql`
**Version**: v1.3.006
**Author**: arkd0ng
**Last Updated**: 2025-10-10

---

## Table of Contents / 목차

1. [Architecture Overview / 아키텍처 개요](#architecture-overview--아키텍처-개요)
2. [Package Structure / 패키지 구조](#package-structure--패키지-구조)
3. [Core Components / 핵심 컴포넌트](#core-components--핵심-컴포넌트)
4. [Internal Implementation / 내부 구현](#internal-implementation--내부-구현)
5. [Design Patterns / 디자인 패턴](#design-patterns--디자인-패턴)
6. [Adding New Features / 새 기능 추가](#adding-new-features--새-기능-추가)
7. [Testing Guide / 테스트 가이드](#testing-guide--테스트-가이드)
8. [Performance / 성능](#performance--성능)
9. [Contributing Guidelines / 기여 가이드라인](#contributing-guidelines--기여-가이드라인)
10. [Code Style / 코드 스타일](#code-style--코드-스타일)

---

## Architecture Overview / 아키텍처 개요

### Design Philosophy / 설계 철학

The MySQL package follows the principle: **"If it's not 10x simpler, don't build it"**

MySQL 패키지는 다음 원칙을 따릅니다: **"10배 간단하지 않으면 만들지 마세요"**

**Core Principles / 핵심 원칙**:
1. **Zero Mental Overhead**: Developers don't worry about connections, retries, or resource management
2. **SQL-Like API**: API syntax stays close to actual SQL
3. **Auto Everything**: All tedious tasks handled automatically
4. **Three-Layer API**: Simple → Query Builder → Raw SQL for different complexity levels

1. **제로 정신적 부담**: 개발자는 연결, 재시도, 리소스 관리에 대해 걱정하지 않음
2. **SQL 같은 API**: API 문법이 실제 SQL과 가깝게 유지됨
3. **모든 것이 자동**: 모든 번거로운 작업이 자동으로 처리됨
4. **3계층 API**: 간단 → 쿼리 빌더 → Raw SQL로 다양한 복잡도 수준 제공

### High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────────┐
│                    User Application                      │
└───────────────────┬─────────────────────────────────────┘
                    │
        ┌───────────┴───────────┬───────────────┐
        │                       │               │
┌───────▼────────┐   ┌─────────▼────────┐   ┌─▼──────────┐
│   Simple API   │   │  Query Builder   │   │  Raw SQL   │
│  (SelectAll,   │   │ (Fluent methods) │   │  (Query,   │
│   Insert...)   │   │                  │   │   Exec)    │
└───────┬────────┘   └─────────┬────────┘   └─┬──────────┘
        │                       │               │
        └───────────┬───────────┴───────────────┘
                    │
        ┌───────────▼────────────┐
        │   Client (core logic)  │
        │   - Connection Pool    │
        │   - Health Check       │
        │   - Credential Rotation│
        └───────────┬────────────┘
                    │
        ┌───────────▼────────────┐
        │   Auto Features Layer  │
        │   - Retry Logic        │
        │   - Error Classification│
        │   - Resource Cleanup   │
        └───────────┬────────────┘
                    │
        ┌───────────▼────────────┐
        │    database/sql        │
        │   (Standard Library)   │
        └───────────┬────────────┘
                    │
        ┌───────────▼────────────┐
        │    MySQL Server        │
        └────────────────────────┘
```

### Component Interaction / 컴포넌트 상호작용

```
User Code
    │
    ▼
SelectAll("users", "age > ?", 18)
    │
    ▼
[Check if closed] ──No──▶ Build SQL query
    │                        │
   Yes                       ▼
    │                   SELECT * FROM users WHERE age > ?
    │                        │
    ▼                        ▼
[Return error]          [executeWithRetry]
                             │
                     ┌───────┴───────┐
                     │               │
                Success          Transient Error?
                     │               │
                     │              Yes ──▶ [Exponential Backoff]
                     │               │              │
                     │               └──────────────┘
                     │                      │
                     │                     Retry (up to max)
                     │                      │
                     ▼                      ▼
                [scanRows]            [Success or Fatal Error]
                     │
                     ▼
              []map[string]interface{}
                     │
                     ▼
                User Code
```

---

## Package Structure / 패키지 구조

### File Organization / 파일 구성

```
database/mysql/
├── client.go           # Main client, connection management
├── config.go           # Configuration structures
├── options.go          # Functional options
├── connection.go       # Health check, connection monitoring
├── rotation.go         # Credential rotation (optional)
├── errors.go           # Error types and classification
├── types.go            # Common types (Tx, CredentialRefreshFunc)
├── retry.go            # Auto retry with exponential backoff
├── scan.go             # Result scanning and type conversion
├── simple.go           # Simple API (SelectAll, Insert, etc.)
├── select_options.go   # SelectWhere with functional options
├── builder.go          # Query Builder with fluent API
├── transaction.go      # Transaction support
└── client_test.go      # Unit tests
```

### File Responsibilities / 파일별 책임

| File / 파일 | Lines / 줄 수 | Responsibility / 책임 |
|-------------|--------------|----------------------|
| `client.go` | ~260 | Client struct, New(), Close(), getCurrentConnection() |
| `config.go` | ~130 | Config struct, validation, default values |
| `options.go` | ~230 | 20+ functional options (WithDSN, WithMaxOpenConns, etc.) |
| `connection.go` | ~75 | Health check goroutine, connection monitoring |
| `rotation.go` | ~85 | Credential rotation goroutine (optional feature) |
| `errors.go` | ~130 | Error types, classification (transient vs fatal) |
| `types.go` | ~73 | Common types (Tx, CredentialRefreshFunc) |
| `retry.go` | ~120 | executeWithRetry(), exponential backoff |
| `scan.go` | ~180 | scanRows(), scanRow(), scanCount(), type conversion |
| `simple.go` | ~370 | SelectAll, SelectOne, Insert, Update, Delete, Count, Exists |
| `select_options.go` | ~360 | SelectWhere, SelectOneWhere, functional options |
| `builder.go` | ~285 | Query Builder, fluent API |
| `transaction.go` | ~230 | Transaction(), Tx methods |
| `client_test.go` | ~120 | Unit tests |

**Total**: ~2,648 lines / 총 ~2,648줄

---

## Core Components / 핵심 컴포넌트

### 1. Client Struct / Client 구조체

The main entry point for all database operations.

모든 데이터베이스 작업의 주요 진입점입니다.

```go
type Client struct {
    config         *Config              // Configuration / 설정
    pools          []*sql.DB            // Connection pools / 연결 풀
    currentPool    atomic.Uint32        // Current pool index / 현재 풀 인덱스
    healthTicker   *time.Ticker         // Health check ticker / 헬스 체크 티커
    rotationTicker *time.Ticker         // Rotation ticker / 순환 티커
    stopChan       chan struct{}        // Stop signal / 중지 신호
    wg             sync.WaitGroup       // Wait group for goroutines / 고루틴용 대기 그룹
    mu             sync.RWMutex         // Protects closed / closed 보호
    closed         bool                 // Client closed flag / 클라이언트 닫힘 플래그
}
```

**Key Methods / 주요 메서드**:
- `New(...Option) (*Client, error)` - Create new client / 새 클라이언트 생성
- `Close() error` - Close all connections / 모든 연결 닫기
- `getCurrentConnection() *sql.DB` - Get current connection pool / 현재 연결 풀 가져오기

### 2. Config Struct / Config 구조체

Holds all configuration settings.

모든 설정을 보유합니다.

```go
type Config struct {
    DSN                 string                 // MySQL DSN
    MaxOpenConns        int                    // Max open connections / 최대 열린 연결 수
    MaxIdleConns        int                    // Max idle connections / 최대 유휴 연결 수
    ConnMaxLifetime     time.Duration          // Connection lifetime / 연결 수명
    ConnMaxIdleTime     time.Duration          // Idle time / 유휴 시간
    RetryMaxAttempts    int                    // Max retry attempts / 최대 재시도 횟수
    RetryInitialInterval time.Duration         // Initial retry interval / 초기 재시도 간격
    RetryMaxInterval    time.Duration          // Max retry interval / 최대 재시도 간격
    RetryMultiplier     float64                // Backoff multiplier / 백오프 배수
    CredentialRefresh   CredentialRefreshFunc  // Credential refresh function / 자격 증명 갱신 함수
    RotationPoolCount   int                    // Number of rotation pools / 순환 풀 수
    RotationInterval    time.Duration          // Rotation interval / 순환 간격
    HealthCheckInterval time.Duration          // Health check interval / 헬스 체크 간격
}
```

**Validation / 검증**:
```go
func (c *Config) Validate() error {
    if c.DSN == "" {
        return ErrInvalidConfig
    }
    // More validation...
    return nil
}
```

### 3. Transaction (Tx) Struct / Transaction (Tx) 구조체

Wrapper for `*sql.Tx` with all Client methods available.

모든 Client 메서드를 사용할 수 있는 `*sql.Tx` 래퍼입니다.

```go
type Tx struct {
    tx     *sql.Tx    // Underlying transaction / 기본 트랜잭션
    client *Client    // Parent client / 부모 클라이언트
}
```

**Key Methods / 주요 메서드**:
- Same as Client: `SelectAll`, `Insert`, `Update`, etc.
- Transaction-specific: `Commit`, `Rollback`

### 4. Query Builder Struct / Query Builder 구조체

Fluent API for building complex queries.

복잡한 쿼리를 빌드하기 위한 Fluent API입니다.

```go
type QueryBuilder struct {
    client       *Client        // Client reference / 클라이언트 참조
    tx           *Tx            // Transaction reference (optional) / 트랜잭션 참조 (선택)
    columns      []string       // SELECT columns / SELECT 컬럼
    table        string         // FROM table / FROM 테이블
    joins        []joinClause   // JOIN clauses / JOIN 절
    whereClauses []whereClause  // WHERE conditions / WHERE 조건
    groupByCols  []string       // GROUP BY columns / GROUP BY 컬럼
    havingClauses []whereClause // HAVING conditions / HAVING 조건
    orderBy      string         // ORDER BY clause / ORDER BY 절
    limitNum     *int           // LIMIT value / LIMIT 값
    offsetNum    *int           // OFFSET value / OFFSET 값
    args         []interface{}  // Query arguments / 쿼리 인자
}
```

### 5. Error Types / 에러 타입

Structured error types for better error handling.

더 나은 에러 처리를 위한 구조화된 에러 타입입니다.

```go
var (
    ErrInvalidConfig    = errors.New("invalid configuration")
    ErrClosed           = errors.New("client is closed")
    ErrNoConnection     = errors.New("no database connection available")
    ErrInvalidDSN       = errors.New("invalid DSN format")
    ErrTransactionFailed = errors.New("transaction failed")
)

// QueryError wraps query errors with context
type QueryError struct {
    Operation string        // Operation name / 작업 이름
    Query     string        // SQL query / SQL 쿼리
    Args      []interface{} // Query arguments / 쿼리 인자
    Err       error         // Original error / 원본 에러
    Duration  time.Duration // Query duration / 쿼리 소요 시간
}
```

---

## Internal Implementation / 내부 구현

### Connection Management Flow / 연결 관리 흐름

```go
// New() creates a client
New(options...)
    │
    ▼
[Apply default config]
    │
    ▼
[Apply user options]
    │
    ▼
[Validate config]
    │
    ▼
[Create connection pool(s)]
    │
    ├──▶ Single pool mode
    │         │
    │         ▼
    │    sql.Open(dsn)
    │         │
    │         ▼
    │    Configure pool settings
    │         │
    │         ▼
    │    db.Ping() to verify
    │
    └──▶ Rotation mode (optional)
              │
              ▼
         Create N pools
              │
              ▼
         Start rotation goroutine
              │
              ▼
         Start health check goroutine
              │
              ▼
         Return client
```

### Query Execution Flow / 쿼리 실행 흐름

```go
SelectAll(ctx, table, condition, args...)
    │
    ▼
[Check if client closed]
    │
    ▼
[Build SQL query]
    query := "SELECT * FROM " + table
    if condition != "" {
        query += " WHERE " + condition
    }
    │
    ▼
[Execute with retry]
    executeWithRetry(ctx, func() error {
        db := client.getCurrentConnection()
        rows, err := db.QueryContext(ctx, query, args...)
        return err
    })
    │
    ├──▶ Success
    │        │
    │        ▼
    │   [Scan rows]
    │        │
    │        ▼
    │   [Convert types]
    │        │
    │        ▼
    │   Return []map[string]interface{}
    │
    └──▶ Error
             │
             ▼
        [Classify error]
             │
      ┌──────┴──────┐
      │             │
  Transient     Fatal
      │             │
      ▼             ▼
  [Retry]    [Return error]
      │
      └──▶ [Exponential backoff]
              │
          ┌───┴────┐
          │        │
     Success   Max attempts
          │        │
          ▼        ▼
    Return    Return error
```

### Retry Logic / 재시도 로직

```go
func (c *Client) executeWithRetry(ctx context.Context, fn func() error) error {
    var lastErr error
    interval := c.config.RetryInitialInterval

    for attempt := 0; attempt < c.config.RetryMaxAttempts; attempt++ {
        // Execute function / 함수 실행
        err := fn()
        if err == nil {
            return nil // Success / 성공
        }

        lastErr = err

        // Check if error is transient / 일시적 에러인지 확인
        if !isTransientError(err) {
            return err // Fatal error, don't retry / 치명적 에러, 재시도 안 함
        }

        // Check context / context 확인
        if ctx.Err() != nil {
            return ctx.Err()
        }

        // Wait before retry / 재시도 전 대기
        if attempt < c.config.RetryMaxAttempts-1 {
            select {
            case <-time.After(interval):
            case <-ctx.Done():
                return ctx.Err()
            }

            // Exponential backoff / 지수 백오프
            interval = time.Duration(float64(interval) * c.config.RetryMultiplier)
            if interval > c.config.RetryMaxInterval {
                interval = c.config.RetryMaxInterval
            }
        }
    }

    return lastErr
}
```

### Transient Error Classification / 일시적 에러 분류

```go
func isTransientError(err error) bool {
    if err == nil {
        return false
    }

    errMsg := err.Error()

    // MySQL error codes / MySQL 에러 코드
    transientErrors := []string{
        "Error 1040",  // Too many connections / 연결이 너무 많음
        "Error 1205",  // Lock wait timeout / 잠금 대기 타임아웃
        "Error 1213",  // Deadlock / 교착상태
        "Error 2002",  // Can't connect / 연결할 수 없음
        "Error 2003",  // Can't connect / 연결할 수 없음
        "Error 2006",  // MySQL server has gone away / MySQL 서버가 사라짐
        "Error 2013",  // Lost connection / 연결 손실
        "bad connection",
        "connection refused",
        "connection reset",
        "broken pipe",
    }

    for _, transient := range transientErrors {
        if strings.Contains(errMsg, transient) {
            return true
        }
    }

    return false
}
```

### Row Scanning and Type Conversion / 행 스캔 및 타입 변환

```go
func scanRows(rows *sql.Rows) ([]map[string]interface{}, error) {
    // Get column names / 컬럼 이름 가져오기
    columns, err := rows.Columns()
    if err != nil {
        return nil, err
    }

    results := []map[string]interface{}{}

    for rows.Next() {
        // Create scan destinations / 스캔 대상 생성
        values := make([]interface{}, len(columns))
        valuePtrs := make([]interface{}, len(columns))
        for i := range columns {
            valuePtrs[i] = &values[i]
        }

        // Scan row / 행 스캔
        if err := rows.Scan(valuePtrs...); err != nil {
            return nil, err
        }

        // Convert to map / 맵으로 변환
        row := make(map[string]interface{})
        for i, col := range columns {
            val := values[i]

            // Type conversion / 타입 변환
            if b, ok := val.([]byte); ok {
                row[col] = string(b)  // []byte → string
            } else {
                row[col] = val
            }
        }

        results = append(results, row)
    }

    // Check for iteration errors / 반복 에러 확인
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return results, nil
}
```

### Health Check Goroutine / 헬스 체크 고루틴

```go
func (c *Client) startHealthCheck() {
    c.wg.Add(1)
    go func() {
        defer c.wg.Done()

        for {
            select {
            case <-c.healthTicker.C:
                // Check all pools / 모든 풀 확인
                for i, pool := range c.pools {
                    if err := pool.Ping(); err != nil {
                        // Log error / 에러 로그
                        // Try to reconnect / 재연결 시도
                        newPool, err := sql.Open("mysql", c.config.DSN)
                        if err == nil {
                            // Replace pool / 풀 교체
                            c.pools[i].Close()
                            c.pools[i] = newPool
                        }
                    }
                }

            case <-c.stopChan:
                return
            }
        }
    }()
}
```

### Credential Rotation (Optional Feature) / 자격 증명 순환 (선택 기능)

```go
func (c *Client) startRotation() {
    c.wg.Add(1)
    go func() {
        defer c.wg.Done()

        rotationIndex := 0

        for {
            select {
            case <-c.rotationTicker.C:
                // Get new credentials / 새 자격 증명 가져오기
                newDSN, err := c.config.CredentialRefresh()
                if err != nil {
                    continue
                }

                // Create new pool with new credentials
                // 새 자격 증명으로 새 풀 생성
                newPool, err := sql.Open("mysql", newDSN)
                if err != nil {
                    continue
                }

                // Configure pool / 풀 설정
                newPool.SetMaxOpenConns(c.config.MaxOpenConns)
                newPool.SetMaxIdleConns(c.config.MaxIdleConns)

                // Test connection / 연결 테스트
                if err := newPool.Ping(); err != nil {
                    newPool.Close()
                    continue
                }

                // Replace old pool / 이전 풀 교체
                oldPool := c.pools[rotationIndex]
                c.pools[rotationIndex] = newPool

                // Close old pool after grace period / 유예 기간 후 이전 풀 닫기
                time.AfterFunc(30*time.Second, func() {
                    oldPool.Close()
                })

                // Move to next pool / 다음 풀로 이동
                rotationIndex = (rotationIndex + 1) % len(c.pools)

            case <-c.stopChan:
                return
            }
        }
    }()
}
```

---

## Design Patterns / 디자인 패턴

### 1. Functional Options Pattern / 함수형 옵션 패턴

**Purpose / 목적**: Provide flexible, extensible configuration without breaking API compatibility.

유연하고 확장 가능한 설정을 API 호환성을 깨지 않고 제공합니다.

**Implementation / 구현**:
```go
// Option is a function that modifies Config
// Option은 Config를 수정하는 함수입니다
type Option func(*Config)

// WithDSN sets the MySQL DSN
// WithDSN은 MySQL DSN을 설정합니다
func WithDSN(dsn string) Option {
    return func(c *Config) {
        c.DSN = dsn
    }
}

// WithMaxOpenConns sets max open connections
// WithMaxOpenConns는 최대 열린 연결 수를 설정합니다
func WithMaxOpenConns(n int) Option {
    return func(c *Config) {
        c.MaxOpenConns = n
    }
}

// New creates a client with options
// New는 옵션으로 클라이언트를 생성합니다
func New(opts ...Option) (*Client, error) {
    config := defaultConfig()
    for _, opt := range opts {
        opt(config)
    }
    // ...
}
```

**Benefits / 장점**:
- Backward compatible / 하위 호환성
- Optional parameters / 선택적 매개변수
- Self-documenting / 자체 문서화
- Extensible / 확장 가능

### 2. Builder Pattern (Query Builder) / 빌더 패턴 (쿼리 빌더)

**Purpose / 목적**: Construct complex queries step-by-step with method chaining.

메서드 체이닝으로 복잡한 쿼리를 단계별로 구성합니다.

**Implementation / 구현**:
```go
type QueryBuilder struct {
    // ... fields
}

func (qb *QueryBuilder) From(table string) *QueryBuilder {
    qb.table = table
    return qb  // Return self for chaining / 체이닝을 위해 self 반환
}

func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
    qb.whereClauses = append(qb.whereClauses, whereClause{condition, args})
    return qb
}

func (qb *QueryBuilder) All(ctx context.Context) ([]map[string]interface{}, error) {
    query, args := qb.buildQuery()
    return qb.executeQuery(ctx, query, args)
}
```

**Benefits / 장점**:
- Fluent interface / 유창한 인터페이스
- IDE autocomplete support / IDE 자동완성 지원
- Type-safe / 타입 안전
- Readable / 가독성 높음

### 3. Singleton Pattern (Global Client) / 싱글톤 패턴 (전역 클라이언트)

**Purpose / 목적**: Optional global client for simple use cases.

간단한 사용 사례를 위한 선택적 전역 클라이언트입니다.

**Implementation / 구현**:
```go
var (
    globalClient *Client
    globalOnce   sync.Once
    globalMu     sync.RWMutex
)

// SetGlobal sets the global client
// SetGlobal은 전역 클라이언트를 설정합니다
func SetGlobal(client *Client) {
    globalMu.Lock()
    defer globalMu.Unlock()
    globalClient = client
}

// Global returns the global client
// Global은 전역 클라이언트를 반환합니다
func Global() *Client {
    globalMu.RLock()
    defer globalMu.RUnlock()
    return globalClient
}
```

### 4. Strategy Pattern (Retry Strategy) / 전략 패턴 (재시도 전략)

**Purpose / 목적**: Different retry strategies for different error types.

다른 에러 타입에 대한 다른 재시도 전략입니다.

**Implementation / 구현**:
```go
type RetryStrategy interface {
    ShouldRetry(err error) bool
    NextInterval(attempt int) time.Duration
}

type ExponentialBackoff struct {
    InitialInterval time.Duration
    MaxInterval     time.Duration
    Multiplier      float64
}

func (e *ExponentialBackoff) NextInterval(attempt int) time.Duration {
    interval := e.InitialInterval
    for i := 0; i < attempt; i++ {
        interval = time.Duration(float64(interval) * e.Multiplier)
        if interval > e.MaxInterval {
            return e.MaxInterval
        }
    }
    return interval
}
```

### 5. Decorator Pattern (Transaction Wrapper) / 데코레이터 패턴 (트랜잭션 래퍼)

**Purpose / 목적**: Wrap `*sql.Tx` to add auto commit/rollback functionality.

자동 커밋/롤백 기능을 추가하기 위해 `*sql.Tx`를 래핑합니다.

**Implementation / 구현**:
```go
type Tx struct {
    tx     *sql.Tx
    client *Client
}

// Transaction wraps user function with auto commit/rollback
// Transaction은 사용자 함수를 자동 커밋/롤백으로 래핑합니다
func (c *Client) Transaction(ctx context.Context, fn func(tx *Tx) error) error {
    // Begin transaction / 트랜잭션 시작
    sqlTx, err := c.getCurrentConnection().BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    tx := &Tx{tx: sqlTx, client: c}

    // Defer rollback on panic / 패닉 시 롤백 지연
    defer func() {
        if p := recover(); p != nil {
            sqlTx.Rollback()
            panic(p)
        }
    }()

    // Execute user function / 사용자 함수 실행
    if err := fn(tx); err != nil {
        sqlTx.Rollback()  // Rollback on error / 에러 시 롤백
        return err
    }

    // Commit on success / 성공 시 커밋
    return sqlTx.Commit()
}
```

### 6. Template Method Pattern (executeWithRetry) / 템플릿 메서드 패턴

**Purpose / 목적**: Define retry algorithm skeleton, let subclasses define specific operations.

재시도 알고리즘 골격을 정의하고 서브클래스가 특정 작업을 정의하도록 합니다.

**Implementation / 구현**:
```go
// Template method / 템플릿 메서드
func (c *Client) executeWithRetry(ctx context.Context, fn func() error) error {
    // ... retry logic (same for all operations)
    // ... 재시도 로직 (모든 작업에 동일)
}

// Specific implementations / 특정 구현
func (c *Client) SelectAll(ctx context.Context, table string, ...) {
    c.executeWithRetry(ctx, func() error {
        // Specific query execution / 특정 쿼리 실행
    })
}

func (c *Client) Insert(ctx context.Context, table string, ...) {
    c.executeWithRetry(ctx, func() error {
        // Specific insert execution / 특정 삽입 실행
    })
}
```

---

## Adding New Features / 새 기능 추가

### Adding a New Simple API Method / 새로운 Simple API 메서드 추가

**Step 1**: Define method in `simple.go` / `simple.go`에 메서드 정의

```go
// SelectDistinct selects distinct rows
// SelectDistinct는 중복 없는 행을 선택합니다
func (c *Client) SelectDistinct(ctx context.Context, table string, columns []string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
    c.mu.RLock()
    if c.closed {
        c.mu.RUnlock()
        return nil, ErrClosed
    }
    c.mu.RUnlock()

    // Build query / 쿼리 빌드
    query := fmt.Sprintf("SELECT DISTINCT %s FROM %s", strings.Join(columns, ", "), table)

    // Add WHERE clause if provided / WHERE 절 추가 (있는 경우)
    var args []interface{}
    if len(conditionAndArgs) > 0 {
        condition := fmt.Sprintf("%v", conditionAndArgs[0])
        query += " WHERE " + condition
        if len(conditionAndArgs) > 1 {
            args = conditionAndArgs[1:]
        }
    }

    start := time.Now()

    // Execute with retry / 재시도로 실행
    var rows *sql.Rows
    err := c.executeWithRetry(ctx, func() error {
        db := c.getCurrentConnection()
        var execErr error
        rows, execErr = db.QueryContext(ctx, query, args...)
        return execErr
    })

    duration := time.Since(start)

    if err != nil {
        c.logQuery(query, args, duration, err)
        return nil, c.wrapError("SelectDistinct", query, args, err, duration)
    }

    // Scan rows / 행 스캔
    results, err := scanRows(rows)
    if err != nil {
        c.logQuery(query, args, duration, err)
        return nil, c.wrapError("SelectDistinct", query, args, err, duration)
    }

    c.logQuery(query, args, duration, nil)
    return results, nil
}
```

**Step 2**: Add transaction support in `transaction.go` / `transaction.go`에 트랜잭션 지원 추가

```go
// SelectDistinct for transactions / 트랜잭션용 SelectDistinct
func (tx *Tx) SelectDistinct(ctx context.Context, table string, columns []string, conditionAndArgs ...interface{}) ([]map[string]interface{}, error) {
    // Similar implementation using tx.tx instead of db
    // tx.tx를 사용하여 유사한 구현
}
```

**Step 3**: Add tests in `client_test.go` / `client_test.go`에 테스트 추가

```go
func TestSelectDistinct(t *testing.T) {
    // Setup test database / 테스트 데이터베이스 설정
    db, _ := mysql.New(mysql.WithDSN(testDSN))
    defer db.Close()

    // Test cases / 테스트 케이스
    results, err := db.SelectDistinct(ctx, "users", []string{"city"})
    if err != nil {
        t.Fatalf("SelectDistinct failed: %v", err)
    }

    // Assertions / 어설션
    if len(results) == 0 {
        t.Error("Expected results")
    }
}
```

**Step 4**: Add documentation / 문서 추가

Update `USER_MANUAL.md` with examples and API reference.

예제와 API 참조로 `USER_MANUAL.md`를 업데이트합니다.

### Adding a New Query Builder Method / 새로운 쿼리 빌더 메서드 추가

**Step 1**: Add method to `QueryBuilder` in `builder.go`

```go
// Union adds UNION clause
// Union은 UNION 절을 추가합니다
func (qb *QueryBuilder) Union(otherQuery *QueryBuilder) *QueryBuilder {
    qb.unionQuery = otherQuery
    return qb
}
```

**Step 2**: Update `buildQuery()` to handle new clause

```go
func (qb *QueryBuilder) buildQuery() (string, []interface{}) {
    // ... existing code ...

    // UNION
    if qb.unionQuery != nil {
        unionSQL, unionArgs := qb.unionQuery.buildQuery()
        parts = append(parts, "UNION", unionSQL)
        args = append(args, unionArgs...)
    }

    return strings.Join(parts, " "), args
}
```

### Adding a New Option / 새로운 옵션 추가

**Step 1**: Add field to `Config` in `config.go`

```go
type Config struct {
    // ... existing fields ...
    CustomTimeout time.Duration  // New field / 새 필드
}
```

**Step 2**: Add option function in `options.go`

```go
// WithCustomTimeout sets custom timeout
// WithCustomTimeout은 사용자 정의 타임아웃을 설정합니다
func WithCustomTimeout(timeout time.Duration) Option {
    return func(c *Config) {
        c.CustomTimeout = timeout
    }
}
```

**Step 3**: Update default config

```go
func defaultConfig() *Config {
    return &Config{
        // ... existing defaults ...
        CustomTimeout: 30 * time.Second,  // Default value / 기본값
    }
}
```

**Step 4**: Use new config in implementation

```go
func (c *Client) someMethod() {
    timeout := c.config.CustomTimeout
    // Use timeout...
}
```

---

## Testing Guide / 테스트 가이드

### Test Structure / 테스트 구조

```
database/mysql/
├── client_test.go      # Main tests / 주요 테스트
├── testdata/           # Test fixtures / 테스트 픽스처
│   ├── schema.sql     # Test database schema / 테스트 데이터베이스 스키마
│   └── data.sql       # Test data / 테스트 데이터
```

### Running Tests / 테스트 실행

```bash
# Run all tests / 모든 테스트 실행
go test ./database/mysql -v

# Run specific test / 특정 테스트 실행
go test ./database/mysql -v -run TestSelectAll

# Run with coverage / 커버리지와 함께 실행
go test ./database/mysql -cover

# Generate coverage report / 커버리지 보고서 생성
go test ./database/mysql -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Test Setup / 테스트 설정

```go
// Setup test database / 테스트 데이터베이스 설정
func setupTestDB(t *testing.T) *mysql.Client {
    // Use test database / 테스트 데이터베이스 사용
    dsn := "root:password@tcp(localhost:3306)/testdb"

    db, err := mysql.New(mysql.WithDSN(dsn))
    if err != nil {
        t.Fatalf("Failed to create client: %v", err)
    }

    // Create test table / 테스트 테이블 생성
    db.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS test_users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(255),
            email VARCHAR(255),
            age INT
        )
    `)

    // Clean up previous test data / 이전 테스트 데이터 정리
    db.Exec(context.Background(), "TRUNCATE TABLE test_users")

    // Insert test data / 테스트 데이터 삽입
    db.Insert(context.Background(), "test_users", map[string]interface{}{
        "name":  "John Doe",
        "email": "john@example.com",
        "age":   30,
    })

    return db
}

// Teardown test database / 테스트 데이터베이스 정리
func teardownTestDB(t *testing.T, db *mysql.Client) {
    db.Exec(context.Background(), "DROP TABLE IF EXISTS test_users")
    db.Close()
}
```

### Writing Test Cases / 테스트 케이스 작성

```go
func TestSelectAll(t *testing.T) {
    db := setupTestDB(t)
    defer teardownTestDB(t, db)

    ctx := context.Background()

    // Test 1: Select all rows / 모든 행 선택
    t.Run("SelectAll", func(t *testing.T) {
        users, err := db.SelectAll(ctx, "test_users")
        if err != nil {
            t.Fatalf("SelectAll failed: %v", err)
        }

        if len(users) != 1 {
            t.Errorf("Expected 1 user, got %d", len(users))
        }

        if users[0]["name"] != "John Doe" {
            t.Errorf("Expected name 'John Doe', got %v", users[0]["name"])
        }
    })

    // Test 2: Select with condition / 조건으로 선택
    t.Run("SelectAllWithCondition", func(t *testing.T) {
        users, err := db.SelectAll(ctx, "test_users", "age > ?", 25)
        if err != nil {
            t.Fatalf("SelectAll failed: %v", err)
        }

        if len(users) != 1 {
            t.Errorf("Expected 1 user, got %d", len(users))
        }
    })

    // Test 3: Empty result / 빈 결과
    t.Run("SelectAllEmpty", func(t *testing.T) {
        users, err := db.SelectAll(ctx, "test_users", "age > ?", 100)
        if err != nil {
            t.Fatalf("SelectAll failed: %v", err)
        }

        if len(users) != 0 {
            t.Errorf("Expected 0 users, got %d", len(users))
        }
    })
}
```

### Testing Transactions / 트랜잭션 테스트

```go
func TestTransaction(t *testing.T) {
    db := setupTestDB(t)
    defer teardownTestDB(t, db)

    ctx := context.Background()

    // Test commit / 커밋 테스트
    t.Run("Commit", func(t *testing.T) {
        err := db.Transaction(ctx, func(tx *mysql.Tx) error {
            tx.Insert(ctx, "test_users", map[string]interface{}{
                "name": "Alice", "email": "alice@example.com", "age": 25,
            })
            return nil // Commit / 커밋
        })

        if err != nil {
            t.Fatalf("Transaction failed: %v", err)
        }

        // Verify data was committed / 데이터가 커밋되었는지 확인
        users, _ := db.SelectAll(ctx, "test_users", "name = ?", "Alice")
        if len(users) != 1 {
            t.Error("Transaction did not commit")
        }
    })

    // Test rollback / 롤백 테스트
    t.Run("Rollback", func(t *testing.T) {
        err := db.Transaction(ctx, func(tx *mysql.Tx) error {
            tx.Insert(ctx, "test_users", map[string]interface{}{
                "name": "Bob", "email": "bob@example.com", "age": 28,
            })
            return errors.New("rollback") // Rollback / 롤백
        })

        if err == nil {
            t.Error("Expected error")
        }

        // Verify data was rolled back / 데이터가 롤백되었는지 확인
        users, _ := db.SelectAll(ctx, "test_users", "name = ?", "Bob")
        if len(users) != 0 {
            t.Error("Transaction did not rollback")
        }
    })
}
```

### Benchmarking / 벤치마킹

```go
func BenchmarkSelectAll(b *testing.B) {
    db := setupTestDB(&testing.T{})
    defer teardownTestDB(&testing.T{}, db)

    ctx := context.Background()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        db.SelectAll(ctx, "test_users")
    }
}

func BenchmarkInsert(b *testing.B) {
    db := setupTestDB(&testing.T{})
    defer teardownTestDB(&testing.T{}, db)

    ctx := context.Background()

    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        db.Insert(ctx, "test_users", map[string]interface{}{
            "name":  fmt.Sprintf("User%d", i),
            "email": fmt.Sprintf("user%d@example.com", i),
            "age":   25,
        })
    }
}
```

### Running Benchmarks / 벤치마크 실행

```bash
# Run benchmarks / 벤치마크 실행
go test ./database/mysql -bench=.

# Run specific benchmark / 특정 벤치마크 실행
go test ./database/mysql -bench=BenchmarkSelectAll

# With memory profiling / 메모리 프로파일링 포함
go test ./database/mysql -bench=. -benchmem

# Output example:
# BenchmarkSelectAll-8    10000    100234 ns/op    2048 B/op    20 allocs/op
```

---

## Performance / 성능

### Performance Characteristics / 성능 특성

| Operation / 작업 | Time Complexity / 시간 복잡도 | Notes / 참고사항 |
|-----------------|------------------------|----------------|
| SelectAll | O(n) | n = number of rows / 행 수 |
| SelectOne | O(1) with index / 인덱스 사용 시 | O(n) without index / 인덱스 없으면 |
| Insert | O(1) | Single row / 단일 행 |
| Update | O(m) | m = affected rows / 영향받은 행 수 |
| Delete | O(m) | m = affected rows / 영향받은 행 수 |
| Transaction | O(sum of operations) / O(작업들의 합) | Depends on operations / 작업에 따라 다름 |

### Optimization Tips / 최적화 팁

#### 1. Use Indexes / 인덱스 사용

```sql
-- Add indexes for frequently queried columns
-- 자주 쿼리되는 컬럼에 인덱스 추가
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_age_city ON users(age, city);
```

#### 2. Select Only Needed Columns / 필요한 컬럼만 선택

```go
// ❌ Bad - select all columns / 나쁜 예 - 모든 컬럼 선택
users, _ := db.SelectAll(ctx, "users")

// ✅ Good - select specific columns / 좋은 예 - 특정 컬럼 선택
users, _ := db.SelectWhere(ctx, "users", "",
    mysql.WithColumns("id", "name", "email"))
```

#### 3. Use Pagination / 페이징 사용

```go
// ❌ Bad - load all data / 나쁜 예 - 모든 데이터 로드
users, _ := db.SelectAll(ctx, "users")

// ✅ Good - paginate / 좋은 예 - 페이징
users, _ := db.SelectWhere(ctx, "users", "",
    mysql.WithLimit(100),
    mysql.WithOffset(0))
```

#### 4. Use Connection Pooling / 연결 풀링 사용

```go
db, _ := mysql.New(
    mysql.WithDSN("..."),
    mysql.WithMaxOpenConns(25),
    mysql.WithMaxIdleConns(10),
)
```

#### 5. Batch Operations / 배치 작업

```go
// ✅ Good - batch in transaction / 좋은 예 - 트랜잭션에서 배치
db.Transaction(ctx, func(tx *mysql.Tx) error {
    for _, item := range items {
        tx.Insert(ctx, "items", item)
    }
    return nil
})
```

### Memory Usage / 메모리 사용량

**Estimated memory per operation / 작업당 예상 메모리**:

- SelectAll: ~2KB base + (rows × avg_row_size)
- Insert: ~1KB
- Update: ~1KB
- Transaction: ~2KB + (operations × avg_op_size)

**Large result set handling / 큰 결과 세트 처리**:

```go
// For very large result sets, use pagination / 매우 큰 결과 세트는 페이징 사용
const batchSize = 10000
for offset := 0; ; offset += batchSize {
    batch, _ := db.SelectWhere(ctx, "large_table", "",
        mysql.WithLimit(batchSize),
        mysql.WithOffset(offset))

    if len(batch) == 0 {
        break
    }

    // Process batch / 배치 처리
    processBatch(batch)
}
```

---

## Contributing Guidelines / 기여 가이드라인

### How to Contribute / 기여 방법

1. **Fork the repository / 저장소 포크**
2. **Create a feature branch / 기능 브랜치 생성**
   ```bash
   git checkout -b feature/new-feature
   ```
3. **Make your changes / 변경사항 작성**
4. **Write tests / 테스트 작성**
5. **Run tests / 테스트 실행**
   ```bash
   go test ./... -v
   ```
6. **Commit with descriptive message / 설명적인 메시지로 커밋**
   ```bash
   git commit -m "Feat: Add new feature X"
   ```
7. **Push to your fork / 포크에 푸시**
   ```bash
   git push origin feature/new-feature
   ```
8. **Create Pull Request / Pull Request 생성**

### Commit Message Format / 커밋 메시지 형식

```
<type>: <subject>

<body>

<footer>
```

**Types / 타입**:
- `Feat`: New feature / 새 기능
- `Fix`: Bug fix / 버그 수정
- `Docs`: Documentation / 문서
- `Refactor`: Code refactoring / 코드 리팩토링
- `Test`: Tests / 테스트
- `Chore`: Build, config / 빌드, 설정

**Examples / 예제**:
```
Feat: Add SelectDistinct method

- Adds new SelectDistinct() method to Simple API
- Includes tests and documentation
- Supports both Client and Tx

Closes #123
```

### Code Review Checklist / 코드 리뷰 체크리스트

- [ ] Code follows package style / 코드가 패키지 스타일을 따름
- [ ] All tests pass / 모든 테스트 통과
- [ ] New features have tests / 새 기능에 테스트 포함
- [ ] Documentation updated / 문서 업데이트됨
- [ ] No breaking changes (or documented) / 호환성 깨는 변경 없음 (또는 문서화됨)
- [ ] Bilingual comments (English/Korean) / 이중 언어 주석 (영문/한글)
- [ ] Error handling included / 에러 처리 포함
- [ ] Performance considered / 성능 고려됨

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

**Variables / 변수**:
```go
// ✅ Good / 좋은 예
userID := 123
userName := "John"
isActive := true

// ❌ Bad / 나쁜 예
user_id := 123  // No snake_case / snake_case 안 됨
UserId := 123   // Not camelCase / camelCase 아님
```

**Functions / 함수**:
```go
// ✅ Exported functions use PascalCase / 내보낸 함수는 PascalCase 사용
func SelectAll(...)
func New(...)

// ✅ Private functions use camelCase / 비공개 함수는 camelCase 사용
func buildQuery(...)
func scanRows(...)
```

**Constants / 상수**:
```go
const (
    // Use PascalCase or SCREAMING_SNAKE_CASE
    // PascalCase 또는 SCREAMING_SNAKE_CASE 사용
    DefaultMaxOpenConns = 25
    DEFAULT_TIMEOUT     = 30
)
```

### Comment Style / 주석 스타일

**Package-level comments / 패키지 레벨 주석**:
```go
// Package mysql provides a simplified MySQL client.
// mysql 패키지는 간단한 MySQL 클라이언트를 제공합니다.
package mysql
```

**Function comments / 함수 주석**:
```go
// SelectAll selects all rows from a table with optional conditions.
// SelectAll은 선택적 조건으로 테이블의 모든 행을 선택합니다.
//
// Example / 예제:
//
//   users, err := db.SelectAll(ctx, "users", "age > ?", 18)
//
func (c *Client) SelectAll(...)
```

**Inline comments / 인라인 주석**:
```go
// Build query / 쿼리 빌드
query := "SELECT * FROM " + table

// Execute with retry / 재시도로 실행
err := c.executeWithRetry(ctx, fn)
```

### Error Handling / 에러 처리

```go
// ✅ Good - descriptive errors / 좋은 예 - 설명적인 에러
if err != nil {
    return fmt.Errorf("failed to select users: %w", err)
}

// ❌ Bad - vague errors / 나쁜 예 - 모호한 에러
if err != nil {
    return err
}
```

### Function Length / 함수 길이

- Keep functions under 50 lines / 함수를 50줄 이하로 유지
- Extract complex logic into helper functions / 복잡한 로직을 헬퍼 함수로 추출
- Single Responsibility Principle / 단일 책임 원칙

### Formatting / 포맷팅

```bash
# Use gofmt / gofmt 사용
gofmt -w .

# Use goimports / goimports 사용
goimports -w .

# Use golangci-lint / golangci-lint 사용
golangci-lint run
```

---

**End of Developer Guide / 개발자 가이드 끝**
