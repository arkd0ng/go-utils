package sliceutil

import (
	"strconv"
	"testing"
)

// TestMap tests the Map function.
// TestMap은 Map 함수를 테스트합니다.
func TestMap(t *testing.T) {
	t.Run("map integers to doubled values", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Map(input, func(n int) int {
			return n * 2
		})
		expected := []int{2, 4, 6, 8, 10}
		if !Equal(result, expected) {
			t.Errorf("Map() = %v, want %v", result, expected)
		}
	})

	t.Run("map integers to strings", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Map(input, func(n int) string {
			return strconv.Itoa(n)
		})
		expected := []string{"1", "2", "3"}
		if !Equal(result, expected) {
			t.Errorf("Map() = %v, want %v", result, expected)
		}
	})

	t.Run("map strings to lengths", func(t *testing.T) {
		input := []string{"hello", "world", "go"}
		result := Map(input, func(s string) int {
			return len(s)
		})
		expected := []int{5, 5, 2}
		if !Equal(result, expected) {
			t.Errorf("Map() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Map(input, func(n int) int {
			return n * 2
		})
		if len(result) != 0 {
			t.Errorf("Map() with empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		result := Map(input, func(n int) int {
			return n * 2
		})
		if result != nil {
			t.Errorf("Map() with nil slice should return nil, got %v", result)
		}
	})
}

// BenchmarkMap benchmarks the Map function.
// BenchmarkMap은 Map 함수를 벤치마크합니다.
func BenchmarkMap(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Map(input, func(n int) int {
			return n * 2
		})
	}
}

// TestFilter tests the Filter function.
// TestFilter는 Filter 함수를 테스트합니다.
func TestFilter(t *testing.T) {
	t.Run("filter even numbers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		result := Filter(input, func(n int) bool {
			return n%2 == 0
		})
		expected := []int{2, 4, 6}
		if !Equal(result, expected) {
			t.Errorf("Filter() = %v, want %v", result, expected)
		}
	})

	t.Run("filter strings by length", func(t *testing.T) {
		input := []string{"apple", "a", "banana", "pear", "ab"}
		result := Filter(input, func(s string) bool {
			return len(s) > 2
		})
		expected := []string{"apple", "banana", "pear"}
		if !Equal(result, expected) {
			t.Errorf("Filter() = %v, want %v", result, expected)
		}
	})

	t.Run("filter all false", func(t *testing.T) {
		input := []int{1, 3, 5, 7}
		result := Filter(input, func(n int) bool {
			return n%2 == 0
		})
		if len(result) != 0 {
			t.Errorf("Filter() with all false should return empty slice, got %v", result)
		}
	})

	t.Run("filter all true", func(t *testing.T) {
		input := []int{2, 4, 6, 8}
		result := Filter(input, func(n int) bool {
			return n%2 == 0
		})
		if !Equal(result, input) {
			t.Errorf("Filter() with all true should return all elements, got %v", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Filter(input, func(n int) bool {
			return n%2 == 0
		})
		if len(result) != 0 {
			t.Errorf("Filter() with empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		result := Filter(input, func(n int) bool {
			return n%2 == 0
		})
		if result != nil {
			t.Errorf("Filter() with nil slice should return nil, got %v", result)
		}
	})
}

// BenchmarkFilter benchmarks the Filter function.
// BenchmarkFilter는 Filter 함수를 벤치마크합니다.
func BenchmarkFilter(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(input, func(n int) bool {
			return n%2 == 0
		})
	}
}

// TestFlatMap tests the FlatMap function.
// TestFlatMap은 FlatMap 함수를 테스트합니다.
func TestFlatMap(t *testing.T) {
	t.Run("flatmap words to characters", func(t *testing.T) {
		input := []string{"hello", "world"}
		result := FlatMap(input, func(s string) []rune {
			return []rune(s)
		})
		expected := []rune{'h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd'}
		if len(result) != len(expected) {
			t.Errorf("FlatMap() length = %v, want %v", len(result), len(expected))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("FlatMap()[%d] = %v, want %v", i, result[i], expected[i])
			}
		}
	})

	t.Run("flatmap numbers to pairs", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := FlatMap(input, func(n int) []int {
			return []int{n, n * 2}
		})
		expected := []int{1, 2, 2, 4, 3, 6}
		if !Equal(result, expected) {
			t.Errorf("FlatMap() = %v, want %v", result, expected)
		}
	})

	t.Run("flatmap with empty results", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := FlatMap(input, func(n int) []int {
			if n%2 == 0 {
				return []int{n}
			}
			return []int{}
		})
		expected := []int{2}
		if !Equal(result, expected) {
			t.Errorf("FlatMap() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := FlatMap(input, func(n int) []int {
			return []int{n, n * 2}
		})
		if len(result) != 0 {
			t.Errorf("FlatMap() with empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		result := FlatMap(input, func(n int) []int {
			return []int{n, n * 2}
		})
		if result != nil {
			t.Errorf("FlatMap() with nil slice should return nil, got %v", result)
		}
	})
}

// BenchmarkFlatMap benchmarks the FlatMap function.
// BenchmarkFlatMap은 FlatMap 함수를 벤치마크합니다.
func BenchmarkFlatMap(b *testing.B) {
	input := make([]int, 100)
	for i := range input {
		input[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FlatMap(input, func(n int) []int {
			return []int{n, n * 2}
		})
	}
}

// TestFlatten tests the Flatten function.
// TestFlatten은 Flatten 함수를 테스트합니다.
func TestFlatten(t *testing.T) {
	t.Run("flatten nested integers", func(t *testing.T) {
		input := [][]int{{1, 2}, {3, 4}, {5}}
		result := Flatten(input)
		expected := []int{1, 2, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Flatten() = %v, want %v", result, expected)
		}
	})

	t.Run("flatten nested strings", func(t *testing.T) {
		input := [][]string{{"hello", "world"}, {"foo", "bar"}}
		result := Flatten(input)
		expected := []string{"hello", "world", "foo", "bar"}
		if !Equal(result, expected) {
			t.Errorf("Flatten() = %v, want %v", result, expected)
		}
	})

	t.Run("flatten with empty inner slices", func(t *testing.T) {
		input := [][]int{{1, 2}, {}, {3}}
		result := Flatten(input)
		expected := []int{1, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("Flatten() = %v, want %v", result, expected)
		}
	})

	t.Run("flatten all empty", func(t *testing.T) {
		input := [][]int{{}, {}, {}}
		result := Flatten(input)
		if len(result) != 0 {
			t.Errorf("Flatten() with all empty should return empty slice, got %v", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := [][]int{}
		result := Flatten(input)
		if len(result) != 0 {
			t.Errorf("Flatten() with empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input [][]int
		result := Flatten(input)
		if result != nil {
			t.Errorf("Flatten() with nil slice should return nil, got %v", result)
		}
	})

	t.Run("single nested slice", func(t *testing.T) {
		input := [][]int{{1, 2, 3, 4, 5}}
		result := Flatten(input)
		expected := []int{1, 2, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Flatten() = %v, want %v", result, expected)
		}
	})
}

// BenchmarkFlatten benchmarks the Flatten function.
// BenchmarkFlatten은 Flatten 함수를 벤치마크합니다.
func BenchmarkFlatten(b *testing.B) {
	input := make([][]int, 100)
	for i := range input {
		input[i] = []int{i, i + 1, i + 2}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Flatten(input)
	}
}

// TestUnique tests the Unique function.
// TestUnique는 Unique 함수를 테스트합니다.
func TestUnique(t *testing.T) {
	t.Run("unique integers", func(t *testing.T) {
		input := []int{1, 2, 2, 3, 3, 3, 4}
		result := Unique(input)
		expected := []int{1, 2, 3, 4}
		if !Equal(result, expected) {
			t.Errorf("Unique() = %v, want %v", result, expected)
		}
	})

	t.Run("unique strings", func(t *testing.T) {
		input := []string{"apple", "banana", "apple", "cherry", "banana"}
		result := Unique(input)
		expected := []string{"apple", "banana", "cherry"}
		if !Equal(result, expected) {
			t.Errorf("Unique() = %v, want %v", result, expected)
		}
	})

	t.Run("no duplicates", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Unique(input)
		if !Equal(result, input) {
			t.Errorf("Unique() with no duplicates = %v, want %v", result, input)
		}
	})

	t.Run("all duplicates", func(t *testing.T) {
		input := []int{1, 1, 1, 1}
		result := Unique(input)
		expected := []int{1}
		if !Equal(result, expected) {
			t.Errorf("Unique() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Unique(input)
		if len(result) != 0 {
			t.Errorf("Unique() with empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		result := Unique(input)
		if result != nil {
			t.Errorf("Unique() with nil slice should return nil, got %v", result)
		}
	})
}

// BenchmarkUnique benchmarks the Unique function.
// BenchmarkUnique는 Unique 함수를 벤치마크합니다.
func BenchmarkUnique(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i % 100 // Create some duplicates
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Unique(input)
	}
}

// TestUniqueBy tests the UniqueBy function.
// TestUniqueBy는 UniqueBy 함수를 테스트합니다.
func TestUniqueBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("unique by name", func(t *testing.T) {
		input := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Alice", 28},
			{"Charlie", 35},
		}
		result := UniqueBy(input, func(p Person) string {
			return p.Name
		})
		if len(result) != 3 {
			t.Errorf("UniqueBy() length = %v, want 3", len(result))
		}
		if result[0].Name != "Alice" || result[1].Name != "Bob" || result[2].Name != "Charlie" {
			t.Errorf("UniqueBy() names incorrect: %v", result)
		}
	})

	t.Run("unique by age", func(t *testing.T) {
		input := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 25},
		}
		result := UniqueBy(input, func(p Person) int {
			return p.Age
		})
		if len(result) != 2 {
			t.Errorf("UniqueBy() length = %v, want 2", len(result))
		}
	})

	t.Run("unique integers by absolute value", func(t *testing.T) {
		input := []int{1, -1, 2, -2, 3}
		result := UniqueBy(input, func(n int) int {
			if n < 0 {
				return -n
			}
			return n
		})
		expected := []int{1, 2, 3}
		if !Equal(result, expected) {
			t.Errorf("UniqueBy() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := UniqueBy(input, func(n int) int {
			return n
		})
		if len(result) != 0 {
			t.Errorf("UniqueBy() with empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		result := UniqueBy(input, func(n int) int {
			return n
		})
		if result != nil {
			t.Errorf("UniqueBy() with nil slice should return nil, got %v", result)
		}
	})
}

// BenchmarkUniqueBy benchmarks the UniqueBy function.
// BenchmarkUniqueBy는 UniqueBy 함수를 벤치마크합니다.
func BenchmarkUniqueBy(b *testing.B) {
	type Person struct {
		Name string
		Age  int
	}
	input := make([]Person, 1000)
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve"}
	for i := range input {
		input[i] = Person{Name: names[i%len(names)], Age: i}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UniqueBy(input, func(p Person) string {
			return p.Name
		})
	}
}

// TestCompact tests the Compact function.
// TestCompact는 Compact 함수를 테스트합니다.
func TestCompact(t *testing.T) {
	t.Run("compact consecutive duplicates", func(t *testing.T) {
		input := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
		result := Compact(input)
		expected := []int{1, 2, 3, 4, 5}
		if !Equal(result, expected) {
			t.Errorf("Compact() = %v, want %v", result, expected)
		}
	})

	t.Run("compact strings", func(t *testing.T) {
		input := []string{"a", "a", "b", "b", "c"}
		result := Compact(input)
		expected := []string{"a", "b", "c"}
		if !Equal(result, expected) {
			t.Errorf("Compact() = %v, want %v", result, expected)
		}
	})

	t.Run("no consecutive duplicates", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Compact(input)
		if !Equal(result, input) {
			t.Errorf("Compact() with no duplicates = %v, want %v", result, input)
		}
	})

	t.Run("all same", func(t *testing.T) {
		input := []int{1, 1, 1, 1}
		result := Compact(input)
		expected := []int{1}
		if !Equal(result, expected) {
			t.Errorf("Compact() = %v, want %v", result, expected)
		}
	})

	t.Run("non-consecutive duplicates", func(t *testing.T) {
		input := []int{1, 2, 1, 2, 1}
		result := Compact(input)
		expected := []int{1, 2, 1, 2, 1}
		if !Equal(result, expected) {
			t.Errorf("Compact() should keep non-consecutive duplicates = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Compact(input)
		if len(result) != 0 {
			t.Errorf("Compact() with empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		result := Compact(input)
		if result != nil {
			t.Errorf("Compact() with nil slice should return nil, got %v", result)
		}
	})
}

// BenchmarkCompact benchmarks the Compact function.
// BenchmarkCompact는 Compact 함수를 벤치마크합니다.
func BenchmarkCompact(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i / 10 // Create consecutive duplicates
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Compact(input)
	}
}

// TestReverse tests the Reverse function.
// TestReverse는 Reverse 함수를 테스트합니다.
func TestReverse(t *testing.T) {
	t.Run("reverse integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Reverse(input)
		expected := []int{5, 4, 3, 2, 1}
		if !Equal(result, expected) {
			t.Errorf("Reverse() = %v, want %v", result, expected)
		}
	})

	t.Run("reverse strings", func(t *testing.T) {
		input := []string{"hello", "world"}
		result := Reverse(input)
		expected := []string{"world", "hello"}
		if !Equal(result, expected) {
			t.Errorf("Reverse() = %v, want %v", result, expected)
		}
	})

	t.Run("reverse single element", func(t *testing.T) {
		input := []int{1}
		result := Reverse(input)
		expected := []int{1}
		if !Equal(result, expected) {
			t.Errorf("Reverse() = %v, want %v", result, expected)
		}
	})

	t.Run("reverse even length", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		result := Reverse(input)
		expected := []int{4, 3, 2, 1}
		if !Equal(result, expected) {
			t.Errorf("Reverse() = %v, want %v", result, expected)
		}
	})

	t.Run("reverse odd length", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		result := Reverse(input)
		expected := []int{5, 4, 3, 2, 1}
		if !Equal(result, expected) {
			t.Errorf("Reverse() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		result := Reverse(input)
		if len(result) != 0 {
			t.Errorf("Reverse() with empty slice should return empty slice, got %v", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var input []int
		result := Reverse(input)
		if result != nil {
			t.Errorf("Reverse() with nil slice should return nil, got %v", result)
		}
	})

	t.Run("original unchanged", func(t *testing.T) {
		input := []int{1, 2, 3}
		original := []int{1, 2, 3}
		Reverse(input)
		if !Equal(input, original) {
			t.Errorf("Reverse() modified original slice: %v", input)
		}
	})
}

// BenchmarkReverse benchmarks the Reverse function.
// BenchmarkReverse는 Reverse 함수를 벤치마크합니다.
func BenchmarkReverse(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(input)
	}
}
