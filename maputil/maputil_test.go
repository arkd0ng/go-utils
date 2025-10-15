package maputil

import (
	"testing"
)

// TestEntry tests the Entry type
// Entry 타입을 테스트합니다
func TestEntry(t *testing.T) {
	entry := Entry[string, int]{Key: "a", Value: 1}

	if entry.Key != "a" {
		t.Errorf("Expected key 'a', got '%s'", entry.Key)
	}

	if entry.Value != 1 {
		t.Errorf("Expected value 1, got %d", entry.Value)
	}
}

// TestVersion tests the version constant
// 버전 상수를 테스트합니다
func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}

	if Version != "v1.8.001" {
		t.Errorf("Expected version 'v1.8.001', got '%s'", Version)
	}
}

// BenchmarkClone benchmarks the Clone function
// Clone 함수를 벤치마크합니다
func BenchmarkClone(b *testing.B) {
	m := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		m[string(rune(i))] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Clone(m)
	}
}

// BenchmarkMerge benchmarks the Merge function
// Merge 함수를 벤치마크합니다
func BenchmarkMerge(b *testing.B) {
	m1 := make(map[string]int, 500)
	m2 := make(map[string]int, 500)
	for i := 0; i < 500; i++ {
		m1[string(rune(i))] = i
		m2[string(rune(i+500))] = i + 500
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Merge(m1, m2)
	}
}

// BenchmarkFilter benchmarks the Filter function
// Filter 함수를 벤치마크합니다
func BenchmarkFilter(b *testing.B) {
	m := make(map[string]int, 1000)
	for i := 0; i < 1000; i++ {
		m[string(rune(i))] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter(m, func(k string, v int) bool {
			return v%2 == 0
		})
	}
}
