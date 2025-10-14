package sliceutil

import (
	"testing"
)

// TestSort tests the Sort function.
// TestSort는 Sort 함수를 테스트합니다.
func TestSort(t *testing.T) {
	t.Run("sort integers", func(t *testing.T) {
		input := []int{3, 1, 4, 1, 5, 9, 2, 6}
		result := Sort(input)
		expected := []int{1, 1, 2, 3, 4, 5, 6, 9}
		if !Equal(result, expected) {
			t.Errorf("Sort() = %v, want %v", result, expected)
		}
		// Verify original is unchanged / 원본이 변경되지 않았는지 확인
		if Equal(input, result) && len(input) > 1 {
			// This might pass if input was already sorted, check first element
			// 입력이 이미 정렬되어 있었다면 통과할 수 있음, 첫 번째 요소 확인
			if input[0] != 3 {
				t.Error("Sort() should not modify original slice")
			}
		}
	})

	t.Run("sort strings", func(t *testing.T) {
		input := []string{"banana", "apple", "cherry", "date"}
		result := Sort(input)
		expected := []string{"apple", "banana", "cherry", "date"}
		if !Equal(result, expected) {
			t.Errorf("Sort() = %v, want %v", result, expected)
		}
	})

	t.Run("sort floats", func(t *testing.T) {
		input := []float64{3.14, 1.41, 2.71, 1.73}
		result := Sort(input)
		expected := []float64{1.41, 1.73, 2.71, 3.14}
		if !Equal(result, expected) {
			t.Errorf("Sort() = %v, want %v", result, expected)
		}
	})

	t.Run("sort already sorted", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Sort(input)
		expected := []int{1, 2, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Sort() = %v, want %v", result, expected)
		}
	})

	t.Run("sort reverse sorted", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		result := Sort(input)
		expected := []int{1, 2, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Sort() = %v, want %v", result, expected)
		}
	})

	t.Run("sort single element", func(t *testing.T) {
		input := []int{42}
		result := Sort(input)
		expected := []int{42}
		if !Equal(result, expected) {
			t.Errorf("Sort() = %v, want %v", result, expected)
		}
	})

	t.Run("sort empty slice", func(t *testing.T) {
		input := []int{}
		result := Sort(input)
		if len(result) != 0 {
			t.Errorf("Sort() with empty slice should return empty, got %v", result)
		}
	})

	t.Run("sort with duplicates", func(t *testing.T) {
		input := []int{5, 2, 8, 2, 9, 1, 5, 5}
		result := Sort(input)
		expected := []int{1, 2, 2, 5, 5, 5, 8, 9}
		if !Equal(result, expected) {
			t.Errorf("Sort() = %v, want %v", result, expected)
		}
	})

	t.Run("sort negative numbers", func(t *testing.T) {
		input := []int{3, -1, 4, -5, 0, 2}
		result := Sort(input)
		expected := []int{-5, -1, 0, 2, 3, 4}
		if !Equal(result, expected) {
			t.Errorf("Sort() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkSort benchmarks the Sort function.
// BenchmarkSort는 Sort 함수를 벤치마크합니다.
func BenchmarkSort(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = 1000 - i // Reverse sorted / 역순 정렬
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sort(slice)
	}
}

// TestSortDesc tests the SortDesc function.
// TestSortDesc는 SortDesc 함수를 테스트합니다.
func TestSortDesc(t *testing.T) {
	t.Run("sort desc integers", func(t *testing.T) {
		input := []int{3, 1, 4, 1, 5, 9, 2, 6}
		result := SortDesc(input)
		expected := []int{9, 6, 5, 4, 3, 2, 1, 1}
		if !Equal(result, expected) {
			t.Errorf("SortDesc() = %v, want %v", result, expected)
		}
	})

	t.Run("sort desc strings", func(t *testing.T) {
		input := []string{"banana", "apple", "cherry", "date"}
		result := SortDesc(input)
		expected := []string{"date", "cherry", "banana", "apple"}
		if !Equal(result, expected) {
			t.Errorf("SortDesc() = %v, want %v", result, expected)
		}
	})

	t.Run("sort desc floats", func(t *testing.T) {
		input := []float64{3.14, 1.41, 2.71, 1.73}
		result := SortDesc(input)
		expected := []float64{3.14, 2.71, 1.73, 1.41}
		if !Equal(result, expected) {
			t.Errorf("SortDesc() = %v, want %v", result, expected)
		}
	})

	t.Run("sort desc already sorted desc", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		result := SortDesc(input)
		expected := []int{5, 4, 3, 2, 1}
		if !Equal(result, expected) {
			t.Errorf("SortDesc() = %v, want %v", result, expected)
		}
	})

	t.Run("sort desc reverse sorted", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := SortDesc(input)
		expected := []int{5, 4, 3, 2, 1}
		if !Equal(result, expected) {
			t.Errorf("SortDesc() = %v, want %v", result, expected)
		}
	})

	t.Run("sort desc single element", func(t *testing.T) {
		input := []int{42}
		result := SortDesc(input)
		expected := []int{42}
		if !Equal(result, expected) {
			t.Errorf("SortDesc() = %v, want %v", result, expected)
		}
	})

	t.Run("sort desc empty slice", func(t *testing.T) {
		input := []int{}
		result := SortDesc(input)
		if len(result) != 0 {
			t.Errorf("SortDesc() with empty slice should return empty, got %v", result)
		}
	})

	t.Run("sort desc with duplicates", func(t *testing.T) {
		input := []int{5, 2, 8, 2, 9, 1, 5, 5}
		result := SortDesc(input)
		expected := []int{9, 8, 5, 5, 5, 2, 2, 1}
		if !Equal(result, expected) {
			t.Errorf("SortDesc() = %v, want %v", result, expected)
		}
	})

	t.Run("sort desc negative numbers", func(t *testing.T) {
		input := []int{3, -1, 4, -5, 0, 2}
		result := SortDesc(input)
		expected := []int{4, 3, 2, 0, -1, -5}
		if !Equal(result, expected) {
			t.Errorf("SortDesc() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkSortDesc benchmarks the SortDesc function.
// BenchmarkSortDesc는 SortDesc 함수를 벤치마크합니다.
func BenchmarkSortDesc(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i // Already sorted ascending / 이미 오름차순 정렬됨
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortDesc(slice)
	}
}

// TestSortBy tests the SortBy function.
// TestSortBy는 SortBy 함수를 테스트합니다.
func TestSortBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("sort by age", func(t *testing.T) {
		input := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
			{"David", 25},
		}
		result := SortBy(input, func(p Person) int { return p.Age })
		if result[0].Age != 25 || result[1].Age != 25 || result[2].Age != 30 || result[3].Age != 35 {
			t.Errorf("SortBy() by age failed, got ages: %d, %d, %d, %d",
				result[0].Age, result[1].Age, result[2].Age, result[3].Age)
		}
	})

	t.Run("sort by name", func(t *testing.T) {
		input := []Person{
			{"Charlie", 35},
			{"Alice", 30},
			{"David", 25},
			{"Bob", 25},
		}
		result := SortBy(input, func(p Person) string { return p.Name })
		expectedNames := []string{"Alice", "Bob", "Charlie", "David"}
		for i, name := range expectedNames {
			if result[i].Name != name {
				t.Errorf("SortBy() by name at index %d = %s, want %s", i, result[i].Name, name)
			}
		}
	})

	t.Run("sort strings by length", func(t *testing.T) {
		input := []string{"banana", "pie", "apple", "a"}
		result := SortBy(input, func(s string) int { return len(s) })
		expected := []string{"a", "pie", "apple", "banana"}
		if !Equal(result, expected) {
			t.Errorf("SortBy() by length = %v, want %v", result, expected)
		}
	})

	t.Run("sort by negative key", func(t *testing.T) {
		input := []int{-3, 1, -5, 2, -1}
		result := SortBy(input, func(n int) int { return n })
		expected := []int{-5, -3, -1, 1, 2}
		if !Equal(result, expected) {
			t.Errorf("SortBy() by negative key = %v, want %v", result, expected)
		}
	})

	t.Run("sort by empty slice", func(t *testing.T) {
		input := []Person{}
		result := SortBy(input, func(p Person) int { return p.Age })
		if len(result) != 0 {
			t.Errorf("SortBy() with empty slice should return empty, got %v", result)
		}
	})

	t.Run("sort by single element", func(t *testing.T) {
		input := []Person{{"Alice", 30}}
		result := SortBy(input, func(p Person) int { return p.Age })
		if len(result) != 1 || result[0].Name != "Alice" {
			t.Errorf("SortBy() with single element failed")
		}
	})
}

// BenchmarkSortBy benchmarks the SortBy function.
// BenchmarkSortBy는 SortBy 함수를 벤치마크합니다.
func BenchmarkSortBy(b *testing.B) {
	type Item struct {
		ID    int
		Value int
	}
	slice := make([]Item, 1000)
	for i := range slice {
		slice[i] = Item{ID: i, Value: 1000 - i}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortBy(slice, func(item Item) int { return item.Value })
	}
}

// TestIsSorted tests the IsSorted function.
// TestIsSorted는 IsSorted 함수를 테스트합니다.
func TestIsSorted(t *testing.T) {
	t.Run("is sorted ascending", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		if !IsSorted(input) {
			t.Error("IsSorted() should return true for sorted slice")
		}
	})

	t.Run("is not sorted", func(t *testing.T) {
		input := []int{1, 3, 2, 4, 5}
		if IsSorted(input) {
			t.Error("IsSorted() should return false for unsorted slice")
		}
	})

	t.Run("is sorted descending", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		if IsSorted(input) {
			t.Error("IsSorted() should return false for descending slice")
		}
	})

	t.Run("is sorted with duplicates", func(t *testing.T) {
		input := []int{1, 2, 2, 3, 3, 3, 4}
		if !IsSorted(input) {
			t.Error("IsSorted() should return true for sorted slice with duplicates")
		}
	})

	t.Run("is sorted single element", func(t *testing.T) {
		input := []int{42}
		if !IsSorted(input) {
			t.Error("IsSorted() should return true for single element")
		}
	})

	t.Run("is sorted empty slice", func(t *testing.T) {
		input := []int{}
		if !IsSorted(input) {
			t.Error("IsSorted() should return true for empty slice")
		}
	})

	t.Run("is sorted strings", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		if !IsSorted(input) {
			t.Error("IsSorted() should return true for sorted strings")
		}
	})

	t.Run("is not sorted strings", func(t *testing.T) {
		input := []string{"banana", "apple", "cherry"}
		if IsSorted(input) {
			t.Error("IsSorted() should return false for unsorted strings")
		}
	})

	t.Run("is sorted with negatives", func(t *testing.T) {
		input := []int{-5, -3, -1, 0, 2, 4}
		if !IsSorted(input) {
			t.Error("IsSorted() should return true for sorted slice with negatives")
		}
	})
}

// BenchmarkIsSorted benchmarks the IsSorted function.
// BenchmarkIsSorted는 IsSorted 함수를 벤치마크합니다.
func BenchmarkIsSorted(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsSorted(slice)
	}
}

// TestIsSortedDesc tests the IsSortedDesc function.
// TestIsSortedDesc는 IsSortedDesc 함수를 테스트합니다.
func TestIsSortedDesc(t *testing.T) {
	t.Run("is sorted descending", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		if !IsSortedDesc(input) {
			t.Error("IsSortedDesc() should return true for descending slice")
		}
	})

	t.Run("is not sorted desc", func(t *testing.T) {
		input := []int{5, 3, 4, 2, 1}
		if IsSortedDesc(input) {
			t.Error("IsSortedDesc() should return false for unsorted slice")
		}
	})

	t.Run("is sorted ascending", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		if IsSortedDesc(input) {
			t.Error("IsSortedDesc() should return false for ascending slice")
		}
	})

	t.Run("is sorted desc with duplicates", func(t *testing.T) {
		input := []int{5, 5, 4, 3, 3, 2, 1}
		if !IsSortedDesc(input) {
			t.Error("IsSortedDesc() should return true for descending slice with duplicates")
		}
	})

	t.Run("is sorted desc single element", func(t *testing.T) {
		input := []int{42}
		if !IsSortedDesc(input) {
			t.Error("IsSortedDesc() should return true for single element")
		}
	})

	t.Run("is sorted desc empty slice", func(t *testing.T) {
		input := []int{}
		if !IsSortedDesc(input) {
			t.Error("IsSortedDesc() should return true for empty slice")
		}
	})

	t.Run("is sorted desc strings", func(t *testing.T) {
		input := []string{"cherry", "banana", "apple"}
		if !IsSortedDesc(input) {
			t.Error("IsSortedDesc() should return true for descending strings")
		}
	})

	t.Run("is not sorted desc strings", func(t *testing.T) {
		input := []string{"banana", "cherry", "apple"}
		if IsSortedDesc(input) {
			t.Error("IsSortedDesc() should return false for unsorted strings")
		}
	})

	t.Run("is sorted desc with negatives", func(t *testing.T) {
		input := []int{4, 2, 0, -1, -3, -5}
		if !IsSortedDesc(input) {
			t.Error("IsSortedDesc() should return true for descending slice with negatives")
		}
	})
}

// BenchmarkIsSortedDesc benchmarks the IsSortedDesc function.
// BenchmarkIsSortedDesc는 IsSortedDesc 함수를 벤치마크합니다.
func BenchmarkIsSortedDesc(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = 1000 - i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsSortedDesc(slice)
	}
}
