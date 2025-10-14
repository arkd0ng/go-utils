# Database/Redis Package - Design Plan / 설계 계획서
# database/redis 패키지 - 설계 계획서

**Version / 버전**: v1.4.x
**Author / 작성자**: arkd0ng
**Created / 작성일**: 2025-10-14
**Status / 상태**: Final Design - Extreme Simplicity / 최종 설계 - 극도의 간결함

---

## Table of Contents / 목차

1. [Why This Package Exists / 왜 이 패키지가 존재하는가](#why-this-package-exists--왜-이-패키지가-존재하는가)
2. [Design Philosophy / 설계 철학](#design-philosophy--설계-철학)
3. [What Users Get / 사용자가 얻는 것](#what-users-get--사용자가-얻는-것)
4. [API Design / API 설계](#api-design--api-설계)
5. [Implementation Architecture / 구현 아키텍처](#implementation-architecture--구현-아키텍처)
6. [File Structure / 파일 구조](#file-structure--파일-구조)
7. [Detailed Features / 상세 기능](#detailed-features--상세-기능)

---

## Why This Package Exists / 왜 이 패키지가 존재하는가

### The Problem / 문제점

Using Redis client libraries (like `go-redis/redis`) requires developers to:

Redis 클라이언트 라이브러리(예: `go-redis/redis`)를 사용하려면 개발자가 다음을 해야 합니다:

1. **Manually manage connections / 수동으로 연결 관리**:
   ```go
   // 매번 연결 상태 확인
   if err := rdb.Ping(ctx).Err(); err != nil {
       // 재접속 로직 작성
       rdb = redis.NewClient(&redis.Options{...})
       // ...
   }
   ```

2. **Write verbose commands for simple operations / 간단한 작업에 장황한 명령어 작성**:
   ```go
   // String 저장
   err := rdb.Set(ctx, "user:123:name", "John", 0).Err()

   // Hash 저장
   err := rdb.HSet(ctx, "user:123", map[string]interface{}{
       "name":  "John",
       "email": "john@example.com",
       "age":   30,
   }).Err()

   // 매번 .Err() 체크, context 전달...
   ```

3. **Handle connection loss manually / 연결 손실 수동 처리**:
   ```go
   // "connection refused" 에러 발생 시
   // 개발자가 직접 재시도 로직 작성
   ```

4. **Complex pipeline and transaction usage / 복잡한 파이프라인 및 트랜잭션 사용**:
   ```go
   // Pipeline
   pipe := rdb.Pipeline()
   pipe.Set(ctx, "key1", "value1", 0)
   pipe.Set(ctx, "key2", "value2", 0)
   _, err := pipe.Exec(ctx)

   // Transaction
   err := rdb.Watch(ctx, func(tx *redis.Tx) error {
       // transaction logic...
   }, "key")
   ```

### The Solution / 해결책

**이 패키지는 위의 모든 번거로움을 제거합니다**:

```go
// 1. 한 번 연결하면 끝 - 자동으로 계속 유지됨
rdb, _ := redis.New(redis.WithAddr("localhost:6379"))

// 2. Simple API - 간단하게
rdb.Set(ctx, "user:123:name", "John")
rdb.Get(ctx, "user:123:name")

// 3. Hash 저장 - map으로 간단하게
rdb.HSetMap(ctx, "user:123", map[string]interface{}{
    "name":  "John",
    "email": "john@example.com",
    "age":   30,
})

// 4. 연결 끊김? 자동 재접속됨
// 5. 에러 처리? 재시도 가능한 에러는 자동 재시도됨
// 6. Pipeline/Transaction? 간단한 API로 제공됨
```

### If It's Not This Simple, Don't Build It / 이 정도로 간단하지 않으면 만들지 마세요

The guiding principle for this package is: **"If it's not dramatically simpler than the standard library, don't build it."**

이 패키지의 기본 원칙은: **"표준 라이브러리보다 극적으로 간단하지 않으면 만들지 마세요."**

---

## Design Philosophy / 설계 철학

### Core Principles / 핵심 원칙

1. **Extreme Simplicity / 극도의 간결함**
   - Reduce 20+ lines of code to 2 lines / 20줄 이상의 코드를 2줄로 축소
   - No verbose command chaining / 장황한 명령어 체이닝 제거
   - Auto-everything: connection, retry, resource cleanup / 모든 것 자동화

2. **Auto Connection Management / 자동 연결 관리**
   - Auto-connect on first use / 첫 사용 시 자동 연결
   - Auto-reconnect on connection loss / 연결 손실 시 자동 재연결
   - Health check in background / 백그라운드 헬스 체크
   - Connection pooling / 연결 풀링

3. **Auto Retry / 자동 재시도**
   - Network errors → auto retry with exponential backoff / 네트워크 에러 → 지수 백오프로 자동 재시도
   - Configurable retry policy / 설정 가능한 재시도 정책
   - Circuit breaker pattern / 서킷 브레이커 패턴

4. **Context Support / Context 지원**
   - All methods accept `context.Context` / 모든 메서드는 `context.Context` 수용
   - Cancellation and timeout support / 취소 및 타임아웃 지원
   - Non-context versions for convenience / 편의를 위한 non-context 버전

5. **Type-Safe / 타입 안전**
   - Strong typing for common data types / 일반적인 데이터 타입에 대한 강력한 타입 지정
   - Generic support for custom types / 사용자 정의 타입을 위한 제네릭 지원
   - JSON serialization for complex types / 복잡한 타입을 위한 JSON 직렬화

---

## What Users Get / 사용자가 얻는 것

### Before vs After / 전후 비교

**❌ Before (standard go-redis/redis):**

```go
import "github.com/redis/go-redis/v9"

// 연결 설정
rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})

// String 저장
ctx := context.Background()
err := rdb.Set(ctx, "key", "value", 0).Err()
if err != nil {
    return err
}

// Hash 저장
err = rdb.HSet(ctx, "user:123", map[string]interface{}{
    "name": "John",
    "age":  30,
}).Err()
if err != nil {
    return err
}

// Hash 전체 가져오기
result := rdb.HGetAll(ctx, "user:123")
if result.Err() != nil {
    return result.Err()
}
data := result.Val()

// List 추가
err = rdb.RPush(ctx, "queue", "item1", "item2").Err()
if err != nil {
    return err
}

// ... 매번 .Err() 체크, verbose한 API ...
```

**✅ After (this package):**

```go
import "github.com/arkd0ng/go-utils/database/redis"

// 연결 설정
rdb, _ := redis.New(redis.WithAddr("localhost:6379"))
defer rdb.Close()

// String 저장
rdb.Set(ctx, "key", "value")

// Hash 저장
rdb.HSetMap(ctx, "user:123", map[string]interface{}{
    "name": "John",
    "age":  30,
})

// Hash 전체 가져오기
data, _ := rdb.HGetAll(ctx, "user:123")

// List 추가
rdb.RPush(ctx, "queue", "item1", "item2")

// 간결하고 읽기 쉬운 API!
```

### Key Benefits / 주요 이점

1. **20 lines → 2 lines** / 20줄 → 2줄
2. **No `.Err()` chaining** / `.Err()` 체이닝 불필요
3. **Auto-retry on network errors** / 네트워크 에러 시 자동 재시도
4. **Auto-reconnect on connection loss** / 연결 손실 시 자동 재연결
5. **Type-safe API with generics** / 제네릭을 사용한 타입 안전 API
6. **Simple transaction support** / 간단한 트랜잭션 지원
7. **Built-in connection pooling** / 내장 연결 풀링
8. **Health check monitoring** / 헬스 체크 모니터링

---

## API Design / API 설계

### 1. Simple API (Core Methods / 핵심 메서드)

#### String Operations / 문자열 작업

```go
// Set a string value / 문자열 값 설정
Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error
SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) error

// Get a string value / 문자열 값 가져오기
Get(ctx context.Context, key string) (string, error)
GetContext(ctx context.Context, key string) (string, error)

// Get and parse to type / 타입으로 파싱하여 가져오기
GetAs[T any](ctx context.Context, key string) (T, error)

// Multiple get/set / 다중 get/set
MGet(ctx context.Context, keys ...string) ([]string, error)
MSet(ctx context.Context, pairs map[string]interface{}) error
```

#### Hash Operations / 해시 작업

```go
// Set hash field / 해시 필드 설정
HSet(ctx context.Context, key, field string, value interface{}) error

// Set multiple hash fields / 여러 해시 필드 설정
HSetMap(ctx context.Context, key string, fields map[string]interface{}) error

// Get hash field / 해시 필드 가져오기
HGet(ctx context.Context, key, field string) (string, error)

// Get all hash fields / 모든 해시 필드 가져오기
HGetAll(ctx context.Context, key string) (map[string]string, error)

// Get all and parse to struct / 구조체로 파싱하여 가져오기
HGetAllAs[T any](ctx context.Context, key string) (T, error)

// Delete hash fields / 해시 필드 삭제
HDel(ctx context.Context, key string, fields ...string) error

// Check if hash field exists / 해시 필드 존재 확인
HExists(ctx context.Context, key, field string) (bool, error)
```

#### List Operations / 리스트 작업

```go
// Push to list / 리스트에 추가
LPush(ctx context.Context, key string, values ...interface{}) error
RPush(ctx context.Context, key string, values ...interface{}) error

// Pop from list / 리스트에서 제거
LPop(ctx context.Context, key string) (string, error)
RPop(ctx context.Context, key string) (string, error)

// Get list range / 리스트 범위 가져오기
LRange(ctx context.Context, key string, start, stop int64) ([]string, error)

// Get list length / 리스트 길이 가져오기
LLen(ctx context.Context, key string) (int64, error)
```

#### Set Operations / 집합 작업

```go
// Add to set / 집합에 추가
SAdd(ctx context.Context, key string, members ...interface{}) error

// Remove from set / 집합에서 제거
SRem(ctx context.Context, key string, members ...interface{}) error

// Get all set members / 모든 집합 멤버 가져오기
SMembers(ctx context.Context, key string) ([]string, error)

// Check if member exists in set / 집합에 멤버 존재 확인
SIsMember(ctx context.Context, key string, member interface{}) (bool, error)

// Get set cardinality / 집합 크기 가져오기
SCard(ctx context.Context, key string) (int64, error)
```

#### Sorted Set Operations / 정렬 집합 작업

```go
// Add to sorted set / 정렬 집합에 추가
ZAdd(ctx context.Context, key string, score float64, member interface{}) error

// Add multiple to sorted set / 정렬 집합에 여러 개 추가
ZAddMultiple(ctx context.Context, key string, members map[string]float64) error

// Get range by score / 점수 범위로 가져오기
ZRangeByScore(ctx context.Context, key string, min, max float64) ([]string, error)

// Get range / 범위 가져오기
ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)

// Remove from sorted set / 정렬 집합에서 제거
ZRem(ctx context.Context, key string, members ...interface{}) error

// Get sorted set cardinality / 정렬 집합 크기 가져오기
ZCard(ctx context.Context, key string) (int64, error)
```

#### Key Operations / 키 작업

```go
// Delete keys / 키 삭제
Del(ctx context.Context, keys ...string) error

// Check if key exists / 키 존재 확인
Exists(ctx context.Context, keys ...string) (int64, error)

// Set expiration / 만료 시간 설정
Expire(ctx context.Context, key string, expiration time.Duration) error

// Get TTL / TTL 가져오기
TTL(ctx context.Context, key string) (time.Duration, error)

// Get keys by pattern / 패턴으로 키 가져오기
Keys(ctx context.Context, pattern string) ([]string, error)
```

### 2. Pipeline API (Batch Operations / 배치 작업)

```go
// Execute multiple commands in a pipeline / 파이프라인에서 여러 명령 실행
Pipeline(ctx context.Context, fn func(pipe Pipeliner) error) error

// Example usage / 사용 예제:
rdb.Pipeline(ctx, func(pipe redis.Pipeliner) error {
    pipe.Set(ctx, "key1", "value1", 0)
    pipe.Set(ctx, "key2", "value2", 0)
    pipe.Incr(ctx, "counter")
    return nil
})
```

### 3. Transaction API (WATCH/MULTI/EXEC)

```go
// Execute commands in a transaction / 트랜잭션에서 명령 실행
Transaction(ctx context.Context, fn func(tx *Tx) error, keys ...string) error

// Example usage / 사용 예제:
rdb.Transaction(ctx, func(tx *redis.Tx) error {
    val, _ := tx.Get(ctx, "counter")
    count, _ := strconv.Atoi(val)
    tx.Set(ctx, "counter", count+1, 0)
    return nil
}, "counter")
```

### 4. Pub/Sub API

```go
// Publish message / 메시지 발행
Publish(ctx context.Context, channel string, message interface{}) error

// Subscribe to channels / 채널 구독
Subscribe(ctx context.Context, channels ...string) (*PubSub, error)

// Pattern subscribe / 패턴 구독
PSubscribe(ctx context.Context, patterns ...string) (*PubSub, error)
```

---

## Implementation Architecture / 구현 아키텍처

### Core Components / 핵심 컴포넌트

```
┌─────────────────────────────────────────────────────────────┐
│                        User Code                            │
│                        사용자 코드                            │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                    Client (client.go)                       │
│  - New(options...)                                          │
│  - Simple API methods (Set, Get, HSet, LPush, etc.)       │
│  - Auto-retry wrapper                                       │
│  클라이언트 - 간단한 API 메서드 및 자동 재시도               │
└────────────────────────┬────────────────────────────────────┘
                         │
           ┌─────────────┼─────────────┐
           │             │             │
           ▼             ▼             ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│ Connection  │  │   Retry     │  │  Pipeline   │
│ Management  │  │   Logic     │  │ Transaction │
│ (conn.go)   │  │ (retry.go)  │  │ (pipeline.go│
│             │  │             │  │  tx.go)     │
│ 연결 관리    │  │ 재시도 로직  │  │ 파이프라인   │
│             │  │             │  │ 트랜잭션     │
└─────────────┘  └─────────────┘  └─────────────┘
           │             │             │
           └─────────────┼─────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│              go-redis/redis (Underlying Library)            │
│              go-redis/redis (기본 라이브러리)                 │
└─────────────────────────────────────────────────────────────┘
```

### Design Patterns / 디자인 패턴

1. **Options Pattern / 옵션 패턴**
   - Functional options for flexible configuration / 유연한 설정을 위한 함수형 옵션
   - `WithAddr()`, `WithPassword()`, `WithDB()`, etc.

2. **Auto-Retry with Circuit Breaker / 서킷 브레이커를 사용한 자동 재시도**
   - Exponential backoff for transient errors / 일시적 오류에 대한 지수 백오프
   - Circuit breaker to prevent cascading failures / 연쇄 실패 방지를 위한 서킷 브레이커

3. **Connection Pool / 연결 풀**
   - Built-in connection pooling from go-redis / go-redis의 내장 연결 풀링
   - Health check and auto-reconnect / 헬스 체크 및 자동 재연결

4. **Type Safety with Generics / 제네릭을 사용한 타입 안전성**
   - `GetAs[T]()`, `HGetAllAs[T]()` for type-safe retrieval / 타입 안전 검색

---

## File Structure / 파일 구조

```
database/redis/
├── client.go           # Core client and Simple API / 핵심 클라이언트 및 Simple API
├── options.go          # Functional options / 함수형 옵션
├── config.go           # Configuration struct / 설정 구조체
├── connection.go       # Connection management / 연결 관리
├── retry.go            # Retry logic / 재시도 로직
├── errors.go           # Error types / 에러 타입
├── string.go           # String operations / 문자열 작업
├── hash.go             # Hash operations / 해시 작업
├── list.go             # List operations / 리스트 작업
├── set.go              # Set operations / 집합 작업
├── zset.go             # Sorted set operations / 정렬 집합 작업
├── key.go              # Key operations / 키 작업
├── pipeline.go         # Pipeline support / 파이프라인 지원
├── transaction.go      # Transaction support / 트랜잭션 지원
├── pubsub.go           # Pub/Sub support / Pub/Sub 지원
├── types.go            # Common types / 공통 타입
├── client_test.go      # Client tests / 클라이언트 테스트
├── integration_test.go # Integration tests / 통합 테스트
└── README.md           # Package documentation / 패키지 문서
```

---

## Detailed Features / 상세 기능

### 1. Auto Connection Management / 자동 연결 관리

```go
type Client struct {
    rdb    *redis.Client
    config *Config
    mu     sync.RWMutex
}

// New creates a new Redis client with auto-connection
// New는 자동 연결이 있는 새 Redis 클라이언트를 생성합니다
func New(opts ...Option) (*Client, error) {
    cfg := defaultConfig()
    for _, opt := range opts {
        opt(cfg)
    }

    rdb := redis.NewClient(&redis.Options{
        Addr:         cfg.Addr,
        Password:     cfg.Password,
        DB:           cfg.DB,
        PoolSize:     cfg.PoolSize,
        MinIdleConns: cfg.MinIdleConns,
        MaxRetries:   cfg.MaxRetries,
    })

    // Health check / 헬스 체크
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := rdb.Ping(ctx).Err(); err != nil {
        return nil, fmt.Errorf("failed to connect to Redis: %w", err)
    }

    client := &Client{
        rdb:    rdb,
        config: cfg,
    }

    // Start health check goroutine / 헬스 체크 고루틴 시작
    if cfg.EnableHealthCheck {
        go client.healthCheck()
    }

    return client, nil
}
```

### 2. Auto Retry with Exponential Backoff / 지수 백오프를 사용한 자동 재시도

```go
// executeWithRetry executes a Redis command with retry logic
// executeWithRetry는 재시도 로직으로 Redis 명령을 실행합니다
func (c *Client) executeWithRetry(ctx context.Context, fn func() error) error {
    var lastErr error

    for i := 0; i < c.config.MaxRetries; i++ {
        if err := fn(); err != nil {
            if !isRetriableError(err) {
                return err
            }
            lastErr = err

            // Exponential backoff / 지수 백오프
            backoff := time.Duration(i+1) * c.config.RetryInterval
            select {
            case <-ctx.Done():
                return ctx.Err()
            case <-time.After(backoff):
                continue
            }
        }
        return nil
    }

    return fmt.Errorf("max retries exceeded: %w", lastErr)
}

// isRetriableError checks if an error is retriable
// isRetriableError는 에러가 재시도 가능한지 확인합니다
func isRetriableError(err error) bool {
    if err == nil {
        return false
    }

    // Network errors are retriable / 네트워크 에러는 재시도 가능
    if errors.Is(err, context.DeadlineExceeded) {
        return true
    }

    errStr := err.Error()
    retriableErrors := []string{
        "connection refused",
        "connection reset",
        "broken pipe",
        "i/o timeout",
    }

    for _, retryErr := range retriableErrors {
        if strings.Contains(errStr, retryErr) {
            return true
        }
    }

    return false
}
```

### 3. Type-Safe Generic Methods / 타입 안전 제네릭 메서드

```go
// GetAs gets a value and unmarshals it to the specified type
// GetAs는 값을 가져와 지정된 타입으로 언마샬합니다
func GetAs[T any](c *Client, ctx context.Context, key string) (T, error) {
    var result T

    val, err := c.Get(ctx, key)
    if err != nil {
        return result, err
    }

    if err := json.Unmarshal([]byte(val), &result); err != nil {
        return result, fmt.Errorf("failed to unmarshal value: %w", err)
    }

    return result, nil
}

// HGetAllAs gets all hash fields and unmarshals to a struct
// HGetAllAs는 모든 해시 필드를 가져와 구조체로 언마샬합니다
func HGetAllAs[T any](c *Client, ctx context.Context, key string) (T, error) {
    var result T

    fields, err := c.HGetAll(ctx, key)
    if err != nil {
        return result, err
    }

    // Convert map to JSON and unmarshal to struct
    // map을 JSON으로 변환하고 구조체로 언마샬
    data, err := json.Marshal(fields)
    if err != nil {
        return result, err
    }

    if err := json.Unmarshal(data, &result); err != nil {
        return result, err
    }

    return result, nil
}
```

### 4. Simple Transaction Support / 간단한 트랜잭션 지원

```go
// Transaction executes a function in a Redis transaction
// Transaction은 Redis 트랜잭션에서 함수를 실행합니다
func (c *Client) Transaction(ctx context.Context, fn func(tx *Tx) error, keys ...string) error {
    return c.executeWithRetry(ctx, func() error {
        return c.rdb.Watch(ctx, func(tx *redis.Tx) error {
            txClient := &Tx{
                tx:     tx,
                client: c,
            }
            return fn(txClient)
        }, keys...)
    })
}

// Tx wraps redis.Tx for simpler API
// Tx는 더 간단한 API를 위해 redis.Tx를 래핑합니다
type Tx struct {
    tx     *redis.Tx
    client *Client
}

// All Simple API methods available on Tx
// Tx에서 사용 가능한 모든 Simple API 메서드
```

---

## Summary / 요약

This Redis package provides:

이 Redis 패키지는 다음을 제공합니다:

1. ✅ **Extreme simplicity**: 20+ lines → 2 lines / 극도의 간결함: 20줄 이상 → 2줄
2. ✅ **Auto-everything**: connection, retry, reconnect / 모든 것 자동화: 연결, 재시도, 재연결
3. ✅ **Type-safe API** with generics / 제네릭을 사용한 타입 안전 API
4. ✅ **Simple transaction** support / 간단한 트랜잭션 지원
5. ✅ **Pipeline support** for batch operations / 배치 작업을 위한 파이프라인 지원
6. ✅ **Pub/Sub support** / Pub/Sub 지원
7. ✅ **Health check** and monitoring / 헬스 체크 및 모니터링
8. ✅ **Context support** everywhere / 모든 곳에서 Context 지원

**Development Principle / 개발 원칙**: If it's not dramatically simpler, don't build it!

**개발 원칙**: 극적으로 간단하지 않으면 만들지 마세요!
