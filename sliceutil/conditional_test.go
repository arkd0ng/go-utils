package sliceutil

import (
	"testing"
)

// TestReplaceIf tests the ReplaceIf function.
// TestReplaceIf는 ReplaceIf 함수를 테스트합니다.
func TestReplaceIf(t *testing.T) {
	t.Run("replace even numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result := ReplaceIf(numbers, func(n int) bool { return n%2 == 0 }, 0)
		expected := []int{1, 0, 3, 0, 5, 0}
		if !Equal(result, expected) {
			t.Errorf("ReplaceIf() = %v, want %v", result, expected)
		}
	})

	t.Run("replace nothing", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7}
		result := ReplaceIf(numbers, func(n int) bool { return n%2 == 0 }, 0)
		if !Equal(result, numbers) {
			t.Errorf("ReplaceIf() = %v, want %v", result, numbers)
		}
	})

	t.Run("replace all", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}
		result := ReplaceIf(numbers, func(n int) bool { return n%2 == 0 }, 0)
		expected := []int{0, 0, 0, 0}
		if !Equal(result, expected) {
			t.Errorf("ReplaceIf() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := ReplaceIf(numbers, func(n int) bool { return true }, 0)
		if len(result) != 0 {
			t.Errorf("ReplaceIf() on empty slice should return empty, got %v", result)
		}
	})

	t.Run("original not modified", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		original := Clone(numbers)
		ReplaceIf(numbers, func(n int) bool { return n%2 == 0 }, 0)
		if !Equal(numbers, original) {
			t.Error("ReplaceIf() should not modify original slice")
		}
	})

	t.Run("strings", func(t *testing.T) {
		words := []string{"apple", "banana", "apricot", "cherry"}
		result := ReplaceIf(words, func(s string) bool { return s[0] == 'a' }, "FRUIT")
		expected := []string{"FRUIT", "banana", "FRUIT", "cherry"}
		if !Equal(result, expected) {
			t.Errorf("ReplaceIf() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkReplaceIf benchmarks the ReplaceIf function.
// BenchmarkReplaceIf는 ReplaceIf 함수를 벤치마크합니다.
func BenchmarkReplaceIf(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReplaceIf(numbers, func(n int) bool { return n%2 == 0 }, 0)
	}
}

// TestReplaceAll tests the ReplaceAll function.
// TestReplaceAll은 ReplaceAll 함수를 테스트합니다.
func TestReplaceAll(t *testing.T) {
	t.Run("replace multiple occurrences", func(t *testing.T) {
		numbers := []int{1, 2, 3, 2, 4, 2}
		result := ReplaceAll(numbers, 2, 99)
		expected := []int{1, 99, 3, 99, 4, 99}
		if !Equal(result, expected) {
			t.Errorf("ReplaceAll() = %v, want %v", result, expected)
		}
	})

	t.Run("replace nothing", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		result := ReplaceAll(numbers, 99, 0)
		if !Equal(result, numbers) {
			t.Errorf("ReplaceAll() = %v, want %v", result, numbers)
		}
	})

	t.Run("replace all", func(t *testing.T) {
		numbers := []int{5, 5, 5, 5}
		result := ReplaceAll(numbers, 5, 0)
		expected := []int{0, 0, 0, 0}
		if !Equal(result, expected) {
			t.Errorf("ReplaceAll() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := ReplaceAll(numbers, 1, 2)
		if len(result) != 0 {
			t.Errorf("ReplaceAll() on empty slice should return empty, got %v", result)
		}
	})

	t.Run("original not modified", func(t *testing.T) {
		numbers := []int{1, 2, 3, 2}
		original := Clone(numbers)
		ReplaceAll(numbers, 2, 99)
		if !Equal(numbers, original) {
			t.Error("ReplaceAll() should not modify original slice")
		}
	})

	t.Run("strings", func(t *testing.T) {
		words := []string{"apple", "banana", "apple", "cherry", "apple"}
		result := ReplaceAll(words, "apple", "orange")
		expected := []string{"orange", "banana", "orange", "cherry", "orange"}
		if !Equal(result, expected) {
			t.Errorf("ReplaceAll() = %v, want %v", result, expected)
		}
	})

	t.Run("replace first", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := ReplaceAll(numbers, 1, 99)
		expected := []int{99, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("ReplaceAll() = %v, want %v", result, expected)
		}
	})

	t.Run("replace last", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := ReplaceAll(numbers, 3, 99)
		expected := []int{1, 2, 99}
		if !Equal(result, expected) {
			t.Errorf("ReplaceAll() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkReplaceAll benchmarks the ReplaceAll function.
// BenchmarkReplaceAll은 ReplaceAll 함수를 벤치마크합니다.
func BenchmarkReplaceAll(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i % 100
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReplaceAll(numbers, 50, 999)
	}
}

// TestUpdateWhere tests the UpdateWhere function.
// TestUpdateWhere는 UpdateWhere 함수를 테스트합니다.
func TestUpdateWhere(t *testing.T) {
	t.Run("update even numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := UpdateWhere(numbers,
			func(n int) bool { return n%2 == 0 },
			func(n int) int { return n * 10 })
		expected := []int{1, 20, 3, 40, 5}
		if !Equal(result, expected) {
			t.Errorf("UpdateWhere() = %v, want %v", result, expected)
		}
	})

	t.Run("update nothing", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7}
		result := UpdateWhere(numbers,
			func(n int) bool { return n%2 == 0 },
			func(n int) int { return n * 10 })
		if !Equal(result, numbers) {
			t.Errorf("UpdateWhere() = %v, want %v", result, numbers)
		}
	})

	t.Run("update all", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}
		result := UpdateWhere(numbers,
			func(n int) bool { return n%2 == 0 },
			func(n int) int { return n * 10 })
		expected := []int{20, 40, 60, 80}
		if !Equal(result, expected) {
			t.Errorf("UpdateWhere() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := UpdateWhere(numbers,
			func(n int) bool { return true },
			func(n int) int { return n * 10 })
		if len(result) != 0 {
			t.Errorf("UpdateWhere() on empty slice should return empty, got %v", result)
		}
	})

	t.Run("original not modified", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		original := Clone(numbers)
		UpdateWhere(numbers,
			func(n int) bool { return n%2 == 0 },
			func(n int) int { return n * 10 })
		if !Equal(numbers, original) {
			t.Error("UpdateWhere() should not modify original slice")
		}
	})

	t.Run("structs", func(t *testing.T) {
		type User struct {
			ID     int
			Active bool
		}
		users := []User{{1, false}, {2, true}, {3, false}}
		result := UpdateWhere(users,
			func(u User) bool { return !u.Active },
			func(u User) User { u.Active = true; return u })

		// Check that inactive users are now active
		for _, u := range result {
			if u.ID == 1 || u.ID == 3 {
				if !u.Active {
					t.Errorf("User %d should be active", u.ID)
				}
			}
		}
		// User 2 should still be active
		if !result[1].Active {
			t.Error("User 2 should remain active")
		}
	})

	t.Run("increment by index", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}
		result := UpdateWhere(numbers,
			func(n int) bool { return n > 25 },
			func(n int) int { return n + 100 })
		expected := []int{10, 20, 130, 140, 150}
		if !Equal(result, expected) {
			t.Errorf("UpdateWhere() = %v, want %v", result, expected)
		}
	})

	t.Run("strings to uppercase", func(t *testing.T) {
		words := []string{"apple", "BANANA", "cherry", "DATE"}
		result := UpdateWhere(words,
			func(s string) bool { return s[0] >= 'a' && s[0] <= 'z' },
			func(s string) string {
				// Convert first char to uppercase
				return string(s[0]-32) + s[1:]
			})
		expected := []string{"Apple", "BANANA", "Cherry", "DATE"}
		if !Equal(result, expected) {
			t.Errorf("UpdateWhere() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkUpdateWhere benchmarks the UpdateWhere function.
// BenchmarkUpdateWhere는 UpdateWhere 함수를 벤치마크합니다.
func BenchmarkUpdateWhere(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UpdateWhere(numbers,
			func(n int) bool { return n%2 == 0 },
			func(n int) int { return n * 10 })
	}
}
