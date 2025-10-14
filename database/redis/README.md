# Redis Package / Redis 패키지

Extreme simplicity Redis client for Go - reduce 20+ lines of code to just 2 lines!

Go를 위한 극도로 간단한 Redis 클라이언트 - 20줄 이상의 코드를 단 2줄로 축소!

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.16-blue)](https://golang.org/)

## Features / 기능

- **Extreme Simplicity / 극도의 간결함**: 20+ lines → 2 lines
- **Auto-Everything / 모든 것 자동화**: Connection, retry, reconnect, resource cleanup
- **Type-Safe API / 타입 안전 API**: Generic support for type-safe operations
- **Auto-Retry / 자동 재시도**: Network errors are automatically retried with exponential backoff
- **Connection Pooling / 연결 풀링**: Built-in connection pooling for high performance
- **Health Check / 헬스 체크**: Background health checking
- **Pipeline Support / 파이프라인 지원**: Batch operations for better performance
- **Transaction Support / 트랜잭션 지원**: Optimistic locking with WATCH/MULTI/EXEC
- **Pub/Sub Support / Pub/Sub 지원**: Message publishing and subscribing

## Installation / 설치

```bash
go get github.com/arkd0ng/go-utils/database/redis
```

## Quick Start / 빠른 시작

### Basic Usage / 기본 사용법

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
    rdb, err := redis.New(redis.WithAddr("localhost:6379"))
    if err != nil {
        log.Fatal(err)
    }
    defer rdb.Close()

    ctx := context.Background()

    // String operations / 문자열 작업
    rdb.Set(ctx, "key", "value")
    val, _ := rdb.Get(ctx, "key")
    fmt.Println(val) // Output: value
}
```

### Before vs After / 전후 비교

**❌ Before (standard go-redis):**

```go
import "github.com/redis/go-redis/v9"

// Setup / 설정
rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})

// String operations / 문자열 작업
ctx := context.Background()
err := rdb.Set(ctx, "key", "value", 0).Err()
if err != nil {
    return err
}

val, err := rdb.Get(ctx, "key").Result()
if err != nil {
    return err
}
// ... 매번 .Err() 또는 .Result() 호출...
```

**✅ After (this package):**

```go
import "github.com/arkd0ng/go-utils/database/redis"

// Setup / 설정
rdb, _ := redis.New(redis.WithAddr("localhost:6379"))
defer rdb.Close()

// String operations / 문자열 작업
ctx := context.Background()
rdb.Set(ctx, "key", "value")
val, _ := rdb.Get(ctx, "key")
// 간결하고 읽기 쉬운 API!
```

## API Reference / API 참조

### Configuration Options / 설정 옵션

```go
rdb, err := redis.New(
    redis.WithAddr("localhost:6379"),        // Redis address / Redis 주소
    redis.WithPassword("password"),          // Password / 비밀번호
    redis.WithDB(0),                         // Database number (0-15) / 데이터베이스 번호
    redis.WithPoolSize(10),                  // Connection pool size / 연결 풀 크기
    redis.WithMinIdleConns(5),               // Minimum idle connections / 최소 유휴 연결
    redis.WithMaxRetries(3),                 // Maximum retries / 최대 재시도 횟수
    redis.WithRetryInterval(100*time.Millisecond), // Retry interval / 재시도 간격
    redis.WithHealthCheck(true),             // Enable health check / 헬스 체크 활성화
    redis.WithHealthCheckInterval(30*time.Second), // Health check interval / 헬스 체크 간격
)
```

### String Operations / 문자열 작업

```go
// Set / 설정
rdb.Set(ctx, "key", "value")
rdb.Set(ctx, "key", "value", 10*time.Second) // With expiration / 만료와 함께

// Get / 가져오기
val, err := rdb.Get(ctx, "key")

// Type-safe get / 타입 안전 가져오기
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
user, err := redis.GetAs[User](rdb, ctx, "user:123")

// Multiple get/set / 다중 get/set
rdb.MSet(ctx, map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
})
vals, _ := rdb.MGet(ctx, "key1", "key2")

// Increment/Decrement / 증가/감소
count, _ := rdb.Incr(ctx, "counter")
rdb.Decr(ctx, "counter")
```

### Hash Operations / 해시 작업

```go
// Set hash field / 해시 필드 설정
rdb.HSet(ctx, "user:123", "name", "John")

// Set multiple fields / 여러 필드 설정
rdb.HSetMap(ctx, "user:123", map[string]interface{}{
    "name":  "John",
    "email": "john@example.com",
    "age":   30,
})

// Get hash field / 해시 필드 가져오기
name, _ := rdb.HGet(ctx, "user:123", "name")

// Get all fields / 모든 필드 가져오기
fields, _ := rdb.HGetAll(ctx, "user:123")

// Type-safe get all / 타입 안전 전체 가져오기
user, err := redis.HGetAllAs[User](rdb, ctx, "user:123")
```

### List Operations / 리스트 작업

```go
// Push / 추가
rdb.RPush(ctx, "queue", "item1", "item2", "item3")
rdb.LPush(ctx, "stack", "item1")

// Pop / 제거
item, _ := rdb.RPop(ctx, "queue")
item, _ := rdb.LPop(ctx, "stack")

// Range / 범위
items, _ := rdb.LRange(ctx, "queue", 0, -1)

// Length / 길이
length, _ := rdb.LLen(ctx, "queue")
```

### Set Operations / 집합 작업

```go
// Add members / 멤버 추가
rdb.SAdd(ctx, "tags", "golang", "redis", "database")

// Check membership / 멤버 확인
exists, _ := rdb.SIsMember(ctx, "tags", "golang")

// Get all members / 모든 멤버 가져오기
members, _ := rdb.SMembers(ctx, "tags")

// Set operations / 집합 작업
union, _ := rdb.SUnion(ctx, "set1", "set2")
inter, _ := rdb.SInter(ctx, "set1", "set2")
diff, _ := rdb.SDiff(ctx, "set1", "set2")
```

### Sorted Set Operations / 정렬 집합 작업

```go
// Add members with scores / 점수와 함께 멤버 추가
rdb.ZAdd(ctx, "leaderboard", 100, "player1")
rdb.ZAdd(ctx, "leaderboard", 200, "player2")

// Add multiple / 여러 개 추가
rdb.ZAddMultiple(ctx, "leaderboard", map[string]float64{
    "player3": 150,
    "player4": 175,
})

// Get range / 범위 가져오기
players, _ := rdb.ZRange(ctx, "leaderboard", 0, -1)

// Get by score / 점수로 가져오기
players, _ := rdb.ZRangeByScore(ctx, "leaderboard", 100, 200)

// Get score / 점수 가져오기
score, _ := rdb.ZScore(ctx, "leaderboard", "player1")
```

### Key Operations / 키 작업

```go
// Delete / 삭제
rdb.Del(ctx, "key1", "key2", "key3")

// Exists / 존재 확인
count, _ := rdb.Exists(ctx, "key1", "key2")

// Set expiration / 만료 설정
rdb.Expire(ctx, "key", 10*time.Second)

// Get TTL / TTL 가져오기
ttl, _ := rdb.TTL(ctx, "key")

// Find keys / 키 찾기
keys, _ := rdb.Keys(ctx, "user:*")

// Scan keys / 키 스캔
keys, cursor, _ := rdb.Scan(ctx, 0, "user:*", 10)
```

### Pipeline Operations / 파이프라인 작업

```go
// Execute multiple commands in one round trip / 한 번의 왕복으로 여러 명령 실행
err := rdb.Pipeline(ctx, func(pipe redis.Pipeliner) error {
    pipe.Set(ctx, "key1", "value1", 0)
    pipe.Set(ctx, "key2", "value2", 0)
    pipe.Incr(ctx, "counter")
    pipe.SAdd(ctx, "set", "member1", "member2")
    return nil
})
```

### Transaction Operations / 트랜잭션 작업

```go
// Transaction with optimistic locking / 낙관적 잠금을 사용한 트랜잭션
err := rdb.Transaction(ctx, func(tx *redis.Tx) error {
    // Get current value / 현재 값 가져오기
    val, err := tx.Get(ctx, "counter")
    if err != nil {
        return err
    }

    count, _ := strconv.Atoi(val)

    // Execute commands in transaction / 트랜잭션에서 명령 실행
    return tx.Exec(ctx, func(pipe redis.Pipeliner) error {
        pipe.Set(ctx, "counter", count+1, 0)
        pipe.Set(ctx, "last_update", time.Now().Unix(), 0)
        return nil
    })
}, "counter") // Watch keys / 키 감시
```

### Pub/Sub Operations / Pub/Sub 작업

```go
// Publish / 발행
rdb.Publish(ctx, "notifications", "Hello, World!")

// Subscribe / 구독
pubsub, _ := rdb.Subscribe(ctx, "notifications", "alerts")
defer pubsub.Close()

// Receive messages / 메시지 받기
ch := pubsub.Channel()
for msg := range ch {
    fmt.Printf("Channel: %s, Message: %s\n", msg.Channel, msg.Payload)
}

// Pattern subscribe / 패턴 구독
pubsub, _ := rdb.PSubscribe(ctx, "notification:*")
```

## Configuration / 설정

### Using Config File / 설정 파일 사용

Create `cfg/redis.yaml`:

```yaml
redis:
  addr: localhost:6379
  password: ""
  db: 0
  pool_size: 10
  min_idle_conns: 5
  max_retries: 3
  retry_interval: 100 # milliseconds
  enable_health_check: true
  health_check_interval: 30 # seconds
```

## Error Handling / 에러 처리

```go
val, err := rdb.Get(ctx, "key")
if err != nil {
    if err == redis.ErrNil {
        // Key doesn't exist / 키가 존재하지 않음
    } else {
        // Other errors / 기타 에러
        log.Printf("Redis error: %v", err)
    }
}
```

## Best Practices / 모범 사례

1. **Always use context / 항상 context 사용**: Pass context for cancellation and timeout
2. **Reuse client / 클라이언트 재사용**: Create one client and reuse it
3. **Close client / 클라이언트 닫기**: Always defer `client.Close()`
4. **Use pipeline for batch operations / 배치 작업에 파이프라인 사용**: Better performance
5. **Use transactions for atomic operations / 원자적 작업에 트랜잭션 사용**: Data consistency
6. **Handle ErrNil / ErrNil 처리**: Check for key existence
7. **Set expiration / 만료 설정**: Prevent memory leaks

## Testing / 테스트

Start Docker Redis:

```bash
./scripts/docker-redis-start.sh
```

Run tests:

```bash
go test ./database/redis -v
```

Stop Docker Redis:

```bash
./scripts/docker-redis-stop.sh
```

## Examples / 예제

See [examples/redis](../../examples/redis/) for complete examples.

완전한 예제는 [examples/redis](../../examples/redis/)를 참조하세요.

## Performance / 성능

- **Connection Pooling / 연결 풀링**: Reuses connections for better performance
- **Pipeline / 파이프라인**: Reduces network round trips
- **Auto-Retry / 자동 재시도**: Handles transient network errors
- **Health Check / 헬스 체크**: Monitors connection health in background

## Contributing / 기여하기

Contributions are welcome! Please see the [Contributing Guidelines](../../CONTRIBUTING.md).

기여를 환영합니다! [기여 가이드라인](../../CONTRIBUTING.md)을 참조하세요.

## License / 라이선스

MIT License - see [LICENSE](../../LICENSE) for details.

MIT 라이선스 - 자세한 내용은 [LICENSE](../../LICENSE)를 참조하세요.
