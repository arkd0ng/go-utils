package maputil

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	result := Map(m, func(k string, v int) string {
		return fmt.Sprintf("%s=%d", k, v)
	})

	if len(result) != 3 {
		t.Errorf("Expected length 3, got %d", len(result))
	}

	// Check that all keys are present
	// 모든 키가 있는지 확인
	for k := range m {
		if _, exists := result[k]; !exists {
			t.Errorf("Key %s should exist in result", k)
		}
	}
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	result := MapKeys(m, func(k string, v int) int {
		return v * 10
	})

	if len(result) != 3 {
		t.Errorf("Expected length 3, got %d", len(result))
	}

	if result[10] != 1 || result[20] != 2 || result[30] != 3 {
		t.Errorf("Unexpected result: %v", result)
	}
}

func TestMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	result := MapValues(m, func(v int) int {
		return v * 2
	})

	if len(result) != 3 {
		t.Errorf("Expected length 3, got %d", len(result))
	}

	if result["a"] != 2 || result["b"] != 4 || result["c"] != 6 {
		t.Errorf("Unexpected result: %v", result)
	}
}

func TestMapEntries(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}

	result := MapEntries(m, func(k string, v int) (int, string) {
		return v, k
	})

	if len(result) != 2 {
		t.Errorf("Expected length 2, got %d", len(result))
	}

	if result[1] != "a" || result[2] != "b" {
		t.Errorf("Unexpected result: %v", result)
	}
}

func TestInvert(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	result := Invert(m)

	if len(result) != 3 {
		t.Errorf("Expected length 3, got %d", len(result))
	}

	if result[1] != "a" || result[2] != "b" || result[3] != "c" {
		t.Errorf("Unexpected result: %v", result)
	}
}

func TestFlatten(t *testing.T) {
	m := map[string]map[string]int{
		"user1": {"age": 25, "score": 100},
		"user2": {"age": 30, "score": 95},
	}

	result := Flatten(m, ".")

	if len(result) != 4 {
		t.Errorf("Expected length 4, got %d", len(result))
	}

	if result["user1.age"] != 25 {
		t.Errorf("Expected user1.age=25, got %d", result["user1.age"])
	}

	if result["user2.score"] != 95 {
		t.Errorf("Expected user2.score=95, got %d", result["user2.score"])
	}
}

func TestUnflatten(t *testing.T) {
	m := map[string]int{
		"user.name":  1,
		"user.age":   25,
		"admin.name": 2,
	}

	result := Unflatten(m, ".")

	if len(result) != 2 {
		t.Errorf("Expected length 2, got %d", len(result))
	}

	userMap, ok := result["user"].(map[string]interface{})
	if !ok {
		t.Error("Expected user to be a map")
	}

	if userMap["age"] != 25 {
		t.Errorf("Expected user.age=25, got %v", userMap["age"])
	}
}

func TestChunk(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	// Chunk size 2
	// 청크 크기 2
	result := Chunk(m, 2)

	if len(result) != 3 {
		t.Errorf("Expected 3 chunks, got %d", len(result))
	}

	// Chunk size 0 (invalid)
	// 청크 크기 0 (유효하지 않음)
	result = Chunk(m, 0)
	if result != nil {
		t.Error("Expected nil for chunk size 0")
	}

	// Chunk size larger than map
	// 맵보다 큰 청크 크기
	result = Chunk(m, 10)
	if len(result) != 1 {
		t.Errorf("Expected 1 chunk, got %d", len(result))
	}
}

func TestPartition(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	trueMap, falseMap := Partition(m, func(k string, v int) bool {
		return v%2 == 0
	})

	if len(trueMap) != 2 {
		t.Errorf("Expected 2 even values, got %d", len(trueMap))
	}

	if len(falseMap) != 2 {
		t.Errorf("Expected 2 odd values, got %d", len(falseMap))
	}

	if trueMap["b"] != 2 || trueMap["d"] != 4 {
		t.Errorf("Unexpected true map: %v", trueMap)
	}
}

func TestCompact(t *testing.T) {
	m := map[string]int{"a": 1, "b": 0, "c": 3, "d": 0, "e": 5}

	result := Compact(m)

	if len(result) != 3 {
		t.Errorf("Expected length 3, got %d", len(result))
	}

	if _, exists := result["b"]; exists {
		t.Error("Key 'b' should be removed (zero value)")
	}

	if _, exists := result["d"]; exists {
		t.Error("Key 'd' should be removed (zero value)")
	}
}

// Benchmark tests
// 벤치마크 테스트

func BenchmarkMap(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Map(m, func(k string, v int) string {
			return fmt.Sprintf("%s=%d", k, v)
		})
	}
}

func BenchmarkMapValues(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapValues(m, func(v int) int {
			return v * 2
		})
	}
}

func BenchmarkInvert(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Invert(m)
	}
}

func BenchmarkPartition(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Partition(m, func(k string, v int) bool {
			return v%2 == 0
		})
	}
}

// Helper function for deep comparison
// 깊은 비교를 위한 헬퍼 함수
func mapEqual(m1, m2 map[string]interface{}) bool {
	return reflect.DeepEqual(m1, m2)
}
