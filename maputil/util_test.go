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

// TestSetMany tests the SetMany function with various scenarios.
// TestSetMany는 다양한 시나리오로 SetMany 함수를 테스트합니다.
func TestSetMany(t *testing.T) {
	t.Run("basic set", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		result := SetMany(m,
			Entry[string, int]{Key: "c", Value: 3},
			Entry[string, int]{Key: "d", Value: 4},
		)

		expected := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		if len(result) != len(expected) {
			t.Errorf("SetMany() length = %d, want %d", len(result), len(expected))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("SetMany() result[%s] = %d, want %d", k, result[k], v)
			}
		}

		// Original map should not be modified
		if len(m) != 2 {
			t.Errorf("SetMany() modified original map, length = %d, want 2", len(m))
		}
	})

	t.Run("update existing keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		result := SetMany(m,
			Entry[string, int]{Key: "a", Value: 10},
			Entry[string, int]{Key: "c", Value: 3},
		)

		expected := map[string]int{"a": 10, "b": 2, "c": 3}
		if len(result) != len(expected) {
			t.Errorf("SetMany() length = %d, want %d", len(result), len(expected))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("SetMany() result[%s] = %d, want %d", k, result[k], v)
			}
		}
	})

	t.Run("empty entries", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		result := SetMany(m)

		if len(result) != len(m) {
			t.Errorf("SetMany() with no entries length = %d, want %d", len(result), len(m))
		}

		for k, v := range m {
			if result[k] != v {
				t.Errorf("SetMany() result[%s] = %d, want %d", k, result[k], v)
			}
		}
	})

	t.Run("single entry", func(t *testing.T) {
		m := map[string]int{"a": 1}
		result := SetMany(m, Entry[string, int]{Key: "b", Value: 2})

		expected := map[string]int{"a": 1, "b": 2}
		if len(result) != len(expected) {
			t.Errorf("SetMany() length = %d, want %d", len(result), len(expected))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("SetMany() result[%s] = %d, want %d", k, result[k], v)
			}
		}
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		result := SetMany(m,
			Entry[string, int]{Key: "a", Value: 1},
			Entry[string, int]{Key: "b", Value: 2},
		)

		expected := map[string]int{"a": 1, "b": 2}
		if len(result) != len(expected) {
			t.Errorf("SetMany() length = %d, want %d", len(result), len(expected))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("SetMany() result[%s] = %d, want %d", k, result[k], v)
			}
		}
	})

	t.Run("duplicate keys in entries", func(t *testing.T) {
		m := map[string]int{"a": 1}
		result := SetMany(m,
			Entry[string, int]{Key: "b", Value: 2},
			Entry[string, int]{Key: "b", Value: 3}, // Duplicate, last one wins
			Entry[string, int]{Key: "c", Value: 4},
		)

		expected := map[string]int{"a": 1, "b": 3, "c": 4}
		if len(result) != len(expected) {
			t.Errorf("SetMany() length = %d, want %d", len(result), len(expected))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("SetMany() result[%s] = %d, want %d", k, result[k], v)
			}
		}
	})

	t.Run("string values", func(t *testing.T) {
		m := map[int]string{1: "one", 2: "two"}
		result := SetMany(m,
			Entry[int, string]{Key: 3, Value: "three"},
			Entry[int, string]{Key: 4, Value: "four"},
		)

		expected := map[int]string{1: "one", 2: "two", 3: "three", 4: "four"}
		if len(result) != len(expected) {
			t.Errorf("SetMany() length = %d, want %d", len(result), len(expected))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("SetMany() result[%d] = %s, want %s", k, result[k], v)
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
		}

		result := SetMany(m,
			Entry[string, User]{Key: "user2", Value: User{Name: "Bob", Age: 30}},
			Entry[string, User]{Key: "user3", Value: User{Name: "Charlie", Age: 35}},
		)

		if len(result) != 3 {
			t.Errorf("SetMany() length = %d, want 3", len(result))
		}

		if result["user1"].Name != "Alice" || result["user1"].Age != 25 {
			t.Errorf("SetMany() result[user1] = %+v, want {Alice 25}", result["user1"])
		}

		if result["user2"].Name != "Bob" || result["user2"].Age != 30 {
			t.Errorf("SetMany() result[user2] = %+v, want {Bob 30}", result["user2"])
		}

		if result["user3"].Name != "Charlie" || result["user3"].Age != 35 {
			t.Errorf("SetMany() result[user3] = %+v, want {Charlie 35}", result["user3"])
		}
	})

	t.Run("large number of entries", func(t *testing.T) {
		m := map[int]int{0: 0}

		entries := make([]Entry[int, int], 100)
		for i := 0; i < 100; i++ {
			entries[i] = Entry[int, int]{Key: i + 1, Value: (i + 1) * 2}
		}

		result := SetMany(m, entries...)

		if len(result) != 101 {
			t.Errorf("SetMany() length = %d, want 101", len(result))
		}

		for i := 0; i <= 100; i++ {
			expected := i * 2
			if result[i] != expected {
				t.Errorf("SetMany() result[%d] = %d, want %d", i, result[i], expected)
			}
		}
	})

	t.Run("immutability check", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		original := make(map[string]int)
		for k, v := range m {
			original[k] = v
		}

		_ = SetMany(m, Entry[string, int]{Key: "c", Value: 3})

		// Original map should remain unchanged
		if len(m) != len(original) {
			t.Errorf("SetMany() modified original map, length = %d, want %d", len(m), len(original))
		}

		for k, v := range original {
			if m[k] != v {
				t.Errorf("SetMany() modified original map[%s] = %d, want %d", k, m[k], v)
			}
		}
	})
}

// BenchmarkSetMany benchmarks the SetMany function.
// BenchmarkSetMany는 SetMany 함수를 벤치마크합니다.
func BenchmarkSetMany(b *testing.B) {
	entryCounts := []int{1, 5, 10, 50, 100}

	for _, entryCount := range entryCounts {
		b.Run(fmt.Sprintf("entries_%d", entryCount), func(b *testing.B) {
			m := make(map[int]int, 10)
			for i := 0; i < 10; i++ {
				m[i] = i * 2
			}

			entries := make([]Entry[int, int], entryCount)
			for i := 0; i < entryCount; i++ {
				entries[i] = Entry[int, int]{Key: i + 100, Value: i * 3}
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = SetMany(m, entries...)
			}
		})
	}
}

// BenchmarkSetManyVsLoop compares SetMany with manual loop.
// BenchmarkSetManyVsLoop은 SetMany와 수동 루프를 비교합니다.
func BenchmarkSetManyVsLoop(b *testing.B) {
	m := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		m[i] = i * 2
	}

	entries := []Entry[int, int]{
		{Key: 100, Value: 300},
		{Key: 101, Value: 303},
		{Key: 102, Value: 306},
		{Key: 103, Value: 309},
		{Key: 104, Value: 312},
	}

	b.Run("SetMany", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SetMany(m, entries...)
		}
	})

	b.Run("ManualLoop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := make(map[int]int, len(m)+len(entries))
			for k, v := range m {
				result[k] = v
			}
			for _, entry := range entries {
				result[entry.Key] = entry.Value
			}
		}
	})
}

// TestTap tests the Tap function with various scenarios.
// TestTap는 다양한 시나리오로 Tap 함수를 테스트합니다.
func TestTap(t *testing.T) {
	t.Run("basic tap", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		var tapped map[string]int

		result := Tap(m, func(m map[string]int) {
			tapped = m
		})

		// Should return the same map
		if len(result) != len(m) {
			t.Errorf("Tap() length = %d, want %d", len(result), len(m))
		}

		for k, v := range m {
			if result[k] != v {
				t.Errorf("Tap() result[%s] = %d, want %d", k, result[k], v)
			}
		}

		// Should have called the function with the map
		if len(tapped) != len(m) {
			t.Errorf("Tap() function not called correctly, tapped length = %d, want %d", len(tapped), len(m))
		}
	})

	t.Run("logging use case", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		var logged []string

		result := Tap(m, func(m map[string]int) {
			for k, v := range m {
				logged = append(logged, fmt.Sprintf("%s=%d", k, v))
			}
		})

		// Map should be unchanged
		if len(result) != 3 {
			t.Errorf("Tap() length = %d, want 3", len(result))
		}

		// Should have logged all entries
		if len(logged) != 3 {
			t.Errorf("Tap() logged %d entries, want 3", len(logged))
		}
	})

	t.Run("chaining pattern", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		var intermediate map[string]int

		// Simulate a chain: Filter -> Tap -> Map
		filtered := Filter(m, func(k string, v int) bool {
			return v > 1
		})

		tapped := Tap(filtered, func(m map[string]int) {
			intermediate = make(map[string]int)
			for k, v := range m {
				intermediate[k] = v
			}
		})

		result := Map(tapped, func(k string, v int) int {
			return v * 2
		})

		// Should have captured intermediate state
		if len(intermediate) != 3 {
			t.Errorf("Tap() intermediate length = %d, want 3", len(intermediate))
		}

		// Final result should be correct
		expected := map[string]int{"b": 4, "c": 6, "d": 8}
		if len(result) != len(expected) {
			t.Errorf("Chained operation length = %d, want %d", len(result), len(expected))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("Chained operation result[%s] = %d, want %d", k, result[k], v)
			}
		}
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[string]int{}
		called := false

		result := Tap(m, func(m map[string]int) {
			called = true
		})

		if !called {
			t.Errorf("Tap() function not called for empty map")
		}

		if len(result) != 0 {
			t.Errorf("Tap() length = %d, want 0", len(result))
		}
	})

	t.Run("statistics collection", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
		var sum, count int

		result := Tap(m, func(m map[string]int) {
			for _, v := range m {
				sum += v
				count++
			}
		})

		// Should have collected statistics
		if sum != 15 {
			t.Errorf("Tap() sum = %d, want 15", sum)
		}

		if count != 5 {
			t.Errorf("Tap() count = %d, want 5", count)
		}

		// Map should be unchanged
		if len(result) != 5 {
			t.Errorf("Tap() length = %d, want 5", len(result))
		}
	})

	t.Run("validation use case", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		var valid bool

		result := Tap(m, func(m map[string]int) {
			// Check if all values are positive
			valid = true
			for _, v := range m {
				if v <= 0 {
					valid = false
					break
				}
			}
		})

		if !valid {
			t.Errorf("Tap() validation failed, all values should be positive")
		}

		// Map should be unchanged
		if len(result) != 3 {
			t.Errorf("Tap() length = %d, want 3", len(result))
		}
	})

	t.Run("nil function", func(t *testing.T) {
		m := map[string]int{"a": 1}

		// Should handle nil function gracefully (will panic as expected)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Tap() with nil function should panic")
			}
		}()

		_ = Tap(m, nil)
	})

	t.Run("different types", func(t *testing.T) {
		m := map[int]string{1: "one", 2: "two", 3: "three"}
		var keys []int

		result := Tap(m, func(m map[int]string) {
			for k := range m {
				keys = append(keys, k)
			}
		})

		// Should have collected keys
		if len(keys) != 3 {
			t.Errorf("Tap() collected %d keys, want 3", len(keys))
		}

		// Map should be unchanged
		if len(result) != 3 {
			t.Errorf("Tap() length = %d, want 3", len(result))
		}

		for k, v := range m {
			if result[k] != v {
				t.Errorf("Tap() result[%d] = %s, want %s", k, result[k], v)
			}
		}
	})
}

// BenchmarkTap benchmarks the Tap function.
// BenchmarkTap는 Tap 함수를 벤치마크합니다.
func BenchmarkTap(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			m := make(map[int]int, size)
			for i := 0; i < size; i++ {
				m[i] = i * 2
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Tap(m, func(m map[int]int) {
					sum := 0
					for _, v := range m {
						sum += v
					}
				})
			}
		})
	}
}

// BenchmarkTapVsInline compares Tap with inline side effect.
// BenchmarkTapVsInline은 Tap과 인라인 부수 효과를 비교합니다.
func BenchmarkTapVsInline(b *testing.B) {
	m := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		m[i] = i * 2
	}

	b.Run("Tap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Tap(m, func(m map[int]int) {
				sum := 0
				for _, v := range m {
					sum += v
				}
			})
		}
	})

	b.Run("Inline", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sum := 0
			for _, v := range m {
				sum += v
			}
			_ = m
		}
	})
}
