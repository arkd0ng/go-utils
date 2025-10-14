package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/arkd0ng/go-utils/database/redis"
	goredis "github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("=== Redis Package Examples ===")
	fmt.Println("=== Redis 패키지 예제 ===\n")

	// Create client / 클라이언트 생성
	rdb, err := redis.New(redis.WithAddr("localhost:6379"))
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer rdb.Close()

	ctx := context.Background()

	// 1. String Operations / 문자열 작업
	fmt.Println("1. String Operations / 문자열 작업")
	stringOperations(ctx, rdb)

	// 2. Hash Operations / 해시 작업
	fmt.Println("\n2. Hash Operations / 해시 작업")
	hashOperations(ctx, rdb)

	// 3. List Operations / 리스트 작업
	fmt.Println("\n3. List Operations / 리스트 작업")
	listOperations(ctx, rdb)

	// 4. Set Operations / 집합 작업
	fmt.Println("\n4. Set Operations / 집합 작업")
	setOperations(ctx, rdb)

	// 5. Sorted Set Operations / 정렬 집합 작업
	fmt.Println("\n5. Sorted Set Operations / 정렬 집합 작업")
	sortedSetOperations(ctx, rdb)

	// 6. Key Operations / 키 작업
	fmt.Println("\n6. Key Operations / 키 작업")
	keyOperations(ctx, rdb)

	// 7. Pipeline Operations / 파이프라인 작업
	fmt.Println("\n7. Pipeline Operations / 파이프라인 작업")
	pipelineOperations(ctx, rdb)

	// 8. Transaction Operations / 트랜잭션 작업
	fmt.Println("\n8. Transaction Operations / 트랜잭션 작업")
	transactionOperations(ctx, rdb)

	fmt.Println("\n=== All examples completed successfully! ===")
	fmt.Println("=== 모든 예제가 성공적으로 완료되었습니다! ===")
}

// stringOperations demonstrates string operations
// stringOperations는 문자열 작업을 시연합니다
func stringOperations(ctx context.Context, rdb *redis.Client) {
	// Set / 설정
	rdb.Set(ctx, "name", "John Doe")
	fmt.Println("SET name = 'John Doe'")

	// Get / 가져오기
	name, _ := rdb.Get(ctx, "name")
	fmt.Printf("GET name = '%s'\n", name)

	// Set with expiration / 만료와 함께 설정
	rdb.Set(ctx, "session", "abc123", 10*time.Second)
	fmt.Println("SET session = 'abc123' with 10s expiration")

	// Increment / 증가
	rdb.Set(ctx, "counter", "0")
	count, _ := rdb.Incr(ctx, "counter")
	fmt.Printf("INCR counter = %d\n", count)

	// Multiple set / 다중 설정
	rdb.MSet(ctx, map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	fmt.Println("MSET key1, key2, key3")

	// Multiple get / 다중 가져오기
	values, _ := rdb.MGet(ctx, "key1", "key2", "key3")
	fmt.Printf("MGET = %v\n", values)

	// Cleanup / 정리
	rdb.Del(ctx, "name", "session", "counter", "key1", "key2", "key3")
}

// hashOperations demonstrates hash operations
// hashOperations는 해시 작업을 시연합니다
func hashOperations(ctx context.Context, rdb *redis.Client) {
	// Set hash field / 해시 필드 설정
	rdb.HSet(ctx, "user:1001", "name", "Alice")
	fmt.Println("HSET user:1001 name = 'Alice'")

	// Set multiple fields / 여러 필드 설정
	rdb.HSetMap(ctx, "user:1001", map[string]interface{}{
		"email": "alice@example.com",
		"age":   "25",
		"city":  "Seoul",
	})
	fmt.Println("HMSET user:1001 email, age, city")

	// Get hash field / 해시 필드 가져오기
	name, _ := rdb.HGet(ctx, "user:1001", "name")
	fmt.Printf("HGET user:1001 name = '%s'\n", name)

	// Get all fields / 모든 필드 가져오기
	fields, _ := rdb.HGetAll(ctx, "user:1001")
	fmt.Printf("HGETALL user:1001 = %v\n", fields)

	// Increment hash field / 해시 필드 증가
	newAge, _ := rdb.HIncrBy(ctx, "user:1001", "age", 1)
	fmt.Printf("HINCRBY user:1001 age 1 = %d\n", newAge)

	// Cleanup / 정리
	rdb.Del(ctx, "user:1001")
}

// listOperations demonstrates list operations
// listOperations는 리스트 작업을 시연합니다
func listOperations(ctx context.Context, rdb *redis.Client) {
	// Push to list / 리스트에 추가
	rdb.RPush(ctx, "queue", "task1", "task2", "task3")
	fmt.Println("RPUSH queue = ['task1', 'task2', 'task3']")

	// Get list length / 리스트 길이 가져오기
	length, _ := rdb.LLen(ctx, "queue")
	fmt.Printf("LLEN queue = %d\n", length)

	// Get range / 범위 가져오기
	items, _ := rdb.LRange(ctx, "queue", 0, -1)
	fmt.Printf("LRANGE queue 0 -1 = %v\n", items)

	// Pop from list / 리스트에서 제거
	item, _ := rdb.LPop(ctx, "queue")
	fmt.Printf("LPOP queue = '%s'\n", item)

	// Cleanup / 정리
	rdb.Del(ctx, "queue")
}

// setOperations demonstrates set operations
// setOperations는 집합 작업을 시연합니다
func setOperations(ctx context.Context, rdb *redis.Client) {
	// Add to set / 집합에 추가
	rdb.SAdd(ctx, "languages", "Go", "Python", "JavaScript", "Rust")
	fmt.Println("SADD languages = ['Go', 'Python', 'JavaScript', 'Rust']")

	// Check membership / 멤버 확인
	exists, _ := rdb.SIsMember(ctx, "languages", "Go")
	fmt.Printf("SISMEMBER languages 'Go' = %v\n", exists)

	// Get all members / 모든 멤버 가져오기
	members, _ := rdb.SMembers(ctx, "languages")
	fmt.Printf("SMEMBERS languages = %v\n", members)

	// Get cardinality / 크기 가져오기
	size, _ := rdb.SCard(ctx, "languages")
	fmt.Printf("SCARD languages = %d\n", size)

	// Set operations / 집합 작업
	rdb.SAdd(ctx, "backend", "Go", "Python", "Java")
	union, _ := rdb.SUnion(ctx, "languages", "backend")
	fmt.Printf("SUNION languages backend = %v\n", union)

	// Cleanup / 정리
	rdb.Del(ctx, "languages", "backend")
}

// sortedSetOperations demonstrates sorted set operations
// sortedSetOperations는 정렬 집합 작업을 시연합니다
func sortedSetOperations(ctx context.Context, rdb *redis.Client) {
	// Add members with scores / 점수와 함께 멤버 추가
	rdb.ZAddMultiple(ctx, "leaderboard", map[string]float64{
		"Alice":   100,
		"Bob":     85,
		"Charlie": 95,
		"David":   90,
	})
	fmt.Println("ZADD leaderboard with scores")

	// Get range (ascending) / 범위 가져오기 (오름차순)
	players, _ := rdb.ZRange(ctx, "leaderboard", 0, -1)
	fmt.Printf("ZRANGE leaderboard 0 -1 = %v\n", players)

	// Get range (descending) / 범위 가져오기 (내림차순)
	topPlayers, _ := rdb.ZRevRange(ctx, "leaderboard", 0, 2)
	fmt.Printf("Top 3 players: %v\n", topPlayers)

	// Get score / 점수 가져오기
	score, _ := rdb.ZScore(ctx, "leaderboard", "Alice")
	fmt.Printf("ZSCORE leaderboard 'Alice' = %.0f\n", score)

	// Increment score / 점수 증가
	newScore, _ := rdb.ZIncrBy(ctx, "leaderboard", 5, "Bob")
	fmt.Printf("ZINCRBY leaderboard 'Bob' 5 = %.0f\n", newScore)

	// Cleanup / 정리
	rdb.Del(ctx, "leaderboard")
}

// keyOperations demonstrates key operations
// keyOperations는 키 작업을 시연합니다
func keyOperations(ctx context.Context, rdb *redis.Client) {
	// Set keys / 키 설정
	rdb.Set(ctx, "app:config:debug", "true")
	rdb.Set(ctx, "app:config:timeout", "30")
	rdb.Set(ctx, "app:user:1001", "Alice")
	fmt.Println("SET multiple app:* keys")

	// Find keys by pattern / 패턴으로 키 찾기
	keys, _ := rdb.Keys(ctx, "app:config:*")
	fmt.Printf("KEYS app:config:* = %v\n", keys)

	// Check existence / 존재 확인
	count, _ := rdb.Exists(ctx, "app:config:debug", "app:config:timeout")
	fmt.Printf("EXISTS app:config:* = %d keys\n", count)

	// Set expiration / 만료 설정
	rdb.Expire(ctx, "app:user:1001", 60*time.Second)
	ttl, _ := rdb.TTL(ctx, "app:user:1001")
	fmt.Printf("TTL app:user:1001 = %v\n", ttl)

	// Delete keys / 키 삭제
	rdb.Del(ctx, "app:config:debug", "app:config:timeout", "app:user:1001")
	fmt.Println("DEL app:* keys")
}

// pipelineOperations demonstrates pipeline operations
// pipelineOperations는 파이프라인 작업을 시연합니다
func pipelineOperations(ctx context.Context, rdb *redis.Client) {
	fmt.Println("Executing multiple commands in pipeline...")
	fmt.Println("파이프라인에서 여러 명령 실행 중...")

	// Execute pipeline / 파이프라인 실행
	err := rdb.Pipeline(ctx, func(pipe goredis.Pipeliner) error {
		pipe.Set(ctx, "batch:1", "value1", 0)
		pipe.Set(ctx, "batch:2", "value2", 0)
		pipe.Set(ctx, "batch:3", "value3", 0)
		pipe.Incr(ctx, "batch:counter")
		pipe.SAdd(ctx, "batch:set", "member1", "member2")
		return nil
	})

	if err != nil {
		fmt.Printf("Pipeline error: %v\n", err)
		return
	}

	fmt.Println("Pipeline executed successfully")
	fmt.Println("파이프라인 성공적으로 실행됨")

	// Verify results / 결과 확인
	val1, _ := rdb.Get(ctx, "batch:1")
	counter, _ := rdb.Get(ctx, "batch:counter")
	fmt.Printf("batch:1 = '%s', batch:counter = '%s'\n", val1, counter)

	// Cleanup / 정리
	rdb.Del(ctx, "batch:1", "batch:2", "batch:3", "batch:counter", "batch:set")
}

// transactionOperations demonstrates transaction operations
// transactionOperations는 트랜잭션 작업을 시연합니다
func transactionOperations(ctx context.Context, rdb *redis.Client) {
	fmt.Println("Executing transaction with optimistic locking...")
	fmt.Println("낙관적 잠금을 사용한 트랜잭션 실행 중...")

	// Initialize counter / 카운터 초기화
	rdb.Set(ctx, "tx:counter", "10")

	// Execute transaction / 트랜잭션 실행
	err := rdb.Transaction(ctx, func(tx *redis.Tx) error {
		// Get current value / 현재 값 가져오기
		val, err := tx.Get(ctx, "tx:counter")
		if err != nil {
			return err
		}

		fmt.Printf("Current counter value: %s\n", val)

		// Execute commands atomically / 명령을 원자적으로 실행
		return tx.Exec(ctx, func(pipe goredis.Pipeliner) error {
			pipe.Incr(ctx, "tx:counter")
			pipe.Set(ctx, "tx:last_update", time.Now().Unix(), 0)
			return nil
		})
	}, "tx:counter") // Watch the counter key / counter 키 감시

	if err != nil {
		fmt.Printf("Transaction error: %v\n", err)
		return
	}

	// Verify result / 결과 확인
	newVal, _ := rdb.Get(ctx, "tx:counter")
	fmt.Printf("New counter value: %s\n", newVal)
	fmt.Println("Transaction completed successfully")
	fmt.Println("트랜잭션 성공적으로 완료됨")

	// Cleanup / 정리
	rdb.Del(ctx, "tx:counter", "tx:last_update")
}
