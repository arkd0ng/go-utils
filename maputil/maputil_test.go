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

	if Version != "v1.8.017" {
		t.Errorf("Expected version 'v1.8.017', got '%s'", Version)
	}
}

// Note: Benchmarks for individual functions are in their respective *_test.go files
// 개별 함수의 벤치마크는 각 *_test.go 파일에 있습니다
