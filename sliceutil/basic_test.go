package sliceutil

import (
	"testing"
)

// TestContains tests the Contains function.
// TestContains는 Contains 함수를 테스트합니다.
func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		item     int
		expected bool
	}{
		{"found in middle", []int{1, 2, 3, 4, 5}, 3, true},
		{"found at start", []int{1, 2, 3, 4, 5}, 1, true},
		{"found at end", []int{1, 2, 3, 4, 5}, 5, true},
		{"not found", []int{1, 2, 3, 4, 5}, 10, false},
		{"empty slice", []int{}, 1, false},
		{"nil slice", nil, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.slice, tt.item)
			if result != tt.expected {
				t.Errorf("Contains(%v, %d) = %v, want %v", tt.slice, tt.item, result, tt.expected)
			}
		})
	}

	// Test with strings / 문자열로 테스트
	t.Run("with strings", func(t *testing.T) {
		fruits := []string{"apple", "banana", "cherry"}
		if !Contains(fruits, "banana") {
			t.Error("Expected to find 'banana'")
		}
		if Contains(fruits, "grape") {
			t.Error("Expected not to find 'grape'")
		}
	})
}

// TestContainsFunc tests the ContainsFunc function.
// TestContainsFunc는 ContainsFunc 함수를 테스트합니다.
func TestContainsFunc(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	tests := []struct {
		name      string
		predicate func(int) bool
		expected  bool
	}{
		{
			"has even number",
			func(n int) bool { return n%2 == 0 },
			true,
		},
		{
			"has number > 3",
			func(n int) bool { return n > 3 },
			true,
		},
		{
			"has number > 10",
			func(n int) bool { return n > 10 },
			false,
		},
		{
			"has negative number",
			func(n int) bool { return n < 0 },
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsFunc(numbers, tt.predicate)
			if result != tt.expected {
				t.Errorf("ContainsFunc() = %v, want %v", result, tt.expected)
			}
		})
	}

	// Test with empty slice / 빈 슬라이스로 테스트
	t.Run("empty slice", func(t *testing.T) {
		empty := []int{}
		result := ContainsFunc(empty, func(n int) bool { return true })
		if result {
			t.Error("Expected false for empty slice")
		}
	})
}

// TestIndexOf tests the IndexOf function.
// TestIndexOf는 IndexOf 함수를 테스트합니다.
func TestIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		item     int
		expected int
	}{
		{"found at index 0", []int{1, 2, 3, 4, 5}, 1, 0},
		{"found at index 2", []int{1, 2, 3, 4, 5}, 3, 2},
		{"found at index 4", []int{1, 2, 3, 4, 5}, 5, 4},
		{"not found", []int{1, 2, 3, 4, 5}, 10, -1},
		{"empty slice", []int{}, 1, -1},
		{"nil slice", nil, 1, -1},
		{"duplicate items", []int{1, 2, 3, 2, 5}, 2, 1}, // Returns first occurrence
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IndexOf(tt.slice, tt.item)
			if result != tt.expected {
				t.Errorf("IndexOf(%v, %d) = %d, want %d", tt.slice, tt.item, result, tt.expected)
			}
		})
	}

	// Test with strings / 문자열로 테스트
	t.Run("with strings", func(t *testing.T) {
		fruits := []string{"apple", "banana", "cherry", "banana"}
		index := IndexOf(fruits, "banana")
		if index != 1 {
			t.Errorf("Expected index 1, got %d", index)
		}
	})
}

// TestLastIndexOf tests the LastIndexOf function.
// TestLastIndexOf는 LastIndexOf 함수를 테스트합니다.
func TestLastIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		item     int
		expected int
	}{
		{"found at last index", []int{1, 2, 3, 4, 5}, 5, 4},
		{"found at middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"not found", []int{1, 2, 3, 4, 5}, 10, -1},
		{"empty slice", []int{}, 1, -1},
		{"nil slice", nil, 1, -1},
		{"duplicate items", []int{1, 2, 3, 2, 5}, 2, 3}, // Returns last occurrence
		{"all same items", []int{2, 2, 2, 2}, 2, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastIndexOf(tt.slice, tt.item)
			if result != tt.expected {
				t.Errorf("LastIndexOf(%v, %d) = %d, want %d", tt.slice, tt.item, result, tt.expected)
			}
		})
	}

	// Test with strings / 문자열로 테스트
	t.Run("with strings", func(t *testing.T) {
		fruits := []string{"apple", "banana", "cherry", "banana"}
		index := LastIndexOf(fruits, "banana")
		if index != 3 {
			t.Errorf("Expected index 3, got %d", index)
		}
	})
}

// TestFind tests the Find function.
// TestFind는 Find 함수를 테스트합니다.
func TestFind(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	tests := []struct {
		name          string
		predicate     func(int) bool
		expectedValue int
		expectedFound bool
	}{
		{
			"find even number",
			func(n int) bool { return n%2 == 0 },
			2,
			true,
		},
		{
			"find number > 3",
			func(n int) bool { return n > 3 },
			4,
			true,
		},
		{
			"find number > 10",
			func(n int) bool { return n > 10 },
			0,
			false,
		},
		{
			"find negative number",
			func(n int) bool { return n < 0 },
			0,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := Find(numbers, tt.predicate)
			if found != tt.expectedFound {
				t.Errorf("Find() found = %v, want %v", found, tt.expectedFound)
			}
			if found && value != tt.expectedValue {
				t.Errorf("Find() value = %d, want %d", value, tt.expectedValue)
			}
		})
	}

	// Test with empty slice / 빈 슬라이스로 테스트
	t.Run("empty slice", func(t *testing.T) {
		empty := []int{}
		_, found := Find(empty, func(n int) bool { return true })
		if found {
			t.Error("Expected not found for empty slice")
		}
	})

	// Test with strings / 문자열로 테스트
	t.Run("with strings", func(t *testing.T) {
		fruits := []string{"apple", "banana", "cherry"}
		fruit, found := Find(fruits, func(s string) bool { return len(s) > 5 })
		if !found || fruit != "banana" {
			t.Errorf("Expected to find 'banana', got '%s', found=%v", fruit, found)
		}
	})
}

// Benchmark tests / 벤치마크 테스트

func BenchmarkContains(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(slice, 500)
	}
}

func BenchmarkContainsFunc(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsFunc(slice, func(n int) bool { return n == 500 })
	}
}

func BenchmarkIndexOf(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IndexOf(slice, 500)
	}
}

func BenchmarkLastIndexOf(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LastIndexOf(slice, 500)
	}
}

func BenchmarkFind(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Find(slice, func(n int) bool { return n == 500 })
	}
}
