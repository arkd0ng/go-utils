package maputil

import (
	"reflect"
	"testing"
)

// TestKeysSorted tests the KeysSorted function / KeysSorted 함수 테스트
func TestKeysSorted(t *testing.T) {
	data := map[string]int{"c": 3, "a": 1, "b": 2}

	result := KeysSorted(data)
	expected := []string{"a", "b", "c"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("KeysSorted() = %v, want %v", result, expected)
	}

	// Test with numbers / 숫자로 테스트
	numData := map[int]string{3: "c", 1: "a", 2: "b"}
	numResult := KeysSorted(numData)
	expectedNum := []int{1, 2, 3}

	if !reflect.DeepEqual(numResult, expectedNum) {
		t.Errorf("KeysSorted(numbers) = %v, want %v", numResult, expectedNum)
	}
}

// TestFindKey tests the FindKey function / FindKey 함수 테스트
func TestFindKey(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Find key where value > 2 / 값이 2보다 큰 키 찾기
	key, found := FindKey(data, func(k string, v int) bool {
		return v > 2
	})

	if !found {
		t.Error("FindKey(v > 2) not found, want found")
	}
	if data[key] <= 2 {
		t.Errorf("FindKey(v > 2) returned key with value %d, want > 2", data[key])
	}

	// Find key that doesn't exist / 존재하지 않는 키 찾기
	key, found = FindKey(data, func(k string, v int) bool {
		return v > 10
	})

	if found {
		t.Error("FindKey(v > 10) found, want not found")
	}
}

// TestFindKeys tests the FindKeys function / FindKeys 함수 테스트
func TestFindKeys(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Find all keys where value is even / 값이 짝수인 모든 키 찾기
	result := FindKeys(data, func(k string, v int) bool {
		return v%2 == 0
	})

	if len(result) != 2 {
		t.Errorf("FindKeys(even) returned %d keys, want 2", len(result))
	}

	// Verify all found keys have even values / 찾은 모든 키가 짝수 값을 가지는지 확인
	for _, k := range result {
		if data[k]%2 != 0 {
			t.Errorf("FindKeys(even) returned key '%s' with odd value %d", k, data[k])
		}
	}
}

// TestRenameKey tests the RenameKey function / RenameKey 함수 테스트
func TestRenameKey(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	result := RenameKey(data, "b", "B")

	// Old key should not exist / 기존 키는 존재하지 않아야 함
	if _, ok := result["b"]; ok {
		t.Error("RenameKey() old key 'b' still exists")
	}

	// New key should exist with same value / 새 키가 같은 값으로 존재해야 함
	if result["B"] != 2 {
		t.Errorf("RenameKey() result['B'] = %d, want 2", result["B"])
	}

	// Other keys should remain / 다른 키들은 유지되어야 함
	if result["a"] != 1 || result["c"] != 3 {
		t.Error("RenameKey() affected other keys")
	}

	// Rename non-existent key should not change map / 존재하지 않는 키 이름 변경은 맵을 변경하지 않아야 함
	result = RenameKey(data, "x", "y")
	if !reflect.DeepEqual(result, data) {
		t.Error("RenameKey(non-existent) changed map")
	}

	// Test immutability / 불변성 테스트
	if _, ok := data["B"]; ok {
		t.Error("RenameKey() modified original map")
	}
}

// TestSwapKeys tests the SwapKeys function / SwapKeys 함수 테스트
func TestSwapKeys(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	result := SwapKeys(data, "a", "c")

	// Values should be swapped / 값이 교환되어야 함
	if result["a"] != 3 {
		t.Errorf("SwapKeys() result['a'] = %d, want 3", result["a"])
	}
	if result["c"] != 1 {
		t.Errorf("SwapKeys() result['c'] = %d, want 1", result["c"])
	}

	// Other keys unchanged / 다른 키는 변경되지 않음
	if result["b"] != 2 {
		t.Errorf("SwapKeys() result['b'] = %d, want 2", result["b"])
	}

	// Swap with non-existent key / 존재하지 않는 키와 교환
	result = SwapKeys(data, "a", "x")
	if !reflect.DeepEqual(result, data) {
		t.Error("SwapKeys(non-existent) changed map")
	}
}

// TestPrefixKeys tests the PrefixKeys function / PrefixKeys 함수 테스트
func TestPrefixKeys(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	result := PrefixKeys(data, "key_")

	expected := map[string]int{"key_a": 1, "key_b": 2, "key_c": 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PrefixKeys() = %v, want %v", result, expected)
	}

	// Test immutability / 불변성 테스트
	if _, ok := data["key_a"]; ok {
		t.Error("PrefixKeys() modified original map")
	}
}

// TestSuffixKeys tests the SuffixKeys function / SuffixKeys 함수 테스트
func TestSuffixKeys(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	result := SuffixKeys(data, "_key")

	expected := map[string]int{"a_key": 1, "b_key": 2, "c_key": 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SuffixKeys() = %v, want %v", result, expected)
	}
}

// TestTransformKeys tests the TransformKeys function / TransformKeys 함수 테스트
func TestTransformKeys(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	// Transform to uppercase / 대문자로 변환
	result := TransformKeys(data, func(k string) string {
		return k + k // Double the key
	})

	expected := map[string]int{"aa": 1, "bb": 2, "cc": 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TransformKeys() = %v, want %v", result, expected)
	}

	// Test with collisions (last value wins) / 충돌 테스트 (마지막 값 우선)
	data2 := map[string]int{"a": 1, "A": 2}
	result2 := TransformKeys(data2, func(k string) string {
		return "x"
	})
	if len(result2) != 1 || result2["x"] == 0 {
		t.Errorf("TransformKeys(collision) = %v, should have 1 key", result2)
	}
}

// Benchmarks / 벤치마크

func BenchmarkKeysSorted(b *testing.B) {
	data := map[string]int{"e": 5, "c": 3, "a": 1, "d": 4, "b": 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KeysSorted(data)
	}
}

func BenchmarkFindKey(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string, v int) bool { return v > 3 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FindKey(data, fn)
	}
}

func BenchmarkFindKeys(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string, v int) bool { return v%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FindKeys(data, fn)
	}
}

func BenchmarkRenameKey(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = RenameKey(data, "c", "C")
	}
}

func BenchmarkSwapKeys(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SwapKeys(data, "a", "e")
	}
}

func BenchmarkPrefixKeys(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PrefixKeys(data, "key_")
	}
}

func BenchmarkSuffixKeys(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SuffixKeys(data, "_key")
	}
}

func BenchmarkTransformKeys(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string) string { return "prefix_" + k }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = TransformKeys(data, fn)
	}
}
