package sliceutil

import "testing"

// TestAll tests the All function.
// TestAll은 All 함수를 테스트합니다.
func TestAll(t *testing.T) {
	t.Run("all true", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}
		result := All(numbers, func(n int) bool { return n%2 == 0 })
		if !result {
			t.Error("All() should return true when all elements satisfy predicate")
		}
	})

	t.Run("not all true", func(t *testing.T) {
		numbers := []int{2, 4, 5, 8}
		result := All(numbers, func(n int) bool { return n%2 == 0 })
		if result {
			t.Error("All() should return false when not all elements satisfy predicate")
		}
	})

	t.Run("all false", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7}
		result := All(numbers, func(n int) bool { return n%2 == 0 })
		if result {
			t.Error("All() should return false when no elements satisfy predicate")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := All(numbers, func(n int) bool { return n%2 == 0 })
		if !result {
			t.Error("All() should return true for empty slice (vacuous truth)")
		}
	})

	t.Run("single element true", func(t *testing.T) {
		numbers := []int{2}
		result := All(numbers, func(n int) bool { return n%2 == 0 })
		if !result {
			t.Error("All() should return true for single element that satisfies predicate")
		}
	})

	t.Run("single element false", func(t *testing.T) {
		numbers := []int{1}
		result := All(numbers, func(n int) bool { return n%2 == 0 })
		if result {
			t.Error("All() should return false for single element that doesn't satisfy predicate")
		}
	})

	t.Run("all strings", func(t *testing.T) {
		words := []string{"hello", "world", "test"}
		result := All(words, func(s string) bool { return len(s) > 0 })
		if !result {
			t.Error("All() should return true when all strings are non-empty")
		}
	})
}

// BenchmarkAll benchmarks the All function.
// BenchmarkAll은 All 함수를 벤치마크합니다.
func BenchmarkAll(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i * 2
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All(numbers, func(n int) bool { return n%2 == 0 })
	}
}

// TestAny tests the Any function.
// TestAny는 Any 함수를 테스트합니다.
func TestAny(t *testing.T) {
	t.Run("has one match", func(t *testing.T) {
		numbers := []int{1, 3, 5, 6}
		result := Any(numbers, func(n int) bool { return n%2 == 0 })
		if !result {
			t.Error("Any() should return true when at least one element satisfies predicate")
		}
	})

	t.Run("has multiple matches", func(t *testing.T) {
		numbers := []int{2, 3, 4, 5}
		result := Any(numbers, func(n int) bool { return n%2 == 0 })
		if !result {
			t.Error("Any() should return true when multiple elements satisfy predicate")
		}
	})

	t.Run("no matches", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7}
		result := Any(numbers, func(n int) bool { return n%2 == 0 })
		if result {
			t.Error("Any() should return false when no elements satisfy predicate")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Any(numbers, func(n int) bool { return n%2 == 0 })
		if result {
			t.Error("Any() should return false for empty slice")
		}
	})

	t.Run("single element true", func(t *testing.T) {
		numbers := []int{2}
		result := Any(numbers, func(n int) bool { return n%2 == 0 })
		if !result {
			t.Error("Any() should return true for single element that satisfies predicate")
		}
	})

	t.Run("single element false", func(t *testing.T) {
		numbers := []int{1}
		result := Any(numbers, func(n int) bool { return n%2 == 0 })
		if result {
			t.Error("Any() should return false for single element that doesn't satisfy predicate")
		}
	})

	t.Run("any strings", func(t *testing.T) {
		words := []string{"", "hello", ""}
		result := Any(words, func(s string) bool { return len(s) > 0 })
		if !result {
			t.Error("Any() should return true when at least one string is non-empty")
		}
	})
}

// BenchmarkAny benchmarks the Any function.
// BenchmarkAny는 Any 함수를 벤치마크합니다.
func BenchmarkAny(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i*2 + 1
	}
	numbers[500] = 1000 // One even number in the middle / 중간에 짝수 하나
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any(numbers, func(n int) bool { return n%2 == 0 })
	}
}

// TestNone tests the None function.
// TestNone은 None 함수를 테스트합니다.
func TestNone(t *testing.T) {
	t.Run("none match", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7}
		result := None(numbers, func(n int) bool { return n%2 == 0 })
		if !result {
			t.Error("None() should return true when no elements satisfy predicate")
		}
	})

	t.Run("one match", func(t *testing.T) {
		numbers := []int{1, 3, 5, 6}
		result := None(numbers, func(n int) bool { return n%2 == 0 })
		if result {
			t.Error("None() should return false when at least one element satisfies predicate")
		}
	})

	t.Run("all match", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}
		result := None(numbers, func(n int) bool { return n%2 == 0 })
		if result {
			t.Error("None() should return false when all elements satisfy predicate")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := None(numbers, func(n int) bool { return n%2 == 0 })
		if !result {
			t.Error("None() should return true for empty slice")
		}
	})

	t.Run("single element true", func(t *testing.T) {
		numbers := []int{1}
		result := None(numbers, func(n int) bool { return n%2 == 0 })
		if !result {
			t.Error("None() should return true for single element that doesn't satisfy predicate")
		}
	})

	t.Run("single element false", func(t *testing.T) {
		numbers := []int{2}
		result := None(numbers, func(n int) bool { return n%2 == 0 })
		if result {
			t.Error("None() should return false for single element that satisfies predicate")
		}
	})

	t.Run("none strings", func(t *testing.T) {
		words := []string{"", "", ""}
		result := None(words, func(s string) bool { return len(s) > 0 })
		if !result {
			t.Error("None() should return true when all strings are empty")
		}
	})
}

// BenchmarkNone benchmarks the None function.
// BenchmarkNone은 None 함수를 벤치마크합니다.
func BenchmarkNone(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i*2 + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		None(numbers, func(n int) bool { return n%2 == 0 })
	}
}

// TestAllEqual tests the AllEqual function.
// TestAllEqual은 AllEqual 함수를 테스트합니다.
func TestAllEqual(t *testing.T) {
	t.Run("all equal", func(t *testing.T) {
		numbers := []int{5, 5, 5, 5}
		if !AllEqual(numbers) {
			t.Error("AllEqual() should return true when all elements are equal")
		}
	})

	t.Run("not all equal", func(t *testing.T) {
		numbers := []int{5, 5, 6, 5}
		if AllEqual(numbers) {
			t.Error("AllEqual() should return false when not all elements are equal")
		}
	})

	t.Run("first different", func(t *testing.T) {
		numbers := []int{1, 5, 5, 5}
		if AllEqual(numbers) {
			t.Error("AllEqual() should return false when first element is different")
		}
	})

	t.Run("last different", func(t *testing.T) {
		numbers := []int{5, 5, 5, 6}
		if AllEqual(numbers) {
			t.Error("AllEqual() should return false when last element is different")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		if !AllEqual(numbers) {
			t.Error("AllEqual() should return true for empty slice")
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		if !AllEqual(numbers) {
			t.Error("AllEqual() should return true for single element")
		}
	})

	t.Run("two equal", func(t *testing.T) {
		numbers := []int{5, 5}
		if !AllEqual(numbers) {
			t.Error("AllEqual() should return true for two equal elements")
		}
	})

	t.Run("two different", func(t *testing.T) {
		numbers := []int{5, 6}
		if AllEqual(numbers) {
			t.Error("AllEqual() should return false for two different elements")
		}
	})

	t.Run("strings all equal", func(t *testing.T) {
		words := []string{"hello", "hello", "hello"}
		if !AllEqual(words) {
			t.Error("AllEqual() should return true when all strings are equal")
		}
	})

	t.Run("strings not all equal", func(t *testing.T) {
		words := []string{"hello", "world", "hello"}
		if AllEqual(words) {
			t.Error("AllEqual() should return false when not all strings are equal")
		}
	})
}

// BenchmarkAllEqual benchmarks the AllEqual function.
// BenchmarkAllEqual은 AllEqual 함수를 벤치마크합니다.
func BenchmarkAllEqual(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = 42
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AllEqual(numbers)
	}
}

// TestIsSortedBy tests the IsSortedBy function.
// TestIsSortedBy는 IsSortedBy 함수를 테스트합니다.
func TestIsSortedBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("sorted by age", func(t *testing.T) {
		people := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 35},
		}
		if !IsSortedBy(people, func(p Person) int { return p.Age }) {
			t.Error("IsSortedBy() should return true when sorted by age")
		}
	})

	t.Run("not sorted by age", func(t *testing.T) {
		people := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}
		if IsSortedBy(people, func(p Person) int { return p.Age }) {
			t.Error("IsSortedBy() should return false when not sorted by age")
		}
	})

	t.Run("sorted by name", func(t *testing.T) {
		people := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}
		if !IsSortedBy(people, func(p Person) string { return p.Name }) {
			t.Error("IsSortedBy() should return true when sorted by name")
		}
	})

	t.Run("sorted with duplicates", func(t *testing.T) {
		people := []Person{
			{"Alice", 25},
			{"Bob", 25},
			{"Charlie", 30},
		}
		if !IsSortedBy(people, func(p Person) int { return p.Age }) {
			t.Error("IsSortedBy() should return true when sorted with duplicates")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		people := []Person{}
		if !IsSortedBy(people, func(p Person) int { return p.Age }) {
			t.Error("IsSortedBy() should return true for empty slice")
		}
	})

	t.Run("single element", func(t *testing.T) {
		people := []Person{{"Alice", 25}}
		if !IsSortedBy(people, func(p Person) int { return p.Age }) {
			t.Error("IsSortedBy() should return true for single element")
		}
	})

	t.Run("sorted strings by length", func(t *testing.T) {
		words := []string{"a", "ab", "abc", "abcd"}
		if !IsSortedBy(words, func(s string) int { return len(s) }) {
			t.Error("IsSortedBy() should return true when sorted by length")
		}
	})

	t.Run("not sorted strings by length", func(t *testing.T) {
		words := []string{"abcd", "ab", "abc", "a"}
		if IsSortedBy(words, func(s string) int { return len(s) }) {
			t.Error("IsSortedBy() should return false when not sorted by length")
		}
	})
}

// BenchmarkIsSortedBy benchmarks the IsSortedBy function.
// BenchmarkIsSortedBy는 IsSortedBy 함수를 벤치마크합니다.
func BenchmarkIsSortedBy(b *testing.B) {
	type Item struct {
		ID    int
		Value int
	}
	items := make([]Item, 1000)
	for i := range items {
		items[i] = Item{ID: i, Value: i}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsSortedBy(items, func(item Item) int { return item.Value })
	}
}

// TestContainsAll tests the ContainsAll function.
// TestContainsAll은 ContainsAll 함수를 테스트합니다.
func TestContainsAll(t *testing.T) {
	t.Run("contains all", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		if !ContainsAll(numbers, 2, 4) {
			t.Error("ContainsAll() should return true when slice contains all items")
		}
	})

	t.Run("contains all single item", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		if !ContainsAll(numbers, 3) {
			t.Error("ContainsAll() should return true when slice contains the item")
		}
	})

	t.Run("does not contain all", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		if ContainsAll(numbers, 2, 6) {
			t.Error("ContainsAll() should return false when slice doesn't contain all items")
		}
	})

	t.Run("missing one item", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		if ContainsAll(numbers, 1, 2, 3, 6) {
			t.Error("ContainsAll() should return false when missing one item")
		}
	})

	t.Run("empty items", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		if !ContainsAll(numbers) {
			t.Error("ContainsAll() should return true when no items specified")
		}
	})

	t.Run("empty slice with items", func(t *testing.T) {
		numbers := []int{}
		if ContainsAll(numbers, 1, 2) {
			t.Error("ContainsAll() should return false for empty slice with items")
		}
	})

	t.Run("empty slice no items", func(t *testing.T) {
		numbers := []int{}
		if !ContainsAll(numbers) {
			t.Error("ContainsAll() should return true for empty slice with no items")
		}
	})

	t.Run("contains all with duplicates in items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		if !ContainsAll(numbers, 2, 2, 3, 3) {
			t.Error("ContainsAll() should return true even with duplicates in items")
		}
	})

	t.Run("contains all strings", func(t *testing.T) {
		words := []string{"apple", "banana", "cherry"}
		if !ContainsAll(words, "apple", "cherry") {
			t.Error("ContainsAll() should return true when slice contains all string items")
		}
	})

	t.Run("does not contain all strings", func(t *testing.T) {
		words := []string{"apple", "banana", "cherry"}
		if ContainsAll(words, "apple", "grape") {
			t.Error("ContainsAll() should return false when slice doesn't contain all string items")
		}
	})
}

// BenchmarkContainsAll benchmarks the ContainsAll function.
// BenchmarkContainsAll은 ContainsAll 함수를 벤치마크합니다.
func BenchmarkContainsAll(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	items := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsAll(numbers, items...)
	}
}
