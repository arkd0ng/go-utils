package maputil

import (
	"testing"
)

// TestEvery tests the Every function
// Every 함수 테스트
func TestEvery(t *testing.T) {
	data := map[string]int{"a": 2, "b": 4, "c": 6}

	// All values are even
	// 모든 값이 짝수
	result := Every(data, func(k string, v int) bool {
		return v%2 == 0
	})
	if !result {
		t.Error("Every(even) = false, want true")
	}

	// Not all values > 3
	// 모든 값이 3보다 큰 것은 아님
	result = Every(data, func(k string, v int) bool {
		return v > 3
	})
	if result {
		t.Error("Every(v > 3) = true, want false")
	}

	// Empty map should return true
	// 빈 맵은 true 반환
	empty := map[string]int{}
	result = Every(empty, func(k string, v int) bool {
		return false
	})
	if !result {
		t.Error("Every(empty) = false, want true")
	}
}

// TestSome tests the Some function
// Some 함수 테스트
func TestSome(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	// At least one even value
	// 최소 한 개의 짝수 값
	result := Some(data, func(k string, v int) bool {
		return v%2 == 0
	})
	if !result {
		t.Error("Some(even) = false, want true")
	}

	// No value > 10
	// 10보다 큰 값 없음
	result = Some(data, func(k string, v int) bool {
		return v > 10
	})
	if result {
		t.Error("Some(v > 10) = true, want false")
	}

	// Empty map should return false
	// 빈 맵은 false 반환
	empty := map[string]int{}
	result = Some(empty, func(k string, v int) bool {
		return true
	})
	if result {
		t.Error("Some(empty) = true, want false")
	}
}

// TestNone tests the None function
// None 함수 테스트
func TestNone(t *testing.T) {
	data := map[string]int{"a": 1, "b": 3, "c": 5}

	// No even values
	// 짝수 값 없음
	result := None(data, func(k string, v int) bool {
		return v%2 == 0
	})
	if !result {
		t.Error("None(even) = false, want true")
	}

	// Some values > 2
	// 일부 값이 2보다 큼
	result = None(data, func(k string, v int) bool {
		return v > 2
	})
	if result {
		t.Error("None(v > 2) = true, want false")
	}
}

// TestHasValue tests the HasValue function
// HasValue 함수 테스트
func TestHasValue(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	if !HasValue(data, 2) {
		t.Error("HasValue(2) = false, want true")
	}

	if HasValue(data, 10) {
		t.Error("HasValue(10) = true, want false")
	}

	// Empty map
	// 빈 맵
	empty := map[string]int{}
	if HasValue(empty, 1) {
		t.Error("HasValue(empty, 1) = true, want false")
	}
}

// TestHasEntry tests the HasEntry function
// HasEntry 함수 테스트
func TestHasEntry(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	if !HasEntry(data, "b", 2) {
		t.Error("HasEntry('b', 2) = false, want true")
	}

	// Key exists but value doesn't match
	// 키는 있지만 값이 다름
	if HasEntry(data, "b", 99) {
		t.Error("HasEntry('b', 99) = true, want false")
	}

	// Key doesn't exist
	// 키가 없음
	if HasEntry(data, "x", 1) {
		t.Error("HasEntry('x', 1) = true, want false")
	}
}

// TestIsSubset tests the IsSubset function
// IsSubset 함수 테스트
func TestIsSubset(t *testing.T) {
	subset := map[string]int{"a": 1, "b": 2}
	superset := map[string]int{"a": 1, "b": 2, "c": 3}

	if !IsSubset(subset, superset) {
		t.Error("IsSubset() = false, want true")
	}

	// Different values
	// 값이 다름
	notSubset := map[string]int{"a": 1, "b": 99}
	if IsSubset(notSubset, superset) {
		t.Error("IsSubset(different values) = true, want false")
	}

	// Extra key in subset
	// subset에 추가 키
	notSubset = map[string]int{"a": 1, "b": 2, "x": 10}
	if IsSubset(notSubset, superset) {
		t.Error("IsSubset(extra key) = true, want false")
	}

	// Empty subset is subset of any map
	// 빈 subset은 모든 맵의 부분집합
	empty := map[string]int{}
	if !IsSubset(empty, superset) {
		t.Error("IsSubset(empty) = false, want true")
	}

	// Same maps are subsets
	// 같은 맵은 부분집합
	if !IsSubset(superset, superset) {
		t.Error("IsSubset(same) = false, want true")
	}
}

// TestIsSuperset tests the IsSuperset function
// IsSuperset 함수 테스트
func TestIsSuperset(t *testing.T) {
	subset := map[string]int{"a": 1, "b": 2}
	superset := map[string]int{"a": 1, "b": 2, "c": 3}

	if !IsSuperset(superset, subset) {
		t.Error("IsSuperset() = false, want true")
	}

	if IsSuperset(subset, superset) {
		t.Error("IsSuperset(reversed) = true, want false")
	}

	// Any map is superset of empty map
	// 모든 맵은 빈 맵의 상위집합
	empty := map[string]int{}
	if !IsSuperset(superset, empty) {
		t.Error("IsSuperset(empty) = false, want true")
	}
}

// Benchmarks
// 벤치마크

func BenchmarkEvery(b *testing.B) {
	data := map[string]int{"a": 2, "b": 4, "c": 6, "d": 8, "e": 10}
	fn := func(k string, v int) bool { return v%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Every(data, fn)
	}
}

func BenchmarkSome(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string, v int) bool { return v%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Some(data, fn)
	}
}

func BenchmarkNone(b *testing.B) {
	data := map[string]int{"a": 1, "b": 3, "c": 5, "d": 7, "e": 9}
	fn := func(k string, v int) bool { return v%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = None(data, fn)
	}
}

func BenchmarkHasValue(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = HasValue(data, 3)
	}
}

func BenchmarkHasEntry(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = HasEntry(data, "c", 3)
	}
}

func BenchmarkIsSubset(b *testing.B) {
	subset := map[string]int{"a": 1, "b": 2, "c": 3}
	superset := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsSubset(subset, superset)
	}
}

func BenchmarkIsSuperset(b *testing.B) {
	subset := map[string]int{"a": 1, "b": 2, "c": 3}
	superset := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsSuperset(superset, subset)
	}
}
