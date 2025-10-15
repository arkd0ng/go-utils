package sliceutil

import (
	"testing"
)

// TestFindIndices tests the FindIndices function.
// TestFindIndices는 FindIndices 함수를 테스트합니다.
func TestFindIndices(t *testing.T) {
	t.Run("find even number indices", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		indices := FindIndices(numbers, func(n int) bool { return n%2 == 0 })
		expected := []int{1, 3, 5}
		if !Equal(indices, expected) {
			t.Errorf("FindIndices() = %v, want %v", indices, expected)
		}
	})

	t.Run("find all", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}
		indices := FindIndices(numbers, func(n int) bool { return n%2 == 0 })
		expected := []int{0, 1, 2, 3}
		if !Equal(indices, expected) {
			t.Errorf("FindIndices() = %v, want %v", indices, expected)
		}
	})

	t.Run("find none", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7}
		indices := FindIndices(numbers, func(n int) bool { return n%2 == 0 })
		if len(indices) != 0 {
			t.Errorf("FindIndices() should return empty slice, got %v", indices)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		indices := FindIndices(numbers, func(n int) bool { return n > 0 })
		if len(indices) != 0 {
			t.Errorf("FindIndices() on empty slice should return empty, got %v", indices)
		}
	})

	t.Run("strings", func(t *testing.T) {
		words := []string{"apple", "apricot", "banana", "avocado"}
		indices := FindIndices(words, func(s string) bool { return s[0] == 'a' })
		expected := []int{0, 1, 3}
		if !Equal(indices, expected) {
			t.Errorf("FindIndices() = %v, want %v", indices, expected)
		}
	})

	t.Run("first element matches", func(t *testing.T) {
		numbers := []int{10, 1, 2, 3}
		indices := FindIndices(numbers, func(n int) bool { return n > 5 })
		expected := []int{0}
		if !Equal(indices, expected) {
			t.Errorf("FindIndices() = %v, want %v", indices, expected)
		}
	})

	t.Run("last element matches", func(t *testing.T) {
		numbers := []int{1, 2, 3, 10}
		indices := FindIndices(numbers, func(n int) bool { return n > 5 })
		expected := []int{3}
		if !Equal(indices, expected) {
			t.Errorf("FindIndices() = %v, want %v", indices, expected)
		}
	})
}

// BenchmarkFindIndices benchmarks the FindIndices function.
// BenchmarkFindIndices는 FindIndices 함수를 벤치마크합니다.
func BenchmarkFindIndices(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindIndices(numbers, func(n int) bool { return n%2 == 0 })
	}
}

// TestAtIndices tests the AtIndices function.
// TestAtIndices는 AtIndices 함수를 테스트합니다.
func TestAtIndices(t *testing.T) {
	t.Run("select specific indices", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}
		selected := AtIndices(numbers, []int{0, 2, 4})
		expected := []int{10, 30, 50}
		if !Equal(selected, expected) {
			t.Errorf("AtIndices() = %v, want %v", selected, expected)
		}
	})

	t.Run("out of bounds indices skipped", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}
		selected := AtIndices(numbers, []int{0, 10, 2})
		expected := []int{10, 30}
		if !Equal(selected, expected) {
			t.Errorf("AtIndices() = %v, want %v", selected, expected)
		}
	})

	t.Run("negative indices skipped", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}
		selected := AtIndices(numbers, []int{-1, 0, 2})
		expected := []int{10, 30}
		if !Equal(selected, expected) {
			t.Errorf("AtIndices() = %v, want %v", selected, expected)
		}
	})

	t.Run("empty indices", func(t *testing.T) {
		numbers := []int{10, 20, 30}
		selected := AtIndices(numbers, []int{})
		if len(selected) != 0 {
			t.Errorf("AtIndices() with empty indices should return empty, got %v", selected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		selected := AtIndices(numbers, []int{0, 1, 2})
		if len(selected) != 0 {
			t.Errorf("AtIndices() on empty slice should return empty, got %v", selected)
		}
	})

	t.Run("duplicate indices", func(t *testing.T) {
		numbers := []int{10, 20, 30}
		selected := AtIndices(numbers, []int{0, 0, 1, 1})
		expected := []int{10, 10, 20, 20}
		if !Equal(selected, expected) {
			t.Errorf("AtIndices() = %v, want %v", selected, expected)
		}
	})

	t.Run("unordered indices", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}
		selected := AtIndices(numbers, []int{4, 0, 2})
		expected := []int{50, 10, 30}
		if !Equal(selected, expected) {
			t.Errorf("AtIndices() = %v, want %v", selected, expected)
		}
	})

	t.Run("strings", func(t *testing.T) {
		words := []string{"a", "b", "c", "d", "e"}
		selected := AtIndices(words, []int{1, 3})
		expected := []string{"b", "d"}
		if !Equal(selected, expected) {
			t.Errorf("AtIndices() = %v, want %v", selected, expected)
		}
	})
}

// BenchmarkAtIndices benchmarks the AtIndices function.
// BenchmarkAtIndices는 AtIndices 함수를 벤치마크합니다.
func BenchmarkAtIndices(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	indices := []int{0, 10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AtIndices(numbers, indices)
	}
}

// TestRemoveIndices tests the RemoveIndices function.
// TestRemoveIndices는 RemoveIndices 함수를 테스트합니다.
func TestRemoveIndices(t *testing.T) {
	t.Run("remove specific indices", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}
		result := RemoveIndices(numbers, []int{1, 3})
		expected := []int{10, 30, 50}
		if !Equal(result, expected) {
			t.Errorf("RemoveIndices() = %v, want %v", result, expected)
		}
	})

	t.Run("original not modified", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}
		original := Clone(numbers)
		RemoveIndices(numbers, []int{1, 3})
		if !Equal(numbers, original) {
			t.Error("RemoveIndices() should not modify original slice")
		}
	})

	t.Run("out of bounds indices skipped", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}
		result := RemoveIndices(numbers, []int{1, 10, 3})
		expected := []int{10, 30, 50}
		if !Equal(result, expected) {
			t.Errorf("RemoveIndices() = %v, want %v", result, expected)
		}
	})

	t.Run("negative indices skipped", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40, 50}
		result := RemoveIndices(numbers, []int{-1, 1, 3})
		expected := []int{10, 30, 50}
		if !Equal(result, expected) {
			t.Errorf("RemoveIndices() = %v, want %v", result, expected)
		}
	})

	t.Run("empty indices", func(t *testing.T) {
		numbers := []int{10, 20, 30}
		result := RemoveIndices(numbers, []int{})
		if !Equal(result, numbers) {
			t.Errorf("RemoveIndices() with empty indices should return copy of slice")
		}
	})

	t.Run("remove all", func(t *testing.T) {
		numbers := []int{10, 20, 30}
		result := RemoveIndices(numbers, []int{0, 1, 2})
		if len(result) != 0 {
			t.Errorf("RemoveIndices() removing all should return empty, got %v", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := RemoveIndices(numbers, []int{0, 1, 2})
		if len(result) != 0 {
			t.Errorf("RemoveIndices() on empty slice should return empty, got %v", result)
		}
	})

	t.Run("duplicate indices", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40}
		result := RemoveIndices(numbers, []int{1, 1, 2})
		expected := []int{10, 40}
		if !Equal(result, expected) {
			t.Errorf("RemoveIndices() = %v, want %v", result, expected)
		}
	})

	t.Run("strings", func(t *testing.T) {
		words := []string{"a", "b", "c", "d", "e"}
		result := RemoveIndices(words, []int{0, 2, 4})
		expected := []string{"b", "d"}
		if !Equal(result, expected) {
			t.Errorf("RemoveIndices() = %v, want %v", result, expected)
		}
	})

	t.Run("remove first", func(t *testing.T) {
		numbers := []int{10, 20, 30}
		result := RemoveIndices(numbers, []int{0})
		expected := []int{20, 30}
		if !Equal(result, expected) {
			t.Errorf("RemoveIndices() = %v, want %v", result, expected)
		}
	})

	t.Run("remove last", func(t *testing.T) {
		numbers := []int{10, 20, 30}
		result := RemoveIndices(numbers, []int{2})
		expected := []int{10, 20}
		if !Equal(result, expected) {
			t.Errorf("RemoveIndices() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkRemoveIndices benchmarks the RemoveIndices function.
// BenchmarkRemoveIndices는 RemoveIndices 함수를 벤치마크합니다.
func BenchmarkRemoveIndices(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	indices := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveIndices(numbers, indices)
	}
}
