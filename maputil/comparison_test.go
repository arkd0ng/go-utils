package maputil

import (
	"reflect"
	"sort"
	"testing"
)

// TestDiff tests the Diff function
// Diff 함수 테스트
func TestDiff(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 20, "c": 3, "d": 4}

	result := Diff(m1, m2)

	// Should contain keys that differ
	// 다른 키들을 포함해야 함
	if _, ok := result["b"]; !ok {
		t.Error("Diff() missing key 'b'")
	}
	if _, ok := result["d"]; !ok {
		t.Error("Diff() missing key 'd'")
	}

	// Should not contain keys with same values
	// 같은 값을 가진 키는 포함하지 않아야 함
	if _, ok := result["a"]; ok {
		t.Error("Diff() should not contain key 'a'")
	}
	if _, ok := result["c"]; ok {
		t.Error("Diff() should not contain key 'c'")
	}
}

// TestDiffKeys tests the DiffKeys function
// DiffKeys 함수 테스트
func TestDiffKeys(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 10, "b": 20, "c": 3, "d": 4}

	result := DiffKeys(m1, m2)

	// Should return keys that exist in m2 but not in m1, or have different values
	// m2에는 있지만 m1에는 없거나 값이 다른 키들을 반환해야 함
	sort.Strings(result)

	// 'a', 'b' have different values, 'd' is new
	// 'a', 'b'는 값이 다르고, 'd'는 새 키
	expectedLen := 3
	if len(result) != expectedLen {
		t.Errorf("DiffKeys() returned %d keys, want %d", len(result), expectedLen)
	}
}

// TestCompare tests the Compare function
// Compare 함수 테스트
func TestCompare(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 20, "d": 4}

	added, removed, modified := Compare(m1, m2)

	// Added: keys in m2 but not in m1
	// 추가됨: m2에는 있지만 m1에는 없는 키
	if !reflect.DeepEqual(added, map[string]int{"d": 4}) {
		t.Errorf("Compare() added = %v, want map[d:4]", added)
	}

	// Removed: keys in m1 but not in m2
	// 제거됨: m1에는 있지만 m2에는 없는 키
	if !reflect.DeepEqual(removed, map[string]int{"c": 3}) {
		t.Errorf("Compare() removed = %v, want map[c:3]", removed)
	}

	// Modified: keys in both but with different values
	// 수정됨: 둘 다 있지만 값이 다른 키
	if !reflect.DeepEqual(modified, map[string]int{"b": 20}) {
		t.Errorf("Compare() modified = %v, want map[b:20]", modified)
	}
}

// TestCompareEmpty tests Compare with empty maps
// 빈 맵으로 Compare 테스트
func TestCompareEmpty(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{}

	added, removed, modified := Compare(m1, m2)

	if len(added) != 0 {
		t.Errorf("Compare(m1, empty) added = %v, want empty", added)
	}
	if !reflect.DeepEqual(removed, m1) {
		t.Errorf("Compare(m1, empty) removed = %v, want %v", removed, m1)
	}
	if len(modified) != 0 {
		t.Errorf("Compare(m1, empty) modified = %v, want empty", modified)
	}

	// Reverse
	// 반대
	added, removed, modified = Compare(m2, m1)

	if !reflect.DeepEqual(added, m1) {
		t.Errorf("Compare(empty, m1) added = %v, want %v", added, m1)
	}
	if len(removed) != 0 {
		t.Errorf("Compare(empty, m1) removed = %v, want empty", removed)
	}
	if len(modified) != 0 {
		t.Errorf("Compare(empty, m1) modified = %v, want empty", modified)
	}
}

// TestCommonKeys tests the CommonKeys function
// CommonKeys 함수 테스트
func TestCommonKeys(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 2, "c": 4, "d": 5}
	m3 := map[string]int{"b": 2, "c": 3, "e": 6}

	result := CommonKeys(m1, m2, m3)

	// Only 'b' and 'c' appear in all three
	// 'b'와 'c'만 세 맵 모두에 존재
	if len(result) != 2 {
		t.Errorf("CommonKeys() returned %d keys, want 2", len(result))
	}

	sort.Strings(result)
	expected := []string{"b", "c"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CommonKeys() = %v, want %v", result, expected)
	}
}

// TestCommonKeysNoCommon tests CommonKeys with no common keys
// 공통 키가 없는 CommonKeys 테스트
func TestCommonKeysNoCommon(t *testing.T) {
	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"b": 2}
	m3 := map[string]int{"c": 3}

	result := CommonKeys(m1, m2, m3)

	if len(result) != 0 {
		t.Errorf("CommonKeys(no common) = %v, want []", result)
	}
}

// TestAllKeys tests the AllKeys function
// AllKeys 함수 테스트
func TestAllKeys(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	m3 := map[string]int{"c": 5, "d": 6}

	result := AllKeys(m1, m2, m3)

	// Should contain all unique keys
	// 모든 고유 키를 포함해야 함
	if len(result) != 4 {
		t.Errorf("AllKeys() returned %d keys, want 4", len(result))
	}

	sort.Strings(result)
	expected := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AllKeys() = %v, want %v", result, expected)
	}
}

// TestAllKeysEmpty tests AllKeys with empty map
// 빈 맵으로 AllKeys 테스트
func TestAllKeysEmpty(t *testing.T) {
	m1 := map[string]int{"a": 1}
	m2 := map[string]int{}

	result := AllKeys(m1, m2)

	if len(result) != 1 || result[0] != "a" {
		t.Errorf("AllKeys(with empty) = %v, want [a]", result)
	}
}

// TestEqualMaps tests the EqualMaps function
// EqualMaps 함수 테스트
func TestEqualMaps(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	m3 := map[string]int{"a": 1, "b": 20, "c": 3}

	if !EqualMaps(m1, m2) {
		t.Error("EqualMaps(same) = false, want true")
	}

	if EqualMaps(m1, m3) {
		t.Error("EqualMaps(different) = true, want false")
	}

	// Different lengths
	// 다른 길이
	m4 := map[string]int{"a": 1, "b": 2}
	if EqualMaps(m1, m4) {
		t.Error("EqualMaps(different length) = true, want false")
	}

	// Empty maps
	// 빈 맵
	empty1 := map[string]int{}
	empty2 := map[string]int{}
	if !EqualMaps(empty1, empty2) {
		t.Error("EqualMaps(empty) = false, want true")
	}
}

// Benchmarks
// 벤치마크

func BenchmarkDiff(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	m2 := map[string]int{"a": 1, "b": 20, "c": 3, "d": 40, "f": 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Diff(m1, m2)
	}
}

func BenchmarkDiffKeys(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	m2 := map[string]int{"a": 1, "b": 20, "c": 3, "d": 40, "f": 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DiffKeys(m1, m2)
	}
}

func BenchmarkCompare(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	m2 := map[string]int{"a": 1, "b": 20, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = Compare(m1, m2)
	}
}

func BenchmarkCommonKeys(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	m2 := map[string]int{"b": 2, "c": 4, "d": 5, "e": 6}
	m3 := map[string]int{"b": 2, "c": 3, "f": 7}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CommonKeys(m1, m2, m3)
	}
}

func BenchmarkAllKeys(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"c": 4, "d": 5, "e": 6}
	m3 := map[string]int{"e": 7, "f": 8, "g": 9}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AllKeys(m1, m2, m3)
	}
}

func BenchmarkEqualMaps(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EqualMaps(m1, m2)
	}
}
