package maputil

import (
	"reflect"
	"testing"
)

// TestMerge tests the Merge function
// Merge 함수 테스트
func TestMerge(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	m3 := map[string]int{"c": 5, "d": 6}

	result := Merge(m1, m2, m3)
	expected := map[string]int{"a": 1, "b": 3, "c": 5, "d": 6}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Merge() = %v, want %v", result, expected)
	}

	// Test immutability
	// 불변성 테스트
	if len(m1) != 2 {
		t.Error("Merge() modified original map m1")
	}
}

// TestMergeEmpty tests Merge with empty maps
// 빈 맵으로 Merge 테스트
func TestMergeEmpty(t *testing.T) {
	result := Merge[string, int]()
	if len(result) != 0 {
		t.Errorf("Merge() with no args = %v, want empty map", result)
	}

	m1 := map[string]int{"a": 1}
	result = Merge(m1, map[string]int{})
	if !reflect.DeepEqual(result, m1) {
		t.Errorf("Merge() with empty map = %v, want %v", result, m1)
	}
}

// TestMergeWith tests the MergeWith function
// MergeWith 함수 테스트
func TestMergeWith(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 3, "c": 1}

	// Sum values on conflict
	// 충돌 시 값 합산
	result := MergeWith(func(v1, v2 int) int { return v1 + v2 }, m1, m2)
	expected := map[string]int{"a": 1, "b": 5, "c": 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeWith() = %v, want %v", result, expected)
	}

	// Max value on conflict
	// 충돌 시 최댓값
	result = MergeWith(func(v1, v2 int) int {
		if v1 > v2 {
			return v1
		}
		return v2
	}, m1, m2)
	expected = map[string]int{"a": 1, "b": 3, "c": 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeWith(max) = %v, want %v", result, expected)
	}
}

// TestDeepMerge tests the DeepMerge function
// DeepMerge 함수 테스트
func TestDeepMerge(t *testing.T) {
	m1 := map[string]interface{}{
		"a": map[string]interface{}{"x": 1, "y": 2},
		"b": map[string]interface{}{"z": 3},
	}
	m2 := map[string]interface{}{
		"a": map[string]interface{}{"y": 3, "z": 4},
		"c": map[string]interface{}{"w": 5},
	}

	result := DeepMerge(m1, m2)

	// Check top-level keys
	// 최상위 레벨 키 확인
	if _, ok := result["a"]; !ok {
		t.Error("DeepMerge() missing key 'a'")
	}
	if _, ok := result["b"]; !ok {
		t.Error("DeepMerge() missing key 'b'")
	}
	if _, ok := result["c"]; !ok {
		t.Error("DeepMerge() missing key 'c'")
	}
}

// TestUnion tests the Union function
// Union 함수 테스트
func TestUnion(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	m3 := map[string]int{"c": 5, "d": 6}

	result := Union(m1, m2, m3)

	// Union should contain all keys
	// Union은 모든 키를 포함해야 함
	if len(result) != 4 {
		t.Errorf("Union() has %d keys, want 4", len(result))
	}

	if _, ok := result["a"]; !ok {
		t.Error("Union() missing key 'a'")
	}
	if _, ok := result["d"]; !ok {
		t.Error("Union() missing key 'd'")
	}
}

// TestIntersection tests the Intersection function
// Intersection 함수 테스트
func TestIntersection(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 20, "c": 30, "d": 40}
	m3 := map[string]int{"c": 100, "e": 50}

	result := Intersection(m1, m2, m3)

	// Only 'c' appears in all three maps
	// 'c'만 모든 맵에 존재
	if len(result) != 1 {
		t.Errorf("Intersection() = %v, want map with 1 key", result)
	}
	// Value should be from first map
	// 값은 첫 번째 맵에서
	if result["c"] != 3 {
		t.Errorf("Intersection()['c'] = %d, want 3", result["c"])
	}
}

// TestIntersectionEmpty tests Intersection with no common keys
// 공통 키가 없는 Intersection 테스트
func TestIntersectionEmpty(t *testing.T) {
	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"b": 2}

	result := Intersection(m1, m2)
	if len(result) != 0 {
		t.Errorf("Intersection() = %v, want empty map", result)
	}
}

// TestDifference tests the Difference function
// Difference 함수 테스트
func TestDifference(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 2, "c": 4, "d": 5}

	result := Difference(m1, m2)
	expected := map[string]int{"a": 1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Difference() = %v, want %v", result, expected)
	}

	// Test empty result
	// 빈 결과 테스트
	m3 := map[string]int{"a": 1, "b": 2, "c": 3}
	m4 := map[string]int{"a": 1, "b": 2, "c": 3}
	result = Difference(m3, m4)
	if len(result) != 0 {
		t.Errorf("Difference(same maps) = %v, want empty map", result)
	}
}

// TestSymmetricDifference tests the SymmetricDifference function
// SymmetricDifference 함수 테스트
func TestSymmetricDifference(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 2, "c": 4, "d": 5}

	result := SymmetricDifference(m1, m2)

	// Should contain keys that exist in only one map (not both)
	// 한 맵에만 존재하는 키들을 포함해야 함 (둘 다 아님)
	if _, ok := result["a"]; !ok {
		t.Error("SymmetricDifference() missing key 'a'")
	}
	if _, ok := result["d"]; !ok {
		t.Error("SymmetricDifference() missing key 'd'")
	}
	// Keys 'b' and 'c' exist in both maps, so they should not be in the result
	// 키 'b'와 'c'는 두 맵 모두에 존재하므로 결과에 없어야 함
	if _, ok := result["b"]; ok {
		t.Error("SymmetricDifference() should not contain key 'b'")
	}
	if _, ok := result["c"]; ok {
		t.Error("SymmetricDifference() should not contain key 'c'")
	}
}

// Benchmarks
// 벤치마크

func BenchmarkMerge(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"d": 4, "e": 5, "f": 6}
	m3 := map[string]int{"g": 7, "h": 8, "i": 9}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Merge(m1, m2, m3)
	}
}

func BenchmarkMergeWith(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 3, "c": 1, "d": 4}
	fn := func(v1, v2 int) int { return v1 + v2 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MergeWith(fn, m1, m2)
	}
}

func BenchmarkUnion(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"d": 4, "e": 5, "f": 6}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Union(m1, m2)
	}
}

func BenchmarkIntersection(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	m2 := map[string]int{"b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Intersection(m1, m2)
	}
}

func BenchmarkDifference(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	m2 := map[string]int{"b": 2, "c": 3, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Difference(m1, m2)
	}
}

func BenchmarkSymmetricDifference(b *testing.B) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 2, "c": 4, "d": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SymmetricDifference(m1, m2)
	}
}
