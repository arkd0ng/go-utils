# Database/Redis Package - Work Plan / 작업 계획서
# database/redis 패키지 - 작업 계획서

**Version / 버전**: v1.4.x
**Author / 작성자**: arkd0ng
**Created / 작성일**: 2025-10-14
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

This work plan outlines the detailed implementation steps for the `database/redis` package. Each phase is broken down into specific tasks with clear acceptance criteria.

이 작업 계획은 `database/redis` 패키지의 상세한 구현 단계를 설명합니다. 각 단계는 명확한 수용 기준과 함께 구체적인 작업으로 나뉩니다.

### Project Timeline / 프로젝트 타임라인

- **Phase 1**: Foundation / 기초 (2-3 작업 단위)
- **Phase 2**: Core Features / 핵심 기능 (5-7 작업 단위)
- **Phase 3**: Advanced Features / 고급 기능 (3-4 작업 단위)
- **Phase 4**: Testing & Documentation / 테스팅 및 문서화 (3-4 작업 단위)
- **Phase 5**: Release / 릴리스 (1-2 작업 단위)

**Total Estimated Work Units / 총 예상 작업 단위**: 14-20 units

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
   mkdir -p database/redis
   mkdir -p examples/redis
   mkdir -p scripts
   ```

2. Create initial package files / 초기 패키지 파일 생성:
   - `database/redis/client.go` - Core client and Simple API
   - `database/redis/connection.go` - Connection management
   - `database/redis/retry.go` - Retry logic
   - `database/redis/string.go` - String operations
   - `database/redis/hash.go` - Hash operations
   - `database/redis/list.go` - List operations
   - `database/redis/set.go` - Set operations
   - `database/redis/zset.go` - Sorted set operations
   - `database/redis/key.go` - Key operations
   - `database/redis/pipeline.go` - Pipeline support
   - `database/redis/transaction.go` - Transaction support
   - `database/redis/pubsub.go` - Pub/Sub support
   - `database/redis/config.go` - Configuration struct
   - `database/redis/options.go` - Functional options
   - `database/redis/errors.go` - Error types
   - `database/redis/types.go` - Common types
   - `database/redis/client_test.go` - Client tests
   - `database/redis/integration_test.go` - Integration tests

3. Add package documentation / 패키지 문서 추가:
   - `database/redis/README.md`

4. Initialize go.mod dependencies / go.mod 의존성 초기화:
   ```bash
   go get github.com/redis/go-redis/v9@latest
   ```

**Acceptance Criteria / 수용 기준**:
- [ ] All directories created / 모든 디렉토리 생성됨
- [ ] All package files exist with package declaration / 모든 패키지 파일에 패키지 선언이 있음
- [ ] Dependencies added to go.mod / 의존성이 go.mod에 추가됨
- [ ] `go build ./database/redis` succeeds / 빌드 성공

**Estimated Effort / 예상 소요 시간**: 0.5 work unit

---

### Task 1.2: Docker Redis Setup / Docker Redis 설정

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Create Docker Compose configuration and setup scripts for Redis testing.

Redis 테스팅을 위한 Docker Compose 설정 및 설정 스크립트를 생성합니다.

**Subtasks / 하위 작업**:

1. Update `docker-compose.yml` to include Redis service / Redis 서비스를 포함하도록 docker-compose.yml 업데이트:
   ```yaml
   services:
     redis:
       image: redis:7-alpine
       container_name: go-utils-redis
       ports:
         - "6379:6379"
       command: redis-server --appendonly yes
       volumes:
         - redis-data:/data

   volumes:
     redis-data:
   ```

2. Create Redis setup scripts / Redis 설정 스크립트 생성:
   - `scripts/docker-redis-start.sh` - Start Docker Redis
   - `scripts/docker-redis-stop.sh` - Stop and cleanup Docker Redis
   - `scripts/docker-redis-logs.sh` - View Redis logs
   - `scripts/docker-redis-cli.sh` - Connect to Redis CLI

3. Create `cfg/database-redis.yaml` configuration file / database-redis.yaml 설정 파일 생성:
   ```yaml
   redis:
     addr: localhost:6379
     password: ""
     db: 0
     pool_size: 10
     min_idle_conns: 5
   ```

**Acceptance Criteria / 수용 기준**:
- [ ] Docker Compose configuration updated / Docker Compose 설정 업데이트됨
- [ ] All setup scripts created and executable / 모든 설정 스크립트 생성되고 실행 가능
- [ ] Redis starts successfully with `./.docker/scripts/docker-redis-start.sh` / 스크립트로 Redis 성공적으로 시작
- [ ] Redis CLI connection works / Redis CLI 연결 작동

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 1.3: Error Types Definition / 에러 타입 정의

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Define all custom error types and error handling utilities.

모든 커스텀 에러 타입과 에러 처리 유틸리티를 정의합니다.

**Implementation / 구현**:

File: `database/redis/errors.go`

```go
package redis

import (
    "errors"
    "fmt"
    "time"
)

// Predefined errors / 사전 정의된 에러
var (
    ErrConnectionFailed = errors.New("redis connection failed")
    ErrCommandFailed    = errors.New("redis command failed")
    ErrTimeout          = errors.New("operation timeout")
    ErrClosed           = errors.New("redis connection closed")
    ErrInvalidAddr      = errors.New("invalid redis address")
    ErrNil              = errors.New("redis: nil")
)

// RedisError represents a Redis operation error
// RedisError는 Redis 작업 에러를 나타냅니다
type RedisError struct {
    Op       string        // Operation name / 작업 이름
    Key      string        // Redis key / Redis 키
    Args     []interface{} // Command arguments / 명령어 인자
    Err      error         // Original error / 원본 에러
    Time     time.Time     // Error timestamp / 에러 타임스탬프
    Duration time.Duration // Operation duration / 작업 소요 시간
}

func (e *RedisError) Error() string {
    return fmt.Sprintf("redis %s failed for key '%s': %v (took %v)",
        e.Op, e.Key, e.Err, e.Duration)
}

func (e *RedisError) Unwrap() error {
    return e.Err
}

// isRetriableError checks if an error is retriable
// isRetriableError는 에러가 재시도 가능한지 확인합니다
func isRetriableError(err error) bool {
    if err == nil {
        return false
    }

    // Network errors are retriable / 네트워크 에러는 재시도 가능
    retriableErrors := []string{
        "connection refused",
        "connection reset",
        "broken pipe",
        "i/o timeout",
        "EOF",
    }

    errStr := err.Error()
    for _, retryErr := range retriableErrors {
        if strings.Contains(errStr, retryErr) {
            return true
        }
    }

    return false
}
```

**Acceptance Criteria / 수용 기준**:
- [ ] All error types defined / 모든 에러 타입 정의됨
- [ ] Error wrapping and unwrapping implemented / 에러 래핑 및 언래핑 구현됨
- [ ] Retriable error detection implemented / 재시도 가능 에러 감지 구현됨
- [ ] Builds without errors / 에러 없이 빌드됨

**Estimated Effort / 예상 소요 시간**: 0.5 work unit

---

## Phase 2: Core Features / 2단계: 핵심 기능

### Task 2.1: Core Client Implementation / 핵심 클라이언트 구현

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement the core Redis client with connection management and options pattern.

연결 관리 및 옵션 패턴을 갖춘 핵심 Redis 클라이언트를 구현합니다.

**Files / 파일**:
- `database/redis/client.go`
- `database/redis/config.go`
- `database/redis/options.go`
- `database/redis/connection.go`

**Key Features / 주요 기능**:
- Options pattern for configuration / 설정을 위한 옵션 패턴
- Connection pooling / 연결 풀링
- Auto-connect on first use / 첫 사용 시 자동 연결
- Health check in background / 백그라운드 헬스 체크
- Thread-safe operations / 스레드 안전 작업

**Acceptance Criteria / 수용 기준**:
- [ ] Client struct with redis.Client wrapper / redis.Client 래퍼를 가진 Client 구조체
- [ ] `New(opts ...Option)` constructor / 생성자
- [ ] Options: `WithAddr()`, `WithPassword()`, `WithDB()`, `WithPoolSize()` / 옵션
- [ ] `Close()` method / Close 메서드
- [ ] `Ping()` method for health check / 헬스 체크를 위한 Ping 메서드
- [ ] Background health check goroutine / 백그라운드 헬스 체크 고루틴
- [ ] Basic unit tests / 기본 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

### Task 2.2: Retry Logic Implementation / 재시도 로직 구현

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement auto-retry logic with exponential backoff for transient errors.

일시적 에러에 대한 지수 백오프를 사용한 자동 재시도 로직을 구현합니다.

**File / 파일**: `database/redis/retry.go`

**Key Features / 주요 기능**:
- Exponential backoff / 지수 백오프
- Configurable max retries / 설정 가능한 최대 재시도 횟수
- Context cancellation support / Context 취소 지원
- Retriable error detection / 재시도 가능 에러 감지

**Acceptance Criteria / 수용 기준**:
- [ ] `executeWithRetry()` helper function / 헬퍼 함수
- [ ] Exponential backoff implemented / 지수 백오프 구현됨
- [ ] Context timeout respected / Context 타임아웃 준수
- [ ] Unit tests for retry logic / 재시도 로직 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 2.3: String Operations / 문자열 작업

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement all string operations (Set, Get, MGet, MSet, etc.).

모든 문자열 작업을 구현합니다 (Set, Get, MGet, MSet 등).

**File / 파일**: `database/redis/string.go`

**Methods / 메서드**:
```go
// Set sets a string value / 문자열 값 설정
Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error

// Get gets a string value / 문자열 값 가져오기
Get(ctx context.Context, key string) (string, error)

// GetAs gets and parses to type / 타입으로 파싱하여 가져오기
GetAs[T any](ctx context.Context, key string) (T, error)

// MGet gets multiple values / 여러 값 가져오기
MGet(ctx context.Context, keys ...string) ([]string, error)

// MSet sets multiple values / 여러 값 설정
MSet(ctx context.Context, pairs map[string]interface{}) error

// Incr increments a counter / 카운터 증가
Incr(ctx context.Context, key string) (int64, error)

// Decr decrements a counter / 카운터 감소
Decr(ctx context.Context, key string) (int64, error)
```

**Acceptance Criteria / 수용 기준**:
- [ ] All string methods implemented / 모든 문자열 메서드 구현됨
- [ ] Auto-retry on network errors / 네트워크 에러 시 자동 재시도
- [ ] Context support / Context 지원
- [ ] JSON serialization for complex types / 복잡한 타입을 위한 JSON 직렬화
- [ ] Unit tests for all methods / 모든 메서드에 대한 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

### Task 2.4: Hash Operations / 해시 작업

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement all hash operations (HSet, HGet, HGetAll, etc.).

모든 해시 작업을 구현합니다 (HSet, HGet, HGetAll 등).

**File / 파일**: `database/redis/hash.go`

**Methods / 메서드**:
```go
// HSet sets a hash field / 해시 필드 설정
HSet(ctx context.Context, key, field string, value interface{}) error

// HSetMap sets multiple hash fields / 여러 해시 필드 설정
HSetMap(ctx context.Context, key string, fields map[string]interface{}) error

// HGet gets a hash field / 해시 필드 가져오기
HGet(ctx context.Context, key, field string) (string, error)

// HGetAll gets all hash fields / 모든 해시 필드 가져오기
HGetAll(ctx context.Context, key string) (map[string]string, error)

// HGetAllAs gets all and parses to struct / 구조체로 파싱하여 가져오기
HGetAllAs[T any](ctx context.Context, key string) (T, error)

// HDel deletes hash fields / 해시 필드 삭제
HDel(ctx context.Context, key string, fields ...string) error

// HExists checks if hash field exists / 해시 필드 존재 확인
HExists(ctx context.Context, key, field string) (bool, error)

// HLen gets hash length / 해시 길이 가져오기
HLen(ctx context.Context, key string) (int64, error)
```

**Acceptance Criteria / 수용 기준**:
- [ ] All hash methods implemented / 모든 해시 메서드 구현됨
- [ ] Struct marshaling/unmarshaling support / 구조체 마샬링/언마샬링 지원
- [ ] Unit tests for all methods / 모든 메서드에 대한 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

### Task 2.5: List Operations / 리스트 작업

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement all list operations (LPush, RPush, LPop, RPop, LRange, etc.).

모든 리스트 작업을 구현합니다.

**File / 파일**: `database/redis/list.go`

**Methods / 메서드**:
```go
LPush(ctx context.Context, key string, values ...interface{}) error
RPush(ctx context.Context, key string, values ...interface{}) error
LPop(ctx context.Context, key string) (string, error)
RPop(ctx context.Context, key string) (string, error)
LRange(ctx context.Context, key string, start, stop int64) ([]string, error)
LLen(ctx context.Context, key string) (int64, error)
LIndex(ctx context.Context, key string, index int64) (string, error)
LSet(ctx context.Context, key string, index int64, value interface{}) error
```

**Acceptance Criteria / 수용 기준**:
- [ ] All list methods implemented / 모든 리스트 메서드 구현됨
- [ ] Unit tests for all methods / 모든 메서드에 대한 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 2.6: Set Operations / 집합 작업

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement all set operations (SAdd, SRem, SMembers, etc.).

모든 집합 작업을 구현합니다.

**File / 파일**: `database/redis/set.go`

**Methods / 메서드**:
```go
SAdd(ctx context.Context, key string, members ...interface{}) error
SRem(ctx context.Context, key string, members ...interface{}) error
SMembers(ctx context.Context, key string) ([]string, error)
SIsMember(ctx context.Context, key string, member interface{}) (bool, error)
SCard(ctx context.Context, key string) (int64, error)
SUnion(ctx context.Context, keys ...string) ([]string, error)
SInter(ctx context.Context, keys ...string) ([]string, error)
SDiff(ctx context.Context, keys ...string) ([]string, error)
```

**Acceptance Criteria / 수용 기준**:
- [ ] All set methods implemented / 모든 집합 메서드 구현됨
- [ ] Unit tests for all methods / 모든 메서드에 대한 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 2.7: Sorted Set Operations / 정렬 집합 작업

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Implement sorted set operations (ZAdd, ZRange, ZRangeByScore, etc.).

정렬 집합 작업을 구현합니다.

**File / 파일**: `database/redis/zset.go`

**Methods / 메서드**:
```go
ZAdd(ctx context.Context, key string, score float64, member interface{}) error
ZAddMultiple(ctx context.Context, key string, members map[string]float64) error
ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)
ZRangeByScore(ctx context.Context, key string, min, max float64) ([]string, error)
ZRem(ctx context.Context, key string, members ...interface{}) error
ZCard(ctx context.Context, key string) (int64, error)
ZScore(ctx context.Context, key string, member string) (float64, error)
```

**Acceptance Criteria / 수용 기준**:
- [ ] All sorted set methods implemented / 모든 정렬 집합 메서드 구현됨
- [ ] Unit tests for all methods / 모든 메서드에 대한 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 2.8: Key Operations / 키 작업

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Implement key operations (Del, Exists, Expire, TTL, Keys, etc.).

키 작업을 구현합니다.

**File / 파일**: `database/redis/key.go`

**Methods / 메서드**:
```go
Del(ctx context.Context, keys ...string) error
Exists(ctx context.Context, keys ...string) (int64, error)
Expire(ctx context.Context, key string, expiration time.Duration) error
TTL(ctx context.Context, key string) (time.Duration, error)
Keys(ctx context.Context, pattern string) ([]string, error)
Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error)
```

**Acceptance Criteria / 수용 기준**:
- [ ] All key methods implemented / 모든 키 메서드 구현됨
- [ ] Unit tests for all methods / 모든 메서드에 대한 단위 테스트

**Estimated Effort / 예상 소요 시간**: 0.5 work unit

---

## Phase 3: Advanced Features / 3단계: 고급 기능

### Task 3.1: Pipeline Support / 파이프라인 지원

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Implement pipeline support for batch operations.

배치 작업을 위한 파이프라인 지원을 구현합니다.

**File / 파일**: `database/redis/pipeline.go`

**Implementation / 구현**:
```go
// Pipeline executes multiple commands in a pipeline
// Pipeline은 파이프라인에서 여러 명령을 실행합니다
func (c *Client) Pipeline(ctx context.Context, fn func(pipe Pipeliner) error) error {
    return c.executeWithRetry(ctx, func() error {
        pipe := c.rdb.Pipeline()
        if err := fn(&pipeliner{pipe: pipe}); err != nil {
            return err
        }
        _, err := pipe.Exec(ctx)
        return err
    })
}

// Pipeliner wraps redis.Pipeliner
type Pipeliner interface {
    Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
    Get(ctx context.Context, key string) *redis.StringCmd
    // ... all other commands
}
```

**Acceptance Criteria / 수용 기준**:
- [ ] Pipeline method implemented / Pipeline 메서드 구현됨
- [ ] All Simple API methods available in pipeline / 모든 Simple API 메서드 파이프라인에서 사용 가능
- [ ] Unit tests / 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 3.2: Transaction Support / 트랜잭션 지원

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Implement transaction support with WATCH/MULTI/EXEC.

WATCH/MULTI/EXEC를 사용한 트랜잭션 지원을 구현합니다.

**File / 파일**: `database/redis/transaction.go`

**Implementation / 구현**:
```go
// Transaction executes commands in a Redis transaction
// Transaction은 Redis 트랜잭션에서 명령을 실행합니다
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
type Tx struct {
    tx     *redis.Tx
    client *Client
}
```

**Acceptance Criteria / 수용 기준**:
- [ ] Transaction method implemented / Transaction 메서드 구현됨
- [ ] WATCH support for optimistic locking / 낙관적 잠금을 위한 WATCH 지원
- [ ] Unit tests / 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 3.3: Pub/Sub Support / Pub/Sub 지원

**Priority / 우선순위**: 🟢 P2

**Description / 설명**:
Implement Pub/Sub support for message publishing and subscribing.

메시지 발행 및 구독을 위한 Pub/Sub 지원을 구현합니다.

**File / 파일**: `database/redis/pubsub.go`

**Methods / 메서드**:
```go
Publish(ctx context.Context, channel string, message interface{}) error
Subscribe(ctx context.Context, channels ...string) (*PubSub, error)
PSubscribe(ctx context.Context, patterns ...string) (*PubSub, error)
```

**Acceptance Criteria / 수용 기준**:
- [ ] Publish method implemented / Publish 메서드 구현됨
- [ ] Subscribe method implemented / Subscribe 메서드 구현됨
- [ ] Pattern subscribe implemented / 패턴 구독 구현됨
- [ ] Unit tests / 단위 테스트

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

## Phase 4: Testing & Documentation / 4단계: 테스팅 및 문서화

### Task 4.1: Comprehensive Unit Tests / 종합 단위 테스트

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Write comprehensive unit tests for all components.

모든 컴포넌트에 대한 종합 단위 테스트를 작성합니다.

**Files / 파일**:
- `database/redis/client_test.go`
- `database/redis/string_test.go`
- `database/redis/hash_test.go`
- `database/redis/list_test.go`
- `database/redis/set_test.go`
- `database/redis/zset_test.go`
- `database/redis/key_test.go`
- `database/redis/pipeline_test.go`
- `database/redis/transaction_test.go`

**Test Coverage Goals / 테스트 커버리지 목표**:
- Minimum 80% code coverage / 최소 80% 코드 커버리지
- All error paths tested / 모든 에러 경로 테스트됨
- Edge cases covered / 엣지 케이스 커버됨

**Acceptance Criteria / 수용 기준**:
- [ ] All unit tests pass / 모든 단위 테스트 통과
- [ ] Code coverage ≥ 80% / 코드 커버리지 ≥ 80%
- [ ] No test flakiness / 테스트 불안정성 없음

**Estimated Effort / 예상 소요 시간**: 2 work units

---

### Task 4.2: Integration Tests / 통합 테스트

**Priority / 우선순위**: 🟡 P1

**Description / 설명**:
Write integration tests against real Docker Redis instance.

실제 Docker Redis 인스턴스에 대한 통합 테스트를 작성합니다.

**File / 파일**: `database/redis/integration_test.go`

**Test Scenarios / 테스트 시나리오**:
- Connection and reconnection / 연결 및 재연결
- Auto-retry on network errors / 네트워크 에러 시 자동 재시도
- Pipeline operations / 파이프라인 작업
- Transaction operations / 트랜잭션 작업
- Pub/Sub messaging / Pub/Sub 메시징
- Concurrent operations / 동시 작업

**Acceptance Criteria / 수용 기준**:
- [ ] All integration tests pass / 모든 통합 테스트 통과
- [ ] Tests run against Docker Redis / Docker Redis에 대해 테스트 실행
- [ ] Concurrent operations tested / 동시 작업 테스트됨

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

### Task 4.3: Examples / 예제

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Create comprehensive examples demonstrating all features.

모든 기능을 시연하는 종합 예제를 생성합니다.

**Files / 파일**:
- `examples/redis/main.go` - Comprehensive example
- `examples/redis/README.md` - Example documentation

**Example Scenarios / 예제 시나리오**:
1. Basic string operations / 기본 문자열 작업
2. Hash operations with structs / 구조체를 사용한 해시 작업
3. List operations (queue) / 리스트 작업 (큐)
4. Set operations / 집합 작업
5. Sorted set operations (leaderboard) / 정렬 집합 작업 (리더보드)
6. Pipeline batch operations / 파이프라인 배치 작업
7. Transaction with optimistic locking / 낙관적 잠금을 사용한 트랜잭션
8. Pub/Sub messaging / Pub/Sub 메시징

**Acceptance Criteria / 수용 기준**:
- [ ] All examples run successfully / 모든 예제 성공적으로 실행
- [ ] Examples well-documented / 예제 잘 문서화됨
- [ ] Examples demonstrate best practices / 예제가 모범 사례 시연

**Estimated Effort / 예상 소요 시간**: 1.5 work units

---

### Task 4.4: Package Documentation / 패키지 문서화

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Write comprehensive package documentation.

종합 패키지 문서를 작성합니다.

**Files / 파일**:
- `database/redis/README.md` - Package overview and quick start
- `docs/database/redis/USER_MANUAL.md` - Complete user manual
- `docs/database/redis/DEVELOPER_GUIDE.md` - Developer guide

**README Sections / README 섹션**:
- Overview / 개요
- Installation / 설치
- Quick Start / 빠른 시작
- Features / 기능
- API Reference / API 참조
- Examples / 예제
- Configuration / 설정
- Error Handling / 에러 처리

**Acceptance Criteria / 수용 기준**:
- [ ] README.md complete with all sections / 모든 섹션이 포함된 README.md 완성
- [ ] All public APIs documented / 모든 공개 API 문서화됨
- [ ] Code examples in documentation / 문서에 코드 예제 포함
- [ ] Bilingual documentation (English/Korean) / 이중 언어 문서 (영문/한글)

**Estimated Effort / 예상 소요 시간**: 2 work units (README + USER_MANUAL)

---

## Phase 5: Release / 5단계: 릴리스

### Task 5.1: Final Testing & QA / 최종 테스팅 및 QA

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Perform final testing and quality assurance.

최종 테스팅 및 품질 보증을 수행합니다.

**Tasks / 작업**:
- [ ] Run all tests (`go test ./database/redis -v`) / 모든 테스트 실행
- [ ] Run integration tests / 통합 테스트 실행
- [ ] Run all examples / 모든 예제 실행
- [ ] Check code coverage / 코드 커버리지 확인
- [ ] Run linters (`golangci-lint run`) / 린터 실행
- [ ] Review documentation completeness / 문서 완성도 검토
- [ ] Test on multiple Go versions / 여러 Go 버전에서 테스트

**Acceptance Criteria / 수용 기준**:
- [ ] All tests pass / 모든 테스트 통과
- [ ] Code coverage ≥ 80% / 코드 커버리지 ≥ 80%
- [ ] No linter errors / 린터 에러 없음
- [ ] All examples run successfully / 모든 예제 성공적으로 실행

**Estimated Effort / 예상 소요 시간**: 1 work unit

---

### Task 5.2: Update Root Documentation / 루트 문서 업데이트

**Priority / 우선순위**: 🔴 P0

**Description / 설명**:
Update root-level documentation to include Redis package.

Redis 패키지를 포함하도록 루트 레벨 문서를 업데이트합니다.

**Files to Update / 업데이트할 파일**:
- `README.md` - Add Redis package to available packages
- `CHANGELOG.md` - Add v1.4.x summary
- `docs/CHANGELOG/CHANGELOG-v1.4.md` - Complete v1.4.x changelog
- `CLAUDE.md` - Add Redis package architecture

**Acceptance Criteria / 수용 기준**:
- [ ] README.md updated / README.md 업데이트됨
- [ ] CHANGELOG.md updated / CHANGELOG.md 업데이트됨
- [ ] CLAUDE.md updated / CLAUDE.md 업데이트됨

**Estimated Effort / 예상 소요 시간**: 0.5 work unit

---

## Task Dependencies / 작업 의존성

```
Phase 1: Foundation
├── Task 1.1: Project Structure Setup
├── Task 1.2: Docker Redis Setup
└── Task 1.3: Error Types Definition

Phase 2: Core Features
├── Task 2.1: Core Client Implementation (depends on 1.1, 1.3)
├── Task 2.2: Retry Logic Implementation (depends on 2.1)
├── Task 2.3: String Operations (depends on 2.1, 2.2)
├── Task 2.4: Hash Operations (depends on 2.1, 2.2)
├── Task 2.5: List Operations (depends on 2.1, 2.2)
├── Task 2.6: Set Operations (depends on 2.1, 2.2)
├── Task 2.7: Sorted Set Operations (depends on 2.1, 2.2)
└── Task 2.8: Key Operations (depends on 2.1, 2.2)

Phase 3: Advanced Features
├── Task 3.1: Pipeline Support (depends on 2.1-2.8)
├── Task 3.2: Transaction Support (depends on 2.1-2.8)
└── Task 3.3: Pub/Sub Support (depends on 2.1)

Phase 4: Testing & Documentation
├── Task 4.1: Comprehensive Unit Tests (depends on 2.1-3.3)
├── Task 4.2: Integration Tests (depends on 1.2, 2.1-3.3)
├── Task 4.3: Examples (depends on 2.1-3.3)
└── Task 4.4: Package Documentation (depends on 2.1-3.3)

Phase 5: Release
├── Task 5.1: Final Testing & QA (depends on 4.1-4.4)
└── Task 5.2: Update Root Documentation (depends on 5.1)
```

---

## Quality Checklist / 품질 체크리스트

### Code Quality / 코드 품질

- [ ] All code follows Go best practices / 모든 코드가 Go 모범 사례를 따름
- [ ] All public APIs have bilingual documentation / 모든 공개 API에 이중 언어 문서 있음
- [ ] Error messages are clear and actionable / 에러 메시지가 명확하고 실행 가능함
- [ ] No TODOs or FIXMEs in production code / 프로덕션 코드에 TODO나 FIXME 없음
- [ ] Code is DRY (Don't Repeat Yourself) / 코드가 DRY 원칙을 따름

### Testing / 테스팅

- [ ] Unit test coverage ≥ 80% / 단위 테스트 커버리지 ≥ 80%
- [ ] All critical paths tested / 모든 중요 경로 테스트됨
- [ ] Edge cases covered / 엣지 케이스 커버됨
- [ ] Integration tests pass / 통합 테스트 통과
- [ ] No test flakiness / 테스트 불안정성 없음

### Documentation / 문서화

- [ ] README.md complete / README.md 완성
- [ ] All public APIs documented / 모든 공개 API 문서화됨
- [ ] Examples comprehensive / 예제가 포괄적임
- [ ] Bilingual (English/Korean) / 이중 언어 (영문/한글)
- [ ] CHANGELOG updated / CHANGELOG 업데이트됨

### Performance / 성능

- [ ] No memory leaks / 메모리 누수 없음
- [ ] Connection pooling efficient / 연결 풀링이 효율적임
- [ ] Auto-retry doesn't cause cascading failures / 자동 재시도가 연쇄 실패를 일으키지 않음
- [ ] Benchmarks run successfully / 벤치마크 성공적으로 실행됨

### Docker / Docker

- [ ] Docker Compose configuration tested / Docker Compose 설정 테스트됨
- [ ] Setup scripts work on macOS and Linux / 설정 스크립트가 macOS와 Linux에서 작동
- [ ] Redis starts successfully / Redis 성공적으로 시작
- [ ] Cleanup scripts work properly / 정리 스크립트가 제대로 작동

---

## Summary / 요약

This work plan provides a comprehensive roadmap for implementing the `database/redis` package. The phased approach ensures:

이 작업 계획은 `database/redis` 패키지 구현을 위한 포괄적인 로드맵을 제공합니다. 단계별 접근 방식은 다음을 보장합니다:

1. ✅ Solid foundation before building features / 기능 구축 전 견고한 기초
2. ✅ Incremental development with clear milestones / 명확한 이정표가 있는 점진적 개발
3. ✅ Comprehensive testing at every stage / 모든 단계에서 종합 테스팅
4. ✅ High-quality documentation / 고품질 문서화
5. ✅ Production-ready release / 프로덕션 준비된 릴리스

**Next Steps / 다음 단계**: Begin Phase 1 - Task 1.1 (Project Structure Setup)

**다음 단계**: 1단계 시작 - Task 1.1 (프로젝트 구조 설정)
