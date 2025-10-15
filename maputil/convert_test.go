package maputil

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

// TestKeys tests the Keys function / Keys 함수 테스트
func TestKeys(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	result := Keys(data)

	if len(result) != 3 {
		t.Errorf("Keys() returned %d keys, want 3", len(result))
	}

	// Check all keys are present (order doesn't matter for maps)
	// 모든 키가 존재하는지 확인 (맵은 순서 무관)
	sort.Strings(result)
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Keys() = %v, want %v", result, expected)
	}
}

// TestKeysEmpty tests Keys with empty map / 빈 맵으로 Keys 테스트
func TestKeysEmpty(t *testing.T) {
	result := Keys(map[string]int{})
	if len(result) != 0 {
		t.Errorf("Keys(empty) = %v, want []", result)
	}
}

// TestValues tests the Values function / Values 함수 테스트
func TestValues(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	result := Values(data)

	if len(result) != 3 {
		t.Errorf("Values() returned %d values, want 3", len(result))
	}

	// Check all values are present / 모든 값이 존재하는지 확인
	sort.Ints(result)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Values() = %v, want %v", result, expected)
	}
}

// TestEntries tests the Entries function / Entries 함수 테스트
func TestEntries(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	result := Entries(data)

	if len(result) != 3 {
		t.Errorf("Entries() returned %d entries, want 3", len(result))
	}

	// Create map from entries to verify / 항목으로부터 맵을 생성하여 검증
	reconstructed := make(map[string]int)
	for _, entry := range result {
		reconstructed[entry.Key] = entry.Value
	}

	if !reflect.DeepEqual(reconstructed, data) {
		t.Errorf("Entries() produced incorrect entries")
	}
}

// TestFromEntries tests the FromEntries function / FromEntries 함수 테스트
func TestFromEntries(t *testing.T) {
	entries := []Entry[string, int]{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	result := FromEntries(entries)
	expected := map[string]int{"a": 1, "b": 2, "c": 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FromEntries() = %v, want %v", result, expected)
	}

	// Test with duplicate keys (should use last value) / 중복 키로 테스트 (마지막 값 사용)
	entries = []Entry[string, int]{
		{"a", 1},
		{"a", 10},
		{"b", 2},
	}
	result = FromEntries(entries)
	if result["a"] != 10 {
		t.Errorf("FromEntries() with duplicate keys: result['a'] = %d, want 10", result["a"])
	}
}

// TestToJSON tests the ToJSON function / ToJSON 함수 테스트
func TestToJSON(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	result, err := ToJSON(data)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}

	// Verify it's valid JSON containing all keys / 모든 키를 포함하는 유효한 JSON인지 확인
	if !strings.Contains(result, `"a"`) ||
		!strings.Contains(result, `"b"`) ||
		!strings.Contains(result, `"c"`) {
		t.Errorf("ToJSON() = %s, missing keys", result)
	}
}

// TestFromJSON tests the FromJSON function / FromJSON 함수 테스트
func TestFromJSON(t *testing.T) {
	jsonStr := `{"a":1,"b":2,"c":3}`

	var result map[string]int
	err := FromJSON(jsonStr, &result)
	if err != nil {
		t.Fatalf("FromJSON() error = %v", err)
	}

	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FromJSON() = %v, want %v", result, expected)
	}

	// Test with invalid JSON / 잘못된 JSON으로 테스트
	err = FromJSON(`{invalid}`, &result)
	if err == nil {
		t.Error("FromJSON() should return error for invalid JSON")
	}
}

// TestToSlice tests the ToSlice function / ToSlice 함수 테스트
func TestToSlice(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	result := ToSlice(data, func(k string, v int) string {
		return k + "=" + string(rune(v+'0'))
	})

	if len(result) != 3 {
		t.Errorf("ToSlice() returned %d elements, want 3", len(result))
	}

	// Sort for predictable comparison / 예측 가능한 비교를 위해 정렬
	sort.Strings(result)
	expected := []string{"a=1", "b=2", "c=3"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ToSlice() = %v, want %v", result, expected)
	}
}

// TestFromSlice tests the FromSlice function / FromSlice 함수 테스트
func TestFromSlice(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}

	// Create map with ID as key / ID를 키로 사용하는 맵 생성
	result := FromSlice(users, func(u User) int {
		return u.ID
	})

	expected := map[int]User{1: {ID: 1, Name: "Alice"}, 2: {ID: 2, Name: "Bob"}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FromSlice() = %v, want %v", result, expected)
	}

	// Test with duplicate IDs / 중복 ID로 테스트
	users = []User{{ID: 1, Name: "Alice"}, {ID: 1, Name: "Andy"}, {ID: 2, Name: "Bob"}}
	result = FromSlice(users, func(u User) int {
		return u.ID
	})
	// ID 1 should have last user / ID 1은 마지막 사용자를 가져야 함
	if result[1].Name != "Andy" {
		t.Errorf("FromSlice() with duplicates: result[1].Name = %s, want Andy", result[1].Name)
	}
}

// TestFromSliceBy tests the FromSliceBy function / FromSliceBy 함수 테스트
func TestFromSliceBy(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}

	// Create map with first character as key and length as value
	// 첫 문자를 키로, 길이를 값으로 사용하는 맵 생성
	result := FromSliceBy(slice,
		func(s string) string { return string(s[0]) },
		func(s string) int { return len(s) },
	)

	expected := map[string]int{"a": 5, "b": 6, "c": 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FromSliceBy() = %v, want %v", result, expected)
	}

	// Test with duplicate keys / 중복 키로 테스트
	slice = []string{"alice", "andy", "bob"}
	result = FromSliceBy(slice,
		func(s string) string { return string(s[0]) },
		func(s string) int { return len(s) },
	)
	// 'a' key should have last value / 'a' 키는 마지막 값을 가져야 함
	if result["a"] != 4 { // "andy" length
		t.Errorf("FromSliceBy() with duplicates: result['a'] = %d, want 4", result["a"])
	}
}

// Benchmarks / 벤치마크

func BenchmarkKeys(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Keys(data)
	}
}

func BenchmarkValues(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Values(data)
	}
}

func BenchmarkEntries(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Entries(data)
	}
}

func BenchmarkFromEntries(b *testing.B) {
	entries := []Entry[string, int]{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"d", 4},
		{"e", 5},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FromEntries(entries)
	}
}

func BenchmarkToJSON(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ToJSON(data)
	}
}

func BenchmarkFromJSON(b *testing.B) {
	jsonStr := `{"a":1,"b":2,"c":3,"d":4,"e":5}`
	var result map[string]int

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FromJSON(jsonStr, &result)
	}
}

func BenchmarkToSlice(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string, v int) string { return k }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToSlice(data, fn)
	}
}

func BenchmarkFromSlice(b *testing.B) {
	type User struct {
		ID   int
		Name string
	}
	users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}, {ID: 3, Name: "Charlie"}}
	fn := func(u User) int { return u.ID }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FromSlice(users, fn)
	}
}

func BenchmarkFromSliceBy(b *testing.B) {
	slice := []string{"apple", "banana", "cherry", "date", "elderberry"}
	keyFn := func(s string) string { return string(s[0]) }
	valueFn := func(s string) int { return len(s) }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FromSliceBy(slice, keyFn, valueFn)
	}
}

// TestToYAML tests the ToYAML function with various scenarios.
// TestToYAML는 다양한 시나리오로 ToYAML 함수를 테스트합니다.
func TestToYAML(t *testing.T) {
	t.Run("SimpleMap", func(t *testing.T) {
		// Test simple map / 간단한 맵 테스트
		data := map[string]int{"a": 1, "b": 2, "c": 3}
		result, err := ToYAML(data)
		if err != nil {
			t.Errorf("ToYAML() error = %v", err)
		}
		if result == "" {
			t.Error("ToYAML() returned empty string")
		}
		// Check that result contains keys
		if !strings.Contains(result, "a:") || !strings.Contains(result, "b:") || !strings.Contains(result, "c:") {
			t.Errorf("ToYAML() = %s, missing keys", result)
		}
	})

	t.Run("NestedMap", func(t *testing.T) {
		// Test nested map / 중첩 맵 테스트
		data := map[string]interface{}{
			"server": map[string]interface{}{
				"host": "localhost",
				"port": 8080,
			},
			"database": map[string]interface{}{
				"host": "localhost",
				"port": 5432,
			},
		}
		result, err := ToYAML(data)
		if err != nil {
			t.Errorf("ToYAML() error = %v", err)
		}
		if !strings.Contains(result, "server:") || !strings.Contains(result, "database:") {
			t.Errorf("ToYAML() = %s, missing nested keys", result)
		}
	})

	t.Run("EmptyMap", func(t *testing.T) {
		// Test empty map / 빈 맵 테스트
		data := map[string]int{}
		result, err := ToYAML(data)
		if err != nil {
			t.Errorf("ToYAML() error = %v", err)
		}
		if strings.TrimSpace(result) != "{}" {
			t.Errorf("ToYAML(empty) = %s, want {}", result)
		}
	})

	t.Run("StringValues", func(t *testing.T) {
		// Test with string values / 문자열 값 테스트
		data := map[string]string{
			"name":  "Alice",
			"email": "alice@example.com",
			"role":  "admin",
		}
		result, err := ToYAML(data)
		if err != nil {
			t.Errorf("ToYAML() error = %v", err)
		}
		if !strings.Contains(result, "Alice") || !strings.Contains(result, "alice@example.com") {
			t.Errorf("ToYAML() = %s, missing values", result)
		}
	})

	t.Run("MixedTypes", func(t *testing.T) {
		// Test with mixed types / 혼합 타입 테스트
		data := map[string]interface{}{
			"name":    "Alice",
			"age":     30,
			"active":  true,
			"balance": 1234.56,
		}
		result, err := ToYAML(data)
		if err != nil {
			t.Errorf("ToYAML() error = %v", err)
		}
		if result == "" {
			t.Error("ToYAML() returned empty string")
		}
	})

	t.Run("ArrayValues", func(t *testing.T) {
		// Test with array values / 배열 값 테스트
		data := map[string]interface{}{
			"numbers": []int{1, 2, 3},
			"strings": []string{"a", "b", "c"},
		}
		result, err := ToYAML(data)
		if err != nil {
			t.Errorf("ToYAML() error = %v", err)
		}
		if result == "" {
			t.Error("ToYAML() returned empty string")
		}
	})
}

// TestFromYAML tests the FromYAML function with various scenarios.
// TestFromYAML는 다양한 시나리오로 FromYAML 함수를 테스트합니다.
func TestFromYAML(t *testing.T) {
	t.Run("SimpleYAML", func(t *testing.T) {
		// Test simple YAML / 간단한 YAML 테스트
		yamlStr := `
a: 1
b: 2
c: 3
`
		result, err := FromYAML(yamlStr)
		if err != nil {
			t.Errorf("FromYAML() error = %v", err)
		}
		if len(result) != 3 {
			t.Errorf("FromYAML() returned %d keys, want 3", len(result))
		}
		if result["a"] != 1 || result["b"] != 2 || result["c"] != 3 {
			t.Errorf("FromYAML() = %v, values incorrect", result)
		}
	})

	t.Run("NestedYAML", func(t *testing.T) {
		// Test nested YAML / 중첩 YAML 테스트
		yamlStr := `
server:
  host: localhost
  port: 8080
database:
  host: localhost
  port: 5432
`
		result, err := FromYAML(yamlStr)
		if err != nil {
			t.Errorf("FromYAML() error = %v", err)
		}
		if len(result) != 2 {
			t.Errorf("FromYAML() returned %d keys, want 2", len(result))
		}

		server, ok := result["server"].(map[string]interface{})
		if !ok {
			t.Error("FromYAML() server is not a map")
		}
		if server["host"] != "localhost" || server["port"] != 8080 {
			t.Errorf("FromYAML() server values incorrect: %v", server)
		}
	})

	t.Run("EmptyYAML", func(t *testing.T) {
		// Test empty YAML / 빈 YAML 테스트
		yamlStr := `{}`
		result, err := FromYAML(yamlStr)
		if err != nil {
			t.Errorf("FromYAML() error = %v", err)
		}
		if len(result) != 0 {
			t.Errorf("FromYAML(empty) = %v, want empty map", result)
		}
	})

	t.Run("StringValues", func(t *testing.T) {
		// Test with string values / 문자열 값 테스트
		yamlStr := `
name: Alice
email: alice@example.com
role: admin
`
		result, err := FromYAML(yamlStr)
		if err != nil {
			t.Errorf("FromYAML() error = %v", err)
		}
		if result["name"] != "Alice" || result["email"] != "alice@example.com" || result["role"] != "admin" {
			t.Errorf("FromYAML() = %v, values incorrect", result)
		}
	})

	t.Run("MixedTypes", func(t *testing.T) {
		// Test with mixed types / 혼합 타입 테스트
		yamlStr := `
name: Alice
age: 30
active: true
balance: 1234.56
`
		result, err := FromYAML(yamlStr)
		if err != nil {
			t.Errorf("FromYAML() error = %v", err)
		}
		if result["name"] != "Alice" || result["age"] != 30 || result["active"] != true {
			t.Errorf("FromYAML() = %v, values incorrect", result)
		}
	})

	t.Run("ArrayValues", func(t *testing.T) {
		// Test with array values / 배열 값 테스트
		yamlStr := `
numbers:
  - 1
  - 2
  - 3
strings:
  - a
  - b
  - c
`
		result, err := FromYAML(yamlStr)
		if err != nil {
			t.Errorf("FromYAML() error = %v", err)
		}

		numbers, ok := result["numbers"].([]interface{})
		if !ok {
			t.Error("FromYAML() numbers is not a slice")
		}
		if len(numbers) != 3 {
			t.Errorf("FromYAML() numbers length = %d, want 3", len(numbers))
		}
	})

	t.Run("InvalidYAML", func(t *testing.T) {
		// Test with invalid YAML / 유효하지 않은 YAML 테스트
		yamlStr := `invalid: yaml: content:`
		_, err := FromYAML(yamlStr)
		if err == nil {
			t.Error("FromYAML(invalid) should return error")
		}
	})

	t.Run("RoundTrip", func(t *testing.T) {
		// Test round-trip conversion / 왕복 변환 테스트
		original := map[string]interface{}{
			"name": "Alice",
			"age":  30,
		}

		// Convert to YAML
		yamlStr, err := ToYAML(original)
		if err != nil {
			t.Errorf("ToYAML() error = %v", err)
		}

		// Convert back to map
		result, err := FromYAML(yamlStr)
		if err != nil {
			t.Errorf("FromYAML() error = %v", err)
		}

		// Check values
		if result["name"] != "Alice" || result["age"] != 30 {
			t.Errorf("Round-trip failed: got %v, want %v", result, original)
		}
	})
}

// BenchmarkToYAML benchmarks the ToYAML function.
// BenchmarkToYAML는 ToYAML 함수를 벤치마크합니다.
func BenchmarkToYAML(b *testing.B) {
	data := map[string]interface{}{
		"server": map[string]interface{}{
			"host": "localhost",
			"port": 8080,
		},
		"database": map[string]interface{}{
			"host": "localhost",
			"port": 5432,
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ToYAML(data)
	}
}

// BenchmarkFromYAML benchmarks the FromYAML function.
// BenchmarkFromYAML는 FromYAML 함수를 벤치마크합니다.
func BenchmarkFromYAML(b *testing.B) {
	yamlStr := `
server:
  host: localhost
  port: 8080
database:
  host: localhost
  port: 5432
`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FromYAML(yamlStr)
	}
}
