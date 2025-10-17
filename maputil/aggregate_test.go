package maputil

import (
	"testing"
)

func TestReduce(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	sum := Reduce(m, 0, func(acc int, k string, v int) int {
		return acc + v
	})

	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

func TestSum(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	total := Sum(m)

	if total != 6 {
		t.Errorf("Expected 6, got %d", total)
	}

	// Empty map
	// 빈 맵
	empty := map[string]int{}
	if Sum(empty) != 0 {
		t.Error("Sum of empty map should be 0")
	}
}

func TestMin(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}

	key, value, ok := Min(m)

	if !ok {
		t.Error("Expected ok to be true")
	}

	if value != 1 {
		t.Errorf("Expected min value 1, got %d", value)
	}

	if key != "b" {
		t.Errorf("Expected key 'b', got '%s'", key)
	}

	// Empty map
	// 빈 맵
	empty := map[string]int{}
	_, _, ok = Min(empty)
	if ok {
		t.Error("Expected ok to be false for empty map")
	}
}

func TestMax(t *testing.T) {
	m := map[string]int{"a": 3, "b": 1, "c": 2}

	key, value, ok := Max(m)

	if !ok {
		t.Error("Expected ok to be true")
	}

	if value != 3 {
		t.Errorf("Expected max value 3, got %d", value)
	}

	if key != "a" {
		t.Errorf("Expected key 'a', got '%s'", key)
	}

	// Empty map
	// 빈 맵
	empty := map[string]int{}
	_, _, ok = Max(empty)
	if ok {
		t.Error("Expected ok to be false for empty map")
	}
}

func TestMinBy(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	m := map[string]User{
		"alice": {Name: "Alice", Age: 30},
		"bob":   {Name: "Bob", Age: 25},
		"carol": {Name: "Carol", Age: 35},
	}

	key, user, ok := MinBy(m, func(u User) float64 {
		return float64(u.Age)
	})

	if !ok {
		t.Error("Expected ok to be true")
	}

	if user.Age != 25 {
		t.Errorf("Expected age 25, got %d", user.Age)
	}

	if key != "bob" {
		t.Errorf("Expected key 'bob', got '%s'", key)
	}
}

func TestMaxBy(t *testing.T) {
	type User struct {
		Name  string
		Score int
	}

	m := map[string]User{
		"alice": {Name: "Alice", Score: 95},
		"bob":   {Name: "Bob", Score: 88},
		"carol": {Name: "Carol", Score: 92},
	}

	key, user, ok := MaxBy(m, func(u User) float64 {
		return float64(u.Score)
	})

	if !ok {
		t.Error("Expected ok to be true")
	}

	if user.Score != 95 {
		t.Errorf("Expected score 95, got %d", user.Score)
	}

	if key != "alice" {
		t.Errorf("Expected key 'alice', got '%s'", key)
	}
}

func TestAverage(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	avg := Average(m)

	if avg != 2.5 {
		t.Errorf("Expected average 2.5, got %f", avg)
	}

	// Empty map
	// 빈 맵
	empty := map[string]int{}
	if Average(empty) != 0 {
		t.Error("Average of empty map should be 0")
	}
}

func TestGroupBy(t *testing.T) {
	type User struct {
		Name string
		City string
	}

	users := []User{
		{Name: "Alice", City: "Seoul"},
		{Name: "Bob", City: "Seoul"},
		{Name: "Charlie", City: "Busan"},
	}

	byCity := GroupBy[string, User, string](users, func(u User) string {
		return u.City
	})

	if len(byCity) != 2 {
		t.Errorf("Expected 2 cities, got %d", len(byCity))
	}

	if len(byCity["Seoul"]) != 2 {
		t.Errorf("Expected 2 users in Seoul, got %d", len(byCity["Seoul"]))
	}

	if len(byCity["Busan"]) != 1 {
		t.Errorf("Expected 1 user in Busan, got %d", len(byCity["Busan"]))
	}
}

func TestCountBy(t *testing.T) {
	type User struct {
		Name string
		City string
	}

	users := []User{
		{Name: "Alice", City: "Seoul"},
		{Name: "Bob", City: "Seoul"},
		{Name: "Charlie", City: "Busan"},
	}

	counts := CountBy[string, User, string](users, func(u User) string {
		return u.City
	})

	if counts["Seoul"] != 2 {
		t.Errorf("Expected 2 users in Seoul, got %d", counts["Seoul"])
	}

	if counts["Busan"] != 1 {
		t.Errorf("Expected 1 user in Busan, got %d", counts["Busan"])
	}
}

// Benchmark tests
// 벤치마크 테스트

func BenchmarkSum(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sum(m)
	}
}

func BenchmarkAverage(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Average(m)
	}
}

func BenchmarkGroupBy(b *testing.B) {
	type User struct {
		Name string
		City string
	}
	users := []User{
		{Name: "Alice", City: "Seoul"},
		{Name: "Bob", City: "Seoul"},
		{Name: "Charlie", City: "Busan"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GroupBy[string, User, string](users, func(u User) string { return u.City })
	}
}

// TestMedian tests the Median function with various scenarios.
// TestMedian는 다양한 시나리오로 Median 함수를 테스트합니다.
func TestMedian(t *testing.T) {
	t.Run("OddLength", func(t *testing.T) {
		// Test with odd number of values
		// 홀수 개의 값 테스트
		scores := map[string]int{
			"Alice":   85,
			"Bob":     90,
			"Charlie": 75,
			"Diana":   95,
			"Eve":     80,
		}
		median, ok := Median(scores)
		if !ok {
			t.Error("Median should return true for non-empty map")
		}
		if median != 85.0 {
			t.Errorf("Median(odd) = %v, want 85.0", median)
		}
	})

	t.Run("EvenLength", func(t *testing.T) {
		// Test with even number of values
		// 짝수 개의 값 테스트
		scores := map[string]int{
			"Alice":   80,
			"Bob":     90,
			"Charlie": 70,
			"Diana":   100,
		}
		median, ok := Median(scores)
		if !ok {
			t.Error("Median should return true for non-empty map")
		}
		if median != 85.0 {
			t.Errorf("Median(even) = %v, want 85.0", median)
		}
	})

	t.Run("EmptyMap", func(t *testing.T) {
		// Test with empty map
		// 빈 맵 테스트
		scores := map[string]int{}
		median, ok := Median(scores)
		if ok {
			t.Error("Median(empty) should return false")
		}
		if median != 0 {
			t.Errorf("Median(empty) = %v, want 0", median)
		}
	})

	t.Run("SingleValue", func(t *testing.T) {
		// Test with single value
		// 단일 값 테스트
		scores := map[string]int{"Alice": 85}
		median, ok := Median(scores)
		if !ok {
			t.Error("Median should return true for non-empty map")
		}
		if median != 85.0 {
			t.Errorf("Median(single) = %v, want 85.0", median)
		}
	})

	t.Run("TwoValues", func(t *testing.T) {
		// Test with two values
		// 두 값 테스트
		scores := map[string]int{
			"Alice": 80,
			"Bob":   90,
		}
		median, ok := Median(scores)
		if !ok {
			t.Error("Median should return true for non-empty map")
		}
		if median != 85.0 {
			t.Errorf("Median(two) = %v, want 85.0", median)
		}
	})

	t.Run("AllSameValues", func(t *testing.T) {
		// Test with all same values
		// 모두 같은 값 테스트
		scores := map[string]int{
			"Alice": 85,
			"Bob":   85,
			"Charlie": 85,
		}
		median, ok := Median(scores)
		if !ok {
			t.Error("Median should return true for non-empty map")
		}
		if median != 85.0 {
			t.Errorf("Median(same) = %v, want 85.0", median)
		}
	})

	t.Run("FloatValues", func(t *testing.T) {
		// Test with float values
		// 소수점 값 테스트
		prices := map[string]float64{
			"item1": 10.5,
			"item2": 20.5,
			"item3": 15.5,
		}
		median, ok := Median(prices)
		if !ok {
			t.Error("Median should return true for non-empty map")
		}
		if median != 15.5 {
			t.Errorf("Median(float) = %v, want 15.5", median)
		}
	})

	t.Run("NegativeValues", func(t *testing.T) {
		// Test with negative values
		// 음수 값 테스트
		temps := map[string]int{
			"day1": -5,
			"day2": -10,
			"day3": 0,
			"day4": 5,
			"day5": 10,
		}
		median, ok := Median(temps)
		if !ok {
			t.Error("Median should return true for non-empty map")
		}
		if median != 0.0 {
			t.Errorf("Median(negative) = %v, want 0.0", median)
		}
	})
}

// TestFrequencies tests the Frequencies function with various scenarios.
// TestFrequencies는 다양한 시나리오로 Frequencies 함수를 테스트합니다.
func TestFrequencies(t *testing.T) {
	t.Run("StringValues", func(t *testing.T) {
		// Test with string values
		// 문자열 값 테스트
		grades := map[string]string{
			"Alice":   "A",
			"Bob":     "B",
			"Charlie": "A",
			"Diana":   "C",
			"Eve":     "B",
			"Frank":   "A",
		}
		freq := Frequencies(grades)
		if len(freq) != 3 {
			t.Errorf("Frequencies() returned %d unique values, want 3", len(freq))
		}
		if freq["A"] != 3 || freq["B"] != 2 || freq["C"] != 1 {
			t.Errorf("Frequencies() = %v, want A:3, B:2, C:1", freq)
		}
	})

	t.Run("IntValues", func(t *testing.T) {
		// Test with int values
		// 정수 값 테스트
		scores := map[string]int{
			"test1": 85,
			"test2": 90,
			"test3": 85,
			"test4": 90,
			"test5": 75,
		}
		freq := Frequencies(scores)
		if len(freq) != 3 {
			t.Errorf("Frequencies() returned %d unique values, want 3", len(freq))
		}
		if freq[85] != 2 || freq[90] != 2 || freq[75] != 1 {
			t.Errorf("Frequencies() = %v, want 85:2, 90:2, 75:1", freq)
		}
	})

	t.Run("EmptyMap", func(t *testing.T) {
		// Test with empty map
		// 빈 맵 테스트
		data := map[string]int{}
		freq := Frequencies(data)
		if len(freq) != 0 {
			t.Errorf("Frequencies(empty) = %v, want empty map", freq)
		}
	})

	t.Run("AllUnique", func(t *testing.T) {
		// Test with all unique values
		// 모두 고유한 값 테스트
		data := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
			"d": 4,
		}
		freq := Frequencies(data)
		if len(freq) != 4 {
			t.Errorf("Frequencies() returned %d unique values, want 4", len(freq))
		}
		for _, count := range freq {
			if count != 1 {
				t.Errorf("All frequencies should be 1, got %d", count)
			}
		}
	})

	t.Run("AllSame", func(t *testing.T) {
		// Test with all same values
		// 모두 같은 값 테스트
		data := map[string]int{
			"a": 5,
			"b": 5,
			"c": 5,
			"d": 5,
		}
		freq := Frequencies(data)
		if len(freq) != 1 {
			t.Errorf("Frequencies() returned %d unique values, want 1", len(freq))
		}
		if freq[5] != 4 {
			t.Errorf("Frequencies()[5] = %d, want 4", freq[5])
		}
	})

	t.Run("SingleEntry", func(t *testing.T) {
		// Test with single entry
		// 단일 항목 테스트
		data := map[string]int{"a": 1}
		freq := Frequencies(data)
		if len(freq) != 1 {
			t.Errorf("Frequencies() returned %d unique values, want 1", len(freq))
		}
		if freq[1] != 1 {
			t.Errorf("Frequencies()[1] = %d, want 1", freq[1])
		}
	})

	t.Run("BooleanValues", func(t *testing.T) {
		// Test with boolean values
		// 부울 값 테스트
		flags := map[string]bool{
			"feature1": true,
			"feature2": false,
			"feature3": true,
			"feature4": true,
			"feature5": false,
		}
		freq := Frequencies(flags)
		if len(freq) != 2 {
			t.Errorf("Frequencies() returned %d unique values, want 2", len(freq))
		}
		if freq[true] != 3 || freq[false] != 2 {
			t.Errorf("Frequencies() = %v, want true:3, false:2", freq)
		}
	})

	t.Run("DuplicateDetection", func(t *testing.T) {
		// Test for duplicate value detection
		// 중복 값 감지 테스트
		emails := map[string]string{
			"user1": "alice@example.com",
			"user2": "bob@example.com",
			"user3": "alice@example.com",
		}
		freq := Frequencies(emails)

		// Find duplicates
		var duplicates []string
		for email, count := range freq {
			if count > 1 {
				duplicates = append(duplicates, email)
			}
		}

		if len(duplicates) != 1 {
			t.Errorf("Found %d duplicates, want 1", len(duplicates))
		}
		if len(duplicates) > 0 && duplicates[0] != "alice@example.com" {
			t.Errorf("Duplicate = %s, want alice@example.com", duplicates[0])
		}
	})
}

// BenchmarkMedian benchmarks the Median function.
// BenchmarkMedian는 Median 함수를 벤치마크합니다.
func BenchmarkMedian(b *testing.B) {
	scores := map[string]int{
		"Alice":   85,
		"Bob":     90,
		"Charlie": 75,
		"Diana":   95,
		"Eve":     80,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Median(scores)
	}
}

// BenchmarkMedian_Large benchmarks Median with large map.
// BenchmarkMedian_Large는 큰 맵으로 Median을 벤치마크합니다.
func BenchmarkMedian_Large(b *testing.B) {
	scores := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Median(scores)
	}
}

// BenchmarkFrequencies benchmarks the Frequencies function.
// BenchmarkFrequencies는 Frequencies 함수를 벤치마크합니다.
func BenchmarkFrequencies(b *testing.B) {
	grades := map[string]string{
		"Alice":   "A",
		"Bob":     "B",
		"Charlie": "A",
		"Diana":   "C",
		"Eve":     "B",
		"Frank":   "A",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Frequencies(grades)
	}
}

// BenchmarkFrequencies_Large benchmarks Frequencies with large map.
// BenchmarkFrequencies_Large는 큰 맵으로 Frequencies를 벤치마크합니다.
func BenchmarkFrequencies_Large(b *testing.B) {
	data := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i % 10 // Creates duplicates
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Frequencies(data)
	}
}
