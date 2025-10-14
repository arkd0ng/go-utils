# Redis Package - User Manual / Redis 패키지 - 사용자 매뉴얼

**Version / 버전**: v1.4.014
**Last Updated / 최종 업데이트**: 2025-10-14

---

## Table of Contents / 목차

1. [Introduction / 소개](#introduction--소개)
2. [Installation / 설치](#installation--설치)
3. [Quick Start / 빠른 시작](#quick-start--빠른-시작)
4. [Configuration Reference / 설정 참조](#configuration-reference--설정-참조)
5. [Core Operations / 핵심 작업](#core-operations--핵심-작업)
   - [String Operations / 문자열 작업](#string-operations--문자열-작업)
   - [Hash Operations / 해시 작업](#hash-operations--해시-작업)
   - [List Operations / 리스트 작업](#list-operations--리스트-작업)
   - [Set Operations / 집합 작업](#set-operations--집합-작업)
   - [Sorted Set Operations / 정렬 집합 작업](#sorted-set-operations--정렬-집합-작업)
   - [Key Operations / 키 작업](#key-operations--키-작업)
6. [Advanced Features / 고급 기능](#advanced-features--고급-기능)
   - [Pipeline / 파이프라인](#pipeline--파이프라인)
   - [Transactions / 트랜잭션](#transactions--트랜잭션)
   - [Pub/Sub / 발행/구독](#pubsub--발행구독)
7. [Usage Patterns / 사용 패턴](#usage-patterns--사용-패턴)
8. [Common Use Cases / 일반적인 사용 사례](#common-use-cases--일반적인-사용-사례)
9. [Best Practices / 모범 사례](#best-practices--모범-사례)
10. [Troubleshooting / 문제 해결](#troubleshooting--문제-해결)
11. [FAQ](#faq)

---

## Introduction / 소개

The Redis package is an extremely simplified Redis client for Go applications with automatic connection management, retry logic, and comprehensive operation support.

Redis 패키지는 자동 연결 관리, 재시도 로직, 포괄적인 작업 지원을 갖춘 Go 애플리케이션용 극도로 간소화된 Redis 클라이언트입니다.

### Key Features / 주요 기능

- **Extreme Simplicity / 극도의 간결함**: 20+ lines → 2 lines of code / 20줄 이상 → 2줄 코드로 축소
- **Auto-Everything / 자동화**: Connection, retry, reconnect, cleanup / 연결, 재시도, 재연결, 정리 모두 자동
- **60+ Methods / 60개 이상 메서드**: Complete Redis operation coverage / 완전한 Redis 작업 지원
- **Type-Safe Generics / 타입 안전 제네릭**: `GetAs[T]`, `HGetAllAs[T]` for automatic type conversion / 자동 타입 변환
- **Connection Pooling / 연결 풀링**: Built-in connection pool management / 내장 연결 풀 관리
- **Auto-Retry / 자동 재시도**: Exponential backoff for network errors / 네트워크 오류에 대한 지수 백오프
- **Health Monitoring / 헬스 모니터링**: Background health check goroutine / 백그라운드 헬스 체크 고루틴
- **Context Support / Context 지원**: All methods support context for cancellation / 모든 메서드에서 취소를 위한 context 지원
- **Clean API / 깨끗한 API**: Single import, no underlying library exposure / 단일 import, 기본 라이브러리 노출 없음

### Use Cases / 사용 사례

This package is ideal for:

이 패키지는 다음과 같은 경우에 이상적입니다:

- **Caching / 캐싱**: Session storage, API response caching / 세션 저장, API 응답 캐싱
- **Real-time Features / 실시간 기능**: Pub/Sub messaging, live updates / Pub/Sub 메시징, 실시간 업데이트
- **Leaderboards / 리더보드**: Sorted sets for rankings / 순위를 위한 정렬 집합
- **Rate Limiting / 속도 제한**: Counter-based rate limiting / 카운터 기반 속도 제한
- **Queue Systems / 큐 시스템**: Task queues with lists / 리스트를 사용한 작업 큐
- **Distributed Locks / 분산 잠금**: SetNX for distributed locking / 분산 잠금을 위한 SetNX

---

## Installation / 설치

### Prerequisites / 전제 조건

- **Go 1.18 or higher / Go 1.18 이상** (for generics support / 제네릭 지원)
- **Redis Server / Redis 서버**: Redis 5.0+ recommended / Redis 5.0+ 권장

### Install Package / 패키지 설치

```bash
go get github.com/arkd0ng/go-utils/database/redis
```

### Import in Your Code / 코드에 임포트

```go
import "github.com/arkd0ng/go-utils/database/redis"
```

**Note / 참고**: You only need to import this package. No need to import `github.com/redis/go-redis/v9` directly.

이 패키지만 import하면 됩니다. `github.com/redis/go-redis/v9`를 직접 import할 필요가 없습니다.

### Redis Server Setup / Redis 서버 설정

#### Option 1: Docker (Recommended) / 옵션 1: Docker (권장)

```bash
# Using docker compose / docker compose 사용
docker compose up -d redis

# Or using helper script / 또는 헬퍼 스크립트 사용
./.docker/scripts/docker-redis-start.sh
```

#### Option 2: Local Installation / 옵션 2: 로컬 설치

```bash
# macOS
brew install redis
brew services start redis

# Ubuntu/Debian
sudo apt-get install redis-server
sudo systemctl start redis

# Verify Redis is running / Redis 실행 확인
redis-cli ping
# Should return: PONG
```

---

## Quick Start / 빠른 시작

### Example 1: Minimal Setup / 최소 설정

The simplest way to get started:

가장 간단한 시작 방법:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/arkd0ng/go-utils/database/redis"
)

func main() {
    // Create client / 클라이언트 생성
    rdb, err := redis.New()
    if err != nil {
        log.Fatal(err)
    }
    defer rdb.Close()

    ctx := context.Background()

    // Set a value / 값 설정
    rdb.Set(ctx, "hello", "world")

    // Get a value / 값 가져오기
    val, _ := rdb.Get(ctx, "hello")
    fmt.Println(val) // Output: world
}
```

### Example 2: Custom Configuration / 사용자 정의 설정

With custom options:

사용자 정의 옵션 사용:

```go
rdb, err := redis.New(
    redis.WithAddr("localhost:6379"),
    redis.WithPassword("your-password"),
    redis.WithDB(1),
    redis.WithPoolSize(20),
    redis.WithMaxRetries(5),
)
if err != nil {
    log.Fatal(err)
}
defer rdb.Close()
```

### Example 3: Type-Safe Operations / 타입 안전 작업

Using generics for automatic type conversion:

자동 타입 변환을 위한 제네릭 사용:

```go
// Store integer / 정수 저장
rdb.Set(ctx, "count", 42)

// Retrieve as integer / 정수로 검색
count, err := redis.GetAs[int](rdb, ctx, "count")
fmt.Println(count) // Output: 42

// Store struct / 구조체 저장
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

user := User{Name: "Alice", Age: 30}
rdb.HSetMap(ctx, "user:1001", map[string]interface{}{
    "name": user.Name,
    "age":  user.Age,
})

// Retrieve as struct / 구조체로 검색
var result User
fields, _ := rdb.HGetAllAs[User](ctx, "user:1001")
fmt.Printf("%+v\n", fields) // Output: map[age:30 name:Alice]
```

---

## Configuration Reference / 설정 참조

### Available Options / 사용 가능한 옵션

| Option / 옵션 | Type / 타입 | Default / 기본값 | Description / 설명 |
|--------------|------------|----------------|-------------------|
| `WithAddr` | `string` | `"localhost:6379"` | Redis server address / Redis 서버 주소 |
| `WithPassword` | `string` | `""` | Redis password / Redis 비밀번호 |
| `WithDB` | `int` | `0` | Redis database number (0-15) / Redis 데이터베이스 번호 |
| `WithPoolSize` | `int` | `10` | Maximum number of socket connections / 최대 소켓 연결 수 |
| `WithMinIdleConns` | `int` | `2` | Minimum idle connections / 최소 유휴 연결 수 |
| `WithDialTimeout` | `time.Duration` | `5s` | Dial timeout / 다이얼 타임아웃 |
| `WithReadTimeout` | `time.Duration` | `3s` | Read timeout / 읽기 타임아웃 |
| `WithWriteTimeout` | `time.Duration` | `3s` | Write timeout / 쓰기 타임아웃 |
| `WithMaxRetries` | `int` | `3` | Maximum retry attempts / 최대 재시도 횟수 |
| `WithRetryDelay` | `time.Duration` | `100ms` | Initial retry delay / 초기 재시도 지연 |

### Configuration File / 설정 파일

Example `cfg/database-redis.yaml`:

```yaml
redis:
  addr: "localhost:6379"
  password: ""
  db: 0
  pool_size: 20
  min_idle_conns: 5
```

Loading from configuration file:

설정 파일에서 로드:

```go
import (
    "os"
    "gopkg.in/yaml.v3"
)

type Config struct {
    Redis struct {
        Addr         string `yaml:"addr"`
        Password     string `yaml:"password"`
        DB           int    `yaml:"db"`
        PoolSize     int    `yaml:"pool_size"`
        MinIdleConns int    `yaml:"min_idle_conns"`
    } `yaml:"redis"`
}

// Load config / 설정 로드
data, _ := os.ReadFile("cfg/database-redis.yaml")
var cfg Config
yaml.Unmarshal(data, &cfg)

// Create client / 클라이언트 생성
rdb, _ := redis.New(
    redis.WithAddr(cfg.Redis.Addr),
    redis.WithPassword(cfg.Redis.Password),
    redis.WithDB(cfg.Redis.DB),
    redis.WithPoolSize(cfg.Redis.PoolSize),
    redis.WithMinIdleConns(cfg.Redis.MinIdleConns),
)
```

---

## Core Operations / 핵심 작업

### String Operations / 문자열 작업

Strings are the most basic Redis data type.

문자열은 가장 기본적인 Redis 데이터 타입입니다.

#### Set / 설정

```go
// Basic set / 기본 설정
err := rdb.Set(ctx, "key", "value")

// Set with expiration / 만료 시간과 함께 설정
err := rdb.Set(ctx, "session", "abc123", 30*time.Minute)

// Set if not exists / 존재하지 않는 경우만 설정
ok, err := rdb.SetNX(ctx, "lock", "token", 10*time.Second)

// Set with expiration / 만료 시간 설정
err := rdb.SetEX(ctx, "temp", "data", 60*time.Second)
```

#### Get / 가져오기

```go
// Basic get / 기본 가져오기
val, err := rdb.Get(ctx, "key")
if err == redis.ErrNil {
    // Key does not exist / 키가 존재하지 않음
}

// Get with type conversion / 타입 변환과 함께 가져오기
count, err := redis.GetAs[int](rdb, ctx, "counter")
price, err := redis.GetAs[float64](rdb, ctx, "price")
isActive, err := redis.GetAs[bool](rdb, ctx, "active")
```

#### Multiple Operations / 다중 작업

```go
// Set multiple keys / 여러 키 설정
err := rdb.MSet(ctx, map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
    "key3": "value3",
})

// Get multiple keys / 여러 키 가져오기
values, err := rdb.MGet(ctx, "key1", "key2", "key3")
// Returns: []interface{}{"value1", "value2", "value3"}
```

#### Numeric Operations / 숫자 작업

```go
// Increment / 증가
newVal, err := rdb.Incr(ctx, "counter")        // +1
newVal, err := rdb.IncrBy(ctx, "counter", 5)   // +5

// Decrement / 감소
newVal, err := rdb.Decr(ctx, "counter")        // -1
newVal, err := rdb.DecrBy(ctx, "counter", 3)   // -3
```

#### String Manipulation / 문자열 조작

```go
// Append to string / 문자열에 추가
length, err := rdb.Append(ctx, "log", "new entry\n")

// Get substring / 부분 문자열 가져오기
substr, err := rdb.GetRange(ctx, "text", 0, 10)
```

### Hash Operations / 해시 작업

Hashes are maps between string fields and string values.

해시는 문자열 필드와 문자열 값 사이의 맵입니다.

#### Set Hash Fields / 해시 필드 설정

```go
// Set single field / 단일 필드 설정
err := rdb.HSet(ctx, "user:1001", "name", "Alice")

// Set multiple fields / 여러 필드 설정
err := rdb.HSetMap(ctx, "user:1001", map[string]interface{}{
    "name":  "Alice",
    "email": "alice@example.com",
    "age":   30,
    "city":  "Seoul",
})
```

#### Get Hash Fields / 해시 필드 가져오기

```go
// Get single field / 단일 필드 가져오기
name, err := rdb.HGet(ctx, "user:1001", "name")

// Get all fields / 모든 필드 가져오기
fields, err := rdb.HGetAll(ctx, "user:1001")
// Returns: map[string]string{"name": "Alice", "email": "...", ...}

// Get all fields with type conversion / 타입 변환과 함께 모든 필드 가져오기
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
    City  string `json:"city"`
}
user, err := rdb.HGetAllAs[User](ctx, "user:1001")
```

#### Hash Utilities / 해시 유틸리티

```go
// Check if field exists / 필드 존재 확인
exists, err := rdb.HExists(ctx, "user:1001", "email")

// Get hash length / 해시 길이 가져오기
length, err := rdb.HLen(ctx, "user:1001")

// Get all keys / 모든 키 가져오기
keys, err := rdb.HKeys(ctx, "user:1001")

// Get all values / 모든 값 가져오기
values, err := rdb.HVals(ctx, "user:1001")

// Delete fields / 필드 삭제
count, err := rdb.HDel(ctx, "user:1001", "temp_field")
```

#### Hash Numeric Operations / 해시 숫자 작업

```go
// Increment integer field / 정수 필드 증가
newAge, err := rdb.HIncrBy(ctx, "user:1001", "age", 1)

// Increment float field / 실수 필드 증가
newScore, err := rdb.HIncrByFloat(ctx, "user:1001", "score", 10.5)
```

### List Operations / 리스트 작업

Lists are simply lists of strings, sorted by insertion order.

리스트는 삽입 순서대로 정렬된 단순한 문자열 리스트입니다.

#### Push / 추가

```go
// Push to left (head) / 왼쪽(머리)에 추가
count, err := rdb.LPush(ctx, "queue", "task1", "task2", "task3")

// Push to right (tail) / 오른쪽(꼬리)에 추가
count, err := rdb.RPush(ctx, "queue", "task4", "task5")
```

#### Pop / 제거

```go
// Pop from left (head) / 왼쪽(머리)에서 제거
item, err := rdb.LPop(ctx, "queue")

// Pop from right (tail) / 오른쪽(꼬리)에서 제거
item, err := rdb.RPop(ctx, "queue")
```

#### Range Operations / 범위 작업

```go
// Get range / 범위 가져오기
items, err := rdb.LRange(ctx, "queue", 0, -1)  // All items / 모든 항목
items, err := rdb.LRange(ctx, "queue", 0, 9)   // First 10 / 처음 10개

// Get length / 길이 가져오기
length, err := rdb.LLen(ctx, "queue")

// Get by index / 인덱스로 가져오기
item, err := rdb.LIndex(ctx, "queue", 0)
```

#### List Modification / 리스트 수정

```go
// Set value at index / 인덱스에 값 설정
err := rdb.LSet(ctx, "queue", 0, "new_value")

// Remove elements / 요소 제거
count, err := rdb.LRem(ctx, "queue", 2, "task1")  // Remove 2 occurrences

// Trim list / 리스트 자르기
err := rdb.LTrim(ctx, "queue", 0, 99)  // Keep first 100 items
```

### Set Operations / 집합 작업

Sets are unordered collections of unique strings.

집합은 고유한 문자열의 정렬되지 않은 컬렉션입니다.

#### Basic Set Operations / 기본 집합 작업

```go
// Add members / 멤버 추가
count, err := rdb.SAdd(ctx, "tags", "go", "redis", "database")

// Get all members / 모든 멤버 가져오기
members, err := rdb.SMembers(ctx, "tags")

// Check membership / 멤버십 확인
exists, err := rdb.SIsMember(ctx, "tags", "go")

// Get cardinality (size) / 크기 가져오기
size, err := rdb.SCard(ctx, "tags")

// Remove members / 멤버 제거
count, err := rdb.SRem(ctx, "tags", "database")
```

#### Set Algebra / 집합 대수

```go
// Union / 합집합
rdb.SAdd(ctx, "set1", "a", "b", "c")
rdb.SAdd(ctx, "set2", "c", "d", "e")
union, err := rdb.SUnion(ctx, "set1", "set2")
// Result: [a, b, c, d, e]

// Intersection / 교집합
inter, err := rdb.SInter(ctx, "set1", "set2")
// Result: [c]

// Difference / 차집합
diff, err := rdb.SDiff(ctx, "set1", "set2")
// Result: [a, b]
```

#### Random Operations / 랜덤 작업

```go
// Pop random member / 랜덤 멤버 제거
member, err := rdb.SPop(ctx, "tags")

// Get random member without removing / 제거하지 않고 랜덤 멤버 가져오기
member, err := rdb.SRandMember(ctx, "tags")
```

### Sorted Set Operations / 정렬 집합 작업

Sorted sets are similar to sets but every member has an associated score.

정렬 집합은 집합과 유사하지만 모든 멤버에 연관된 점수가 있습니다.

#### Add Members / 멤버 추가

```go
// Add single member / 단일 멤버 추가
count, err := rdb.ZAdd(ctx, "leaderboard", 100, "Alice")

// Add multiple members / 여러 멤버 추가
count, err := rdb.ZAddMultiple(ctx, "leaderboard", map[string]float64{
    "Alice":   100,
    "Bob":     85,
    "Charlie": 95,
    "David":   90,
})
```

#### Get Members by Rank / 순위로 멤버 가져오기

```go
// Get range (ascending order) / 범위 가져오기 (오름차순)
members, err := rdb.ZRange(ctx, "leaderboard", 0, -1)    // All
members, err := rdb.ZRange(ctx, "leaderboard", 0, 2)     // Top 3 (lowest scores)

// Get range (descending order) / 범위 가져오기 (내림차순)
topPlayers, err := rdb.ZRevRange(ctx, "leaderboard", 0, 2)  // Top 3 (highest scores)
```

#### Get Members by Score / 점수로 멤버 가져오기

```go
// Get members in score range / 점수 범위의 멤버 가져오기
members, err := rdb.ZRangeByScore(ctx, "leaderboard", 80, 100)
```

#### Score Operations / 점수 작업

```go
// Get member score / 멤버 점수 가져오기
score, err := rdb.ZScore(ctx, "leaderboard", "Alice")

// Get member rank / 멤버 순위 가져오기
rank, err := rdb.ZRank(ctx, "leaderboard", "Alice")        // 0-based, ascending
rank, err := rdb.ZRevRank(ctx, "leaderboard", "Alice")     // 0-based, descending

// Increment score / 점수 증가
newScore, err := rdb.ZIncrBy(ctx, "leaderboard", 10, "Bob")

// Get size / 크기 가져오기
size, err := rdb.ZCard(ctx, "leaderboard")

// Remove members / 멤버 제거
count, err := rdb.ZRem(ctx, "leaderboard", "Charlie")
```

### Key Operations / 키 작업

Operations that work on keys of any type.

모든 타입의 키에 작동하는 작업입니다.

#### Basic Key Operations / 기본 키 작업

```go
// Delete keys / 키 삭제
count, err := rdb.Del(ctx, "key1", "key2", "key3")

// Check existence / 존재 확인
count, err := rdb.Exists(ctx, "key1", "key2")  // Returns: number of keys that exist

// Get key type / 키 타입 가져오기
keyType, err := rdb.Type(ctx, "mykey")  // Returns: "string", "list", "set", etc.

// Rename key / 키 이름 변경
err := rdb.Rename(ctx, "oldkey", "newkey")

// Rename if new key doesn't exist / 새 키가 존재하지 않는 경우만 이름 변경
ok, err := rdb.RenameNX(ctx, "oldkey", "newkey")
```

#### Expiration / 만료

```go
// Set expiration / 만료 설정
err := rdb.Expire(ctx, "session", 30*time.Minute)

// Set expiration at specific time / 특정 시간에 만료 설정
expireAt := time.Now().Add(24 * time.Hour)
err := rdb.ExpireAt(ctx, "session", expireAt)

// Get time-to-live / 남은 시간 가져오기
ttl, err := rdb.TTL(ctx, "session")
// Returns: time.Duration (-2 = key doesn't exist, -1 = no expiration)

// Remove expiration / 만료 제거
err := rdb.Persist(ctx, "session")
```

#### Key Search / 키 검색

```go
// Find keys by pattern / 패턴으로 키 찾기
keys, err := rdb.Keys(ctx, "user:*")          // All user keys
keys, err := rdb.Keys(ctx, "session:*")       // All session keys

// Scan keys (better for large datasets) / 키 스캔 (대용량 데이터셋에 더 적합)
cursor, keys, err := rdb.Scan(ctx, 0, "user:*", 100)
```

---

## Advanced Features / 고급 기능

### Pipeline / 파이프라인

Pipelines allow you to send multiple commands at once, reducing network round-trips.

파이프라인을 사용하면 여러 명령을 한 번에 보낼 수 있어 네트워크 왕복을 줄일 수 있습니다.

```go
err := rdb.Pipeline(ctx, func(pipe redis.Pipeliner) error {
    pipe.Set(ctx, "key1", "value1", 0)
    pipe.Set(ctx, "key2", "value2", 0)
    pipe.Set(ctx, "key3", "value3", 0)
    pipe.Incr(ctx, "counter")
    pipe.SAdd(ctx, "myset", "member1", "member2")
    return nil
})

if err != nil {
    log.Printf("Pipeline error: %v", err)
}
```

**Benefits / 장점:**
- Reduced network latency / 네트워크 지연 감소
- Improved throughput / 처리량 개선
- Atomic execution (all or nothing) / 원자적 실행 (전체 또는 아무것도)

### Transactions / 트랜잭션

Transactions provide optimistic locking using WATCH/MULTI/EXEC.

트랜잭션은 WATCH/MULTI/EXEC를 사용하여 낙관적 잠금을 제공합니다.

```go
err := rdb.Transaction(ctx, func(tx *redis.Tx) error {
    // Read current value / 현재 값 읽기
    val, err := tx.Get(ctx, "counter")
    if err != nil {
        return err
    }

    // Execute commands atomically / 명령을 원자적으로 실행
    return tx.Exec(ctx, func(pipe redis.Pipeliner) error {
        pipe.Incr(ctx, "counter")
        pipe.Set(ctx, "last_update", time.Now().Unix(), 0)
        return nil
    })
}, "counter") // Watch keys / 키 감시

if err != nil {
    log.Printf("Transaction error: %v", err)
}
```

**Use Cases / 사용 사례:**
- Atomic counter updates / 원자적 카운터 업데이트
- Account balance transfers / 계좌 잔액 이체
- Inventory management / 재고 관리
- Preventing race conditions / 경쟁 조건 방지

### Pub/Sub / 발행/구독

Publish/Subscribe messaging for real-time communication.

실시간 통신을 위한 발행/구독 메시징입니다.

#### Publisher / 발행자

```go
// Publish message / 메시지 발행
count, err := rdb.Publish(ctx, "news", "Breaking news!")
// Returns: number of subscribers who received the message
```

#### Subscriber / 구독자

```go
// Subscribe to channels / 채널 구독
pubsub, err := rdb.Subscribe(ctx, "news", "updates")
if err != nil {
    log.Fatal(err)
}
defer pubsub.Close()

// Receive messages / 메시지 수신
ch := pubsub.Channel()
for msg := range ch {
    fmt.Printf("Received from %s: %s\n", msg.Channel, msg.Payload)
}
```

#### Pattern Subscribe / 패턴 구독

```go
// Subscribe to pattern / 패턴 구독
pubsub, err := rdb.PSubscribe(ctx, "news:*", "updates:*")
if err != nil {
    log.Fatal(err)
}
defer pubsub.Close()

// Receive messages / 메시지 수신
ch := pubsub.Channel()
for msg := range ch {
    fmt.Printf("Pattern: %s, Channel: %s, Payload: %s\n",
        msg.Pattern, msg.Channel, msg.Payload)
}
```

---

## Usage Patterns / 사용 패턴

### Pattern 1: Session Storage / 세션 저장

```go
// Store session / 세션 저장
func StoreSession(ctx context.Context, rdb *redis.Client, sessionID string, data map[string]interface{}) error {
    key := fmt.Sprintf("session:%s", sessionID)
    err := rdb.HSetMap(ctx, key, data)
    if err != nil {
        return err
    }
    // Set 30 minute expiration / 30분 만료 설정
    return rdb.Expire(ctx, key, 30*time.Minute)
}

// Get session / 세션 가져오기
func GetSession(ctx context.Context, rdb *redis.Client, sessionID string) (map[string]string, error) {
    key := fmt.Sprintf("session:%s", sessionID)
    return rdb.HGetAll(ctx, key)
}

// Delete session / 세션 삭제
func DeleteSession(ctx context.Context, rdb *redis.Client, sessionID string) error {
    key := fmt.Sprintf("session:%s", sessionID)
    _, err := rdb.Del(ctx, key)
    return err
}
```

### Pattern 2: Cache with TTL / TTL이 있는 캐시

```go
// Cache with automatic expiration / 자동 만료가 있는 캐시
func CacheData(ctx context.Context, rdb *redis.Client, key string, data interface{}, ttl time.Duration) error {
    jsonData, err := json.Marshal(data)
    if err != nil {
        return err
    }
    return rdb.Set(ctx, key, string(jsonData), ttl)
}

// Get from cache / 캐시에서 가져오기
func GetCachedData(ctx context.Context, rdb *redis.Client, key string, result interface{}) error {
    data, err := rdb.Get(ctx, key)
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(data), result)
}
```

### Pattern 3: Rate Limiting / 속도 제한

```go
// Simple rate limiter / 간단한 속도 제한기
func CheckRateLimit(ctx context.Context, rdb *redis.Client, userID string, maxRequests int, window time.Duration) (bool, error) {
    key := fmt.Sprintf("ratelimit:%s", userID)

    // Increment counter / 카운터 증가
    count, err := rdb.Incr(ctx, key)
    if err != nil {
        return false, err
    }

    // Set expiration on first request / 첫 요청 시 만료 설정
    if count == 1 {
        rdb.Expire(ctx, key, window)
    }

    return count <= int64(maxRequests), nil
}
```

### Pattern 4: Distributed Lock / 분산 잠금

```go
// Acquire lock / 잠금 획득
func AcquireLock(ctx context.Context, rdb *redis.Client, resource string, ttl time.Duration) (bool, error) {
    lockKey := fmt.Sprintf("lock:%s", resource)
    token := uuid.New().String()
    return rdb.SetNX(ctx, lockKey, token, ttl)
}

// Release lock / 잠금 해제
func ReleaseLock(ctx context.Context, rdb *redis.Client, resource string) error {
    lockKey := fmt.Sprintf("lock:%s", resource)
    _, err := rdb.Del(ctx, lockKey)
    return err
}
```

### Pattern 5: Task Queue / 작업 큐

```go
// Enqueue task / 작업 큐에 추가
func EnqueueTask(ctx context.Context, rdb *redis.Client, queueName string, task string) error {
    _, err := rdb.RPush(ctx, queueName, task)
    return err
}

// Dequeue task / 작업 큐에서 제거
func DequeueTask(ctx context.Context, rdb *redis.Client, queueName string) (string, error) {
    return rdb.LPop(ctx, queueName)
}

// Get queue length / 큐 길이 가져오기
func GetQueueLength(ctx context.Context, rdb *redis.Client, queueName string) (int64, error) {
    return rdb.LLen(ctx, queueName)
}
```

### Pattern 6: Leaderboard / 리더보드

```go
// Update player score / 플레이어 점수 업데이트
func UpdateScore(ctx context.Context, rdb *redis.Client, playerID string, points float64) error {
    _, err := rdb.ZIncrBy(ctx, "leaderboard", points, playerID)
    return err
}

// Get top players / 상위 플레이어 가져오기
func GetTopPlayers(ctx context.Context, rdb *redis.Client, count int) ([]string, error) {
    return rdb.ZRevRange(ctx, "leaderboard", 0, int64(count-1))
}

// Get player rank / 플레이어 순위 가져오기
func GetPlayerRank(ctx context.Context, rdb *redis.Client, playerID string) (int64, error) {
    rank, err := rdb.ZRevRank(ctx, "leaderboard", playerID)
    if err != nil {
        return 0, err
    }
    return rank + 1, nil // Convert to 1-based rank
}
```

---

## Common Use Cases / 일반적인 사용 사례

### Use Case 1: API Response Caching / API 응답 캐싱

```go
func GetUserData(ctx context.Context, rdb *redis.Client, userID string) (*User, error) {
    cacheKey := fmt.Sprintf("user:%s", userID)

    // Try cache first / 먼저 캐시 시도
    cachedData, err := rdb.Get(ctx, cacheKey)
    if err == nil {
        var user User
        json.Unmarshal([]byte(cachedData), &user)
        return &user, nil
    }

    // Cache miss - fetch from database / 캐시 미스 - 데이터베이스에서 가져오기
    user, err := fetchUserFromDB(userID)
    if err != nil {
        return nil, err
    }

    // Store in cache / 캐시에 저장
    userData, _ := json.Marshal(user)
    rdb.Set(ctx, cacheKey, string(userData), 1*time.Hour)

    return user, nil
}
```

### Use Case 2: Shopping Cart / 장바구니

```go
// Add item to cart / 장바구니에 항목 추가
func AddToCart(ctx context.Context, rdb *redis.Client, userID, productID string, quantity int) error {
    cartKey := fmt.Sprintf("cart:%s", userID)
    return rdb.HSet(ctx, cartKey, productID, quantity)
}

// Get cart / 장바구니 가져오기
func GetCart(ctx context.Context, rdb *redis.Client, userID string) (map[string]string, error) {
    cartKey := fmt.Sprintf("cart:%s", userID)
    return rdb.HGetAll(ctx, cartKey)
}

// Remove item from cart / 장바구니에서 항목 제거
func RemoveFromCart(ctx context.Context, rdb *redis.Client, userID, productID string) error {
    cartKey := fmt.Sprintf("cart:%s", userID)
    _, err := rdb.HDel(ctx, cartKey, productID)
    return err
}

// Clear cart / 장바구니 비우기
func ClearCart(ctx context.Context, rdb *redis.Client, userID string) error {
    cartKey := fmt.Sprintf("cart:%s", userID)
    _, err := rdb.Del(ctx, cartKey)
    return err
}
```

### Use Case 3: Real-time Notifications / 실시간 알림

```go
// Notification broadcaster / 알림 브로드캐스터
func BroadcastNotification(ctx context.Context, rdb *redis.Client, channel, message string) error {
    _, err := rdb.Publish(ctx, channel, message)
    return err
}

// Notification listener / 알림 리스너
func ListenForNotifications(ctx context.Context, rdb *redis.Client, channels ...string) error {
    pubsub, err := rdb.Subscribe(ctx, channels...)
    if err != nil {
        return err
    }
    defer pubsub.Close()

    ch := pubsub.Channel()
    for {
        select {
        case msg := <-ch:
            handleNotification(msg.Channel, msg.Payload)
        case <-ctx.Done():
            return ctx.Err()
        }
    }
}

func handleNotification(channel, payload string) {
    log.Printf("Notification on %s: %s", channel, payload)
}
```

### Use Case 4: Job Queue with Priority / 우선순위가 있는 작업 큐

```go
// Enqueue job with priority / 우선순위와 함께 작업 큐에 추가
func EnqueueJob(ctx context.Context, rdb *redis.Client, jobID string, priority float64) error {
    _, err := rdb.ZAdd(ctx, "job_queue", priority, jobID)
    return err
}

// Dequeue highest priority job / 최우선순위 작업 큐에서 제거
func DequeueJob(ctx context.Context, rdb *redis.Client) (string, error) {
    jobs, err := rdb.ZRange(ctx, "job_queue", 0, 0)
    if err != nil || len(jobs) == 0 {
        return "", err
    }

    jobID := jobs[0]
    _, err = rdb.ZRem(ctx, "job_queue", jobID)
    return jobID, err
}
```

### Use Case 5: Online Users Tracking / 온라인 사용자 추적

```go
// Mark user as online / 사용자를 온라인으로 표시
func MarkUserOnline(ctx context.Context, rdb *redis.Client, userID string) error {
    _, err := rdb.SAdd(ctx, "online_users", userID)
    return err
}

// Mark user as offline / 사용자를 오프라인으로 표시
func MarkUserOffline(ctx context.Context, rdb *redis.Client, userID string) error {
    _, err := rdb.SRem(ctx, "online_users", userID)
    return err
}

// Get online users count / 온라인 사용자 수 가져오기
func GetOnlineUsersCount(ctx context.Context, rdb *redis.Client) (int64, error) {
    return rdb.SCard(ctx, "online_users")
}

// Check if user is online / 사용자가 온라인인지 확인
func IsUserOnline(ctx context.Context, rdb *redis.Client, userID string) (bool, error) {
    return rdb.SIsMember(ctx, "online_users", userID)
}
```

---

## Best Practices / 모범 사례

### 1. Always Use Context / 항상 Context 사용

```go
// Good / 좋음
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
val, err := rdb.Get(ctx, "key")

// Bad / 나쁨
val, err := rdb.Get(context.Background(), "key")  // No timeout
```

### 2. Handle Errors Properly / 에러를 적절히 처리

```go
val, err := rdb.Get(ctx, "key")
if err == redis.ErrNil {
    // Key doesn't exist - this is not an error / 키가 존재하지 않음 - 에러가 아님
    log.Println("Key not found, using default value")
} else if err != nil {
    // Actual error / 실제 에러
    return fmt.Errorf("redis error: %w", err)
}
```

### 3. Use Connection Pooling / 연결 풀링 사용

```go
// Set appropriate pool size / 적절한 풀 크기 설정
rdb, _ := redis.New(
    redis.WithPoolSize(20),        // Max connections / 최대 연결 수
    redis.WithMinIdleConns(5),     // Minimum idle / 최소 유휴 연결
)
```

### 4. Set TTL on Cache Keys / 캐시 키에 TTL 설정

```go
// Always set expiration on cache keys / 캐시 키에 항상 만료 설정
rdb.Set(ctx, "cache:user:123", data, 1*time.Hour)  // Good
rdb.Set(ctx, "cache:user:123", data)               // Bad - never expires
```

### 5. Use Pipelines for Bulk Operations / 대량 작업에 파이프라인 사용

```go
// Good - using pipeline / 좋음 - 파이프라인 사용
rdb.Pipeline(ctx, func(pipe redis.Pipeliner) error {
    for i := 0; i < 1000; i++ {
        pipe.Set(ctx, fmt.Sprintf("key:%d", i), i, 0)
    }
    return nil
})

// Bad - individual commands / 나쁨 - 개별 명령
for i := 0; i < 1000; i++ {
    rdb.Set(ctx, fmt.Sprintf("key:%d", i), i)
}
```

### 6. Use Consistent Key Naming / 일관된 키 명명 사용

```go
// Good - structured naming / 좋음 - 구조화된 명명
"user:1001:profile"
"user:1001:sessions"
"cache:api:user:1001"
"queue:email:high"

// Bad - inconsistent / 나쁨 - 일관성 없음
"user1001"
"userProfile1001"
"user_sessions_1001"
```

### 7. Monitor Connection Health / 연결 상태 모니터링

```go
// Ping periodically to check connection / 연결 확인을 위해 주기적으로 Ping
err := rdb.Ping(ctx)
if err != nil {
    log.Printf("Redis connection lost: %v", err)
    // Implement reconnection logic / 재연결 로직 구현
}
```

### 8. Use Transactions for Critical Operations / 중요한 작업에 트랜잭션 사용

```go
// For operations that must be atomic / 원자적이어야 하는 작업
err := rdb.Transaction(ctx, func(tx *redis.Tx) error {
    balance, _ := tx.Get(ctx, "balance")
    // ... check and update balance
    return tx.Exec(ctx, func(pipe redis.Pipeliner) error {
        pipe.Set(ctx, "balance", newBalance, 0)
        return nil
    })
}, "balance")
```

### 9. Properly Close Connections / 연결을 적절히 닫기

```go
func main() {
    rdb, err := redis.New()
    if err != nil {
        log.Fatal(err)
    }
    defer rdb.Close()  // Always defer Close / 항상 Close를 defer

    // ... use rdb
}
```

### 10. Use Appropriate Data Structures / 적절한 데이터 구조 사용

```go
// For user profiles - use Hash / 사용자 프로필 - Hash 사용
rdb.HSetMap(ctx, "user:1001", userData)

// For rankings - use Sorted Set / 순위 - Sorted Set 사용
rdb.ZAdd(ctx, "leaderboard", score, playerID)

// For unique tags - use Set / 고유 태그 - Set 사용
rdb.SAdd(ctx, "tags", "go", "redis", "cache")

// For queues - use List / 큐 - List 사용
rdb.LPush(ctx, "queue", task)
```

---

## Troubleshooting / 문제 해결

### Problem 1: Connection Refused / 연결 거부

**Symptom / 증상:**
```
dial tcp 127.0.0.1:6379: connect: connection refused
```

**Solutions / 해결책:**

1. Check if Redis is running / Redis가 실행 중인지 확인:
   ```bash
   redis-cli ping
   # Should return: PONG
   ```

2. Start Redis if not running / 실행 중이 아니면 Redis 시작:
   ```bash
   # Docker
   docker compose up -d redis

   # Local
   redis-server
   ```

3. Verify the address / 주소 확인:
   ```go
   rdb, _ := redis.New(redis.WithAddr("localhost:6379"))
   ```

### Problem 2: Authentication Failed / 인증 실패

**Symptom / 증상:**
```
NOAUTH Authentication required
```

**Solutions / 해결책:**

Set password in configuration / 설정에서 비밀번호 설정:
```go
rdb, _ := redis.New(
    redis.WithAddr("localhost:6379"),
    redis.WithPassword("your-password"),
)
```

### Problem 3: Key Not Found / 키를 찾을 수 없음

**Symptom / 증상:**
```
redis: nil
```

**Solutions / 해결책:**

```go
val, err := rdb.Get(ctx, "key")
if err == redis.ErrNil {
    // Key doesn't exist - handle gracefully / 키가 존재하지 않음 - 우아하게 처리
    log.Println("Using default value")
    val = "default"
} else if err != nil {
    // Real error / 실제 에러
    return err
}
```

### Problem 4: Timeout Errors / 타임아웃 에러

**Symptom / 증상:**
```
context deadline exceeded
i/o timeout
```

**Solutions / 해결책:**

1. Increase timeout / 타임아웃 증가:
   ```go
   rdb, _ := redis.New(
       redis.WithDialTimeout(10*time.Second),
       redis.WithReadTimeout(5*time.Second),
       redis.WithWriteTimeout(5*time.Second),
   )
   ```

2. Use context with timeout / 타임아웃이 있는 context 사용:
   ```go
   ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
   defer cancel()
   ```

### Problem 5: Memory Issues / 메모리 문제

**Symptom / 증상:**
```
OOM command not allowed when used memory > 'maxmemory'
```

**Solutions / 해결책:**

1. Set TTL on keys / 키에 TTL 설정:
   ```go
   rdb.Set(ctx, "key", "value", 1*time.Hour)
   ```

2. Use eviction policies / 축출 정책 사용:
   ```bash
   # In redis.conf
   maxmemory 256mb
   maxmemory-policy allkeys-lru
   ```

3. Clean up old keys / 오래된 키 정리:
   ```go
   keys, _ := rdb.Keys(ctx, "cache:*")
   if len(keys) > 0 {
       rdb.Del(ctx, keys...)
   }
   ```

### Problem 6: Slow Operations / 느린 작업

**Solutions / 해결책:**

1. Use Pipeline for bulk operations / 대량 작업에 파이프라인 사용
2. Increase pool size / 풀 크기 증가
3. Use Scan instead of Keys / Keys 대신 Scan 사용:
   ```go
   // Bad for large datasets / 대용량 데이터셋에 나쁨
   keys, _ := rdb.Keys(ctx, "*")

   // Good / 좋음
   cursor, keys, _ := rdb.Scan(ctx, 0, "*", 100)
   ```

---

## FAQ

### Q1: Do I need to import `github.com/redis/go-redis/v9`? / `github.com/redis/go-redis/v9`를 import해야 하나요?

**A:** No! You only need to import `github.com/arkd0ng/go-utils/database/redis`. All necessary types are exported from our package.

아니요! `github.com/arkd0ng/go-utils/database/redis`만 import하면 됩니다. 필요한 모든 타입이 우리 패키지에서 export됩니다.

### Q2: How do I handle nil values? / nil 값을 어떻게 처리하나요?

**A:** Check for `redis.ErrNil` error:

```go
val, err := rdb.Get(ctx, "key")
if err == redis.ErrNil {
    // Key doesn't exist / 키가 존재하지 않음
} else if err != nil {
    // Other error / 다른 에러
}
```

### Q3: What's the difference between Set and SetNX? / Set과 SetNX의 차이는 무엇인가요?

**A:**
- `Set`: Always sets the value / 항상 값을 설정
- `SetNX`: Only sets if key doesn't exist (useful for locks) / 키가 존재하지 않는 경우만 설정 (잠금에 유용)

### Q4: How do I implement distributed locking? / 분산 잠금을 어떻게 구현하나요?

**A:** Use `SetNX` with expiration:

```go
token := uuid.New().String()
acquired, _ := rdb.SetNX(ctx, "lock:resource", token, 10*time.Second)
if acquired {
    // Got the lock / 잠금 획득
    defer rdb.Del(ctx, "lock:resource")
    // ... do work
}
```

### Q5: What's the recommended pool size? / 권장 풀 크기는 얼마인가요?

**A:** Depends on your workload, but start with:
- Pool Size: 10-20 for most applications / 대부분의 애플리케이션에 10-20
- Min Idle: 2-5 connections / 2-5개 연결

```go
rdb, _ := redis.New(
    redis.WithPoolSize(20),
    redis.WithMinIdleConns(5),
)
```

### Q6: How do I handle connection failures? / 연결 실패를 어떻게 처리하나요?

**A:** The package automatically retries with exponential backoff. You can configure:

패키지는 자동으로 지수 백오프로 재시도합니다. 설정할 수 있습니다:

```go
rdb, _ := redis.New(
    redis.WithMaxRetries(5),
    redis.WithRetryDelay(100*time.Millisecond),
)
```

### Q7: Can I use this with Redis Cluster? / Redis 클러스터와 함께 사용할 수 있나요?

**A:** Currently, this package is designed for single-node Redis or Redis Sentinel. Redis Cluster support may be added in future versions.

현재 이 패키지는 단일 노드 Redis 또는 Redis Sentinel용으로 설계되었습니다. Redis 클러스터 지원은 향후 버전에 추가될 수 있습니다.

### Q8: How do I debug slow queries? / 느린 쿼리를 어떻게 디버그하나요?

**A:** Use Redis SLOWLOG command:

```bash
redis-cli SLOWLOG GET 10
```

Or monitor in real-time:

```bash
redis-cli MONITOR
```

### Q9: What's the difference between Pipeline and Transaction? / Pipeline과 Transaction의 차이는 무엇인가요?

**A:**
- **Pipeline**: Reduces network round-trips, no atomicity guarantee / 네트워크 왕복 감소, 원자성 보장 없음
- **Transaction**: Provides atomicity with optimistic locking (WATCH) / 낙관적 잠금으로 원자성 제공 (WATCH)

### Q10: How do I expire keys automatically? / 키를 자동으로 만료시키려면 어떻게 하나요?

**A:** Set TTL when creating keys:

```go
// With Set / Set과 함께
rdb.Set(ctx, "key", "value", 1*time.Hour)

// Or separately / 또는 별도로
rdb.Set(ctx, "key", "value")
rdb.Expire(ctx, "key", 1*time.Hour)
```

### Q11: Can I use generics with all operations? / 모든 작업에 제네릭을 사용할 수 있나요?

**A:** Type-safe generics are available for:
- `GetAs[T]`: String values / 문자열 값
- `HGetAllAs[T]`: Hash fields / 해시 필드

Other operations return `string` or `[]string` which you can convert manually.

다른 작업은 수동으로 변환할 수 있는 `string` 또는 `[]string`을 반환합니다.

### Q12: How do I monitor connection health? / 연결 상태를 어떻게 모니터링하나요?

**A:** Use the Ping method:

```go
err := rdb.Ping(ctx)
if err != nil {
    log.Printf("Redis health check failed: %v", err)
}
```

The package also runs automatic background health checks.

패키지는 또한 자동 백그라운드 헬스 체크를 실행합니다.

### Q13: What happens if Redis goes down? / Redis가 다운되면 어떻게 되나요?

**A:** The package will:
1. Return error immediately / 즉시 에러 반환
2. Automatically retry with exponential backoff / 지수 백오프로 자동 재시도
3. Continue retrying until max retries reached / 최대 재시도 횟수에 도달할 때까지 재시도 계속

### Q14: How do I batch operations efficiently? / 작업을 효율적으로 일괄 처리하려면 어떻게 하나요?

**A:** Use Pipeline:

```go
rdb.Pipeline(ctx, func(pipe redis.Pipeliner) error {
    for i := 0; i < 1000; i++ {
        pipe.Set(ctx, fmt.Sprintf("key:%d", i), i, 0)
    }
    return nil
})
```

This sends all 1000 commands in one network round-trip.

이는 1000개의 모든 명령을 하나의 네트워크 왕복으로 보냅니다.

### Q15: Where can I find more examples? / 더 많은 예제는 어디에서 찾을 수 있나요?

**A:** Check the following resources:
- `examples/redis/main.go`: 8 comprehensive examples / 8개의 포괄적인 예제
- `database/redis/README.md`: API reference / API 참조
- `docs/database/redis/DEVELOPER_GUIDE.md`: Advanced topics / 고급 주제

---

## Additional Resources / 추가 리소스

- **Package Documentation / 패키지 문서**: [database/redis/README.md](../../../database/redis/README.md)
- **Developer Guide / 개발자 가이드**: [DEVELOPER_GUIDE.md](./DEVELOPER_GUIDE.md)
- **Examples / 예제**: [examples/redis/](../../../examples/redis/)
- **Redis Official Documentation / Redis 공식 문서**: https://redis.io/documentation
- **go-redis Documentation / go-redis 문서**: https://redis.uptrace.dev/

---

## License / 라이선스

MIT License - see [LICENSE](../../../LICENSE) file for details

MIT 라이선스 - 자세한 내용은 [LICENSE](../../../LICENSE) 파일 참조

---

**Last Updated / 최종 업데이트**: 2025-10-14
**Version / 버전**: v1.4.014
