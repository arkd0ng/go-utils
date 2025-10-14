package sliceutil

import (
	"testing"
)

// TestChunk tests the Chunk function.
// TestChunk는 Chunk 함수를 테스트합니다.
func TestChunk(t *testing.T) {
	t.Run("chunk by 3", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		result := Chunk(numbers, 3)
		if len(result) != 3 {
			t.Errorf("Chunk() should have 3 chunks, got %v", len(result))
		}
		if !Equal(result[0], []int{1, 2, 3}) {
			t.Errorf("Chunk() chunk 0 = %v, want [1, 2, 3]", result[0])
		}
		if !Equal(result[1], []int{4, 5, 6}) {
			t.Errorf("Chunk() chunk 1 = %v, want [4, 5, 6]", result[1])
		}
		if !Equal(result[2], []int{7}) {
			t.Errorf("Chunk() chunk 2 = %v, want [7]", result[2])
		}
	})

	t.Run("chunk by 2", func(t *testing.T) {
		words := []string{"a", "b", "c", "d", "e"}
		result := Chunk(words, 2)
		if len(result) != 3 {
			t.Errorf("Chunk() should have 3 chunks, got %v", len(result))
		}
	})

	t.Run("chunk exact division", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result := Chunk(numbers, 2)
		if len(result) != 3 {
			t.Errorf("Chunk() should have 3 chunks, got %v", len(result))
		}
	})

	t.Run("chunk size larger than slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Chunk(numbers, 10)
		if len(result) != 1 {
			t.Errorf("Chunk() should have 1 chunk, got %v", len(result))
		}
		if !Equal(result[0], numbers) {
			t.Errorf("Chunk() chunk 0 = %v, want %v", result[0], numbers)
		}
	})

	t.Run("chunk size 1", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Chunk(numbers, 1)
		if len(result) != 3 {
			t.Errorf("Chunk() should have 3 chunks, got %v", len(result))
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Chunk(numbers, 3)
		if len(result) != 0 {
			t.Errorf("Chunk() with empty slice should return empty, got %v", result)
		}
	})

	t.Run("invalid size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Chunk(numbers, 0)
		if len(result) != 0 {
			t.Errorf("Chunk() with size 0 should return empty, got %v", result)
		}
	})
}

// BenchmarkChunk benchmarks the Chunk function.
// BenchmarkChunk는 Chunk 함수를 벤치마크합니다.
func BenchmarkChunk(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Chunk(numbers, 10)
	}
}

// TestTake tests the Take function.
// TestTake는 Take 함수를 테스트합니다.
func TestTake(t *testing.T) {
	t.Run("take first 3", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Take(numbers, 3)
		expected := []int{1, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("Take() = %v, want %v", result, expected)
		}
	})

	t.Run("take more than length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Take(numbers, 10)
		if !Equal(result, numbers) {
			t.Errorf("Take() = %v, want %v", result, numbers)
		}
	})

	t.Run("take 0", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Take(numbers, 0)
		if len(result) != 0 {
			t.Errorf("Take(0) should return empty slice, got %v", result)
		}
	})

	t.Run("take negative", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Take(numbers, -1)
		if len(result) != 0 {
			t.Errorf("Take(negative) should return empty slice, got %v", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Take(numbers, 3)
		if len(result) != 0 {
			t.Errorf("Take() with empty slice should return empty, got %v", result)
		}
	})
}

// BenchmarkTake benchmarks the Take function.
// BenchmarkTake는 Take 함수를 벤치마크합니다.
func BenchmarkTake(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Take(numbers, 100)
	}
}

// TestTakeLast tests the TakeLast function.
// TestTakeLast는 TakeLast 함수를 테스트합니다.
func TestTakeLast(t *testing.T) {
	t.Run("take last 3", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := TakeLast(numbers, 3)
		expected := []int{3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("TakeLast() = %v, want %v", result, expected)
		}
	})

	t.Run("take last more than length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := TakeLast(numbers, 10)
		if !Equal(result, numbers) {
			t.Errorf("TakeLast() = %v, want %v", result, numbers)
		}
	})

	t.Run("take last 0", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := TakeLast(numbers, 0)
		if len(result) != 0 {
			t.Errorf("TakeLast(0) should return empty slice, got %v", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := TakeLast(numbers, 3)
		if len(result) != 0 {
			t.Errorf("TakeLast() with empty slice should return empty, got %v", result)
		}
	})
}

// BenchmarkTakeLast benchmarks the TakeLast function.
// BenchmarkTakeLast는 TakeLast 함수를 벤치마크합니다.
func BenchmarkTakeLast(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TakeLast(numbers, 100)
	}
}

// TestDrop tests the Drop function.
// TestDrop은 Drop 함수를 테스트합니다.
func TestDrop(t *testing.T) {
	t.Run("drop first 2", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Drop(numbers, 2)
		expected := []int{3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Drop() = %v, want %v", result, expected)
		}
	})

	t.Run("drop more than length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Drop(numbers, 10)
		if len(result) != 0 {
			t.Errorf("Drop() should return empty slice, got %v", result)
		}
	})

	t.Run("drop 0", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Drop(numbers, 0)
		if !Equal(result, numbers) {
			t.Errorf("Drop(0) = %v, want %v", result, numbers)
		}
	})

	t.Run("drop negative", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Drop(numbers, -1)
		if !Equal(result, numbers) {
			t.Errorf("Drop(negative) = %v, want %v", result, numbers)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Drop(numbers, 3)
		if len(result) != 0 {
			t.Errorf("Drop() with empty slice should return empty, got %v", result)
		}
	})
}

// BenchmarkDrop benchmarks the Drop function.
// BenchmarkDrop은 Drop 함수를 벤치마크합니다.
func BenchmarkDrop(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Drop(numbers, 100)
	}
}

// TestDropLast tests the DropLast function.
// TestDropLast는 DropLast 함수를 테스트합니다.
func TestDropLast(t *testing.T) {
	t.Run("drop last 2", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := DropLast(numbers, 2)
		expected := []int{1, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("DropLast() = %v, want %v", result, expected)
		}
	})

	t.Run("drop last more than length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := DropLast(numbers, 10)
		if len(result) != 0 {
			t.Errorf("DropLast() should return empty slice, got %v", result)
		}
	})

	t.Run("drop last 0", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := DropLast(numbers, 0)
		if !Equal(result, numbers) {
			t.Errorf("DropLast(0) = %v, want %v", result, numbers)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := DropLast(numbers, 3)
		if len(result) != 0 {
			t.Errorf("DropLast() with empty slice should return empty, got %v", result)
		}
	})
}

// BenchmarkDropLast benchmarks the DropLast function.
// BenchmarkDropLast는 DropLast 함수를 벤치마크합니다.
func BenchmarkDropLast(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DropLast(numbers, 100)
	}
}

// TestSlice tests the Slice function.
// TestSlice는 Slice 함수를 테스트합니다.
func TestSlice(t *testing.T) {
	t.Run("slice middle", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Slice(numbers, 1, 4)
		expected := []int{2, 3, 4}
		if !Equal(result, expected) {
			t.Errorf("Slice() = %v, want %v", result, expected)
		}
	})

	t.Run("slice with negative indices", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Slice(numbers, -3, -1)
		expected := []int{3, 4}
		if !Equal(result, expected) {
			t.Errorf("Slice() = %v, want %v", result, expected)
		}
	})

	t.Run("slice entire", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Slice(numbers, 0, 3)
		if !Equal(result, numbers) {
			t.Errorf("Slice() = %v, want %v", result, numbers)
		}
	})

	t.Run("slice beyond bounds", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Slice(numbers, 0, 10)
		if !Equal(result, numbers) {
			t.Errorf("Slice() = %v, want %v", result, numbers)
		}
	})

	t.Run("slice invalid range", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Slice(numbers, 5, 2)
		if len(result) != 0 {
			t.Errorf("Slice() with invalid range should return empty, got %v", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Slice(numbers, 0, 5)
		if len(result) != 0 {
			t.Errorf("Slice() with empty slice should return empty, got %v", result)
		}
	})
}

// BenchmarkSlice benchmarks the Slice function.
// BenchmarkSlice는 Slice 함수를 벤치마크합니다.
func BenchmarkSlice(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slice(numbers, 100, 200)
	}
}

// TestSample tests the Sample function.
// TestSample은 Sample 함수를 테스트합니다.
func TestSample(t *testing.T) {
	t.Run("sample 3 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := Sample(numbers, 3)
		if len(result) != 3 {
			t.Errorf("Sample() should return 3 elements, got %v", len(result))
		}
		// Check all elements are from original slice
		for _, v := range result {
			if !Contains(numbers, v) {
				t.Errorf("Sample() contains element not in original slice: %v", v)
			}
		}
		// Check no duplicates
		if len(Unique(result)) != len(result) {
			t.Errorf("Sample() should not have duplicates: %v", result)
		}
	})

	t.Run("sample more than length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sample(numbers, 10)
		if len(result) != 3 {
			t.Errorf("Sample() should return all elements, got %v", len(result))
		}
		// Should contain all original elements
		for _, v := range numbers {
			if !Contains(result, v) {
				t.Errorf("Sample() missing element: %v", v)
			}
		}
	})

	t.Run("sample 0", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sample(numbers, 0)
		if len(result) != 0 {
			t.Errorf("Sample(0) should return empty slice, got %v", result)
		}
	})

	t.Run("sample negative", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sample(numbers, -1)
		if len(result) != 0 {
			t.Errorf("Sample(negative) should return empty slice, got %v", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Sample(numbers, 3)
		if len(result) != 0 {
			t.Errorf("Sample() with empty slice should return empty, got %v", result)
		}
	})
}

// BenchmarkSample benchmarks the Sample function.
// BenchmarkSample은 Sample 함수를 벤치마크합니다.
func BenchmarkSample(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sample(numbers, 100)
	}
}
