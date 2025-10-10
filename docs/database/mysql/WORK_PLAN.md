# Database/MySQL Package - Work Plan / 작업 계획서
# database/mysql 패키지 - 작업 계획서

**Version / 버전**: v1.3.x
**Author / 작성자**: arkd0ng
**Created / 작성일**: 2025-10-10
**Status / 상태**: Planning / 계획 중

---

## Table of Contents / 목차

1. [Overview / 개요](#overview--개요)
2. [Work Phases / 작업 단계](#work-phases--작업-단계)
3. [Phase 1: Foundation / 1단계: 기초](#phase-1-foundation--1단계-기초)
4. [Phase 2: Core Features / 2단계: 핵심 기능](#phase-2-core-features--2단계-핵심-기능)
5. [Phase 3: Advanced Features / 3단계: 고급 기능](#phase-3-advanced-features--3단계-고급-기능)
6. [Phase 4: Testing & Documentation / 4단계: 테스팅 및 문서화](#phase-4-testing--documentation--4단계-테스팅-및-문서화)
7. [Phase 5: Release / 5단계: 릴리스](#phase-5-release--5단계-릴리스)
8. [Task Dependencies / 작업 의존성](#task-dependencies--작업-의존성)
9. [Quality Checklist / 품질 체크리스트](#quality-checklist--품질-체크리스트)

---

## Overview / 개요

This work plan outlines the detailed implementation steps for the `database/mysql` package. Each phase is broken down into specific tasks with clear acceptance criteria.

이 작업 계획은 `database/mysql` 패키지의 상세한 구현 단계를 설명합니다. 각 단계는 명확한 수용 기준과 함께 구체적인 작업으로 나뉩니다.

### Project Timeline / 프로젝트 타임라인

- **Phase 1**: Foundation / 기초 (2-3 작업 단위)
- **Phase 2**: Core Features / 핵심 기능 (4-6 작업 단위)
- **Phase 3**: Advanced Features / 고급 기능 (3-4 작업 단위)
- **Phase 4**: Testing & Documentation / 테스팅 및 문서화 (3-4 작업 단위)
- **Phase 5**: Release / 릴리스 (1-2 작업 단위)

**Total Estimated Work Units / 총 예상 작업 단위**: 13-19 units

---

## Work Phases / 작업 단계

### Priority Legend / 우선순위 범례

- 🔴 **P0**: Critical / 필수 - Must have for MVP / MVP를 위해 반드시 필요
- 🟡 **P1**: High / 높음 - Important for production readiness / 프로덕션 준비를 위해 중요
- 🟢 **P2**: Medium / 보통 - Nice to have / 있으면 좋음
- 🔵 **P3**: Low / 낮음 - Future enhancement / 향후 개선사항

---

## Phase 1: Foundation / 1단계: 기초

### Task 1.1: Project Structure Setup / 프로젝트 구조 설정

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Create the basic directory structure and initialize the package files.

기본 디렉토리 구조를 생성하고 패키지 파일을 초기화합니다.

**Subtasks / 하위 작업**:

1. Create directory structure / 디렉토리 구조 생성:
   ```bash
   mkdir -p database/mysql
   mkdir -p examples/mysql
   ```

2. Create initial package files / 초기 패키지 파일 생성:
   - `database/mysql/client.go`
   - `database/mysql/connection.go`
   - `database/mysql/rotation.go`
   - `database/mysql/simple.go`
   - `database/mysql/builder.go`
   - `database/mysql/transaction.go`
   - `database/mysql/retry.go`
   - `database/mysql/scan.go`
   - `database/mysql/config.go`
   - `database/mysql/options.go`
   - `database/mysql/errors.go`
   - `database/mysql/types.go`
   - `database/mysql/client_test.go`
   - `database/mysql/rotation_test.go`

3. Add package documentation / 패키지 문서 추가:
   - `database/mysql/doc.go` (package-level documentation)
   - `database/mysql/README.md`

4. Initialize go.mod dependencies / go.mod 의존성 초기화:
   ```bash
   go get github.com/go-sql-driver/mysql@latest
   # Note: Vault integration is user's responsibility / Vault 통합은 사용자 책임
   ```

**Acceptance Criteria / 수용 기준**:
- [ ] All directories created / 모든 디렉토리 생성됨
- [ ] All package files exist with package declaration / 모든 패키지 파일에 패키지 선언이 있음
- [ ] Dependencies added to go.mod / 의존성이 go.mod에 추가됨
- [ ] `go build ./database/mysql` succeeds / 빌드 성공

**Estimated Effort / 예상 소요 시간**: 0.5 work unit

---

### Task 1.2: Error Types Definition / 에러 타입 정의

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Define all custom error types and error handling utilities.

모든 커스텀 에러 타입과 에러 처리 유틸리티를 정의합니다.

**Implementation / 구현**:

File: `database/mysql/errors.go`

```go
package mysql

import (
    "errors"
    "fmt"
    "time"
)

// Predefined errors / 사전 정의된 에러
var (
    ErrConnectionFailed  = errors.New("database connection failed")
    ErrQueryFailed       = errors.New("query execution failed")
    ErrTransactionFailed = errors.New("transaction failed")
    ErrTimeout           = errors.New("operation timeout")
    ErrClosed            = errors.New("database connection closed")
    ErrInvalidDSN        = errors.New("invalid DSN format")
    ErrNoRows            = errors.New("no rows in result set")
)

// DBError represents a database operation error
// DBError는 데이터베이스 작업 에러를 나타냅니다
type DBError struct {
    Op       string        // Operation name / 작업 이름
    Query    string        // SQL query / SQL 쿼리
    Args     []interface{} // Query arguments / 쿼리 인자
    Err      error         // Original error / 원본 에러
    Time     time.Time     // Error timestamp / 에러 타임스탬프
    Duration time.Duration // Operation duration / 작업 소요 시간
}

func (e *DBError) Error() string
func (e *DBError) Unwrap() error
func isRetryableError(err error) bool
func isConnectionError(err error) bool
func isQueryError(err error) bool
```

**Subtasks / 하위 작업**:

1. Define error constants / 에러 상수 정의
2. Implement DBError struct / DBError 구조체 구현
3. Implement Error() method / Error() 메서드 구현
4. Implement Unwrap() method / Unwrap() 메서드 구현
5. Implement error classification functions / 에러 분류 함수 구현:
   - `isRetryableError()`
   - `isConnectionError()`
   - `isQueryError()`

**Acceptance Criteria / 수용 기준**:
- [ ] All error types defined with bilingual comments / 모든 에러 타입이 이중 언어 주석과 함께 정의됨
- [ ] DBError implements error interface / DBError가 error 인터페이스 구현
- [ ] Error classification functions work correctly / 에러 분류 함수가 올바르게 작동
- [ ] Unit tests for all error functions / 모든 에러 함수에 대한 유닛 테스트
- [ ] `go test ./database/mysql -run TestError` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 1.3: Configuration Structure / 설정 구조

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement the configuration structure and default values.

설정 구조와 기본값을 구현합니다.

**Implementation / 구현**:

File: `database/mysql/config.go`

```go
package mysql

import (
    "crypto/tls"
    "time"

    "github.com/arkd0ng/go-utils/logging"
)

// config holds all configuration options for the MySQL client
// config는 MySQL 클라이언트의 모든 설정 옵션을 보유합니다
type config struct {
    // Connection settings / 연결 설정
    dsn             string
    maxOpenConns    int
    maxIdleConns    int
    connMaxLifetime time.Duration
    connMaxIdleTime time.Duration

    // Credential rotation (optional) / 자격 증명 순환 (선택)
    credRefreshFunc    CredentialRefreshFunc  // User function / 사용자 함수
    poolCount          int                     // Number of connection pools / 연결 풀 개수
    rotationInterval   time.Duration           // Rotation interval / 순환 주기

    // Timeout settings / 타임아웃 설정
    connectTimeout  time.Duration
    queryTimeout    time.Duration

    // Retry settings / 재시도 설정
    maxRetries      int
    retryDelay      time.Duration

    // Logging settings / 로깅 설정
    logger             *logging.Logger
    logQueries         bool
    logSlowQueries     bool
    slowQueryThreshold time.Duration

    // Health check settings / 헬스 체크 설정
    enableHealthCheck   bool
    healthCheckInterval time.Duration

    // Security settings / 보안 설정
    enableTLS bool
    tlsConfig *tls.Config
}

func defaultConfig() *config
func (c *config) validate() error
```

**Subtasks / 하위 작업**:

1. Define config struct with all fields / 모든 필드가 있는 config 구조체 정의
2. Implement defaultConfig() function / defaultConfig() 함수 구현
3. Implement validate() method / validate() 메서드 구현
4. Add bilingual comments for all fields / 모든 필드에 이중 언어 주석 추가

**Acceptance Criteria / 수용 기준**:
- [ ] All configuration fields defined / 모든 설정 필드 정의됨
- [ ] Default values are production-ready / 기본값이 프로덕션에 적합함
- [ ] validate() checks all required fields / validate()가 모든 필수 필드 확인
- [ ] Unit tests for config validation / 설정 검증 유닛 테스트
- [ ] `go test ./database/mysql -run TestConfig` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

## Phase 2: Core Features / 2단계: 핵심 기능

### Task 2.1: Functional Options Implementation / 함수형 옵션 구현

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement all functional option functions for flexible configuration.

유연한 설정을 위한 모든 함수형 옵션 함수를 구현합니다.

**Implementation / 구현**:

File: `database/mysql/options.go`

```go
package mysql

// Option is a function that configures the MySQL client
// Option은 MySQL 클라이언트를 설정하는 함수입니다
type Option func(*config) error

// Connection options / 연결 옵션
func WithDSN(dsn string) Option
func WithMaxOpenConns(n int) Option
func WithMaxIdleConns(n int) Option
func WithConnMaxLifetime(d time.Duration) Option
func WithConnMaxIdleTime(d time.Duration) Option

// Timeout options / 타임아웃 옵션
func WithConnectTimeout(d time.Duration) Option
func WithQueryTimeout(d time.Duration) Option

// Retry options / 재시도 옵션
func WithMaxRetries(n int) Option
func WithRetryDelay(d time.Duration) Option

// Logging options / 로깅 옵션
func WithLogger(logger *logging.Logger) Option
func WithQueryLogging(enable bool) Option
func WithSlowQueryLogging(enable bool) Option
func WithSlowQueryThreshold(d time.Duration) Option

// Health check options / 헬스 체크 옵션
func WithHealthCheck(enable bool) Option
func WithHealthCheckInterval(d time.Duration) Option

// Security options / 보안 옵션
func WithTLS(tlsConfig *tls.Config) Option
```

**Subtasks / 하위 작업**:

1. Define Option type / Option 타입 정의
2. Implement all connection option functions / 모든 연결 옵션 함수 구현
3. Implement all timeout option functions / 모든 타임아웃 옵션 함수 구현
4. Implement all retry option functions / 모든 재시도 옵션 함수 구현
5. Implement all logging option functions / 모든 로깅 옵션 함수 구현
6. Implement all health check option functions / 모든 헬스 체크 옵션 함수 구현
7. Implement all security option functions / 모든 보안 옵션 함수 구현
8. Add comprehensive bilingual documentation / 포괄적인 이중 언어 문서 추가
9. Add usage examples in comments / 주석에 사용 예제 추가

**Acceptance Criteria / 수용 기준**:
- [ ] All option functions implemented / 모든 옵션 함수 구현됨
- [ ] Each function has bilingual comments / 각 함수에 이중 언어 주석이 있음
- [ ] Each function has usage example / 각 함수에 사용 예제가 있음
- [ ] Option functions validate input / 옵션 함수가 입력 검증
- [ ] Unit tests for all options / 모든 옵션에 대한 유닛 테스트
- [ ] `go test ./database/mysql -run TestOptions` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

### Task 2.2: Client Core Implementation / 클라이언트 핵심 구현

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement the main Client struct and basic connection management.

메인 Client 구조체와 기본 연결 관리를 구현합니다.

**Implementation / 구현**:

File: `database/mysql/client.go`

```go
package mysql

import (
    "context"
    "database/sql"
    "sync"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/arkd0ng/go-utils/logging"
)

// Client represents a MySQL database client
// Client는 MySQL 데이터베이스 클라이언트를 나타냅니다
type Client struct {
    // Connection pool / 연결 풀
    connections      []*sql.DB         // Multiple connections for rotation / 순환을 위한 여러 연결
    currentIdx       int               // Current connection index / 현재 연결 인덱스
    connectionsMu    sync.RWMutex      // Connection array synchronization / 연결 배열 동기화

    // Configuration / 설정
    config           *config           // Configuration / 설정
    credProvider     CredentialProvider // Credential provider / 자격 증명 제공자

    // State / 상태
    logger           *logging.Logger   // Logger (optional) / 로거 (선택)
    healthy          bool              // Connection state / 연결 상태

    // Background tasks / 백그라운드 작업
    stopChan         chan struct{}     // Stop signal / 종료 신호
    healthCheckStop  chan struct{}     // Health check stop / 헬스 체크 중지
    rotationStop     chan struct{}     // Rotation stop / 순환 중지

    // Synchronization / 동기화
    mu               sync.RWMutex      // General synchronization / 일반 동기화
    closed           bool              // Closed flag / 종료 플래그
}

func New(opts ...Option) (*Client, error)
func (c *Client) Close() error
func (c *Client) Ping(ctx context.Context) error
func (c *Client) Stats() sql.DBStats
func (c *Client) wrapError(op, query string, args []interface{}, err error, duration time.Duration) error
```

**Subtasks / 하위 작업**:

1. Define Client struct / Client 구조체 정의
2. Implement New() constructor / New() 생성자 구현:
   - Apply all options / 모든 옵션 적용
   - Validate configuration / 설정 검증
   - Open database connection / 데이터베이스 연결 오픈
   - Configure connection pool / 연결 풀 설정
   - Start health check goroutine (if enabled) / 헬스 체크 goroutine 시작
3. Implement Close() method / Close() 메서드 구현:
   - Stop health check / 헬스 체크 중지
   - Wait for health check to finish / 헬스 체크 완료 대기
   - Close database connection / 데이터베이스 연결 종료
4. Implement Ping() method / Ping() 메서드 구현
5. Implement Stats() method / Stats() 메서드 구현
6. Implement wrapError() helper / wrapError() 헬퍼 구현

**Acceptance Criteria / 수용 기준**:
- [ ] Client struct properly defined / Client 구조체가 올바르게 정의됨
- [ ] New() creates working client / New()가 작동하는 클라이언트 생성
- [ ] Connection pool configured correctly / 연결 풀이 올바르게 설정됨
- [ ] Close() gracefully shuts down / Close()가 우아하게 종료
- [ ] Thread-safe with mutex / 뮤텍스로 스레드 안전
- [ ] Unit tests for client lifecycle / 클라이언트 생명주기 유닛 테스트
- [ ] `go test ./database/mysql -run TestClient` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 2 work units

---

### Task 2.3: Query Execution Methods / 쿼리 실행 메서드

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement all query execution methods with context support.

Context 지원이 있는 모든 쿼리 실행 메서드를 구현합니다.

**Implementation / 구현**:

File: `database/mysql/client.go` (continued)

```go
func (c *Client) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
func (c *Client) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
func (c *Client) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
func (c *Client) executeWithRetry(ctx context.Context, op string, fn func() error) error
func (c *Client) logQuery(query string, args []interface{}, duration time.Duration, err error)
```

**Subtasks / 하위 작업**:

1. Implement Query() method / Query() 메서드 구현:
   - Context support / Context 지원
   - Query logging / 쿼리 로깅
   - Error wrapping / 에러 래핑
   - Execution time tracking / 실행 시간 추적

2. Implement QueryRow() method / QueryRow() 메서드 구현:
   - Similar to Query() but for single row / Query()와 유사하지만 단일 행용

3. Implement Exec() method / Exec() 메서드 구현:
   - For INSERT, UPDATE, DELETE / INSERT, UPDATE, DELETE용
   - Return sql.Result / sql.Result 반환

4. Implement executeWithRetry() helper / executeWithRetry() 헬퍼 구현:
   - Retry logic with exponential backoff / 지수 백오프를 사용한 재시도 로직
   - Only retry on transient errors / 일시적인 에러에만 재시도
   - Respect context cancellation / Context 취소 존중

5. Implement logQuery() helper / logQuery() 헬퍼 구현:
   - Log all queries (if enabled) / 모든 쿼리 로깅 (활성화된 경우)
   - Log slow queries / 느린 쿼리 로깅
   - Include execution time / 실행 시간 포함

**Acceptance Criteria / 수용 기준**:
- [ ] All query methods implemented / 모든 쿼리 메서드 구현됨
- [ ] Context timeout works / Context 타임아웃 작동
- [ ] Context cancellation works / Context 취소 작동
- [ ] Retry logic works for transient errors / 일시적 에러에 대한 재시도 로직 작동
- [ ] Query logging works / 쿼리 로깅 작동
- [ ] Slow query logging works / 느린 쿼리 로깅 작동
- [ ] Unit tests with sqlmock / sqlmock를 사용한 유닛 테스트
- [ ] `go test ./database/mysql -run TestQuery` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 2 work units

---

### Task 2.4: Transaction Support / 트랜잭션 지원

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement transaction support with commit and rollback.

커밋과 롤백을 지원하는 트랜잭션을 구현합니다.

**Implementation / 구현**:

File: `database/mysql/transaction.go`

```go
package mysql

import (
    "context"
    "database/sql"
    "sync"

    "github.com/arkd0ng/go-utils/logging"
)

// Transaction represents a database transaction
// Transaction은 데이터베이스 트랜잭션을 나타냅니다
type Transaction struct {
    tx       *sql.Tx
    client   *Client
    logger   *logging.Logger
    mu       sync.Mutex
    finished bool
}

func (c *Client) Begin(ctx context.Context) (*Transaction, error)
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Transaction, error)
func (t *Transaction) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
func (t *Transaction) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
func (t *Transaction) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
func (t *Transaction) Commit() error
func (t *Transaction) Rollback() error
func (t *Transaction) logQuery(query string, args []interface{}, duration time.Duration, err error)
```

**Subtasks / 하위 작업**:

1. Define Transaction struct / Transaction 구조체 정의

2. Implement Begin() method / Begin() 메서드 구현:
   - Start transaction with default options / 기본 옵션으로 트랜잭션 시작
   - Return Transaction wrapper / Transaction 래퍼 반환

3. Implement BeginTx() method / BeginTx() 메서드 구현:
   - Support custom transaction options / 커스텀 트랜잭션 옵션 지원
   - Support isolation levels / 격리 수준 지원

4. Implement Transaction query methods / Transaction 쿼리 메서드 구현:
   - Query()
   - QueryRow()
   - Exec()
   - All with context support / 모두 Context 지원

5. Implement Commit() method / Commit() 메서드 구현:
   - Check if already finished / 이미 완료되었는지 확인
   - Log commit / 커밋 로깅

6. Implement Rollback() method / Rollback() 메서드 구현:
   - Check if already finished / 이미 완료되었는지 확인
   - Log rollback / 롤백 로깅

7. Implement logQuery() helper / logQuery() 헬퍼 구현

**Acceptance Criteria / 수용 기준**:
- [ ] Transaction struct properly defined / Transaction 구조체가 올바르게 정의됨
- [ ] Begin() creates working transaction / Begin()이 작동하는 트랜잭션 생성
- [ ] BeginTx() supports isolation levels / BeginTx()가 격리 수준 지원
- [ ] All query methods work within transaction / 모든 쿼리 메서드가 트랜잭션 내에서 작동
- [ ] Commit() successfully commits / Commit()이 성공적으로 커밋
- [ ] Rollback() successfully rolls back / Rollback()이 성공적으로 롤백
- [ ] Cannot commit/rollback twice / 두 번 커밋/롤백 불가
- [ ] Unit tests for all transaction scenarios / 모든 트랜잭션 시나리오 유닛 테스트
- [ ] `go test ./database/mysql -run TestTransaction` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

## Phase 3: Advanced Features / 3단계: 고급 기능

### Task 3.1: Health Check Implementation / 헬스 체크 구현

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Implement periodic health checks and automatic reconnection.

주기적 헬스 체크와 자동 재연결을 구현합니다.

**Implementation / 구현**:

File: `database/mysql/health.go`

```go
package mysql

import (
    "context"
    "time"
)

func (c *Client) startHealthCheck()
func (c *Client) stopHealthCheck()
func (c *Client) performHealthCheck(ctx context.Context) error
func (c *Client) IsHealthy(ctx context.Context) bool
func (c *Client) reconnect(ctx context.Context) error
```

**Subtasks / 하위 작업**:

1. Implement startHealthCheck() / startHealthCheck() 구현:
   - Start goroutine / goroutine 시작
   - Periodic ping to database / 데이터베이스에 주기적 ping
   - Call performHealthCheck() / performHealthCheck() 호출

2. Implement stopHealthCheck() / stopHealthCheck() 구현:
   - Send stop signal / 중지 신호 전송
   - Wait for goroutine to finish / goroutine 완료 대기

3. Implement performHealthCheck() / performHealthCheck() 구현:
   - Ping database / 데이터베이스 ping
   - Log health status / 헬스 상태 로깅
   - Trigger reconnect on failure / 실패 시 재연결 트리거

4. Implement IsHealthy() / IsHealthy() 구현:
   - Check if connection is healthy / 연결이 정상인지 확인
   - Return boolean / boolean 반환

5. Implement reconnect() / reconnect() 구현:
   - Close existing connection / 기존 연결 종료
   - Create new connection / 새 연결 생성
   - Reconfigure connection pool / 연결 풀 재설정

**Acceptance Criteria / 수용 기준**:
- [ ] Health check runs periodically / 헬스 체크가 주기적으로 실행됨
- [ ] Health check detects connection issues / 헬스 체크가 연결 문제 감지
- [ ] Automatic reconnection works / 자동 재연결 작동
- [ ] IsHealthy() returns correct status / IsHealthy()가 올바른 상태 반환
- [ ] Health check stops on Close() / Close() 시 헬스 체크 중지
- [ ] Unit tests for health check / 헬스 체크 유닛 테스트
- [ ] `go test ./database/mysql -run TestHealth` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

### Task 3.2: Connection Pool Metrics / 연결 풀 메트릭

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Implement connection pool metrics and statistics.

연결 풀 메트릭과 통계를 구현합니다.

**Implementation / 구현**:

File: `database/mysql/metrics.go`

```go
package mysql

import (
    "database/sql"
    "encoding/json"
)

// Metrics represents database connection pool metrics
// Metrics는 데이터베이스 연결 풀 메트릭을 나타냅니다
type Metrics struct {
    MaxOpenConnections int           // Maximum number of open connections / 최대 오픈 연결 수
    OpenConnections    int           // Current open connections / 현재 오픈 연결 수
    InUse              int           // Connections in use / 사용 중인 연결
    Idle               int           // Idle connections / 유휴 연결
    WaitCount          int64         // Total wait count / 총 대기 횟수
    WaitDuration       time.Duration // Total wait duration / 총 대기 시간
    MaxIdleClosed      int64         // Connections closed due to idle / 유휴로 인해 종료된 연결
    MaxLifetimeClosed  int64         // Connections closed due to lifetime / 수명으로 인해 종료된 연결
}

func (c *Client) GetMetrics() Metrics
func (m Metrics) String() string
func (m Metrics) JSON() ([]byte, error)
```

**Subtasks / 하위 작업**:

1. Define Metrics struct / Metrics 구조체 정의
2. Implement GetMetrics() method / GetMetrics() 메서드 구현:
   - Get stats from sql.DB / sql.DB에서 통계 가져오기
   - Convert to Metrics struct / Metrics 구조체로 변환
3. Implement String() method / String() 메서드 구현:
   - Human-readable format / 사람이 읽을 수 있는 형식
4. Implement JSON() method / JSON() 메서드 구현:
   - JSON format for APIs / API를 위한 JSON 형식

**Acceptance Criteria / 수용 기준**:
- [ ] Metrics struct defined / Metrics 구조체 정의됨
- [ ] GetMetrics() returns accurate data / GetMetrics()가 정확한 데이터 반환
- [ ] String() returns readable format / String()이 읽기 쉬운 형식 반환
- [ ] JSON() returns valid JSON / JSON()이 유효한 JSON 반환
- [ ] Unit tests for metrics / 메트릭 유닛 테스트
- [ ] `go test ./database/mysql -run TestMetrics` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 3.3: TLS/SSL Support / TLS/SSL 지원

**Priority / 우선순위**: 🟢 P2

**Description / 설명**:
Add TLS/SSL support for secure database connections.

안전한 데이터베이스 연결을 위한 TLS/SSL 지원을 추가합니다.

**Implementation / 구현**:

Update: `database/mysql/client.go`

```go
import "crypto/tls"
import "github.com/go-sql-driver/mysql"

func (c *Client) configureTLS() error {
    if !c.config.enableTLS {
        return nil
    }

    // Register custom TLS config / 커스텀 TLS 설정 등록
    mysql.RegisterTLSConfig("custom", c.config.tlsConfig)

    // Update DSN to use TLS / DSN을 TLS 사용하도록 업데이트
    // ...

    return nil
}
```

**Subtasks / 하위 작업**:

1. Add TLS configuration to config struct / config 구조체에 TLS 설정 추가
2. Implement WithTLS() option / WithTLS() 옵션 구현
3. Implement configureTLS() helper / configureTLS() 헬퍼 구현
4. Update DSN to include TLS parameter / TLS 매개변수를 포함하도록 DSN 업데이트
5. Add documentation for TLS setup / TLS 설정 문서 추가

**Acceptance Criteria / 수용 기준**:
- [ ] WithTLS() option works / WithTLS() 옵션 작동
- [ ] TLS configuration registered / TLS 설정 등록됨
- [ ] Connection uses TLS when enabled / 활성화 시 연결이 TLS 사용
- [ ] Unit tests for TLS configuration / TLS 설정 유닛 테스트
- [ ] Integration test with TLS (optional) / TLS를 사용한 통합 테스트 (선택사항)
- [ ] `go test ./database/mysql -run TestTLS` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

## Phase 4: Testing & Documentation / 4단계: 테스팅 및 문서화

### Task 4.1: Comprehensive Unit Tests / 포괄적인 유닛 테스트

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Write comprehensive unit tests for all package components.

모든 패키지 컴포넌트에 대한 포괄적인 유닛 테스트를 작성합니다.

**Test Files / 테스트 파일**:
- `client_test.go`
- `transaction_test.go`
- `options_test.go`
- `errors_test.go`
- `health_test.go`
- `metrics_test.go`

**Subtasks / 하위 작업**:

1. Setup test infrastructure / 테스트 인프라 설정:
   - Use `sqlmock` for database mocking / 데이터베이스 모킹에 `sqlmock` 사용
   - Create test helpers / 테스트 헬퍼 생성

2. Write unit tests for Client / Client 유닛 테스트 작성:
   - New() constructor / New() 생성자
   - Close() method / Close() 메서드
   - Ping() method / Ping() 메서드
   - Query methods / 쿼리 메서드
   - Error scenarios / 에러 시나리오

3. Write unit tests for Transaction / Transaction 유닛 테스트 작성:
   - Begin() / BeginTx()
   - Query methods within transaction / 트랜잭션 내 쿼리 메서드
   - Commit() / Rollback()
   - Error scenarios / 에러 시나리오

4. Write unit tests for Options / Options 유닛 테스트 작성:
   - All option functions / 모든 옵션 함수
   - Option validation / 옵션 검증

5. Write unit tests for Errors / Errors 유닛 테스트 작성:
   - Error wrapping / 에러 래핑
   - Error classification / 에러 분류

6. Write unit tests for Health Check / Health Check 유닛 테스트 작성:
   - Health check goroutine / 헬스 체크 goroutine
   - Reconnection logic / 재연결 로직

7. Write unit tests for Metrics / Metrics 유닛 테스트 작성:
   - GetMetrics() / String() / JSON()

**Acceptance Criteria / 수용 기준**:
- [ ] Test coverage > 80% / 테스트 커버리지 > 80%
- [ ] All public methods tested / 모든 공개 메서드 테스트됨
- [ ] Error scenarios tested / 에러 시나리오 테스트됨
- [ ] Table-driven tests used / 테이블 기반 테스트 사용
- [ ] All tests have bilingual comments / 모든 테스트에 이중 언어 주석
- [ ] `go test ./database/mysql -v` passes / 테스트 통과
- [ ] `go test ./database/mysql -cover` shows > 80% / 커버리지 > 80%

**Estimated Effort / 예상 소요 시간**: 2.5 work units

---

### Task 4.2: Integration Tests / 통합 테스트

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Write integration tests with real MySQL instance.

실제 MySQL 인스턴스로 통합 테스트를 작성합니다.

**Implementation / 구현**:

File: `database/mysql/integration_test.go`

```go
// +build integration

package mysql_test

import "testing"

func TestIntegration_Connection(t *testing.T)
func TestIntegration_Query(t *testing.T)
func TestIntegration_Transaction(t *testing.T)
func TestIntegration_HealthCheck(t *testing.T)
```

**Subtasks / 하위 작업**:

1. Create integration test file with build tag / 빌드 태그가 있는 통합 테스트 파일 생성

2. Setup MySQL Docker container for tests / 테스트를 위한 MySQL Docker 컨테이너 설정:
   - Create docker-compose.yml / docker-compose.yml 생성
   - Setup test database / 테스트 데이터베이스 설정

3. Write integration tests / 통합 테스트 작성:
   - Connection test / 연결 테스트
   - CRUD operations / CRUD 작업
   - Transaction test / 트랜잭션 테스트
   - Health check test / 헬스 체크 테스트
   - Reconnection test / 재연결 테스트

4. Add CI/CD integration (optional) / CI/CD 통합 추가 (선택사항)

**Acceptance Criteria / 수용 기준**:
- [ ] Integration tests run with Docker / Docker로 통합 테스트 실행
- [ ] All major scenarios tested / 모든 주요 시나리오 테스트됨
- [ ] Tests can be skipped without Docker / Docker 없이 테스트 건너뛸 수 있음
- [ ] `go test -tags=integration ./database/mysql -v` passes / 테스트 통과

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

### Task 4.3: Package Documentation / 패키지 문서화

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Create comprehensive package documentation.

포괄적인 패키지 문서를 작성합니다.

**Documentation Files / 문서 파일**:
- `database/mysql/README.md`
- `database/mysql/doc.go`
- `docs/database/mysql/USER_MANUAL.md`
- `docs/database/mysql/DEVELOPER_GUIDE.md`

**Subtasks / 하위 작업**:

1. Create package README / 패키지 README 작성:
   - Overview / 개요
   - Installation / 설치
   - Quick start / 빠른 시작
   - Basic examples / 기본 예제
   - Configuration options / 설정 옵션

2. Create doc.go / doc.go 작성:
   - Package-level documentation / 패키지 레벨 문서
   - Usage examples / 사용 예제

3. Create USER_MANUAL.md / USER_MANUAL.md 작성:
   - Introduction / 소개
   - Installation / 설치
   - Quick Start / 빠른 시작
   - Configuration Reference / 설정 참조
   - Usage Patterns / 사용 패턴
   - Common Use Cases / 일반적인 사용 사례
   - Best Practices / 모범 사례
   - Troubleshooting / 문제 해결
   - FAQ

4. Create DEVELOPER_GUIDE.md / DEVELOPER_GUIDE.md 작성:
   - Architecture Overview / 아키텍처 개요
   - Package Structure / 패키지 구조
   - Core Components / 핵심 컴포넌트
   - Internal Implementation / 내부 구현
   - Design Patterns / 디자인 패턴
   - Adding New Features / 새 기능 추가
   - Testing Guide / 테스트 가이드
   - Performance / 성능
   - Contributing Guidelines / 기여 가이드라인
   - Code Style / 코드 스타일

**Acceptance Criteria / 수용 기준**:
- [ ] All documentation files created / 모든 문서 파일 생성됨
- [ ] All documentation is bilingual (English/Korean) / 모든 문서가 이중 언어
- [ ] Code examples are tested and working / 코드 예제가 테스트되고 작동
- [ ] Documentation follows CLAUDE.md standards / 문서가 CLAUDE.md 표준 준수
- [ ] `go doc github.com/arkd0ng/go-utils/database/mysql` works / 작동함

**Estimated Effort / 예상 소요 시간**: 2 work units

---

### Task 4.4: Usage Examples / 사용 예제

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Create comprehensive usage examples.

포괄적인 사용 예제를 작성합니다.

**Implementation / 구현**:

File: `examples/mysql/main.go`

**Subtasks / 하위 작업**:

1. Create basic connection example / 기본 연결 예제 작성
2. Create query examples / 쿼리 예제 작성:
   - SELECT queries / SELECT 쿼리
   - INSERT queries / INSERT 쿼리
   - UPDATE queries / UPDATE 쿼리
   - DELETE queries / DELETE 쿼리
3. Create transaction example / 트랜잭션 예제 작성
4. Create error handling example / 에러 처리 예제 작성
5. Create health check example / 헬스 체크 예제 작성
6. Create metrics example / 메트릭 예제 작성
7. Create configuration examples / 설정 예제 작성
8. Add bilingual comments / 이중 언어 주석 추가

**Acceptance Criteria / 수용 기준**:
- [ ] All major features demonstrated / 모든 주요 기능 시연됨
- [ ] Examples are well-commented / 예제에 주석이 잘 달려 있음
- [ ] Examples can be run / 예제를 실행할 수 있음
- [ ] `go run examples/mysql/main.go` works / 작동함

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

## Phase 5: Release / 5단계: 릴리스

### Task 5.1: Final Review and Polish / 최종 검토 및 마무리

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Final code review, polish, and quality checks.

최종 코드 검토, 마무리 및 품질 확인.

**Subtasks / 하위 작업**:

1. Code review / 코드 검토:
   - Check all bilingual comments / 모든 이중 언어 주석 확인
   - Verify error handling / 에러 처리 검증
   - Check for code smells / 코드 스멜 확인

2. Run all tests / 모든 테스트 실행:
   - `go test ./database/mysql -v`
   - `go test -tags=integration ./database/mysql -v`
   - `go test ./database/mysql -cover`

3. Run benchmarks / 벤치마크 실행:
   - `go test ./database/mysql -bench=.`

4. Run static analysis / 정적 분석 실행:
   - `go vet ./database/mysql`
   - `golint ./database/mysql`
   - `staticcheck ./database/mysql`

5. Check documentation / 문서 확인:
   - Verify all examples work / 모든 예제 작동 확인
   - Check for typos / 오타 확인
   - Verify bilingual consistency / 이중 언어 일관성 확인

**Acceptance Criteria / 수용 기준**:
- [ ] All tests pass / 모든 테스트 통과
- [ ] Test coverage > 80% / 테스트 커버리지 > 80%
- [ ] No linter warnings / linter 경고 없음
- [ ] All documentation complete / 모든 문서 완성
- [ ] All examples work / 모든 예제 작동

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 5.2: Update Root Documentation / 루트 문서 업데이트

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Update root README and CHANGELOG.

루트 README와 CHANGELOG를 업데이트합니다.

**Subtasks / 하위 작업**:

1. Update root README.md / 루트 README.md 업데이트:
   - Add database/mysql to package list / 패키지 목록에 database/mysql 추가
   - Add quick example / 빠른 예제 추가

2. Update CHANGELOG.md / CHANGELOG.md 업데이트:
   - Add v1.3.x entry / v1.3.x 항목 추가
   - Link to detailed changelog / 상세 changelog 링크

3. Update CHANGELOG-v1.3.md / CHANGELOG-v1.3.md 업데이트:
   - List all changes / 모든 변경사항 나열
   - List all new features / 모든 새 기능 나열

4. Update CLAUDE.md / CLAUDE.md 업데이트:
   - Add database/mysql architecture / database/mysql 아키텍처 추가
   - Update dependencies / 의존성 업데이트
   - Update examples / 예제 업데이트

**Acceptance Criteria / 수용 기준**:
- [ ] Root README updated / 루트 README 업데이트됨
- [ ] CHANGELOG updated / CHANGELOG 업데이트됨
- [ ] CHANGELOG-v1.3.md complete / CHANGELOG-v1.3.md 완성됨
- [ ] CLAUDE.md updated / CLAUDE.md 업데이트됨
- [ ] All documentation consistent / 모든 문서 일관성 있음

**Estimated Effort / 예상 소요 시간**: 0.5 work unit

---

### Task 5.3: Git Commit and Push / Git 커밋 및 푸시

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Commit all changes and push to GitHub.

모든 변경사항을 커밋하고 GitHub에 푸시합니다.

**Subtasks / 하위 작업**:

1. Update version to v1.3.001 / 버전을 v1.3.001로 업데이트

2. Stage all files / 모든 파일 스테이징:
   ```bash
   git add .
   ```

3. Commit with descriptive message / 설명이 포함된 메시지로 커밋:
   ```bash
   git commit -m "Feat: Add database/mysql package with connection pooling, transactions, and health checks"
   ```

4. Push to GitHub / GitHub에 푸시:
   ```bash
   git push origin main
   ```

5. Create Git tag / Git 태그 생성:
   ```bash
   git tag v1.3.001
   git push origin v1.3.001
   ```

**Acceptance Criteria / 수용 기준**:
- [ ] All files committed / 모든 파일 커밋됨
- [ ] Commit message follows convention / 커밋 메시지가 규칙 준수
- [ ] Pushed to GitHub / GitHub에 푸시됨
- [ ] Git tag created / Git 태그 생성됨

**Estimated Effort / 예상 소요 시간**: 0.5 work unit

---

## Task Dependencies / 작업 의존성

### Dependency Graph / 의존성 그래프

```
Phase 1: Foundation
├─ 1.1 Project Structure ──────┐
├─ 1.2 Error Types ────────────┤
└─ 1.3 Configuration ──────────┤
                               │
Phase 2: Core Features         │
├─ 2.1 Options ◄───────────────┤
├─ 2.2 Client Core ◄───────────┼──────┐
├─ 2.3 Query Methods ◄─────────┘      │
└─ 2.4 Transactions ◄─────────────────┘
                                       │
Phase 3: Advanced Features             │
├─ 3.1 Health Check ◄──────────────────┤
├─ 3.2 Metrics ◄───────────────────────┤
└─ 3.3 TLS ◄───────────────────────────┘
                                       │
Phase 4: Testing & Documentation       │
├─ 4.1 Unit Tests ◄────────────────────┤
├─ 4.2 Integration Tests ◄─────────────┤
├─ 4.3 Documentation ◄─────────────────┤
└─ 4.4 Examples ◄──────────────────────┘
                                       │
Phase 5: Release                       │
├─ 5.1 Review ◄────────────────────────┤
├─ 5.2 Update Docs ◄───────────────────┤
└─ 5.3 Commit & Push ◄─────────────────┘
```

### Critical Path / 중요 경로

The critical path for MVP (Minimum Viable Product):

MVP를 위한 중요 경로:

1. Task 1.1 → 1.2 → 1.3 (Foundation)
2. Task 2.1 → 2.2 → 2.3 → 2.4 (Core Features)
3. Task 4.1 (Unit Tests)
4. Task 4.3 (Documentation)
5. Task 5.1 → 5.2 → 5.3 (Release)

**Optional for MVP / MVP에 선택사항**:
- Task 3.1, 3.2, 3.3 (Can be added post-MVP)
- Task 4.2 (Integration tests nice to have)
- Task 4.4 (Examples nice to have)

---

## Quality Checklist / 품질 체크리스트

### Code Quality / 코드 품질

- [ ] All code has bilingual comments (English/Korean) / 모든 코드에 이중 언어 주석
- [ ] Follows Go best practices / Go 모범 사례 준수
- [ ] No linter warnings / linter 경고 없음
- [ ] No race conditions / race condition 없음
- [ ] Proper error handling / 적절한 에러 처리
- [ ] Thread-safe with mutexes / 뮤텍스로 스레드 안전

### Testing / 테스팅

- [ ] Unit test coverage > 80% / 유닛 테스트 커버리지 > 80%
- [ ] All public methods tested / 모든 공개 메서드 테스트됨
- [ ] Error scenarios tested / 에러 시나리오 테스트됨
- [ ] Table-driven tests used / 테이블 기반 테스트 사용
- [ ] Integration tests pass / 통합 테스트 통과
- [ ] Benchmarks run successfully / 벤치마크 성공적으로 실행

### Documentation / 문서화

- [ ] Package README complete / 패키지 README 완성
- [ ] doc.go with examples / 예제가 있는 doc.go
- [ ] USER_MANUAL.md complete / USER_MANUAL.md 완성
- [ ] DEVELOPER_GUIDE.md complete / DEVELOPER_GUIDE.md 완성
- [ ] All examples work / 모든 예제 작동
- [ ] Bilingual consistency / 이중 언어 일관성

### Security / 보안

- [ ] No SQL injection vulnerabilities / SQL 인젝션 취약점 없음
- [ ] Always use prepared statements / 항상 prepared statement 사용
- [ ] No hardcoded credentials / 하드코딩된 자격 증명 없음
- [ ] TLS support implemented / TLS 지원 구현
- [ ] Secure error messages / 안전한 에러 메시지

### Performance / 성능

- [ ] Connection pooling works efficiently / 연결 풀링이 효율적으로 작동
- [ ] No connection leaks / 연결 누수 없음
- [ ] Proper timeout handling / 적절한 타임아웃 처리
- [ ] Efficient retry logic / 효율적인 재시도 로직
- [ ] Slow query logging works / 느린 쿼리 로깅 작동

### Release / 릴리스

- [ ] Version updated in cfg/app.yaml / cfg/app.yaml에 버전 업데이트됨
- [ ] CHANGELOG updated / CHANGELOG 업데이트됨
- [ ] Root README updated / 루트 README 업데이트됨
- [ ] CLAUDE.md updated / CLAUDE.md 업데이트됨
- [ ] All tests pass / 모든 테스트 통과
- [ ] Git tag created / Git 태그 생성됨

---

## Conclusion / 결론

This work plan provides a detailed roadmap for implementing the `database/mysql` package. By following these phases and tasks in order, we ensure a systematic, high-quality implementation that aligns with the project's design principles and standards.

이 작업 계획은 `database/mysql` 패키지 구현을 위한 상세한 로드맵을 제공합니다. 이러한 단계와 작업을 순서대로 따르면 프로젝트의 설계 원칙과 표준에 부합하는 체계적이고 고품질의 구현을 보장할 수 있습니다.

**Key Success Factors / 주요 성공 요소**:

1. **Incremental Development / 점진적 개발**: Build and test each component before moving to the next / 다음으로 이동하기 전에 각 컴포넌트를 빌드하고 테스트
2. **Quality First / 품질 우선**: Never compromise on tests and documentation / 테스트와 문서화를 절대 타협하지 않음
3. **Consistency / 일관성**: Follow established patterns from random and logging packages / random 및 logging 패키지의 확립된 패턴 준수
4. **Bilingual Excellence / 이중 언어 우수성**: Maintain high-quality English and Korean documentation / 고품질 영문 및 한글 문서 유지
5. **Security Mindset / 보안 마인드셋**: Security considerations at every step / 모든 단계에서 보안 고려사항

**Next Action / 다음 조치**:

Begin with Phase 1, Task 1.1: Project Structure Setup

1단계, 작업 1.1부터 시작: 프로젝트 구조 설정

---

**Document Version / 문서 버전**: 1.0.0
**Last Updated / 최종 업데이트**: 2025-10-10
**Status / 상태**: Ready for Implementation / 구현 준비 완료
