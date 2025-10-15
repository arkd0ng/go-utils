package maputil

import (
	"testing"
)

func TestGet(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// Test existing key / 존재하는 키 테스트
	value, ok := Get(m, "a")
	if !ok || value != 1 {
		t.Errorf("Expected (1, true), got (%d, %v)", value, ok)
	}

	// Test non-existing key / 존재하지 않는 키 테스트
	value, ok = Get(m, "d")
	if ok || value != 0 {
		t.Errorf("Expected (0, false), got (%d, %v)", value, ok)
	}
}

func TestGetOr(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}

	// Test existing key / 존재하는 키 테스트
	value := GetOr(m, "a", 10)
	if value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}

	// Test non-existing key / 존재하지 않는 키 테스트
	value = GetOr(m, "c", 10)
	if value != 10 {
		t.Errorf("Expected 10, got %d", value)
	}
}

func TestSet(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}

	// Add new key / 새 키 추가
	result := Set(m, "c", 3)
	if len(result) != 3 {
		t.Errorf("Expected length 3, got %d", len(result))
	}
	if result["c"] != 3 {
		t.Errorf("Expected c=3, got %d", result["c"])
	}

	// Original map should be unchanged / 원본 맵은 변경되지 않아야 함
	if len(m) != 2 {
		t.Errorf("Original map should have length 2, got %d", len(m))
	}

	// Update existing key / 기존 키 업데이트
	result = Set(m, "a", 10)
	if result["a"] != 10 {
		t.Errorf("Expected a=10, got %d", result["a"])
	}
}

func TestDelete(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Delete existing keys / 존재하는 키 삭제
	result := Delete(m, "b", "d")
	if len(result) != 2 {
		t.Errorf("Expected length 2, got %d", len(result))
	}
	if _, exists := result["b"]; exists {
		t.Error("Key 'b' should be deleted")
	}
	if _, exists := result["d"]; exists {
		t.Error("Key 'd' should be deleted")
	}

	// Original map should be unchanged / 원본 맵은 변경되지 않아야 함
	if len(m) != 4 {
		t.Errorf("Original map should have length 4, got %d", len(m))
	}

	// Delete non-existing keys / 존재하지 않는 키 삭제
	result = Delete(m, "x", "y")
	if len(result) != 4 {
		t.Errorf("Expected length 4, got %d", len(result))
	}

	// Delete with no keys / 키 없이 삭제
	result = Delete(m)
	if len(result) != 4 {
		t.Errorf("Expected length 4, got %d", len(result))
	}
}

func TestHas(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}

	if !Has(m, "a") {
		t.Error("Expected Has('a') to be true")
	}

	if Has(m, "c") {
		t.Error("Expected Has('c') to be false")
	}
}

func TestIsEmpty(t *testing.T) {
	empty := map[string]int{}
	nonEmpty := map[string]int{"a": 1}

	if !IsEmpty(empty) {
		t.Error("Expected IsEmpty to be true for empty map")
	}

	if IsEmpty(nonEmpty) {
		t.Error("Expected IsEmpty to be false for non-empty map")
	}
}

func TestIsNotEmpty(t *testing.T) {
	empty := map[string]int{}
	nonEmpty := map[string]int{"a": 1}

	if IsNotEmpty(empty) {
		t.Error("Expected IsNotEmpty to be false for empty map")
	}

	if !IsNotEmpty(nonEmpty) {
		t.Error("Expected IsNotEmpty to be true for non-empty map")
	}
}

func TestLen(t *testing.T) {
	m1 := map[string]int{}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}

	if Len(m1) != 0 {
		t.Errorf("Expected length 0, got %d", Len(m1))
	}

	if Len(m2) != 3 {
		t.Errorf("Expected length 3, got %d", Len(m2))
	}
}

func TestClear(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	result := Clear(m)
	if len(result) != 0 {
		t.Errorf("Expected empty map, got length %d", len(result))
	}

	// Original map should be unchanged / 원본 맵은 변경되지 않아야 함
	if len(m) != 3 {
		t.Errorf("Original map should have length 3, got %d", len(m))
	}
}

func TestClone(t *testing.T) {
	// Test nil map / nil 맵 테스트
	var nilMap map[string]int
	clonedNil := Clone(nilMap)
	if clonedNil != nil {
		t.Error("Clone of nil map should be nil")
	}

	// Test non-nil map / 비nil 맵 테스트
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	cloned := Clone(m)

	if len(cloned) != len(m) {
		t.Errorf("Expected length %d, got %d", len(m), len(cloned))
	}

	// Verify all entries / 모든 항목 확인
	for k, v := range m {
		if cloned[k] != v {
			t.Errorf("Expected %s=%d, got %d", k, v, cloned[k])
		}
	}

	// Modify clone / 복제본 수정
	cloned["d"] = 4
	if len(m) != 3 {
		t.Error("Original map should not be modified")
	}
}

func TestEqual(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	m3 := map[string]int{"a": 1, "b": 2}
	m4 := map[string]int{"a": 1, "b": 3, "c": 3}

	// Equal maps / 동일한 맵
	if !Equal(m1, m2) {
		t.Error("Expected maps to be equal")
	}

	// Different lengths / 다른 길이
	if Equal(m1, m3) {
		t.Error("Expected maps to be different (different lengths)")
	}

	// Different values / 다른 값
	if Equal(m1, m4) {
		t.Error("Expected maps to be different (different values)")
	}

	// Empty maps / 빈 맵
	empty1 := map[string]int{}
	empty2 := map[string]int{}
	if !Equal(empty1, empty2) {
		t.Error("Expected empty maps to be equal")
	}
}

// Benchmark tests / 벤치마크 테스트

func BenchmarkGet(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Get(m, "a")
	}
}

func BenchmarkSet(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Set(m, "c", 3)
	}
}

func BenchmarkDelete(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Delete(m, "b", "d")
	}
}

func BenchmarkCloneBasic(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Clone(m)
	}
}

func BenchmarkEqual(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Equal(m1, m2)
	}
}
