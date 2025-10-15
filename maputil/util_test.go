package maputil

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// TestForEach tests the ForEach function with various scenarios.
// TestForEach는 다양한 시나리오로 ForEach 함수를 테스트합니다.
func TestForEach(t *testing.T) {
	t.Run("basic iteration", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		var result []string

		ForEach(m, func(k string, v int) {
			result = append(result, fmt.Sprintf("%s=%d", k, v))
		})

		// Sort for consistent comparison (map iteration order is random)
		sort.Strings(result)

		expected := []string{"a=1", "b=2", "c=3"}
		if len(result) != len(expected) {
			t.Errorf("ForEach() length = %d, want %d", len(result), len(expected))
		}

		for i, v := range result {
			if v != expected[i] {
				t.Errorf("ForEach() result[%d] = %s, want %s", i, v, expected[i])
			}
		}
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		count := 0

		ForEach(m, func(k string, v int) {
			count++
		})

		if count != 0 {
			t.Errorf("ForEach() on empty map called function %d times, want 0", count)
		}
	})

	t.Run("single entry", func(t *testing.T) {
		m := map[string]int{"only": 42}
		var key string
		var value int

		ForEach(m, func(k string, v int) {
			key = k
			value = v
		})

		if key != "only" || value != 42 {
			t.Errorf("ForEach() got (%s, %d), want (only, 42)", key, value)
		}
	})

	t.Run("accumulate values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		sum := 0

		ForEach(m, func(k string, v int) {
			sum += v
		})

		if sum != 10 {
			t.Errorf("ForEach() sum = %d, want 10", sum)
		}
	})

	t.Run("collect keys", func(t *testing.T) {
		m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
		var keys []string

		ForEach(m, func(k string, v int) {
			keys = append(keys, k)
		})

		sort.Strings(keys)
		expected := []string{"apple", "banana", "cherry"}

		if len(keys) != len(expected) {
			t.Errorf("ForEach() collected %d keys, want %d", len(keys), len(expected))
		}

		for i, k := range keys {
			if k != expected[i] {
				t.Errorf("ForEach() keys[%d] = %s, want %s", i, k, expected[i])
			}
		}
	})

	t.Run("different types", func(t *testing.T) {
		m := map[int]string{1: "one", 2: "two", 3: "three"}
		count := 0

		ForEach(m, func(k int, v string) {
			count++
			if k < 1 || k > 3 {
				t.Errorf("ForEach() unexpected key %d", k)
			}
		})

		if count != 3 {
			t.Errorf("ForEach() called %d times, want 3", count)
		}
	})

	t.Run("complex values", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}

		m := map[string]User{
			"user1": {Name: "Alice", Age: 25},
			"user2": {Name: "Bob", Age: 30},
		}

		var names []string
		ForEach(m, func(k string, u User) {
			names = append(names, u.Name)
		})

		sort.Strings(names)
		expected := []string{"Alice", "Bob"}

		if len(names) != len(expected) {
			t.Errorf("ForEach() collected %d names, want %d", len(names), len(expected))
		}

		for i, name := range names {
			if name != expected[i] {
				t.Errorf("ForEach() names[%d] = %s, want %s", i, name, expected[i])
			}
		}
	})
}

// TestForEachLogging demonstrates logging use case.
// TestForEachLogging은 로깅 사용 사례를 시연합니다.
func TestForEachLogging(t *testing.T) {
	m := map[string]int{"errors": 5, "warnings": 12, "info": 100}
	var logs []string

	ForEach(m, func(level string, count int) {
		logs = append(logs, fmt.Sprintf("Log level %s: %d messages", level, count))
	})

	if len(logs) != 3 {
		t.Errorf("ForEach() logging collected %d logs, want 3", len(logs))
	}

	// Verify all levels were logged
	logStr := strings.Join(logs, "|")
	for level := range m {
		if !strings.Contains(logStr, level) {
			t.Errorf("ForEach() missing log for level %s", level)
		}
	}
}

// BenchmarkForEach benchmarks the ForEach function.
// BenchmarkForEach는 ForEach 함수를 벤치마크합니다.
func BenchmarkForEach(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := make(map[int]int, size)
			for i := 0; i < size; i++ {
				m[i] = i * 2
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sum := 0
				ForEach(m, func(k, v int) {
					sum += v
				})
			}
		})
	}
}

// BenchmarkForEachVsRange compares ForEach with native range.
// BenchmarkForEachVsRange는 ForEach와 네이티브 range를 비교합니다.
func BenchmarkForEachVsRange(b *testing.B) {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i * 2
	}

	b.Run("ForEach", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sum := 0
			ForEach(m, func(k, v int) {
				sum += v
			})
		}
	})

	b.Run("NativeRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sum := 0
			for _, v := range m {
				sum += v
			}
		}
	})
}

// TestGetMany tests the GetMany function with various scenarios.
// TestGetMany는 다양한 시나리오로 GetMany 함수를 테스트합니다.
func TestGetMany(t *testing.T) {
	t.Run("basic retrieval", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		values := GetMany(m, "a", "c", "d")

		expected := []int{1, 3, 4}
		if len(values) != len(expected) {
			t.Errorf("GetMany() length = %d, want %d", len(values), len(expected))
		}

		for i, v := range values {
			if v != expected[i] {
				t.Errorf("GetMany() values[%d] = %d, want %d", i, v, expected[i])
			}
		}
	})

	t.Run("non-existent keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		values := GetMany(m, "a", "c", "d")

		expected := []int{1, 0, 0} // c and d don't exist, return zero values
		if len(values) != len(expected) {
			t.Errorf("GetMany() length = %d, want %d", len(values), len(expected))
		}

		for i, v := range values {
			if v != expected[i] {
				t.Errorf("GetMany() values[%d] = %d, want %d", i, v, expected[i])
			}
		}
	})

	t.Run("empty keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		values := GetMany(m)

		if len(values) != 0 {
			t.Errorf("GetMany() with no keys returned %d values, want 0", len(values))
		}
	})

	t.Run("single key", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		values := GetMany(m, "a")

		if len(values) != 1 {
			t.Errorf("GetMany() length = %d, want 1", len(values))
		}

		if values[0] != 1 {
			t.Errorf("GetMany() value = %d, want 1", values[0])
		}
	})

	t.Run("duplicate keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		values := GetMany(m, "a", "a", "b")

		expected := []int{1, 1, 2}
		if len(values) != len(expected) {
			t.Errorf("GetMany() length = %d, want %d", len(values), len(expected))
		}

		for i, v := range values {
			if v != expected[i] {
				t.Errorf("GetMany() values[%d] = %d, want %d", i, v, expected[i])
			}
		}
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		values := GetMany(m, "a", "b", "c")

		expected := []int{0, 0, 0}
		if len(values) != len(expected) {
			t.Errorf("GetMany() length = %d, want %d", len(values), len(expected))
		}

		for i, v := range values {
			if v != expected[i] {
				t.Errorf("GetMany() values[%d] = %d, want %d", i, v, expected[i])
			}
		}
	})

	t.Run("string values", func(t *testing.T) {
		m := map[int]string{1: "one", 2: "two", 3: "three"}
		values := GetMany(m, 1, 3, 5)

		expected := []string{"one", "three", ""} // 5 doesn't exist, returns empty string
		if len(values) != len(expected) {
			t.Errorf("GetMany() length = %d, want %d", len(values), len(expected))
		}

		for i, v := range values {
			if v != expected[i] {
				t.Errorf("GetMany() values[%d] = %s, want %s", i, v, expected[i])
			}
		}
	})

	t.Run("complex values", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}

		m := map[string]User{
			"user1": {Name: "Alice", Age: 25},
			"user2": {Name: "Bob", Age: 30},
		}

		values := GetMany(m, "user1", "user2", "user3")

		if len(values) != 3 {
			t.Errorf("GetMany() length = %d, want 3", len(values))
		}

		if values[0].Name != "Alice" || values[0].Age != 25 {
			t.Errorf("GetMany() values[0] = %+v, want {Alice 25}", values[0])
		}

		if values[1].Name != "Bob" || values[1].Age != 30 {
			t.Errorf("GetMany() values[1] = %+v, want {Bob 30}", values[1])
		}

		// user3 doesn't exist, should return zero value
		if values[2].Name != "" || values[2].Age != 0 {
			t.Errorf("GetMany() values[2] = %+v, want zero value", values[2])
		}
	})

	t.Run("large number of keys", func(t *testing.T) {
		m := make(map[int]int, 1000)
		for i := 0; i < 1000; i++ {
			m[i] = i * 2
		}

		keys := make([]int, 100)
		for i := 0; i < 100; i++ {
			keys[i] = i * 10
		}

		values := GetMany(m, keys...)

		if len(values) != 100 {
			t.Errorf("GetMany() length = %d, want 100", len(values))
		}

		for i, v := range values {
			expected := keys[i] * 2
			if v != expected {
				t.Errorf("GetMany() values[%d] = %d, want %d", i, v, expected)
			}
		}
	})
}

// BenchmarkGetMany benchmarks the GetMany function.
// BenchmarkGetMany는 GetMany 함수를 벤치마크합니다.
func BenchmarkGetMany(b *testing.B) {
	keyCounts := []int{1, 5, 10, 50, 100}

	for _, keyCount := range keyCounts {
		b.Run(fmt.Sprintf("keys_%d", keyCount), func(b *testing.B) {
			m := make(map[int]int, 1000)
			for i := 0; i < 1000; i++ {
				m[i] = i * 2
			}

			keys := make([]int, keyCount)
			for i := 0; i < keyCount; i++ {
				keys[i] = i
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = GetMany(m, keys...)
			}
		})
	}
}

// BenchmarkGetManyVsLoop compares GetMany with manual loop.
// BenchmarkGetManyVsLoop은 GetMany와 수동 루프를 비교합니다.
func BenchmarkGetManyVsLoop(b *testing.B) {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i * 2
	}

	keys := make([]int, 10)
	for i := 0; i < 10; i++ {
		keys[i] = i
	}

	b.Run("GetMany", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = GetMany(m, keys...)
		}
	})

	b.Run("ManualLoop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := make([]int, len(keys))
			for j, key := range keys {
				result[j] = m[key]
			}
		}
	})
}
