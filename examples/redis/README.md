# Redis Package Examples / Redis 패키지 예제

This directory contains comprehensive examples demonstrating all features of the go-utils Redis package.

이 디렉토리는 go-utils Redis 패키지의 모든 기능을 시연하는 포괄적인 예제를 포함합니다.

## Running Examples / 예제 실행

```bash
cd examples/redis
go run .
```

The examples will automatically:
- Connect to Redis at localhost:6379
- Run all 8 examples
- Clean up all test data

예제는 자동으로:
- localhost:6379의 Redis에 연결
- 8개 예제 모두 실행
- 모든 테스트 데이터 정리

## Examples Covered / 포함된 예제

### 1. String Operations / 문자열 작업

Demonstrates basic string operations:
- `Set` - Set string value with optional expiration
- `Get` - Get string value
- `MSet` - Set multiple string values at once
- `MGet` - Get multiple string values at once
- `Incr` - Increment integer value
- Auto-cleanup of test keys

기본 문자열 작업 시연:
- `Set` - 선택적 만료와 함께 문자열 값 설정
- `Get` - 문자열 값 가져오기
- `MSet` - 여러 문자열 값 한 번에 설정
- `MGet` - 여러 문자열 값 한 번에 가져오기
- `Incr` - 정수 값 증가
- 테스트 키 자동 정리

### 2. Hash Operations / 해시 작업

Demonstrates hash (map) operations:
- `HSet` - Set hash field
- `HGet` - Get hash field
- `HSetMap` - Set multiple hash fields from map
- `HGetAll` - Get all hash fields
- `HGetAllAs[T]` - Type-safe hash retrieval with generics
- `HIncrBy` - Increment hash field by integer
- Auto-cleanup of test hashes

해시(맵) 작업 시연:
- `HSet` - 해시 필드 설정
- `HGet` - 해시 필드 가져오기
- `HSetMap` - 맵에서 여러 해시 필드 설정
- `HGetAll` - 모든 해시 필드 가져오기
- `HGetAllAs[T]` - 제네릭을 사용한 타입 안전 해시 검색
- `HIncrBy` - 해시 필드를 정수로 증가
- 테스트 해시 자동 정리

### 3. List Operations / 리스트 작업

Demonstrates list operations:
- `LPush` - Push to head of list
- `RPush` - Push to tail of list
- `LPop` - Pop from head of list
- `RPop` - Pop from tail of list
- `LRange` - Get range of list elements
- `LLen` - Get list length
- Auto-cleanup of test lists

리스트 작업 시연:
- `LPush` - 리스트 머리에 추가
- `RPush` - 리스트 꼬리에 추가
- `LPop` - 리스트 머리에서 제거
- `RPop` - 리스트 꼬리에서 제거
- `LRange` - 리스트 요소 범위 가져오기
- `LLen` - 리스트 길이 가져오기
- 테스트 리스트 자동 정리

### 4. Set Operations / 집합 작업

Demonstrates set operations:
- `SAdd` - Add members to set
- `SMembers` - Get all members
- `SCard` - Get set cardinality (size)
- `SIsMember` - Check if member exists
- Set operations: `SUnion`, `SInter`, `SDiff`
- Auto-cleanup of test sets

집합 작업 시연:
- `SAdd` - 집합에 멤버 추가
- `SMembers` - 모든 멤버 가져오기
- `SCard` - 집합 크기 가져오기
- `SIsMember` - 멤버 존재 확인
- 집합 연산: `SUnion`, `SInter`, `SDiff`
- 테스트 집합 자동 정리

### 5. Sorted Set Operations / 정렬 집합 작업

Demonstrates sorted set operations:
- `ZAdd` - Add members with scores
- `ZAddMultiple` - Add multiple members with scores
- `ZRange` - Get members by rank range
- `ZRangeByScore` - Get members by score range
- `ZScore` - Get member score
- `ZRank` - Get member rank
- Auto-cleanup of test sorted sets

정렬 집합 작업 시연:
- `ZAdd` - 점수와 함께 멤버 추가
- `ZAddMultiple` - 점수와 함께 여러 멤버 추가
- `ZRange` - 순위 범위로 멤버 가져오기
- `ZRangeByScore` - 점수 범위로 멤버 가져오기
- `ZScore` - 멤버 점수 가져오기
- `ZRank` - 멤버 순위 가져오기
- 테스트 정렬 집합 자동 정리

### 6. Key Operations / 키 작업

Demonstrates key management operations:
- `Del` - Delete keys
- `Exists` - Check if key exists
- `Expire` - Set key expiration
- `TTL` - Get key time-to-live
- `Keys` - Find keys by pattern
- `Type` - Get key data type
- Auto-cleanup of test keys

키 관리 작업 시연:
- `Del` - 키 삭제
- `Exists` - 키 존재 확인
- `Expire` - 키 만료 설정
- `TTL` - 키 남은 시간 가져오기
- `Keys` - 패턴으로 키 찾기
- `Type` - 키 데이터 타입 가져오기
- 테스트 키 자동 정리

### 7. Pipeline Operations / 파이프라인 작업

Demonstrates pipeline for batch operations:
- Create pipeline with `Pipeline()`
- Queue multiple commands
- Execute all commands at once with `Exec()`
- Retrieve results from each command
- Benefits: Reduced network round-trips
- Auto-cleanup of test data

배치 작업을 위한 파이프라인 시연:
- `Pipeline()`으로 파이프라인 생성
- 여러 명령 대기열에 추가
- `Exec()`으로 모든 명령 한 번에 실행
- 각 명령의 결과 검색
- 장점: 네트워크 왕복 감소
- 테스트 데이터 자동 정리

### 8. Transaction Operations / 트랜잭션 작업

Demonstrates transaction with optimistic locking:
- Use `Transaction()` with WATCH/MULTI/EXEC
- Optimistic locking prevents race conditions
- Automatic retry on transaction failure
- All-or-nothing execution
- Auto-cleanup of test data

낙관적 잠금을 사용한 트랜잭션 시연:
- WATCH/MULTI/EXEC와 함께 `Transaction()` 사용
- 낙관적 잠금으로 경쟁 조건 방지
- 트랜잭션 실패 시 자동 재시도
- 전체 또는 아무것도 실행되지 않음
- 테스트 데이터 자동 정리

## Prerequisites / 전제 조건

### Option 1: Docker Redis (Recommended) / 옵션 1: Docker Redis (권장)

```bash
# Start Redis using Docker / Docker를 사용하여 Redis 시작
./scripts/docker-redis-start.sh

# Stop Redis / Redis 중지
./scripts/docker-redis-stop.sh

# View logs / 로그 보기
./scripts/docker-redis-logs.sh

# Connect to Redis CLI / Redis CLI에 연결
./scripts/docker-redis-cli.sh
```

### Option 2: Local Redis / 옵션 2: 로컬 Redis

```bash
# macOS
brew install redis
brew services start redis

# Linux
sudo apt-get install redis-server
sudo systemctl start redis
```

### Requirements / 요구사항

- Docker Desktop installed and running (for Option 1)
- Go 1.24.6 or higher
- Redis running at localhost:6379

Docker Desktop이 설치되어 실행 중이어야 합니다 (옵션 1의 경우).

## Configuration / 설정

The examples use default Redis configuration:
- Host: localhost
- Port: 6379
- Password: (none)
- Database: 0

예제는 기본 Redis 설정을 사용합니다:
- 호스트: localhost
- 포트: 6379
- 비밀번호: (없음)
- 데이터베이스: 0

To customize, edit the connection in `main.go`:

사용자 정의하려면 `main.go`의 연결을 편집하세요:

```go
rdb, err := redis.New(
    redis.WithAddr("localhost:6379"),
    redis.WithPassword("your-password"),
    redis.WithDB(0),
)
```

## Output Example / 출력 예제

```
=== Redis Package Examples ===
=== Redis 패키지 예제 ===

1. String Operations / 문자열 작업
SET name = 'John Doe'
GET name = 'John Doe'
SET session = 'abc123' with 10s expiration
INCR counter = 1
MSET key1, key2, key3
MGET = [value1 value2 value3]

2. Hash Operations / 해시 작업
HSET user:1001 name = 'Alice'
HGET user:1001 name = 'Alice'
HSET user:1001 age = '30'
HGETALL user:1001 = map[age:30 name:Alice]
Type-safe HGetAllAs: User{Name:Alice Age:30}
HINCRBY user:1001 age = 31

3. List Operations / 리스트 작업
...

=== All examples completed successfully! ===
=== 모든 예제가 성공적으로 완료되었습니다! ===
```

## Cleanup / 정리

All examples automatically clean up their test data. No manual cleanup is required.

모든 예제는 자동으로 테스트 데이터를 정리합니다. 수동 정리가 필요하지 않습니다.

However, if you want to flush all Redis data manually:

그러나 수동으로 모든 Redis 데이터를 플러시하려면:

```bash
# Connect to Redis CLI / Redis CLI에 연결
./scripts/docker-redis-cli.sh

# Flush current database / 현재 데이터베이스 플러시
> FLUSHDB

# Flush all databases / 모든 데이터베이스 플러시
> FLUSHALL
```

## Troubleshooting / 문제 해결

### Cannot connect to Redis / Redis에 연결할 수 없음

**Problem / 문제**: `Failed to connect to Redis: dial tcp [::1]:6379: connect: connection refused`

**Solution / 해결책**:
1. Check if Redis is running / Redis가 실행 중인지 확인
   ```bash
   # For Docker
   docker ps | grep go-utils-redis

   # For local Redis
   redis-cli ping
   ```

2. Start Redis if not running / 실행 중이 아니면 Redis 시작
   ```bash
   # For Docker
   ./scripts/docker-redis-start.sh

   # For local Redis
   brew services start redis
   ```

### Port already in use / 포트가 이미 사용 중

**Problem / 문제**: Port 6379 is already in use by another process

**Solution / 해결책**:
1. Check what's using port 6379 / 6379 포트를 사용하는 프로세스 확인
   ```bash
   lsof -i :6379
   ```

2. Stop the conflicting process or use a different port
   충돌하는 프로세스를 중지하거나 다른 포트 사용

### Slow operations / 느린 작업

**Problem / 문제**: Operations are slower than expected

**Solution / 해결책**:
- Redis is designed for in-memory speed. If slow, check:
  - Network latency (use localhost for best performance)
  - Redis configuration (check maxmemory, persistence settings)
  - System resources (CPU, memory)

Redis는 메모리 내 속도를 위해 설계되었습니다. 느린 경우 확인:
- 네트워크 지연 (최상의 성능을 위해 localhost 사용)
- Redis 설정 (maxmemory, 지속성 설정 확인)
- 시스템 리소스 (CPU, 메모리)

## Additional Resources / 추가 리소스

- [Redis Package Documentation](../../database/redis/README.md)
- [Redis Package User Manual](../../docs/database/redis/USER_MANUAL.md)
- [Redis Package Developer Guide](../../docs/database/redis/DEVELOPER_GUIDE.md)
- [Redis Official Documentation](https://redis.io/documentation)
- [go-redis Documentation](https://redis.uptrace.dev/)

## License / 라이선스

MIT License - see [LICENSE](../../LICENSE) file for details

MIT 라이선스 - 자세한 내용은 [LICENSE](../../LICENSE) 파일 참조
