package redis

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

// TestNew tests creating a new Redis client
// TestNew는 새로운 Redis 클라이언트 생성을 테스트합니다
func TestNew(t *testing.T) {
	// Skip if Redis is not available / Redis가 사용 불가능하면 건너뜀
	client, err := New(WithAddr("localhost:6379"))
	if err != nil {
		t.Skip("Redis is not available, skipping test")
		return
	}
	defer client.Close()

	// Test ping / Ping 테스트
	ctx := context.Background()
	if err := client.Ping(ctx); err != nil {
		t.Fatalf("Failed to ping Redis: %v", err)
	}
}

// TestStringOperations tests string operations
// TestStringOperations는 문자열 작업을 테스트합니다
func TestStringOperations(t *testing.T) {
	client, err := New(WithAddr("localhost:6379"))
	if err != nil {
		t.Skip("Redis is not available, skipping test")
		return
	}
	defer client.Close()

	ctx := context.Background()
	key := "test:string:key"

	// Set / 설정
	if err := client.Set(ctx, key, "hello"); err != nil {
		t.Fatalf("Failed to set key: %v", err)
	}

	// Get / 가져오기
	val, err := client.Get(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get key: %v", err)
	}
	if val != "hello" {
		t.Errorf("Expected 'hello', got '%s'", val)
	}

	// Delete / 삭제
	if err := client.Del(ctx, key); err != nil {
		t.Fatalf("Failed to delete key: %v", err)
	}
}

// TestHashOperations tests hash operations
// TestHashOperations는 해시 작업을 테스트합니다
func TestHashOperations(t *testing.T) {
	client, err := New(WithAddr("localhost:6379"))
	if err != nil {
		t.Skip("Redis is not available, skipping test")
		return
	}
	defer client.Close()

	ctx := context.Background()
	key := "test:hash:key"

	// HSet / 해시 설정
	if err := client.HSet(ctx, key, "field1", "value1"); err != nil {
		t.Fatalf("Failed to set hash field: %v", err)
	}

	// HGet / 해시 가져오기
	val, err := client.HGet(ctx, key, "field1")
	if err != nil {
		t.Fatalf("Failed to get hash field: %v", err)
	}
	if val != "value1" {
		t.Errorf("Expected 'value1', got '%s'", val)
	}

	// HGetAll / 모든 해시 가져오기
	fields := map[string]interface{}{
		"field2": "value2",
		"field3": "value3",
	}
	if err := client.HSetMap(ctx, key, fields); err != nil {
		t.Fatalf("Failed to set hash fields: %v", err)
	}

	allFields, err := client.HGetAll(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get all hash fields: %v", err)
	}
	if len(allFields) != 3 {
		t.Errorf("Expected 3 fields, got %d", len(allFields))
	}

	// Cleanup / 정리
	if err := client.Del(ctx, key); err != nil {
		t.Fatalf("Failed to delete key: %v", err)
	}
}

// TestListOperations tests list operations
// TestListOperations는 리스트 작업을 테스트합니다
func TestListOperations(t *testing.T) {
	client, err := New(WithAddr("localhost:6379"))
	if err != nil {
		t.Skip("Redis is not available, skipping test")
		return
	}
	defer client.Close()

	ctx := context.Background()
	key := "test:list:key"

	// RPush / 오른쪽에 추가
	if err := client.RPush(ctx, key, "item1", "item2", "item3"); err != nil {
		t.Fatalf("Failed to push to list: %v", err)
	}

	// LLen / 리스트 길이
	length, err := client.LLen(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get list length: %v", err)
	}
	if length != 3 {
		t.Errorf("Expected length 3, got %d", length)
	}

	// LRange / 범위 가져오기
	items, err := client.LRange(ctx, key, 0, -1)
	if err != nil {
		t.Fatalf("Failed to get list range: %v", err)
	}
	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}

	// Cleanup / 정리
	if err := client.Del(ctx, key); err != nil {
		t.Fatalf("Failed to delete key: %v", err)
	}
}

// TestSetOperations tests set operations
// TestSetOperations는 집합 작업을 테스트합니다
func TestSetOperations(t *testing.T) {
	client, err := New(WithAddr("localhost:6379"))
	if err != nil {
		t.Skip("Redis is not available, skipping test")
		return
	}
	defer client.Close()

	ctx := context.Background()
	key := "test:set:key"

	// SAdd / 집합에 추가
	if err := client.SAdd(ctx, key, "member1", "member2", "member3"); err != nil {
		t.Fatalf("Failed to add to set: %v", err)
	}

	// SCard / 집합 크기
	size, err := client.SCard(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get set size: %v", err)
	}
	if size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}

	// SIsMember / 멤버 확인
	exists, err := client.SIsMember(ctx, key, "member1")
	if err != nil {
		t.Fatalf("Failed to check member: %v", err)
	}
	if !exists {
		t.Error("Expected member1 to exist")
	}

	// Cleanup / 정리
	if err := client.Del(ctx, key); err != nil {
		t.Fatalf("Failed to delete key: %v", err)
	}
}

// TestExpiration tests key expiration
// TestExpiration는 키 만료를 테스트합니다
func TestExpiration(t *testing.T) {
	client, err := New(WithAddr("localhost:6379"))
	if err != nil {
		t.Skip("Redis is not available, skipping test")
		return
	}
	defer client.Close()

	ctx := context.Background()
	key := "test:expiration:key"

	// Set with expiration / 만료와 함께 설정
	if err := client.Set(ctx, key, "value", 1*time.Second); err != nil {
		t.Fatalf("Failed to set key with expiration: %v", err)
	}

	// Check TTL / TTL 확인
	ttl, err := client.TTL(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get TTL: %v", err)
	}
	if ttl <= 0 {
		t.Errorf("Expected positive TTL, got %v", ttl)
	}

	// Wait for expiration / 만료 대기
	time.Sleep(2 * time.Second)

	// Key should be gone / 키가 사라져야 함
	_, err = client.Get(ctx, key)
	if err != ErrNil {
		t.Errorf("Expected ErrNil after expiration, got %v", err)
	}
}

// TestPipeline tests pipeline operations
// TestPipeline는 파이프라인 작업을 테스트합니다
func TestPipeline(t *testing.T) {
	client, err := New(WithAddr("localhost:6379"))
	if err != nil {
		t.Skip("Redis is not available, skipping test")
		return
	}
	defer client.Close()

	ctx := context.Background()

	// Execute multiple commands in pipeline / 파이프라인에서 여러 명령 실행
	err = client.Pipeline(ctx, func(pipe redis.Pipeliner) error {
		pipe.Set(ctx, "key1", "value1", 0)
		pipe.Set(ctx, "key2", "value2", 0)
		pipe.Set(ctx, "key3", "value3", 0)
		return nil
	})

	if err != nil {
		t.Fatalf("Failed to execute pipeline: %v", err)
	}

	// Verify values / 값 확인
	val1, _ := client.Get(ctx, "key1")
	val2, _ := client.Get(ctx, "key2")
	val3, _ := client.Get(ctx, "key3")

	if val1 != "value1" || val2 != "value2" || val3 != "value3" {
		t.Error("Pipeline values don't match expected")
	}

	// Cleanup / 정리
	client.Del(ctx, "key1", "key2", "key3")
}
