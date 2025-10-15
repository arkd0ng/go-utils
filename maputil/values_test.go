package maputil

import (
	"reflect"
	"testing"
)

// TestValuesSorted tests the ValuesSorted function / ValuesSorted 함수 테스트
func TestValuesSorted(t *testing.T) {
	data := map[string]int{"a": 3, "b": 1, "c": 2}

	result := ValuesSorted(data)
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ValuesSorted() = %v, want %v", result, expected)
	}

	// Test with strings / 문자열로 테스트
	strData := map[int]string{1: "c", 2: "a", 3: "b"}
	strResult := ValuesSorted(strData)
	expectedStr := []string{"a", "b", "c"}

	if !reflect.DeepEqual(strResult, expectedStr) {
		t.Errorf("ValuesSorted(strings) = %v, want %v", strResult, expectedStr)
	}
}

// TestUniqueValues tests the UniqueValues function / UniqueValues 함수 테스트
func TestUniqueValues(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3, "e": 2}

	result := UniqueValues(data)

	if len(result) != 3 {
		t.Errorf("UniqueValues() returned %d values, want 3", len(result))
	}

	// Check all unique values are present / 모든 고유 값이 존재하는지 확인
	hasOne := false
	hasTwo := false
	hasThree := false
	for _, v := range result {
		if v == 1 {
			hasOne = true
		}
		if v == 2 {
			hasTwo = true
		}
		if v == 3 {
			hasThree = true
		}
	}

	if !hasOne || !hasTwo || !hasThree {
		t.Errorf("UniqueValues() = %v, missing some unique values", result)
	}
}

// TestUniqueValuesEmpty tests UniqueValues with all same values / 모두 같은 값으로 UniqueValues 테스트
func TestUniqueValuesAllSame(t *testing.T) {
	data := map[string]int{"a": 1, "b": 1, "c": 1}

	result := UniqueValues(data)

	if len(result) != 1 || result[0] != 1 {
		t.Errorf("UniqueValues(all same) = %v, want [1]", result)
	}
}

// TestReplaceValue tests the ReplaceValue function / ReplaceValue 함수 테스트
func TestReplaceValue(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3, "e": 2}

	// Replace all 1s with 10 / 모든 1을 10으로 교체
	result := ReplaceValue(data, 1, 10)

	if result["a"] != 10 || result["c"] != 10 {
		t.Errorf("ReplaceValue(1, 10) failed to replace all 1s")
	}
	if result["b"] != 2 || result["d"] != 3 || result["e"] != 2 {
		t.Error("ReplaceValue(1, 10) changed other values")
	}

	// Test immutability / 불변성 테스트
	if data["a"] != 1 {
		t.Error("ReplaceValue() modified original map")
	}

	// Replace non-existent value / 존재하지 않는 값 교체
	result = ReplaceValue(data, 99, 100)
	if !reflect.DeepEqual(result, data) {
		t.Error("ReplaceValue(non-existent) changed map")
	}
}

// TestUpdateValues tests the UpdateValues function / UpdateValues 함수 테스트
func TestUpdateValues(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3}

	// Double all values / 모든 값을 2배로
	result := UpdateValues(data, func(k string, v int) int {
		return v * 2
	})

	expected := map[string]int{"a": 2, "b": 4, "c": 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("UpdateValues(*2) = %v, want %v", result, expected)
	}

	// Test immutability / 불변성 테스트
	if data["a"] != 1 {
		t.Error("UpdateValues() modified original map")
	}
}

// TestMinValue tests the MinValue function / MinValue 함수 테스트
func TestMinValue(t *testing.T) {
	data := map[string]int{"a": 3, "b": 1, "c": 2}

	min, found := MinValue(data)
	if !found {
		t.Fatal("MinValue() not found")
	}
	if min != 1 {
		t.Errorf("MinValue() = %d, want 1", min)
	}

	// Empty map / 빈 맵
	empty := map[string]int{}
	_, found = MinValue(empty)
	if found {
		t.Error("MinValue(empty) found, want not found")
	}
}

// TestMaxValue tests the MaxValue function / MaxValue 함수 테스트
func TestMaxValue(t *testing.T) {
	data := map[string]int{"a": 3, "b": 1, "c": 2}

	max, found := MaxValue(data)
	if !found {
		t.Fatal("MaxValue() not found")
	}
	if max != 3 {
		t.Errorf("MaxValue() = %d, want 3", max)
	}

	// Empty map / 빈 맵
	empty := map[string]int{}
	_, found = MaxValue(empty)
	if found {
		t.Error("MaxValue(empty) found, want not found")
	}
}

// TestSumValues tests the SumValues function / SumValues 함수 테스트
func TestSumValues(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	sum := SumValues(data)
	if sum != 10 {
		t.Errorf("SumValues() = %d, want 10", sum)
	}

	// Empty map / 빈 맵
	empty := map[string]int{}
	sum = SumValues(empty)
	if sum != 0 {
		t.Errorf("SumValues(empty) = %d, want 0", sum)
	}

	// Float values / 실수 값
	floatData := map[string]float64{"a": 1.5, "b": 2.5, "c": 3.0}
	floatSum := SumValues(floatData)
	if floatSum != 7.0 {
		t.Errorf("SumValues(float) = %f, want 7.0", floatSum)
	}
}

// Benchmarks / 벤치마크

func BenchmarkValuesSorted(b *testing.B) {
	data := map[string]int{"a": 5, "b": 3, "c": 1, "d": 4, "e": 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ValuesSorted(data)
	}
}

func BenchmarkUniqueValues(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3, "e": 2, "f": 1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = UniqueValues(data)
	}
}

func BenchmarkReplaceValue(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 1, "d": 3, "e": 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ReplaceValue(data, 1, 10)
	}
}

func BenchmarkUpdateValues(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string, v int) int { return v * 2 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = UpdateValues(data, fn)
	}
}

func BenchmarkMinValue(b *testing.B) {
	data := map[string]int{"a": 5, "b": 3, "c": 1, "d": 4, "e": 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MinValue(data)
	}
}

func BenchmarkMaxValue(b *testing.B) {
	data := map[string]int{"a": 5, "b": 3, "c": 1, "d": 4, "e": 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MaxValue(data)
	}
}

func BenchmarkSumValues(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumValues(data)
	}
}
