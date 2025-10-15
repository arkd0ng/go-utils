package sliceutil

import (
	"testing"
)

// TestReduce tests the Reduce function.
// TestReduce는 Reduce 함수를 테스트합니다.
func TestReduce(t *testing.T) {
	t.Run("reduce sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Reduce(numbers, 0, func(acc, n int) int {
			return acc + n
		})
		expected := 15
		if result != expected {
			t.Errorf("Reduce() sum = %v, want %v", result, expected)
		}
	})

	t.Run("reduce product", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Reduce(numbers, 1, func(acc, n int) int {
			return acc * n
		})
		expected := 120
		if result != expected {
			t.Errorf("Reduce() product = %v, want %v", result, expected)
		}
	})

	t.Run("reduce concat strings", func(t *testing.T) {
		words := []string{"hello", " ", "world"}
		result := Reduce(words, "", func(acc, w string) string {
			return acc + w
		})
		expected := "hello world"
		if result != expected {
			t.Errorf("Reduce() concat = %v, want %v", result, expected)
		}
	})

	t.Run("reduce max", func(t *testing.T) {
		numbers := []int{3, 7, 2, 9, 1}
		result := Reduce(numbers, numbers[0], func(acc, n int) int {
			if n > acc {
				return n
			}
			return acc
		})
		expected := 9
		if result != expected {
			t.Errorf("Reduce() max = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Reduce(numbers, 10, func(acc, n int) int {
			return acc + n
		})
		if result != 10 {
			t.Errorf("Reduce() with empty slice should return initial value, got %v", result)
		}
	})
}

// BenchmarkReduce benchmarks the Reduce function.
// BenchmarkReduce는 Reduce 함수를 벤치마크합니다.
func BenchmarkReduce(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(numbers, 0, func(acc, n int) int {
			return acc + n
		})
	}
}

// TestSum tests the Sum function.
// TestSum은 Sum 함수를 테스트합니다.
func TestSum(t *testing.T) {
	t.Run("sum integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expected := 15
		if result != expected {
			t.Errorf("Sum() = %v, want %v", result, expected)
		}
	})

	t.Run("sum floats", func(t *testing.T) {
		floats := []float64{1.5, 2.5, 3.0}
		result := Sum(floats)
		expected := 7.0
		if result != expected {
			t.Errorf("Sum() = %v, want %v", result, expected)
		}
	})

	t.Run("sum negative numbers", func(t *testing.T) {
		numbers := []int{-1, -2, -3}
		result := Sum(numbers)
		expected := -6
		if result != expected {
			t.Errorf("Sum() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Sum(numbers)
		if result != 0 {
			t.Errorf("Sum() with empty slice should return 0, got %v", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		result := Sum(numbers)
		if result != 42 {
			t.Errorf("Sum() = %v, want 42", result)
		}
	})
}

// BenchmarkSum benchmarks the Sum function.
// BenchmarkSum은 Sum 함수를 벤치마크합니다.
func BenchmarkSum(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

// TestMin tests the Min function.
// TestMin은 Min 함수를 테스트합니다.
func TestMin(t *testing.T) {
	t.Run("min integers", func(t *testing.T) {
		numbers := []int{3, 1, 4, 1, 5}
		result, err := Min(numbers)
		if err != nil {
			t.Errorf("Min() unexpected error: %v", err)
		}
		if result != 1 {
			t.Errorf("Min() = %v, want 1", result)
		}
	})

	t.Run("min strings", func(t *testing.T) {
		words := []string{"banana", "apple", "cherry"}
		result, err := Min(words)
		if err != nil {
			t.Errorf("Min() unexpected error: %v", err)
		}
		if result != "apple" {
			t.Errorf("Min() = %v, want apple", result)
		}
	})

	t.Run("min negative numbers", func(t *testing.T) {
		numbers := []int{-1, -5, -3}
		result, err := Min(numbers)
		if err != nil {
			t.Errorf("Min() unexpected error: %v", err)
		}
		if result != -5 {
			t.Errorf("Min() = %v, want -5", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		result, err := Min(numbers)
		if err != nil {
			t.Errorf("Min() unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("Min() = %v, want 42", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		_, err := Min(numbers)
		if err == nil {
			t.Error("Min() with empty slice should return error")
		}
	})
}

// BenchmarkMin benchmarks the Min function.
// BenchmarkMin은 Min 함수를 벤치마크합니다.
func BenchmarkMin(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(numbers)
	}
}

// TestMax tests the Max function.
// TestMax는 Max 함수를 테스트합니다.
func TestMax(t *testing.T) {
	t.Run("max integers", func(t *testing.T) {
		numbers := []int{3, 1, 4, 1, 5}
		result, err := Max(numbers)
		if err != nil {
			t.Errorf("Max() unexpected error: %v", err)
		}
		if result != 5 {
			t.Errorf("Max() = %v, want 5", result)
		}
	})

	t.Run("max strings", func(t *testing.T) {
		words := []string{"banana", "apple", "cherry"}
		result, err := Max(words)
		if err != nil {
			t.Errorf("Max() unexpected error: %v", err)
		}
		if result != "cherry" {
			t.Errorf("Max() = %v, want cherry", result)
		}
	})

	t.Run("max negative numbers", func(t *testing.T) {
		numbers := []int{-1, -5, -3}
		result, err := Max(numbers)
		if err != nil {
			t.Errorf("Max() unexpected error: %v", err)
		}
		if result != -1 {
			t.Errorf("Max() = %v, want -1", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		result, err := Max(numbers)
		if err != nil {
			t.Errorf("Max() unexpected error: %v", err)
		}
		if result != 42 {
			t.Errorf("Max() = %v, want 42", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		_, err := Max(numbers)
		if err == nil {
			t.Error("Max() with empty slice should return error")
		}
	})
}

// BenchmarkMax benchmarks the Max function.
// BenchmarkMax는 Max 함수를 벤치마크합니다.
func BenchmarkMax(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(numbers)
	}
}

// TestAverage tests the Average function.
// TestAverage는 Average 함수를 테스트합니다.
func TestAverage(t *testing.T) {
	t.Run("average integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Average(numbers)
		expected := 3.0
		if result != expected {
			t.Errorf("Average() = %v, want %v", result, expected)
		}
	})

	t.Run("average floats", func(t *testing.T) {
		floats := []float64{1.5, 2.5, 3.0}
		result := Average(floats)
		expected := 2.3333333333333335
		if result != expected {
			t.Errorf("Average() = %v, want %v", result, expected)
		}
	})

	t.Run("average negative numbers", func(t *testing.T) {
		numbers := []int{-2, -4, -6}
		result := Average(numbers)
		expected := -4.0
		if result != expected {
			t.Errorf("Average() = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := Average(numbers)
		if result != 0 {
			t.Errorf("Average() with empty slice should return 0, got %v", result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		result := Average(numbers)
		if result != 42.0 {
			t.Errorf("Average() = %v, want 42.0", result)
		}
	})
}

// BenchmarkAverage benchmarks the Average function.
// BenchmarkAverage는 Average 함수를 벤치마크합니다.
func BenchmarkAverage(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Average(numbers)
	}
}

// TestGroupBy tests the GroupBy function.
// TestGroupBy는 GroupBy 함수를 테스트합니다.
func TestGroupBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("group by age", func(t *testing.T) {
		people := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 25},
			{"David", 30},
		}
		result := GroupBy(people, func(p Person) int {
			return p.Age
		})
		if len(result) != 2 {
			t.Errorf("GroupBy() should have 2 groups, got %v", len(result))
		}
		if len(result[25]) != 2 {
			t.Errorf("GroupBy() age 25 should have 2 people, got %v", len(result[25]))
		}
		if len(result[30]) != 2 {
			t.Errorf("GroupBy() age 30 should have 2 people, got %v", len(result[30]))
		}
	})

	t.Run("group by even/odd", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result := GroupBy(numbers, func(n int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		})
		if len(result) != 2 {
			t.Errorf("GroupBy() should have 2 groups, got %v", len(result))
		}
		if len(result["even"]) != 3 {
			t.Errorf("GroupBy() even should have 3 numbers, got %v", len(result["even"]))
		}
		if len(result["odd"]) != 3 {
			t.Errorf("GroupBy() odd should have 3 numbers, got %v", len(result["odd"]))
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := GroupBy(numbers, func(n int) int {
			return n % 2
		})
		if len(result) != 0 {
			t.Errorf("GroupBy() with empty slice should return empty map, got %v", result)
		}
	})

	t.Run("all same group", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}
		result := GroupBy(numbers, func(n int) string {
			return "even"
		})
		if len(result) != 1 {
			t.Errorf("GroupBy() should have 1 group, got %v", len(result))
		}
		if len(result["even"]) != 4 {
			t.Errorf("GroupBy() even should have 4 numbers, got %v", len(result["even"]))
		}
	})
}

// BenchmarkGroupBy benchmarks the GroupBy function.
// BenchmarkGroupBy는 GroupBy 함수를 벤치마크합니다.
func BenchmarkGroupBy(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GroupBy(numbers, func(n int) int {
			return n % 10
		})
	}
}

// TestPartition tests the Partition function.
// TestPartition은 Partition 함수를 테스트합니다.
func TestPartition(t *testing.T) {
	t.Run("partition even/odd", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		evens, odds := Partition(numbers, func(n int) bool {
			return n%2 == 0
		})
		expectedEvens := []int{2, 4, 6}
		expectedOdds := []int{1, 3, 5}
		if !Equal(evens, expectedEvens) {
			t.Errorf("Partition() evens = %v, want %v", evens, expectedEvens)
		}
		if !Equal(odds, expectedOdds) {
			t.Errorf("Partition() odds = %v, want %v", odds, expectedOdds)
		}
	})

	t.Run("partition by length", func(t *testing.T) {
		words := []string{"apple", "a", "banana", "pear", "ab"}
		long, short := Partition(words, func(s string) bool {
			return len(s) > 2
		})
		if len(long) != 3 {
			t.Errorf("Partition() long should have 3 words, got %v", len(long))
		}
		if len(short) != 2 {
			t.Errorf("Partition() short should have 2 words, got %v", len(short))
		}
	})

	t.Run("all true", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}
		evens, odds := Partition(numbers, func(n int) bool {
			return n%2 == 0
		})
		if !Equal(evens, numbers) {
			t.Errorf("Partition() all true should return all in first slice")
		}
		if len(odds) != 0 {
			t.Errorf("Partition() all true should return empty second slice, got %v", odds)
		}
	})

	t.Run("all false", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7}
		evens, odds := Partition(numbers, func(n int) bool {
			return n%2 == 0
		})
		if len(evens) != 0 {
			t.Errorf("Partition() all false should return empty first slice, got %v", evens)
		}
		if !Equal(odds, numbers) {
			t.Errorf("Partition() all false should return all in second slice")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		trueSlice, falseSlice := Partition(numbers, func(n int) bool {
			return n%2 == 0
		})
		if len(trueSlice) != 0 || len(falseSlice) != 0 {
			t.Errorf("Partition() with empty slice should return two empty slices")
		}
	})
}

// BenchmarkPartition benchmarks the Partition function.
// BenchmarkPartition은 Partition 함수를 벤치마크합니다.
func BenchmarkPartition(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Partition(numbers, func(n int) bool {
			return n%2 == 0
		})
	}
}

// TestReduceRight tests the ReduceRight function.
// TestReduceRight는 ReduceRight 함수를 테스트합니다.
func TestReduceRight(t *testing.T) {
	t.Run("reduce right sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := ReduceRight(numbers, 0, func(acc, n int) int {
			return acc + n
		})
		expected := 15
		if result != expected {
			t.Errorf("ReduceRight() = %v, want %v", result, expected)
		}
	})

	t.Run("reduce right concat strings", func(t *testing.T) {
		words := []string{"a", "b", "c"}
		result := ReduceRight(words, "", func(acc, w string) string {
			return acc + w
		})
		expected := "cba" // Reversed order
		if result != expected {
			t.Errorf("ReduceRight() concat = %v, want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		numbers := []int{}
		result := ReduceRight(numbers, 10, func(acc, n int) int {
			return acc + n
		})
		if result != 10 {
			t.Errorf("ReduceRight() with empty slice should return initial value, got %v", result)
		}
	})
}

// TestCountBy tests the CountBy function.
// TestCountBy는 CountBy 함수를 테스트합니다.
func TestCountBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
		City string
	}

	t.Run("count by city", func(t *testing.T) {
		people := []Person{
			{"Alice", 30, "Seoul"},
			{"Bob", 25, "Busan"},
			{"Charlie", 35, "Seoul"},
			{"Dave", 28, "Seoul"},
			{"Eve", 32, "Busan"},
		}
		result := CountBy(people, func(p Person) string { return p.City })
		if result["Seoul"] != 3 {
			t.Errorf("CountBy() Seoul count = %v, want 3", result["Seoul"])
		}
		if result["Busan"] != 2 {
			t.Errorf("CountBy() Busan count = %v, want 2", result["Busan"])
		}
	})

	t.Run("count by age group", func(t *testing.T) {
		people := []Person{
			{"Alice", 25, "Seoul"},
			{"Bob", 35, "Busan"},
			{"Charlie", 28, "Seoul"},
			{"Dave", 32, "Seoul"},
		}
		result := CountBy(people, func(p Person) string {
			if p.Age < 30 {
				return "20s"
			}
			return "30s"
		})
		if result["20s"] != 2 {
			t.Errorf("CountBy() 20s count = %v, want 2", result["20s"])
		}
		if result["30s"] != 2 {
			t.Errorf("CountBy() 30s count = %v, want 2", result["30s"])
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		people := []Person{}
		result := CountBy(people, func(p Person) string { return p.City })
		if len(result) != 0 {
			t.Errorf("CountBy() with empty slice should return empty map, got %v", result)
		}
	})
}

// TestMinBy tests the MinBy function.
// TestMinBy는 MinBy 함수를 테스트합니다.
func TestMinBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("find youngest person", func(t *testing.T) {
		people := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}
		result, err := MinBy(people, func(p Person) int { return p.Age })
		if err != nil {
			t.Fatalf("MinBy() unexpected error: %v", err)
		}
		if result.Name != "Bob" || result.Age != 25 {
			t.Errorf("MinBy() = %+v, want Bob age 25", result)
		}
	})

	t.Run("find shortest name", func(t *testing.T) {
		people := []Person{
			{"Alexander", 30},
			{"Bo", 25},
			{"Charlie", 35},
		}
		result, err := MinBy(people, func(p Person) int { return len(p.Name) })
		if err != nil {
			t.Fatalf("MinBy() unexpected error: %v", err)
		}
		if result.Name != "Bo" {
			t.Errorf("MinBy() = %+v, want Bo", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		people := []Person{}
		_, err := MinBy(people, func(p Person) int { return p.Age })
		if err == nil {
			t.Error("MinBy() with empty slice should return error")
		}
	})
}

// TestMaxBy tests the MaxBy function.
// TestMaxBy는 MaxBy 함수를 테스트합니다.
func TestMaxBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("find oldest person", func(t *testing.T) {
		people := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}
		result, err := MaxBy(people, func(p Person) int { return p.Age })
		if err != nil {
			t.Fatalf("MaxBy() unexpected error: %v", err)
		}
		if result.Name != "Charlie" || result.Age != 35 {
			t.Errorf("MaxBy() = %+v, want Charlie age 35", result)
		}
	})

	t.Run("find longest name", func(t *testing.T) {
		people := []Person{
			{"Alexander", 30},
			{"Bo", 25},
			{"Charlie", 35},
		}
		result, err := MaxBy(people, func(p Person) int { return len(p.Name) })
		if err != nil {
			t.Fatalf("MaxBy() unexpected error: %v", err)
		}
		if result.Name != "Alexander" {
			t.Errorf("MaxBy() = %+v, want Alexander", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		people := []Person{}
		_, err := MaxBy(people, func(p Person) int { return p.Age })
		if err == nil {
			t.Error("MaxBy() with empty slice should return error")
		}
	})
}
