package sliceutil

import (
	"fmt"
	"testing"
)

// TestScan tests the Scan function.
// TestScan은 Scan 함수를 테스트합니다.
func TestScan(t *testing.T) {
	t.Run("cumulative sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Scan(numbers, 0, func(acc, n int) int {
			return acc + n
		})
		expected := []int{0, 1, 3, 6, 10, 15}
		if !Equal(result, expected) {
			t.Errorf("Scan() = %v, want %v", result, expected)
		}
	})

	t.Run("running product", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Scan(numbers, 1, func(acc, n int) int {
			return acc * n
		})
		expected := []int{1, 1, 2, 6, 24, 120}
		if !Equal(result, expected) {
			t.Errorf("Scan() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Scan(numbers, 10, func(acc, n int) int {
			return acc + n
		})
		expected := []int{10}
		if !Equal(result, expected) {
			t.Errorf("Scan() = %v, want %v", result, expected)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{5}
		result := Scan(numbers, 10, func(acc, n int) int {
			return acc + n
		})
		expected := []int{10, 15}
		if !Equal(result, expected) {
			t.Errorf("Scan() = %v, want %v", result, expected)
		}
	})

	t.Run("string concatenation", func(t *testing.T) {
		words := []string{"a", "b", "c"}
		result := Scan(words, "", func(acc, s string) string {
			return acc + s
		})
		expected := []string{"", "a", "ab", "abc"}
		if !Equal(result, expected) {
			t.Errorf("Scan() = %v, want %v", result, expected)
		}
	})

	t.Run("max accumulator", func(t *testing.T) {
		numbers := []int{3, 1, 4, 1, 5, 9, 2}
		result := Scan(numbers, 0, func(acc, n int) int {
			if n > acc {
				return n
			}
			return acc
		})
		expected := []int{0, 3, 3, 4, 4, 5, 9, 9}
		if !Equal(result, expected) {
			t.Errorf("Scan() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkScan benchmarks the Scan function.
// BenchmarkScan은 Scan 함수를 벤치마크합니다.
func BenchmarkScan(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Scan(numbers, 0, func(acc, n int) int {
			return acc + n
		})
	}
}

// TestZipWith tests the ZipWith function.
// TestZipWith는 ZipWith 함수를 테스트합니다.
func TestZipWith(t *testing.T) {
	t.Run("combine numbers and strings", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		strings := []string{"a", "b", "c"}
		result := ZipWith(numbers, strings, func(n int, s string) string {
			return fmt.Sprintf("%d:%s", n, s)
		})
		expected := []string{"1:a", "2:b", "3:c"}
		if !Equal(result, expected) {
			t.Errorf("ZipWith() = %v, want %v", result, expected)
		}
	})

	t.Run("sum corresponding elements", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := []int{10, 20, 30}
		result := ZipWith(a, b, func(x, y int) int {
			return x + y
		})
		expected := []int{11, 22, 33}
		if !Equal(result, expected) {
			t.Errorf("ZipWith() = %v, want %v", result, expected)
		}
	})

	t.Run("first slice shorter", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{10, 20, 30, 40}
		result := ZipWith(a, b, func(x, y int) int {
			return x + y
		})
		expected := []int{11, 22}
		if !Equal(result, expected) {
			t.Errorf("ZipWith() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		a := []int{}
		b := []int{1, 2, 3}
		result := ZipWith(a, b, func(x, y int) int {
			return x + y
		})
		if len(result) != 0 {
			t.Errorf("ZipWith() should return empty, got %v", result)
		}
	})

	t.Run("both empty", func(t *testing.T) {
		a := []int{}
		b := []int{}
		result := ZipWith(a, b, func(x, y int) int {
			return x + y
		})
		if len(result) != 0 {
			t.Errorf("ZipWith() should return empty, got %v", result)
		}
	})

	t.Run("multiply floats", func(t *testing.T) {
		a := []float64{1.5, 2.0, 3.5}
		b := []float64{2.0, 3.0, 4.0}
		result := ZipWith(a, b, func(x, y float64) float64 {
			return x * y
		})
		expected := []float64{3.0, 6.0, 14.0}
		if len(result) != len(expected) {
			t.Fatalf("ZipWith() length = %v, want %v", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("ZipWith()[%d] = %v, want %v", i, result[i], expected[i])
			}
		}
	})
}

// BenchmarkZipWith benchmarks the ZipWith function.
// BenchmarkZipWith는 ZipWith 함수를 벤치마크합니다.
func BenchmarkZipWith(b *testing.B) {
	a := make([]int, 1000)
	b2 := make([]int, 1000)
	for i := range a {
		a[i] = i
		b2[i] = i * 2
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ZipWith(a, b2, func(x, y int) int {
			return x + y
		})
	}
}

// TestRotateLeft tests the RotateLeft function.
// TestRotateLeft는 RotateLeft 함수를 테스트합니다.
func TestRotateLeft(t *testing.T) {
	t.Run("rotate by 2", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := RotateLeft(numbers, 2)
		expected := []int{3, 4, 5, 1, 2}
		if !Equal(result, expected) {
			t.Errorf("RotateLeft() = %v, want %v", result, expected)
		}
	})

	t.Run("rotate by 0", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := RotateLeft(numbers, 0)
		if !Equal(result, numbers) {
			t.Errorf("RotateLeft(0) = %v, want %v", result, numbers)
		}
	})

	t.Run("rotate by length", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := RotateLeft(numbers, 5)
		if !Equal(result, numbers) {
			t.Errorf("RotateLeft(len) = %v, want %v", result, numbers)
		}
	})

	t.Run("rotate by more than length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := RotateLeft(numbers, 7) // 7 % 3 = 1
		expected := []int{2, 3, 1}
		if !Equal(result, expected) {
			t.Errorf("RotateLeft() = %v, want %v", result, expected)
		}
	})

	t.Run("rotate by negative (right)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := RotateLeft(numbers, -1)
		expected := []int{5, 1, 2, 3, 4}
		if !Equal(result, expected) {
			t.Errorf("RotateLeft(-1) = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := RotateLeft(numbers, 3)
		if len(result) != 0 {
			t.Errorf("RotateLeft() on empty should return empty, got %v", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		result := RotateLeft(numbers, 5)
		expected := []int{42}
		if !Equal(result, expected) {
			t.Errorf("RotateLeft() = %v, want %v", result, expected)
		}
	})

	t.Run("original not modified", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		original := Clone(numbers)
		RotateLeft(numbers, 2)
		if !Equal(numbers, original) {
			t.Error("RotateLeft() should not modify original slice")
		}
	})

	t.Run("strings", func(t *testing.T) {
		words := []string{"a", "b", "c", "d"}
		result := RotateLeft(words, 1)
		expected := []string{"b", "c", "d", "a"}
		if !Equal(result, expected) {
			t.Errorf("RotateLeft() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkRotateLeft benchmarks the RotateLeft function.
// BenchmarkRotateLeft는 RotateLeft 함수를 벤치마크합니다.
func BenchmarkRotateLeft(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RotateLeft(numbers, 100)
	}
}

// TestRotateRight tests the RotateRight function.
// TestRotateRight는 RotateRight 함수를 테스트합니다.
func TestRotateRight(t *testing.T) {
	t.Run("rotate by 2", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := RotateRight(numbers, 2)
		expected := []int{4, 5, 1, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("RotateRight() = %v, want %v", result, expected)
		}
	})

	t.Run("rotate by 0", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := RotateRight(numbers, 0)
		if !Equal(result, numbers) {
			t.Errorf("RotateRight(0) = %v, want %v", result, numbers)
		}
	})

	t.Run("rotate by length", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := RotateRight(numbers, 5)
		if !Equal(result, numbers) {
			t.Errorf("RotateRight(len) = %v, want %v", result, numbers)
		}
	})

	t.Run("rotate by more than length", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := RotateRight(numbers, 7) // 7 % 3 = 1
		expected := []int{3, 1, 2}
		if !Equal(result, expected) {
			t.Errorf("RotateRight() = %v, want %v", result, expected)
		}
	})

	t.Run("rotate by negative (left)", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := RotateRight(numbers, -1)
		expected := []int{2, 3, 4, 5, 1}
		if !Equal(result, expected) {
			t.Errorf("RotateRight(-1) = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := RotateRight(numbers, 3)
		if len(result) != 0 {
			t.Errorf("RotateRight() on empty should return empty, got %v", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		result := RotateRight(numbers, 5)
		expected := []int{42}
		if !Equal(result, expected) {
			t.Errorf("RotateRight() = %v, want %v", result, expected)
		}
	})

	t.Run("original not modified", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		original := Clone(numbers)
		RotateRight(numbers, 2)
		if !Equal(numbers, original) {
			t.Error("RotateRight() should not modify original slice")
		}
	})

	t.Run("equivalence with rotate left", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		rightResult := RotateRight(numbers, 2)
		leftResult := RotateLeft(numbers, -2)
		if !Equal(rightResult, leftResult) {
			t.Error("RotateRight(n) should equal RotateLeft(-n)")
		}
	})
}

// BenchmarkRotateRight benchmarks the RotateRight function.
// BenchmarkRotateRight는 RotateRight 함수를 벤치마크합니다.
func BenchmarkRotateRight(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RotateRight(numbers, 100)
	}
}
