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
