package maputil

import (
	"testing"
)

// TestGetOrSet tests the GetOrSet function with various scenarios.
// TestGetOrSet는 다양한 시나리오로 GetOrSet 함수를 테스트합니다.
func TestGetOrSet(t *testing.T) {
	t.Run("GetExistingKey", func(t *testing.T) {
		// Test getting an existing key
		// 기존 키 가져오기 테스트
		m := map[string]int{"a": 1, "b": 2}
		result := GetOrSet(m, "a", 10)
		if result != 1 {
			t.Errorf("GetOrSet(existing key) = %v, want 1", result)
		}
		if m["a"] != 1 {
			t.Errorf("Map value should remain unchanged: got %v, want 1", m["a"])
		}
	})

	t.Run("SetNewKey", func(t *testing.T) {
		// Test setting a new key
		// 새 키 설정 테스트
		m := map[string]int{"a": 1, "b": 2}
		result := GetOrSet(m, "c", 10)
		if result != 10 {
			t.Errorf("GetOrSet(new key) = %v, want 10", result)
		}
		if m["c"] != 10 {
			t.Errorf("Map should have new key: got %v, want 10", m["c"])
		}
	})

	t.Run("EmptyMap", func(t *testing.T) {
		// Test with empty map
		// 빈 맵 테스트
		m := map[string]int{}
		result := GetOrSet(m, "x", 100)
		if result != 100 {
			t.Errorf("GetOrSet(empty map) = %v, want 100", result)
		}
		if len(m) != 1 || m["x"] != 100 {
			t.Errorf("Map should have one entry: got %v", m)
		}
	})

	t.Run("ZeroValue", func(t *testing.T) {
		// Test with zero value as default
		// 0 값을 기본값으로 테스트
		m := map[string]int{"a": 5}
		result := GetOrSet(m, "b", 0)
		if result != 0 {
			t.Errorf("GetOrSet(zero default) = %v, want 0", result)
		}
		if m["b"] != 0 {
			t.Errorf("Map should have zero value: got %v, want 0", m["b"])
		}
	})

	t.Run("StringMap", func(t *testing.T) {
		// Test with string map
		// 문자열 맵 테스트
		m := map[string]string{"name": "Alice"}
		result := GetOrSet(m, "age", "25")
		if result != "25" {
			t.Errorf("GetOrSet(string map) = %v, want '25'", result)
		}
		if m["age"] != "25" {
			t.Errorf("Map should have age: got %v, want '25'", m["age"])
		}
	})

	t.Run("StructMap", func(t *testing.T) {
		// Test with struct values
		// 구조체 값 테스트
		type User struct {
			Name string
			Age  int
		}
		m := map[int]User{1: {Name: "Alice", Age: 30}}
		defaultUser := User{Name: "Bob", Age: 25}
		result := GetOrSet(m, 2, defaultUser)
		if result.Name != "Bob" || result.Age != 25 {
			t.Errorf("GetOrSet(struct) = %+v, want %+v", result, defaultUser)
		}
	})

	t.Run("Cache", func(t *testing.T) {
		// Test cache usage pattern
		// 캐시 사용 패턴 테스트
		cache := make(map[string]int)

		// First access - should set value
		// 첫 번째 접근 - 값 설정
		val1 := GetOrSet(cache, "compute", 42)
		if val1 != 42 {
			t.Errorf("First access = %v, want 42", val1)
		}

		// Second access - should get cached value
		// 두 번째 접근 - 캐시된 값 가져오기
		val2 := GetOrSet(cache, "compute", 100)
		if val2 != 42 {
			t.Errorf("Cached access = %v, want 42", val2)
		}
	})
}

// TestSetDefault tests the SetDefault function with various scenarios.
// TestSetDefault는 다양한 시나리오로 SetDefault 함수를 테스트합니다.
func TestSetDefault(t *testing.T) {
	t.Run("SetNewKey", func(t *testing.T) {
		// Test setting a new key
		// 새 키 설정 테스트
		m := map[string]string{"host": "localhost"}
		wasSet := SetDefault(m, "port", "8080")
		if !wasSet {
			t.Error("SetDefault should return true for new key")
		}
		if m["port"] != "8080" {
			t.Errorf("Map should have port: got %v, want '8080'", m["port"])
		}
	})

	t.Run("ExistingKey", func(t *testing.T) {
		// Test with existing key (should not overwrite)
		// 기존 키 테스트 (덮어쓰지 않아야 함)
		m := map[string]string{"host": "localhost"}
		wasSet := SetDefault(m, "host", "0.0.0.0")
		if wasSet {
			t.Error("SetDefault should return false for existing key")
		}
		if m["host"] != "localhost" {
			t.Errorf("Map value should remain: got %v, want 'localhost'", m["host"])
		}
	})

	t.Run("EmptyMap", func(t *testing.T) {
		// Test with empty map
		// 빈 맵 테스트
		m := map[int]int{}
		wasSet := SetDefault(m, 1, 100)
		if !wasSet {
			t.Error("SetDefault should return true for empty map")
		}
		if m[1] != 100 {
			t.Errorf("Map should have value: got %v, want 100", m[1])
		}
	})

	t.Run("ZeroValue", func(t *testing.T) {
		// Test with zero value
		// 0 값 테스트
		m := map[string]int{"a": 5}
		wasSet := SetDefault(m, "b", 0)
		if !wasSet {
			t.Error("SetDefault should return true for new key with zero value")
		}
		if _, exists := m["b"]; !exists {
			t.Error("Map should have key 'b'")
		}
		if m["b"] != 0 {
			t.Errorf("Map should have zero value: got %v, want 0", m["b"])
		}
	})

	t.Run("MultipleDefaults", func(t *testing.T) {
		// Test setting multiple defaults
		// 여러 기본값 설정 테스트
		m := map[string]int{"a": 1}
		SetDefault(m, "b", 2)
		SetDefault(m, "c", 3)
		SetDefault(m, "a", 10) // Should not overwrite / 덮어쓰지 않아야 함

		if len(m) != 3 {
			t.Errorf("Map should have 3 entries: got %v", len(m))
		}
		if m["a"] != 1 || m["b"] != 2 || m["c"] != 3 {
			t.Errorf("Map values incorrect: got %v", m)
		}
	})

	t.Run("NilDefault", func(t *testing.T) {
		// Test with nil as default value
		// nil을 기본값으로 테스트
		m := map[string]*int{}
		wasSet := SetDefault(m, "x", nil)
		if !wasSet {
			t.Error("SetDefault should return true for new key with nil")
		}
		if _, exists := m["x"]; !exists {
			t.Error("Map should have key 'x'")
		}
	})

	t.Run("Configuration", func(t *testing.T) {
		// Test configuration usage pattern
		// 설정 사용 패턴 테스트
		config := map[string]string{
			"host": "localhost",
		}

		SetDefault(config, "port", "8080")
		SetDefault(config, "timeout", "30s")
		SetDefault(config, "host", "0.0.0.0") // Should not change / 변경되지 않아야 함

		if config["host"] != "localhost" {
			t.Error("Existing config should not be overwritten")
		}
		if config["port"] != "8080" || config["timeout"] != "30s" {
			t.Error("New config values should be set")
		}
	})
}

// TestDefaults tests the Defaults function with various scenarios.
// TestDefaults는 다양한 시나리오로 Defaults 함수를 테스트합니다.
func TestDefaults(t *testing.T) {
	t.Run("BasicMerge", func(t *testing.T) {
		// Test basic merge with defaults
		// 기본값과 기본 병합 테스트
		config := map[string]string{"host": "localhost"}
		defaults := map[string]string{
			"host":    "0.0.0.0",
			"port":    "8080",
			"timeout": "30s",
		}
		result := Defaults(config, defaults)

		if result["host"] != "localhost" {
			t.Errorf("Original value should be preserved: got %v", result["host"])
		}
		if result["port"] != "8080" || result["timeout"] != "30s" {
			t.Errorf("Default values should be added: got %v", result)
		}
	})

	t.Run("EmptyOriginal", func(t *testing.T) {
		// Test with empty original map
		// 빈 원본 맵 테스트
		config := map[string]int{}
		defaults := map[string]int{"a": 1, "b": 2, "c": 3}
		result := Defaults(config, defaults)

		if len(result) != 3 {
			t.Errorf("Result should have 3 entries: got %v", len(result))
		}
		if result["a"] != 1 || result["b"] != 2 || result["c"] != 3 {
			t.Errorf("Result should equal defaults: got %v", result)
		}
	})

	t.Run("EmptyDefaults", func(t *testing.T) {
		// Test with empty defaults map
		// 빈 기본값 맵 테스트
		config := map[string]int{"a": 1, "b": 2}
		defaults := map[string]int{}
		result := Defaults(config, defaults)

		if len(result) != 2 {
			t.Errorf("Result should have 2 entries: got %v", len(result))
		}
		if result["a"] != 1 || result["b"] != 2 {
			t.Errorf("Result should equal config: got %v", result)
		}
	})

	t.Run("BothEmpty", func(t *testing.T) {
		// Test with both maps empty
		// 두 맵 모두 빈 경우 테스트
		config := map[string]int{}
		defaults := map[string]int{}
		result := Defaults(config, defaults)

		if len(result) != 0 {
			t.Errorf("Result should be empty: got %v", result)
		}
	})

	t.Run("OriginalNotModified", func(t *testing.T) {
		// Test that original maps are not modified
		// 원본 맵이 수정되지 않는지 테스트
		config := map[string]int{"a": 1}
		defaults := map[string]int{"b": 2}
		result := Defaults(config, defaults)

		result["c"] = 3

		if len(config) != 1 {
			t.Errorf("Original config should not be modified: got %v", config)
		}
		if len(defaults) != 1 {
			t.Errorf("Original defaults should not be modified: got %v", defaults)
		}
	})

	t.Run("ComplexTypes", func(t *testing.T) {
		// Test with complex value types
		// 복잡한 값 타입 테스트
		type Config struct {
			Value   string
			Enabled bool
		}
		config := map[string]Config{
			"feature1": {Value: "custom", Enabled: true},
		}
		defaults := map[string]Config{
			"feature1": {Value: "default1", Enabled: false},
			"feature2": {Value: "default2", Enabled: true},
		}
		result := Defaults(config, defaults)

		if result["feature1"].Value != "custom" {
			t.Error("Original value should be preserved")
		}
		if result["feature2"].Value != "default2" {
			t.Error("Default value should be added")
		}
	})

	t.Run("Precedence", func(t *testing.T) {
		// Test that original values take precedence
		// 원본 값이 우선하는지 테스트
		config := map[string]int{"a": 100, "b": 200}
		defaults := map[string]int{"a": 1, "b": 2, "c": 3}
		result := Defaults(config, defaults)

		if result["a"] != 100 || result["b"] != 200 {
			t.Error("Original values should take precedence")
		}
		if result["c"] != 3 {
			t.Error("Default value should be added")
		}
	})

	t.Run("UserPreferences", func(t *testing.T) {
		// Test user preferences pattern
		// 사용자 기본 설정 패턴 테스트
		userPrefs := map[string]string{
			"theme": "dark",
		}
		systemDefaults := map[string]string{
			"theme":    "light",
			"language": "en",
			"timezone": "UTC",
		}
		finalPrefs := Defaults(userPrefs, systemDefaults)

		if finalPrefs["theme"] != "dark" {
			t.Error("User preference should override default")
		}
		if finalPrefs["language"] != "en" || finalPrefs["timezone"] != "UTC" {
			t.Error("System defaults should be applied")
		}
	})
}

// BenchmarkGetOrSet benchmarks the GetOrSet function.
// BenchmarkGetOrSet는 GetOrSet 함수를 벤치마크합니다.
func BenchmarkGetOrSet(b *testing.B) {
	m := map[int]int{1: 10, 2: 20, 3: 30}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetOrSet(m, i%10, 100)
	}
}

// BenchmarkGetOrSet_Existing benchmarks GetOrSet with existing keys.
// BenchmarkGetOrSet_Existing은 기존 키로 GetOrSet을 벤치마크합니다.
func BenchmarkGetOrSet_Existing(b *testing.B) {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetOrSet(m, i%1000, 100)
	}
}

// BenchmarkSetDefault benchmarks the SetDefault function.
// BenchmarkSetDefault는 SetDefault 함수를 벤치마크합니다.
func BenchmarkSetDefault(b *testing.B) {
	m := map[int]int{1: 10, 2: 20, 3: 30}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SetDefault(m, i%10, 100)
	}
}

// BenchmarkSetDefault_Existing benchmarks SetDefault with existing keys.
// BenchmarkSetDefault_Existing은 기존 키로 SetDefault를 벤치마크합니다.
func BenchmarkSetDefault_Existing(b *testing.B) {
	m := make(map[int]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SetDefault(m, i%1000, 100)
	}
}

// BenchmarkDefaults benchmarks the Defaults function.
// BenchmarkDefaults는 Defaults 함수를 벤치마크합니다.
func BenchmarkDefaults(b *testing.B) {
	config := map[string]int{"a": 1, "b": 2, "c": 3}
	defaults := map[string]int{"a": 10, "d": 4, "e": 5, "f": 6}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Defaults(config, defaults)
	}
}

// BenchmarkDefaults_Large benchmarks Defaults with large maps.
// BenchmarkDefaults_Large는 큰 맵으로 Defaults를 벤치마크합니다.
func BenchmarkDefaults_Large(b *testing.B) {
	config := make(map[int]int, 100)
	defaults := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		config[i] = i
		defaults[i+50] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Defaults(config, defaults)
	}
}
