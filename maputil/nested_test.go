package maputil

import (
	"strings"
	"testing"
)

// TestGetNested tests the GetNested function with various scenarios.
// TestGetNested는 다양한 시나리오로 GetNested 함수를 테스트합니다.
func TestGetNested(t *testing.T) {
	t.Run("SimpleNested", func(t *testing.T) {
		// Test simple nested access
		// 간단한 중첩 접근 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		val, ok := GetNested(data, "user", "name")
		if !ok || val != "Alice" {
			t.Errorf("GetNested(user, name) = %v, %v; want Alice, true", val, ok)
		}
	})

	t.Run("DeepNested", func(t *testing.T) {
		// Test deep nested access
		// 깊은 중첩 접근 테스트
		data := map[string]interface{}{
			"server": map[string]interface{}{
				"database": map[string]interface{}{
					"connection": map[string]interface{}{
						"host": "localhost",
						"port": 5432,
					},
				},
			},
		}
		val, ok := GetNested(data, "server", "database", "connection", "host")
		if !ok || val != "localhost" {
			t.Errorf("GetNested(deep path) = %v, %v; want localhost, true", val, ok)
		}
	})

	t.Run("MissingKey", func(t *testing.T) {
		// Test with missing key
		// 누락된 키 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		val, ok := GetNested(data, "user", "email")
		if ok || val != nil {
			t.Errorf("GetNested(missing key) = %v, %v; want nil, false", val, ok)
		}
	})

	t.Run("MissingIntermediateKey", func(t *testing.T) {
		// Test with missing intermediate key
		// 중간 키 누락 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		val, ok := GetNested(data, "admin", "name")
		if ok || val != nil {
			t.Errorf("GetNested(missing intermediate) = %v, %v; want nil, false", val, ok)
		}
	})

	t.Run("EmptyPath", func(t *testing.T) {
		// Test with empty path
		// 빈 경로 테스트
		data := map[string]interface{}{"key": "value"}
		val, ok := GetNested(data)
		if ok || val != nil {
			t.Errorf("GetNested(empty path) = %v, %v; want nil, false", val, ok)
		}
	})

	t.Run("NonMapIntermediate", func(t *testing.T) {
		// Test when intermediate value is not a map
		// 중간 값이 맵이 아닌 경우 테스트
		data := map[string]interface{}{
			"user": "Alice",
		}
		val, ok := GetNested(data, "user", "name")
		if ok || val != nil {
			t.Errorf("GetNested(non-map intermediate) = %v, %v; want nil, false", val, ok)
		}
	})

	t.Run("DifferentTypes", func(t *testing.T) {
		// Test with different value types
		// 다양한 값 타입 테스트
		data := map[string]interface{}{
			"config": map[string]interface{}{
				"timeout": 30,
				"enabled": true,
				"servers": []string{"s1", "s2"},
			},
		}
		timeout, ok := GetNested(data, "config", "timeout")
		if !ok || timeout != 30 {
			t.Errorf("GetNested(int) = %v, %v; want 30, true", timeout, ok)
		}
		enabled, ok := GetNested(data, "config", "enabled")
		if !ok || enabled != true {
			t.Errorf("GetNested(bool) = %v, %v; want true, true", enabled, ok)
		}
		_, ok = GetNested(data, "config", "servers")
		if !ok {
			t.Errorf("GetNested(slice) ok = %v; want true", ok)
		}
	})

	t.Run("SingleKey", func(t *testing.T) {
		// Test with single key path
		// 단일 키 경로 테스트
		data := map[string]interface{}{
			"name": "Alice",
		}
		val, ok := GetNested(data, "name")
		if !ok || val != "Alice" {
			t.Errorf("GetNested(single key) = %v, %v; want Alice, true", val, ok)
		}
	})
}

// TestSetNested tests the SetNested function with various scenarios.
// TestSetNested는 다양한 시나리오로 SetNested 함수를 테스트합니다.
func TestSetNested(t *testing.T) {
	t.Run("CreateNewPath", func(t *testing.T) {
		// Test creating new nested path
		// 새 중첩 경로 생성 테스트
		data := map[string]interface{}{}
		result := SetNested(data, "Seoul", "user", "address", "city")

		val, ok := GetNested(result, "user", "address", "city")
		if !ok || val != "Seoul" {
			t.Errorf("SetNested(new path) failed: got %v, want Seoul", val)
		}
	})

	t.Run("UpdateExisting", func(t *testing.T) {
		// Test updating existing value
		// 기존 값 업데이트 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		result := SetNested(data, "Bob", "user", "name")

		val, ok := GetNested(result, "user", "name")
		if !ok || val != "Bob" {
			t.Errorf("SetNested(update) = %v; want Bob", val)
		}
	})

	t.Run("Immutability", func(t *testing.T) {
		// Test that original map is not modified
		// 원본 맵이 수정되지 않는지 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		result := SetNested(data, "Bob", "user", "name")

		origVal, _ := GetNested(data, "user", "name")
		newVal, _ := GetNested(result, "user", "name")

		if origVal == newVal {
			t.Error("SetNested should not modify original map")
		}
		if origVal != "Alice" {
			t.Errorf("Original value = %v; want Alice", origVal)
		}
	})

	t.Run("OverwriteNonMap", func(t *testing.T) {
		// Test overwriting non-map intermediate value
		// 맵이 아닌 중간 값 덮어쓰기 테스트
		data := map[string]interface{}{
			"user": "Alice",
		}
		result := SetNested(data, "alice@example.com", "user", "email")

		val, ok := GetNested(result, "user", "email")
		if !ok || val != "alice@example.com" {
			t.Errorf("SetNested(overwrite) failed: got %v", val)
		}
	})

	t.Run("EmptyPath", func(t *testing.T) {
		// Test with empty path
		// 빈 경로 테스트
		data := map[string]interface{}{"key": "value"}
		result := SetNested(data, "new")

		if len(result) != 1 || result["key"] != "value" {
			t.Error("SetNested(empty path) should return original map")
		}
	})

	t.Run("SingleKey", func(t *testing.T) {
		// Test with single key
		// 단일 키 테스트
		data := map[string]interface{}{}
		result := SetNested(data, "Alice", "name")

		if result["name"] != "Alice" {
			t.Errorf("SetNested(single key) = %v; want Alice", result["name"])
		}
	})

	t.Run("DifferentTypes", func(t *testing.T) {
		// Test setting different value types
		// 다양한 값 타입 설정 테스트
		data := map[string]interface{}{}
		result := SetNested(data, 42, "config", "timeout")
		result = SetNested(result, true, "config", "enabled")
		result = SetNested(result, []string{"s1", "s2"}, "config", "servers")

		timeout, _ := GetNested(result, "config", "timeout")
		enabled, _ := GetNested(result, "config", "enabled")
		servers, _ := GetNested(result, "config", "servers")

		if timeout != 42 || enabled != true || len(servers.([]string)) != 2 {
			t.Error("SetNested(different types) failed")
		}
	})

	t.Run("DeepNesting", func(t *testing.T) {
		// Test deep nesting
		// 깊은 중첩 테스트
		data := map[string]interface{}{}
		result := SetNested(data, "value", "a", "b", "c", "d", "e")

		val, ok := GetNested(result, "a", "b", "c", "d", "e")
		if !ok || val != "value" {
			t.Errorf("SetNested(deep) failed: got %v", val)
		}
	})
}

// TestHasNested tests the HasNested function with various scenarios.
// TestHasNested는 다양한 시나리오로 HasNested 함수를 테스트합니다.
func TestHasNested(t *testing.T) {
	t.Run("ExistingPath", func(t *testing.T) {
		// Test with existing path
		// 기존 경로 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name":  "Alice",
				"email": "alice@example.com",
			},
		}
		if !HasNested(data, "user", "name") {
			t.Error("HasNested(existing) should return true")
		}
	})

	t.Run("MissingPath", func(t *testing.T) {
		// Test with missing path
		// 누락된 경로 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		if HasNested(data, "user", "phone") {
			t.Error("HasNested(missing) should return false")
		}
	})

	t.Run("MissingIntermediate", func(t *testing.T) {
		// Test with missing intermediate key
		// 중간 키 누락 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		if HasNested(data, "admin", "name") {
			t.Error("HasNested(missing intermediate) should return false")
		}
	})

	t.Run("EmptyPath", func(t *testing.T) {
		// Test with empty path
		// 빈 경로 테스트
		data := map[string]interface{}{"key": "value"}
		if HasNested(data) {
			t.Error("HasNested(empty path) should return false")
		}
	})

	t.Run("NonMapIntermediate", func(t *testing.T) {
		// Test when intermediate is not a map
		// 중간 값이 맵이 아닌 경우 테스트
		data := map[string]interface{}{
			"user": "Alice",
		}
		if HasNested(data, "user", "name") {
			t.Error("HasNested(non-map) should return false")
		}
	})

	t.Run("SingleKey", func(t *testing.T) {
		// Test with single key
		// 단일 키 테스트
		data := map[string]interface{}{
			"name": "Alice",
		}
		if !HasNested(data, "name") {
			t.Error("HasNested(single key) should return true")
		}
	})

	t.Run("DeepPath", func(t *testing.T) {
		// Test with deep path
		// 깊은 경로 테스트
		data := map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": map[string]interface{}{
						"d": "value",
					},
				},
			},
		}
		if !HasNested(data, "a", "b", "c", "d") {
			t.Error("HasNested(deep) should return true")
		}
		if HasNested(data, "a", "b", "c", "e") {
			t.Error("HasNested(deep missing) should return false")
		}
	})
}

// TestDeleteNested tests the DeleteNested function with various scenarios.
// TestDeleteNested는 다양한 시나리오로 DeleteNested 함수를 테스트합니다.
func TestDeleteNested(t *testing.T) {
	t.Run("DeleteExisting", func(t *testing.T) {
		// Test deleting existing value
		// 기존 값 삭제 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name":     "Alice",
				"password": "secret",
			},
		}
		result := DeleteNested(data, "user", "password")

		if HasNested(result, "user", "password") {
			t.Error("DeleteNested should remove the key")
		}
		if !HasNested(result, "user", "name") {
			t.Error("DeleteNested should not affect other keys")
		}
	})

	t.Run("Immutability", func(t *testing.T) {
		// Test that original map is not modified
		// 원본 맵이 수정되지 않는지 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		result := DeleteNested(data, "user", "name")

		if !HasNested(data, "user", "name") {
			t.Error("DeleteNested should not modify original map")
		}
		if HasNested(result, "user", "name") {
			t.Error("DeleteNested should remove key from result")
		}
	})

	t.Run("DeleteMissing", func(t *testing.T) {
		// Test deleting non-existent key
		// 존재하지 않는 키 삭제 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		result := DeleteNested(data, "user", "email")

		if !HasNested(result, "user", "name") {
			t.Error("DeleteNested(missing) should not affect other keys")
		}
	})

	t.Run("DeleteMissingPath", func(t *testing.T) {
		// Test deleting with missing intermediate path
		// 중간 경로 누락으로 삭제 테스트
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"name": "Alice",
			},
		}
		result := DeleteNested(data, "admin", "name")

		// Should return unchanged
		if !HasNested(result, "user", "name") {
			t.Error("DeleteNested(missing path) should not affect map")
		}
	})

	t.Run("EmptyPath", func(t *testing.T) {
		// Test with empty path
		// 빈 경로 테스트
		data := map[string]interface{}{"key": "value"}
		result := DeleteNested(data)

		if len(result) != 1 || result["key"] != "value" {
			t.Error("DeleteNested(empty path) should return original map")
		}
	})

	t.Run("SingleKey", func(t *testing.T) {
		// Test with single key
		// 단일 키 테스트
		data := map[string]interface{}{
			"name": "Alice",
			"age":  30,
		}
		result := DeleteNested(data, "name")

		if _, exists := result["name"]; exists {
			t.Error("DeleteNested(single key) should remove key")
		}
		if result["age"] != 30 {
			t.Error("DeleteNested should not affect other keys")
		}
	})

	t.Run("DeepNesting", func(t *testing.T) {
		// Test deep nesting
		// 깊은 중첩 테스트
		data := map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": "value",
					"d": "keep",
				},
			},
		}
		result := DeleteNested(data, "a", "b", "c")

		if HasNested(result, "a", "b", "c") {
			t.Error("DeleteNested(deep) should remove nested key")
		}
		if !HasNested(result, "a", "b", "d") {
			t.Error("DeleteNested should not affect sibling keys")
		}
	})
}

// TestSafeGet tests the SafeGet function with various scenarios.
// TestSafeGet는 다양한 시나리오로 SafeGet 함수를 테스트합니다.
func TestSafeGet(t *testing.T) {
	t.Run("ValidPath", func(t *testing.T) {
		// Test with valid path
		// 유효한 경로 테스트
		data := map[string]interface{}{
			"server": map[string]interface{}{
				"host": "localhost",
				"port": 8080,
			},
		}
		val, err := SafeGet(data, "server", "host")
		if err != nil {
			t.Errorf("SafeGet(valid) error = %v; want nil", err)
		}
		if val != "localhost" {
			t.Errorf("SafeGet(valid) = %v; want localhost", val)
		}
	})

	t.Run("MissingKey", func(t *testing.T) {
		// Test with missing key
		// 누락된 키 테스트
		data := map[string]interface{}{
			"server": map[string]interface{}{
				"host": "localhost",
			},
		}
		val, err := SafeGet(data, "server", "port")
		if err == nil {
			t.Error("SafeGet(missing) should return error")
		}
		if val != nil {
			t.Errorf("SafeGet(missing) value = %v; want nil", val)
		}
		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("Error should mention 'not found': %v", err)
		}
	})

	t.Run("NonMapIntermediate", func(t *testing.T) {
		// Test when intermediate is not a map
		// 중간 값이 맵이 아닌 경우 테스트
		data := map[string]interface{}{
			"server": "localhost",
		}
		val, err := SafeGet(data, "server", "host")
		if err == nil {
			t.Error("SafeGet(non-map) should return error")
		}
		if val != nil {
			t.Errorf("SafeGet(non-map) value = %v; want nil", val)
		}
		if !strings.Contains(err.Error(), "not a map") {
			t.Errorf("Error should mention 'not a map': %v", err)
		}
	})

	t.Run("EmptyPath", func(t *testing.T) {
		// Test with empty path
		// 빈 경로 테스트
		data := map[string]interface{}{"key": "value"}
		val, err := SafeGet(data)
		if err == nil {
			t.Error("SafeGet(empty) should return error")
		}
		if val != nil {
			t.Errorf("SafeGet(empty) value = %v; want nil", val)
		}
		if !strings.Contains(err.Error(), "empty") {
			t.Errorf("Error should mention 'empty': %v", err)
		}
	})

	t.Run("DifferentTypes", func(t *testing.T) {
		// Test with different value types
		// 다양한 값 타입 테스트
		data := map[string]interface{}{
			"config": map[string]interface{}{
				"timeout": 30,
				"enabled": true,
			},
		}
		timeout, err := SafeGet(data, "config", "timeout")
		if err != nil || timeout != 30 {
			t.Errorf("SafeGet(int) = %v, %v; want 30, nil", timeout, err)
		}
		enabled, err := SafeGet(data, "config", "enabled")
		if err != nil || enabled != true {
			t.Errorf("SafeGet(bool) = %v, %v; want true, nil", enabled, err)
		}
	})

	t.Run("SingleKey", func(t *testing.T) {
		// Test with single key
		// 단일 키 테스트
		data := map[string]interface{}{
			"name": "Alice",
		}
		val, err := SafeGet(data, "name")
		if err != nil {
			t.Errorf("SafeGet(single) error = %v; want nil", err)
		}
		if val != "Alice" {
			t.Errorf("SafeGet(single) = %v; want Alice", val)
		}
	})

	t.Run("DeepPath", func(t *testing.T) {
		// Test with deep path
		// 깊은 경로 테스트
		data := map[string]interface{}{
			"a": map[string]interface{}{
				"b": map[string]interface{}{
					"c": "value",
				},
			},
		}
		val, err := SafeGet(data, "a", "b", "c")
		if err != nil {
			t.Errorf("SafeGet(deep) error = %v; want nil", err)
		}
		if val != "value" {
			t.Errorf("SafeGet(deep) = %v; want value", val)
		}
	})

	t.Run("NilInput", func(t *testing.T) {
		// Test with nil input
		// nil 입력 테스트
		var data interface{} = nil
		val, err := SafeGet(data, "key")
		if err == nil {
			t.Error("SafeGet(nil) should return error")
		}
		if val != nil {
			t.Errorf("SafeGet(nil) value = %v; want nil", val)
		}
	})
}

// BenchmarkGetNested benchmarks the GetNested function.
// BenchmarkGetNested는 GetNested 함수를 벤치마크합니다.
func BenchmarkGetNested(b *testing.B) {
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": "value",
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetNested(data, "a", "b", "c")
	}
}

// BenchmarkSetNested benchmarks the SetNested function.
// BenchmarkSetNested는 SetNested 함수를 벤치마크합니다.
func BenchmarkSetNested(b *testing.B) {
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SetNested(data, "value", "a", "b", "c")
	}
}

// BenchmarkHasNested benchmarks the HasNested function.
// BenchmarkHasNested는 HasNested 함수를 벤치마크합니다.
func BenchmarkHasNested(b *testing.B) {
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": "value",
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasNested(data, "a", "b", "c")
	}
}

// BenchmarkDeleteNested benchmarks the DeleteNested function.
// BenchmarkDeleteNested는 DeleteNested 함수를 벤치마크합니다.
func BenchmarkDeleteNested(b *testing.B) {
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": "value",
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DeleteNested(data, "a", "b", "c")
	}
}

// BenchmarkSafeGet benchmarks the SafeGet function.
// BenchmarkSafeGet는 SafeGet 함수를 벤치마크합니다.
func BenchmarkSafeGet(b *testing.B) {
	data := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": "value",
			},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SafeGet(data, "a", "b", "c")
	}
}
