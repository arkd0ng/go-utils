# Redis Package - Developer Guide / 개발자 가이드

**Package**: `github.com/arkd0ng/go-utils/database/redis`
**Version**: v1.4.014
**Author**: arkd0ng
**Last Updated**: 2025-10-14

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

The Redis package follows the principle: **"If it's not dramatically simpler, don't build it"**

Redis 패키지는 다음 원칙을 따릅니다: **"극적으로 간단하지 않으면 만들지 마세요"**

**Core Principles / 핵심 원칙**:
1. **Extreme Simplicity**: Reduce 20+ lines of boilerplate to 2 lines of code
2. **Auto Everything**: Connection, retry, reconnect, cleanup all handled automatically
3. **Type Safety**: Generic methods for compile-time type checking
4. **Context Support**: All methods support context for cancellation and timeout
5. **Transparent Abstraction**: Simple API that doesn't hide Redis functionality

1. **극도의 간결함**: 20줄 이상의 보일러플레이트를 2줄의 코드로 축소
2. **모든 것이 자동**: 연결, 재시도, 재연결, 정리가 모두 자동으로 처리됨
3. **타입 안전성**: 컴파일 타임 타입 검사를 위한 제네릭 메서드
4. **Context 지원**: 모든 메서드가 취소 및 타임아웃을 위한 context 지원
5. **투명한 추상화**: Redis 기능을 숨기지 않는 간단한 API

### High-Level Architecture / 상위 수준 아키텍처

```
┌─────────────────────────────────────────────────────────┐
│                    User Application                      │
└───────────────────┬─────────────────────────────────────┘
                    │
        ┌───────────┴───────────┬───────────────┐
        │                       │               │
┌───────▼────────┐   ┌─────────▼────────┐   ┌─▼──────────┐
│  Simple API    │   │  Pipeline/Txn    │   │  Pub/Sub   │
│ (Set, Get,     │   │   (Batch ops)    │   │ (Messaging)│
│  HSet, LPush)  │   │                  │   │            │
└───────┬────────┘   └─────────┬────────┘   └─┬──────────┘
        │                       │               │
        └───────────┬───────────┴───────────────┘
                    │
        ┌───────────▼────────────┐
        │   Client (core logic)  │
        │   - Connection Pool    │
        │   - Health Check       │
        │   - Generic Methods    │
        └───────────┬────────────┘
                    │
        ┌───────────▼────────────┐
        │   Auto Features Layer  │
        │   - Retry Logic        │
        │   - Error Classification│
        │   - Context Handling   │
        └───────────┬────────────┘
                    │
        ┌───────────▼────────────┐
        │    go-redis/v9         │
        │  (redis/go-redis)      │
        └───────────┬────────────┘
                    │
        ┌───────────▼────────────┐
        │    Redis Server        │
        └────────────────────────┘
```

### Component Interaction / 컴포넌트 상호작용

```
User Code
    │
    ▼
Set(ctx, "key", "value")
    │
    ▼
[Check if closed] ──No──▶ Marshal value (if needed)
    │                        │
   Yes                       ▼
    │                   client.rdb.Set(ctx, key, value, 0)
    │                        │
    ▼                        ▼
[Return error]          [executeWithRetry]
                             │
                     ┌───────┴───────┐
                     │               │
                Success          Network Error?
                     │               │
                     │              Yes ──▶ [Exponential Backoff]
                     │               │              │
                     │               └──────────────┘
                     │                      │
                     │                     Retry (up to max)
                     │                      │
                     ▼                      ▼
                [Return nil]        [Success or Fatal Error]
                     │
                     ▼
                User Code
```

---

## Package Structure / 패키지 구조

### File Organization / 파일 구성

```
database/redis/
├── client.go           # Main client, connection management (55 lines)
├── config.go           # Configuration structures (76 lines)
├── options.go          # Functional options (105 lines)
├── connection.go       # Health check, connection monitoring (46 lines)
├── errors.go           # Error types and classification (79 lines)
├── types.go            # Common types (Client, Tx, Pipeliner) (47 lines)
├── retry.go            # Auto retry with exponential backoff (61 lines)
├── string.go           # String operations (210 lines)
├── hash.go             # Hash operations (183 lines)
├── list.go             # List operations (128 lines)
├── set.go              # Set operations (161 lines)
├── zset.go             # Sorted Set operations (174 lines)
├── key.go              # Key operations (148 lines)
├── pipeline.go         # Pipeline support (31 lines)
├── transaction.go      # Transaction support (57 lines)
├── pubsub.go           # Pub/Sub support (59 lines)
└── client_test.go      # Unit tests (268 lines)

Total: ~1,888 lines of production code
```

### File Responsibilities / 파일별 책임

| File / 파일 | Lines / 줄 수 | Responsibility / 책임 |
|-------------|--------------|----------------------|
| `client.go` | 55 | Client struct, New(), Close(), getCurrentConnection() |
| `config.go` | 76 | Config struct, default config, WithXXX options |
| `options.go` | 105 | Functional options pattern (11 options) |
| `connection.go` | 46 | Health check goroutine, connection monitoring |
| `errors.go` | 79 | Error types, classification, isRetryableError() |
| `types.go` | 47 | Client, Tx, Pipeliner types, generic helpers |
| `retry.go` | 61 | executeWithRetry(), exponential backoff logic |
| `string.go` | 210 | String operations (9 methods) |
| `hash.go` | 183 | Hash operations (8 methods + generic) |
| `list.go` | 128 | List operations (9 methods) |
| `set.go` | 161 | Set operations (8 methods) |
| `zset.go` | 174 | Sorted Set operations (8 methods) |
| `key.go` | 148 | Key operations (8 methods) |
| `pipeline.go` | 31 | Pipeline() method |
| `transaction.go` | 57 | Transaction struct, Watch(), TxPipeline(), Exec() |
| `pubsub.go` | 59 | Pub/Sub methods (Publish, Subscribe, PSubscribe) |
| `client_test.go` | 268 | Unit tests (8 test functions) |

---

## Core Components / 핵심 컴포넌트

### 1. Client Structure / 클라이언트 구조

**File**: `client.go`

```go
type Client struct {
    rdb    *redis.Client    // Underlying go-redis client / 기본 go-redis 클라이언트
    config *Config          // Configuration / 설정
    done   chan struct{}    // Shutdown signal / 종료 신호
    closed bool             // Close flag / 종료 플래그
    mu     sync.RWMutex     // Protects closed / closed 보호
}
```

**Key Methods / 주요 메서드**:
- `New(opts ...Option) (*Client, error)`: Create new client / 새 클라이언트 생성
- `Close()`: Cleanup resources / 리소스 정리
- `Ping(ctx)`: Health check / 헬스 체크
- `executeWithRetry(ctx, fn)`: Retry wrapper / 재시도 래퍼

### 2. Configuration / 설정

**File**: `config.go`

```go
type Config struct {
    Addr              string        // Redis server address / Redis 서버 주소
    Password          string        // Password / 비밀번호
    DB                int           // Database number / 데이터베이스 번호
    PoolSize          int           // Connection pool size / 연결 풀 크기
    MinIdleConns      int           // Minimum idle connections / 최소 유휴 연결
    DialTimeout       time.Duration // Connection timeout / 연결 타임아웃
    ReadTimeout       time.Duration // Read timeout / 읽기 타임아웃
    WriteTimeout      time.Duration // Write timeout / 쓰기 타임아웃
    MaxRetries        int           // Maximum retries / 최대 재시도
    RetryInterval     time.Duration // Base retry interval / 기본 재시도 간격
    EnableHealthCheck bool          // Enable background health check / 백그라운드 헬스 체크 활성화
}
```

**Default Configuration / 기본 설정**:
```go
func defaultConfig() *Config {
    return &Config{
        Addr:              "localhost:6379",
        Password:          "",
        DB:                0,
        PoolSize:          20,
        MinIdleConns:      5,
        DialTimeout:       5 * time.Second,
        ReadTimeout:       3 * time.Second,
        WriteTimeout:      3 * time.Second,
        MaxRetries:        3,
        RetryInterval:     100 * time.Millisecond,
        EnableHealthCheck: true,
    }
}
```

### 3. Options Pattern / 옵션 패턴

**File**: `options.go`

```go
type Option func(*Config)

func WithAddr(addr string) Option {
    return func(c *Config) {
        c.Addr = addr
    }
}

// 11 total options available / 총 11개 옵션 사용 가능
```

**Available Options / 사용 가능한 옵션**:
1. `WithAddr(addr)` - Redis server address / Redis 서버 주소
2. `WithPassword(password)` - Authentication / 인증
3. `WithDB(db)` - Database selection / 데이터베이스 선택
4. `WithPoolSize(size)` - Connection pool size / 연결 풀 크기
5. `WithMinIdleConns(n)` - Minimum idle connections / 최소 유휴 연결
6. `WithDialTimeout(timeout)` - Connection timeout / 연결 타임아웃
7. `WithReadTimeout(timeout)` - Read timeout / 읽기 타임아웃
8. `WithWriteTimeout(timeout)` - Write timeout / 쓰기 타임아웃
9. `WithMaxRetries(n)` - Maximum retries / 최대 재시도
10. `WithRetryInterval(interval)` - Retry interval / 재시도 간격
11. `WithHealthCheck(enable)` - Health check / 헬스 체크

### 4. Error Handling / 에러 처리

**File**: `errors.go`

```go
var (
    ErrInvalidAddr     = errors.New("redis: invalid address")
    ErrClientClosed    = errors.New("redis: client is closed")
    ErrKeyNotFound     = redis.Nil
    ErrNilValue        = errors.New("redis: nil value")
    ErrInvalidType     = errors.New("redis: invalid type")
    ErrMarshalFailed   = errors.New("redis: marshal failed")
    ErrUnmarshalFailed = errors.New("redis: unmarshal failed")
)

func isRetryableError(err error) bool {
    // Network errors, timeout errors are retryable
    // 네트워크 에러, 타임아웃 에러는 재시도 가능
}
```

**Error Classification / 에러 분류**:
- **Retryable Errors / 재시도 가능 에러**: Network errors, timeouts, connection errors
- **Fatal Errors / 치명적 에러**: Invalid address, client closed, authentication failure
- **Not Found Errors / 찾을 수 없음 에러**: `redis.Nil` (key doesn't exist)

### 5. Retry Logic / 재시도 로직

**File**: `retry.go`

```go
func (c *Client) executeWithRetry(ctx context.Context, fn func() error) error {
    var lastErr error

    for attempt := 0; attempt <= c.config.MaxRetries; attempt++ {
        // Check context / context 확인
        if err := ctx.Err(); err != nil {
            return err
        }

        // Execute function / 함수 실행
        lastErr = fn()
        if lastErr == nil {
            return nil // Success / 성공
        }

        // Check if retryable / 재시도 가능 여부 확인
        if !isRetryableError(lastErr) {
            return lastErr // Fatal error / 치명적 에러
        }

        // Exponential backoff / 지수 백오프
        backoff := time.Duration(1<<uint(attempt)) * c.config.RetryInterval

        select {
        case <-time.After(backoff):
            // Continue to next attempt / 다음 시도로 계속
        case <-ctx.Done():
            return ctx.Err()
        }
    }

    return lastErr
}
```

**Backoff Strategy / 백오프 전략**:
- Attempt 0: 100ms (base interval)
- Attempt 1: 200ms (2^1 * 100ms)
- Attempt 2: 400ms (2^2 * 100ms)
- Attempt 3: 800ms (2^3 * 100ms)

### 6. Health Check / 헬스 체크

**File**: `connection.go`

```go
func (c *Client) healthCheck() {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
            if err := c.rdb.Ping(ctx).Err(); err != nil {
                // Log error or trigger reconnection
                // 에러 로그 또는 재연결 트리거
            }
            cancel()
        case <-c.done:
            return // Shutdown / 종료
        }
    }
}
```

**Health Check Features / 헬스 체크 기능**:
- Runs in background goroutine / 백그라운드 고루틴에서 실행
- Checks connection every 10 seconds / 10초마다 연결 확인
- Can be disabled with `WithHealthCheck(false)` / `WithHealthCheck(false)`로 비활성화 가능
- Graceful shutdown via `done` channel / `done` 채널을 통한 우아한 종료

---

## Internal Implementation / 내부 구현

### 1. String Operations Flow / String 작업 흐름

**File**: `string.go`

```
Set(ctx, key, value, ttl)
    │
    ▼
[Marshal value to string if needed]
    │
    ▼
executeWithRetry(ctx, func() {
    return c.rdb.Set(ctx, key, value, ttl).Err()
})
    │
    ▼
[Return error or nil]
```

**Key Implementation / 주요 구현**:
```go
func (c *Client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
    // Check if closed / 종료 확인
    c.mu.RLock()
    if c.closed {
        c.mu.RUnlock()
        return ErrClientClosed
    }
    c.mu.RUnlock()

    // Execute with retry / 재시도와 함께 실행
    return c.executeWithRetry(ctx, func() error {
        return c.rdb.Set(ctx, key, value, ttl).Err()
    })
}
```

### 2. Generic Type-Safe Methods / 제네릭 타입 안전 메서드

**File**: `string.go`, `hash.go`

```go
// GetAs retrieves a value and unmarshals it to type T
// GetAs는 값을 조회하고 T 타입으로 역직렬화합니다
func GetAs[T any](c *Client, ctx context.Context, key string) (T, error) {
    var result T

    // Get string value / 문자열 값 가져오기
    val, err := c.Get(ctx, key)
    if err != nil {
        return result, err
    }

    // Try direct type assertion first / 먼저 직접 타입 단언 시도
    if v, ok := any(val).(T); ok {
        return v, nil
    }

    // Try unmarshal from JSON / JSON에서 역직렬화 시도
    if err := json.Unmarshal([]byte(val), &result); err != nil {
        return result, fmt.Errorf("%w: %v", ErrUnmarshalFailed, err)
    }

    return result, nil
}
```

**Type Safety Benefits / 타입 안전성 장점**:
- Compile-time type checking / 컴파일 타임 타입 검사
- No need for type assertions / 타입 단언 불필요
- Automatic JSON marshaling/unmarshaling / 자동 JSON 직렬화/역직렬화
- Reduces boilerplate code / 보일러플레이트 코드 감소

### 3. Pipeline Operations / 파이프라인 작업

**File**: `pipeline.go`

```go
func (c *Client) Pipeline(ctx context.Context, fn func(pipe Pipeliner) error) error {
    // Check if closed / 종료 확인
    c.mu.RLock()
    if c.closed {
        c.mu.RUnlock()
        return ErrClientClosed
    }
    c.mu.RUnlock()

    // Execute with retry / 재시도와 함께 실행
    return c.executeWithRetry(ctx, func() error {
        pipe := c.rdb.Pipeline()
        if err := fn(pipe); err != nil {
            return err
        }
        _, err := pipe.Exec(ctx)
        return err
    })
}
```

**Pipeline Benefits / 파이프라인 장점**:
- Batch multiple commands in one network round-trip / 여러 명령을 하나의 네트워크 왕복으로 배치
- Reduces network latency / 네트워크 지연 감소
- Atomic execution / 원자적 실행
- Type alias `Pipeliner = redis.Pipeliner` hides go-redis dependency / 타입 별칭으로 go-redis 의존성 숨김

### 4. Transaction Operations / 트랜잭션 작업

**File**: `transaction.go`

```go
type Tx struct {
    tx *redis.Tx
}

func (c *Client) Watch(ctx context.Context, keys []string, fn func(*Tx) error) error {
    // Check if closed / 종료 확인
    c.mu.RLock()
    if c.closed {
        c.mu.RUnlock()
        return ErrClientClosed
    }
    c.mu.RUnlock()

    // Execute with retry / 재시도와 함께 실행
    return c.executeWithRetry(ctx, func() error {
        return c.rdb.Watch(ctx, func(tx *redis.Tx) error {
            return fn(&Tx{tx: tx})
        }, keys...)
    })
}

func (tx *Tx) Exec(ctx context.Context, fn func(pipe Pipeliner) error) error {
    pipe := tx.tx.TxPipeline()
    if err := fn(pipe); err != nil {
        return err
    }
    _, err := pipe.Exec(ctx)
    return err
}
```

**Transaction Pattern / 트랜잭션 패턴**:
1. **Watch keys** for optimistic locking / 낙관적 잠금을 위해 키 감시
2. **Execute function** with transaction pipeline / 트랜잭션 파이프라인으로 함수 실행
3. **Exec commits** all commands atomically / Exec이 모든 명령을 원자적으로 커밋
4. **Retry** if watch key was modified / 감시 키가 수정된 경우 재시도

### 5. Pub/Sub Operations / Pub/Sub 작업

**File**: `pubsub.go`

```go
func (c *Client) Publish(ctx context.Context, channel string, message interface{}) error {
    // Check if closed / 종료 확인
    c.mu.RLock()
    if c.closed {
        c.mu.RUnlock()
        return ErrClientClosed
    }
    c.mu.RUnlock()

    // Marshal message if needed / 필요시 메시지 직렬화
    var msg string
    switch v := message.(type) {
    case string:
        msg = v
    case []byte:
        msg = string(v)
    default:
        data, err := json.Marshal(message)
        if err != nil {
            return fmt.Errorf("%w: %v", ErrMarshalFailed, err)
        }
        msg = string(data)
    }

    // Execute with retry / 재시도와 함께 실행
    return c.executeWithRetry(ctx, func() error {
        return c.rdb.Publish(ctx, channel, msg).Err()
    })
}
```

**Pub/Sub Pattern / Pub/Sub 패턴**:
- **Publisher**: Publishes messages to channels / 채널에 메시지 발행
- **Subscriber**: Subscribes to channels and receives messages / 채널 구독 및 메시지 수신
- **Pattern Subscribe**: Subscribe to channels matching patterns / 패턴에 일치하는 채널 구독

---

## Design Patterns / 디자인 패턴

### 1. Functional Options Pattern / 함수형 옵션 패턴

**Why / 이유**:
- Makes API flexible and extensible / API를 유연하고 확장 가능하게 만듦
- Avoids long parameter lists / 긴 매개변수 목록 방지
- Provides sensible defaults / 합리적인 기본값 제공
- Allows optional parameters / 선택적 매개변수 허용

**Implementation / 구현**:
```go
// Option is a function that modifies Config
// Option은 Config를 수정하는 함수입니다
type Option func(*Config)

// New creates a client with options
// New는 옵션과 함께 클라이언트를 생성합니다
func New(opts ...Option) (*Client, error) {
    cfg := defaultConfig()
    for _, opt := range opts {
        opt(cfg)
    }
    // ... create client
}
```

**Usage / 사용법**:
```go
client, err := redis.New(
    redis.WithAddr("localhost:6379"),
    redis.WithDB(1),
    redis.WithPoolSize(50),
)
```

### 2. Type Alias Pattern / 타입 별칭 패턴

**Why / 이유**:
- Hides underlying library dependency / 기본 라이브러리 의존성 숨김
- Single import point for users / 사용자를 위한 단일 import 지점
- Non-breaking change (maintains compatibility) / 중단되지 않는 변경 (호환성 유지)

**Implementation / 구현**:
```go
// Pipeliner is a type alias for redis.Pipeliner
// Pipeliner는 redis.Pipeliner의 타입 별칭입니다
type Pipeliner = redis.Pipeliner
```

**Benefits / 장점**:
- Users only import `github.com/arkd0ng/go-utils/database/redis`
- No need to import `github.com/redis/go-redis/v9`
- Cleaner API and dependency management / 더 깨끗한 API 및 의존성 관리

### 3. Retry with Exponential Backoff / 지수 백오프를 사용한 재시도

**Why / 이유**:
- Handles transient network errors gracefully / 일시적인 네트워크 에러를 우아하게 처리
- Prevents overwhelming the server / 서버 과부하 방지
- Increases success rate of operations / 작업 성공률 증가

**Implementation / 구현**:
```go
func (c *Client) executeWithRetry(ctx context.Context, fn func() error) error {
    var lastErr error

    for attempt := 0; attempt <= c.config.MaxRetries; attempt++ {
        lastErr = fn()
        if lastErr == nil {
            return nil
        }

        if !isRetryableError(lastErr) {
            return lastErr
        }

        // Exponential backoff: 100ms, 200ms, 400ms, 800ms
        backoff := time.Duration(1<<uint(attempt)) * c.config.RetryInterval

        select {
        case <-time.After(backoff):
        case <-ctx.Done():
            return ctx.Err()
        }
    }

    return lastErr
}
```

### 4. Health Check Pattern / 헬스 체크 패턴

**Why / 이유**:
- Proactively detects connection issues / 연결 문제를 사전에 감지
- Enables automatic reconnection / 자동 재연결 활성화
- Improves reliability / 안정성 향상

**Implementation / 구현**:
```go
// Start background health check goroutine
// 백그라운드 헬스 체크 고루틴 시작
if cfg.EnableHealthCheck {
    go client.healthCheck()
}

func (c *Client) healthCheck() {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            // Ping Redis server / Redis 서버 ping
        case <-c.done:
            return
        }
    }
}
```

### 5. Generic Methods Pattern / 제네릭 메서드 패턴

**Why / 이유**:
- Type safety at compile time / 컴파일 타임 타입 안전성
- Reduces type assertion boilerplate / 타입 단언 보일러플레이트 감소
- Automatic marshaling/unmarshaling / 자동 직렬화/역직렬화

**Implementation / 구현**:
```go
func GetAs[T any](c *Client, ctx context.Context, key string) (T, error) {
    var result T

    val, err := c.Get(ctx, key)
    if err != nil {
        return result, err
    }

    // Try direct type assertion / 직접 타입 단언 시도
    if v, ok := any(val).(T); ok {
        return v, nil
    }

    // Try JSON unmarshal / JSON 역직렬화 시도
    if err := json.Unmarshal([]byte(val), &result); err != nil {
        return result, fmt.Errorf("%w: %v", ErrUnmarshalFailed, err)
    }

    return result, nil
}
```

**Usage / 사용법**:
```go
// Type-safe retrieval / 타입 안전 조회
count, err := redis.GetAs[int](rdb, ctx, "count")
user, err := redis.GetAs[User](rdb, ctx, "user:123")
```

---

## Adding New Features / 새 기능 추가

### Step-by-Step Guide / 단계별 가이드

#### 1. Adding New String Operation / 새로운 String 작업 추가

**File**: `string.go`

**Example: Add GetRange method / 예시: GetRange 메서드 추가**

```go
// GetRange returns substring of string stored at key
// GetRange는 키에 저장된 문자열의 부분 문자열을 반환합니다
func (c *Client) GetRange(ctx context.Context, key string, start, end int64) (string, error) {
    // 1. Check if closed / 종료 확인
    c.mu.RLock()
    if c.closed {
        c.mu.RUnlock()
        return "", ErrClientClosed
    }
    c.mu.RUnlock()

    // 2. Execute with retry / 재시도와 함께 실행
    var result string
    err := c.executeWithRetry(ctx, func() error {
        var err error
        result, err = c.rdb.GetRange(ctx, key, start, end).Result()
        return err
    })

    return result, err
}
```

**Testing / 테스트**:
```go
func TestGetRange(t *testing.T) {
    // Setup / 설정
    client, err := New(WithAddr("localhost:6379"))
    if err != nil {
        t.Fatal(err)
    }
    defer client.Close()

    ctx := context.Background()

    // Set value / 값 설정
    err = client.Set(ctx, "mykey", "Hello World", 0)
    if err != nil {
        t.Fatal(err)
    }

    // Test GetRange / GetRange 테스트
    result, err := client.GetRange(ctx, "mykey", 0, 4)
    if err != nil {
        t.Fatal(err)
    }

    if result != "Hello" {
        t.Errorf("Expected 'Hello', got '%s'", result)
    }
}
```

#### 2. Adding New Hash Operation / 새로운 Hash 작업 추가

**File**: `hash.go`

**Example: Add HKeys method / 예시: HKeys 메서드 추가**

```go
// HKeys returns all field names in hash stored at key
// HKeys는 키에 저장된 해시의 모든 필드 이름을 반환합니다
func (c *Client) HKeys(ctx context.Context, key string) ([]string, error) {
    // 1. Check if closed / 종료 확인
    c.mu.RLock()
    if c.closed {
        c.mu.RUnlock()
        return nil, ErrClientClosed
    }
    c.mu.RUnlock()

    // 2. Execute with retry / 재시도와 함께 실행
    var result []string
    err := c.executeWithRetry(ctx, func() error {
        var err error
        result, err = c.rdb.HKeys(ctx, key).Result()
        return err
    })

    return result, err
}
```

#### 3. Adding New Configuration Option / 새로운 설정 옵션 추가

**File**: `config.go`, `options.go`

**Step 1: Add field to Config / Config에 필드 추가**:
```go
type Config struct {
    // ... existing fields
    MaxConnAge time.Duration // New: maximum connection age / 새로운: 최대 연결 수명
}

func defaultConfig() *Config {
    return &Config{
        // ... existing defaults
        MaxConnAge: 30 * time.Minute, // Default value / 기본값
    }
}
```

**Step 2: Add option function / 옵션 함수 추가**:
```go
// WithMaxConnAge sets the maximum connection age
// WithMaxConnAge는 최대 연결 수명을 설정합니다
func WithMaxConnAge(age time.Duration) Option {
    return func(c *Config) {
        c.MaxConnAge = age
    }
}
```

**Step 3: Use in client creation / 클라이언트 생성에 사용**:
```go
rdb := redis.NewClient(&redis.Options{
    // ... existing options
    MaxConnAge: cfg.MaxConnAge,
})
```

#### 4. Adding New Error Type / 새로운 에러 타입 추가

**File**: `errors.go`

```go
var (
    // ... existing errors
    ErrKeyExpired = errors.New("redis: key has expired")
)

func isRetryableError(err error) bool {
    if err == nil {
        return false
    }

    // ... existing checks

    // New check / 새로운 확인
    if errors.Is(err, ErrKeyExpired) {
        return false // Not retryable / 재시도 불가
    }

    return false
}
```

### Development Checklist / 개발 체크리스트

When adding a new feature, ensure you: / 새 기능을 추가할 때 다음을 확인하세요:

- [ ] Add bilingual documentation (English/Korean) / 이중 언어 문서 추가 (영문/한글)
- [ ] Implement closed state check / 종료 상태 확인 구현
- [ ] Use `executeWithRetry` for network operations / 네트워크 작업에 `executeWithRetry` 사용
- [ ] Add comprehensive error handling / 포괄적인 에러 처리 추가
- [ ] Support context for cancellation / 취소를 위한 context 지원
- [ ] Write unit tests with edge cases / 엣지 케이스와 함께 유닛 테스트 작성
- [ ] Add examples in `examples/redis/main.go` / `examples/redis/main.go`에 예제 추가
- [ ] Update `database/redis/README.md` / `database/redis/README.md` 업데이트
- [ ] Update `docs/database/redis/USER_MANUAL.md` / `docs/database/redis/USER_MANUAL.md` 업데이트
- [ ] Update CHANGELOG / CHANGELOG 업데이트

---

## Testing Guide / 테스트 가이드

### Running Tests / 테스트 실행

```bash
# Run all Redis package tests / 모든 Redis 패키지 테스트 실행
go test ./database/redis -v

# Run specific test / 특정 테스트 실행
go test ./database/redis -v -run TestStringOperations

# Run with coverage / 커버리지와 함께 실행
go test ./database/redis -cover
go test ./database/redis -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run with race detector / 경쟁 감지기와 함께 실행
go test ./database/redis -race
```

### Test Structure / 테스트 구조

**File**: `client_test.go`

```go
func TestStringOperations(t *testing.T) {
    // 1. Setup / 설정
    client, err := New(WithAddr("localhost:6379"))
    if err != nil {
        t.Skipf("Redis not available: %v", err)
    }
    defer client.Close()

    ctx := context.Background()

    // 2. Test cases / 테스트 케이스
    t.Run("Set and Get", func(t *testing.T) {
        // Test implementation / 테스트 구현
    })

    t.Run("Get non-existent key", func(t *testing.T) {
        // Test implementation / 테스트 구현
    })
}
```

### Test Categories / 테스트 카테고리

**Current Tests / 현재 테스트**:
1. `TestStringOperations` - String operations (Set, Get, MSet, MGet, Incr)
2. `TestHashOperations` - Hash operations (HSet, HGet, HGetAll, HGetAllAs[T])
3. `TestListOperations` - List operations (LPush, RPush, LPop, RPop, LRange)
4. `TestSetOperations` - Set operations (SAdd, SMembers, SUnion, SInter, SDiff)
5. `TestSortedSetOperations` - Sorted Set operations (ZAdd, ZRange, ZRangeByScore)
6. `TestKeyOperations` - Key operations (Del, Exists, Expire, TTL, Keys)
7. `TestPipelineOperations` - Pipeline operations
8. `TestTransactionOperations` - Transaction operations

### Test Best Practices / 테스트 모범 사례

1. **Setup and Teardown / 설정 및 정리**:
```go
func TestExample(t *testing.T) {
    // Setup / 설정
    client, err := New(WithAddr("localhost:6379"))
    if err != nil {
        t.Skipf("Redis not available: %v", err)
    }
    defer client.Close() // Teardown / 정리

    // Clean up test data / 테스트 데이터 정리
    ctx := context.Background()
    client.Del(ctx, "test:key")
}
```

2. **Use Subtests / 서브테스트 사용**:
```go
func TestStringOperations(t *testing.T) {
    client, _ := setupClient(t)
    defer client.Close()

    t.Run("Set", func(t *testing.T) {
        // Test Set / Set 테스트
    })

    t.Run("Get", func(t *testing.T) {
        // Test Get / Get 테스트
    })
}
```

3. **Test Edge Cases / 엣지 케이스 테스트**:
```go
func TestEdgeCases(t *testing.T) {
    client, _ := setupClient(t)
    defer client.Close()

    ctx := context.Background()

    // Non-existent key / 존재하지 않는 키
    _, err := client.Get(ctx, "non-existent")
    if !errors.Is(err, redis.ErrKeyNotFound) {
        t.Error("Expected ErrKeyNotFound")
    }

    // Expired key / 만료된 키
    client.Set(ctx, "expired", "value", 1*time.Millisecond)
    time.Sleep(2 * time.Millisecond)
    _, err = client.Get(ctx, "expired")
    if !errors.Is(err, redis.ErrKeyNotFound) {
        t.Error("Expected ErrKeyNotFound for expired key")
    }
}
```

4. **Test Context Cancellation / Context 취소 테스트**:
```go
func TestContextCancellation(t *testing.T) {
    client, _ := setupClient(t)
    defer client.Close()

    ctx, cancel := context.WithCancel(context.Background())
    cancel() // Cancel immediately / 즉시 취소

    err := client.Set(ctx, "key", "value", 0)
    if !errors.Is(err, context.Canceled) {
        t.Error("Expected context.Canceled error")
    }
}
```

### Integration Testing / 통합 테스트

**Docker Setup for Testing / 테스트를 위한 Docker 설정**:

```bash
# Start Redis with Docker Compose / Docker Compose로 Redis 시작
cd /Users/shlee/go-utils
docker compose up -d

# Run tests / 테스트 실행
go test ./database/redis -v

# Stop Redis / Redis 중지
docker compose down
```

**Alternative: Use scripts / 대안: 스크립트 사용**:

```bash
# Start Redis / Redis 시작
./.docker/scripts/docker-redis-start.sh

# Run tests / 테스트 실행
go test ./database/redis -v

# Stop Redis / Redis 중지
./.docker/scripts/docker-redis-stop.sh
```

---

## Performance / 성능

### Benchmarking / 벤치마킹

**Running Benchmarks / 벤치마크 실행**:

```bash
# Run all benchmarks / 모든 벤치마크 실행
go test ./database/redis -bench=.

# Run specific benchmark / 특정 벤치마크 실행
go test ./database/redis -bench=BenchmarkSet

# With memory allocation stats / 메모리 할당 통계와 함께
go test ./database/redis -bench=. -benchmem
```

**Example Benchmarks / 벤치마크 예시**:

```go
func BenchmarkSet(b *testing.B) {
    client, _ := New(WithAddr("localhost:6379"))
    defer client.Close()

    ctx := context.Background()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        client.Set(ctx, fmt.Sprintf("key%d", i), "value", 0)
    }
}

func BenchmarkGet(b *testing.B) {
    client, _ := New(WithAddr("localhost:6379"))
    defer client.Close()

    ctx := context.Background()
    client.Set(ctx, "benchmark-key", "value", 0)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        client.Get(ctx, "benchmark-key")
    }
}

func BenchmarkPipeline(b *testing.B) {
    client, _ := New(WithAddr("localhost:6379"))
    defer client.Close()

    ctx := context.Background()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        client.Pipeline(ctx, func(pipe redis.Pipeliner) error {
            for j := 0; j < 10; j++ {
                pipe.Set(ctx, fmt.Sprintf("key%d", j), "value", 0)
            }
            return nil
        })
    }
}
```

### Performance Optimization Tips / 성능 최적화 팁

#### 1. Use Connection Pooling / 연결 풀링 사용

```go
// Good: Configure appropriate pool size / 좋음: 적절한 풀 크기 설정
client, _ := redis.New(
    redis.WithPoolSize(50),      // 50 connections / 50개 연결
    redis.WithMinIdleConns(10),  // 10 idle connections / 10개 유휴 연결
)
```

**Recommended Pool Sizes / 권장 풀 크기**:
- **Low traffic / 낮은 트래픽**: PoolSize=10, MinIdleConns=2
- **Medium traffic / 중간 트래픽**: PoolSize=20, MinIdleConns=5
- **High traffic / 높은 트래픽**: PoolSize=50, MinIdleConns=10

#### 2. Use Pipeline for Batch Operations / 배치 작업에 파이프라인 사용

```go
// Bad: Individual commands (5 round-trips) / 나쁨: 개별 명령 (5번 왕복)
for i := 0; i < 5; i++ {
    client.Set(ctx, fmt.Sprintf("key%d", i), "value", 0)
}

// Good: Pipeline (1 round-trip) / 좋음: 파이프라인 (1번 왕복)
client.Pipeline(ctx, func(pipe redis.Pipeliner) error {
    for i := 0; i < 5; i++ {
        pipe.Set(ctx, fmt.Sprintf("key%d", i), "value", 0)
    }
    return nil
})
```

**Performance Improvement / 성능 향상**:
- 5x fewer network round-trips / 5배 적은 네트워크 왕복
- ~5x faster for batch operations / 배치 작업이 약 5배 빠름

#### 3. Use MSet/MGet for Multiple Keys / 여러 키에 MSet/MGet 사용

```go
// Bad: Multiple Get calls / 나쁨: 여러 Get 호출
val1, _ := client.Get(ctx, "key1")
val2, _ := client.Get(ctx, "key2")
val3, _ := client.Get(ctx, "key3")

// Good: Single MGet call / 좋음: 단일 MGet 호출
values, _ := client.MGet(ctx, "key1", "key2", "key3")
```

#### 4. Use Appropriate Timeouts / 적절한 타임아웃 사용

```go
client, _ := redis.New(
    redis.WithDialTimeout(5*time.Second),   // Connection timeout / 연결 타임아웃
    redis.WithReadTimeout(3*time.Second),   // Read timeout / 읽기 타임아웃
    redis.WithWriteTimeout(3*time.Second),  // Write timeout / 쓰기 타임아웃
)
```

**Timeout Guidelines / 타임아웃 가이드라인**:
- **Fast operations / 빠른 작업**: 1-3 seconds / 1-3초
- **Normal operations / 일반 작업**: 3-5 seconds / 3-5초
- **Long operations / 긴 작업**: 5-10 seconds / 5-10초

#### 5. Minimize Retry Attempts / 재시도 횟수 최소화

```go
// For latency-sensitive operations / 지연에 민감한 작업
client, _ := redis.New(
    redis.WithMaxRetries(1),                // Only 1 retry / 1번만 재시도
    redis.WithRetryInterval(50*time.Millisecond),
)

// For reliability-critical operations / 안정성이 중요한 작업
client, _ := redis.New(
    redis.WithMaxRetries(5),                // Up to 5 retries / 최대 5번 재시도
    redis.WithRetryInterval(100*time.Millisecond),
)
```

### Performance Monitoring / 성능 모니터링

**Monitor Pool Statistics / 풀 통계 모니터링**:

```go
stats := client.PoolStats()
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Timeouts: %d\n", stats.Timeouts)
fmt.Printf("TotalConns: %d\n", stats.TotalConns)
fmt.Printf("IdleConns: %d\n", stats.IdleConns)
```

**Expected Performance / 예상 성능**:
- **Single operation / 단일 작업**: ~0.5-2ms (localhost)
- **Pipeline (10 ops) / 파이프라인 (10 작업)**: ~1-3ms (localhost)
- **Batch operations / 배치 작업**: 5-10x faster than individual / 개별보다 5-10배 빠름

---

## Contributing Guidelines / 기여 가이드라인

### Code Contribution Process / 코드 기여 프로세스

1. **Fork the repository / 저장소 포크**
2. **Create a feature branch / 기능 브랜치 생성**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make changes / 변경사항 작성**
   - Follow code style guidelines / 코드 스타일 가이드라인 준수
   - Add bilingual documentation / 이중 언어 문서 추가
   - Write comprehensive tests / 포괄적인 테스트 작성

4. **Run tests / 테스트 실행**
   ```bash
   go test ./database/redis -v
   go test ./database/redis -race
   ```

5. **Update documentation / 문서 업데이트**
   - Update README.md
   - Update USER_MANUAL.md
   - Update CHANGELOG.md

6. **Commit changes / 변경사항 커밋**
   ```bash
   git commit -m "Feat: Add GetRange method for string operations"
   ```

7. **Push to fork / 포크에 푸시**
   ```bash
   git push origin feature/your-feature-name
   ```

8. **Create Pull Request / Pull Request 생성**

### Commit Message Format / 커밋 메시지 형식

```
Type: Brief description

Detailed explanation (optional)
```

**Types / 타입**:
- `Feat`: New feature / 새로운 기능
- `Fix`: Bug fix / 버그 수정
- `Docs`: Documentation changes / 문서 변경
- `Refactor`: Code refactoring / 코드 리팩토링
- `Test`: Test additions/changes / 테스트 추가/변경
- `Chore`: Build, config changes / 빌드, 설정 변경

**Examples / 예시**:
```
Feat: Add GetRange method for substring operations

- Implements GETRANGE command
- Adds comprehensive tests
- Updates documentation

Fixes #123
```

### Quality Checklist / 품질 체크리스트

Before submitting a Pull Request, ensure: / Pull Request 제출 전 확인:

- [ ] Code follows project style guidelines / 코드가 프로젝트 스타일 가이드라인을 준수함
- [ ] All tests pass / 모든 테스트 통과
- [ ] No race conditions (`go test -race`) / 경쟁 조건 없음
- [ ] Code coverage maintained or improved / 코드 커버리지 유지 또는 개선
- [ ] Documentation is bilingual (English/Korean) / 문서는 이중 언어 (영문/한글)
- [ ] CHANGELOG.md updated / CHANGELOG.md 업데이트됨
- [ ] Examples added/updated if needed / 필요시 예제 추가/업데이트
- [ ] No breaking changes (or clearly documented) / 중단 변경 없음 (또는 명확히 문서화됨)
- [ ] Error handling is comprehensive / 에러 처리가 포괄적임
- [ ] Context support for cancellation / 취소를 위한 Context 지원

---

## Code Style / 코드 스타일

### Naming Conventions / 명명 규칙

**1. Packages / 패키지**:
- Use lowercase, single word / 소문자, 단일 단어 사용
- Example: `redis`, `mysql`, `logging`

**2. Files / 파일**:
- Use lowercase with underscores / 소문자와 언더스코어 사용
- Group related functionality / 관련 기능 그룹화
- Examples: `string.go`, `hash.go`, `client_test.go`

**3. Types and Structs / 타입 및 구조체**:
- Use PascalCase / PascalCase 사용
- Examples: `Client`, `Config`, `Option`

**4. Functions and Methods / 함수 및 메서드**:
- Use PascalCase for exported / 내보내기용 PascalCase 사용
- Use camelCase for unexported / 내보내지 않는 용 camelCase 사용
- Examples: `New()`, `Set()`, `executeWithRetry()`

**5. Variables / 변수**:
- Use camelCase / camelCase 사용
- Short names for short scopes / 짧은 범위에는 짧은 이름
- Examples: `ctx`, `err`, `client`, `result`

**6. Constants / 상수**:
- Use PascalCase or ALL_CAPS / PascalCase 또는 ALL_CAPS 사용
- Examples: `ErrClientClosed`, `DefaultPoolSize`

### Documentation Style / 문서화 스타일

**1. Package Documentation / 패키지 문서**:
```go
// Package redis provides a simple, auto-everything Redis client.
// Package redis는 간단하고 모든 것이 자동인 Redis 클라이언트를 제공합니다.
package redis
```

**2. Function Documentation / 함수 문서**:
```go
// Set stores a value in Redis with optional TTL.
// Set은 선택적 TTL과 함께 Redis에 값을 저장합니다.
//
// The value can be any type and will be automatically marshaled.
// 값은 모든 타입이 될 수 있으며 자동으로 직렬화됩니다.
//
// Example / 예시:
//
//     err := client.Set(ctx, "key", "value", 10*time.Minute)
func (c *Client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
```

**3. Type Documentation / 타입 문서**:
```go
// Client is a Redis client with auto-retry and connection pooling.
// Client는 자동 재시도 및 연결 풀링이 있는 Redis 클라이언트입니다.
type Client struct {
    rdb    *redis.Client  // Underlying go-redis client / 기본 go-redis 클라이언트
    config *Config        // Configuration / 설정
    done   chan struct{}  // Shutdown signal / 종료 신호
}
```

### Code Organization / 코드 구성

**1. File Structure / 파일 구조**:
```go
// 1. Package declaration / 패키지 선언
package redis

// 2. Imports (standard library first, then external) / Import (표준 라이브러리 먼저, 그 다음 외부)
import (
    "context"
    "fmt"

    "github.com/redis/go-redis/v9"
)

// 3. Constants / 상수
const (
    DefaultPoolSize = 20
)

// 4. Types / 타입
type Client struct {
    // ...
}

// 5. Constructor / 생성자
func New(opts ...Option) (*Client, error) {
    // ...
}

// 6. Methods / 메서드
func (c *Client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
    // ...
}

// 7. Helper functions / 헬퍼 함수
func marshalValue(value interface{}) (string, error) {
    // ...
}
```

**2. Error Handling / 에러 처리**:
```go
// Good: Check errors explicitly / 좋음: 에러를 명시적으로 확인
result, err := client.Get(ctx, "key")
if err != nil {
    if errors.Is(err, redis.ErrKeyNotFound) {
        // Handle not found / 찾을 수 없음 처리
    }
    return err
}

// Good: Wrap errors with context / 좋음: Context와 함께 에러 래핑
if err := client.Set(ctx, key, value, 0); err != nil {
    return fmt.Errorf("failed to set key %s: %w", key, err)
}
```

**3. Context Usage / Context 사용**:
```go
// Good: Accept context as first parameter / 좋음: context를 첫 번째 매개변수로 받음
func (c *Client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
    // Check context / context 확인
    if err := ctx.Err(); err != nil {
        return err
    }

    // Use context in operations / 작업에 context 사용
    return c.rdb.Set(ctx, key, value, ttl).Err()
}
```

**4. Comments / 주석**:
```go
// Good: Bilingual comments / 좋음: 이중 언어 주석
// Set stores a value in Redis
// Set은 Redis에 값을 저장합니다

// Good: Explain why, not what / 좋음: 무엇이 아닌 이유 설명
// Use exponential backoff to avoid overwhelming the server
// 서버 과부하를 피하기 위해 지수 백오프 사용

// Bad: Obvious comment / 나쁨: 명백한 주석
// i++ // Increment i / i 증가
```

### Testing Style / 테스트 스타일

**1. Test Function Names / 테스트 함수 이름**:
```go
// Good: Descriptive names / 좋음: 설명적인 이름
func TestStringOperations(t *testing.T)
func TestHashOperations(t *testing.T)
func TestContextCancellation(t *testing.T)
```

**2. Table-Driven Tests / 테이블 주도 테스트**:
```go
func TestSet(t *testing.T) {
    tests := []struct {
        name    string
        key     string
        value   interface{}
        ttl     time.Duration
        wantErr bool
    }{
        {"Valid string", "key1", "value", 0, false},
        {"Valid int", "key2", 42, 0, false},
        {"With TTL", "key3", "value", 10*time.Second, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation / 테스트 구현
        })
    }
}
```

**3. Test Assertions / 테스트 단언**:
```go
// Good: Clear error messages / 좋음: 명확한 에러 메시지
if result != expected {
    t.Errorf("Expected %v, got %v", expected, result)
}

// Good: Use errors.Is for error comparison / 좋음: 에러 비교에 errors.Is 사용
if !errors.Is(err, redis.ErrKeyNotFound) {
    t.Errorf("Expected ErrKeyNotFound, got %v", err)
}
```

---

## Appendix / 부록

### Glossary / 용어집

| Term / 용어 | Definition / 정의 |
|------------|------------------|
| **Pipeline** | Batch multiple commands in one network round-trip / 여러 명령을 하나의 네트워크 왕복으로 배치 |
| **Transaction** | Atomic execution of multiple commands / 여러 명령의 원자적 실행 |
| **TTL** | Time To Live - expiration time for keys / Time To Live - 키의 만료 시간 |
| **Pub/Sub** | Publish/Subscribe - messaging pattern / 발행/구독 - 메시징 패턴 |
| **Pipeliner** | Interface for pipeline operations / 파이프라인 작업을 위한 인터페이스 |
| **Context** | Carries deadlines, cancellation signals / 데드라인, 취소 신호 전달 |
| **Exponential Backoff** | Retry strategy with increasing delays / 지연 증가 재시도 전략 |
| **Health Check** | Background monitoring of connection status / 연결 상태의 백그라운드 모니터링 |
| **Generic** | Type parameter for compile-time type safety / 컴파일 타임 타입 안전성을 위한 타입 매개변수 |
| **Type Alias** | Alternative name for existing type / 기존 타입의 대체 이름 |

### Useful Links / 유용한 링크

- **Redis Documentation**: https://redis.io/docs/
- **go-redis Documentation**: https://redis.uptrace.dev/
- **Go Generics**: https://go.dev/doc/tutorial/generics
- **Functional Options**: https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

### Common Redis Commands Reference / 일반 Redis 명령 참조

| Command / 명령 | Package Method / 패키지 메서드 | Description / 설명 |
|---------------|-------------------------------|-------------------|
| SET | Set(key, value, ttl) | Set string value / 문자열 값 설정 |
| GET | Get(key) | Get string value / 문자열 값 가져오기 |
| HSET | HSet(key, field, value) | Set hash field / 해시 필드 설정 |
| HGET | HGet(key, field) | Get hash field / 해시 필드 가져오기 |
| LPUSH | LPush(key, values...) | Push to list head / 리스트 헤드에 푸시 |
| RPUSH | RPush(key, values...) | Push to list tail / 리스트 테일에 푸시 |
| SADD | SAdd(key, members...) | Add to set / 집합에 추가 |
| ZADD | ZAdd(key, members...) | Add to sorted set / 정렬 집합에 추가 |
| DEL | Del(key) | Delete key / 키 삭제 |
| EXPIRE | Expire(key, ttl) | Set expiration / 만료 설정 |
| KEYS | Keys(pattern) | Find keys by pattern / 패턴으로 키 찾기 |
| PUBLISH | Publish(channel, msg) | Publish message / 메시지 발행 |

---

**Document Version / 문서 버전**: v1.4.014
**Last Updated / 마지막 업데이트**: 2025-10-14
**Maintained by / 관리자**: arkd0ng
**License / 라이선스**: MIT
