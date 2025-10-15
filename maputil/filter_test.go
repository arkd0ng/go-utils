package maputil

import (
	"reflect"
	"testing"
)

// TestFilter tests the Filter function / Filter 함수 테스트
func TestFilter(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Filter values > 2 / 2보다 큰 값 필터링
	result := Filter(data, func(k string, v int) bool {
		return v > 2
	})
	expected := map[string]int{"c": 3, "d": 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter(v > 2) = %v, want %v", result, expected)
	}

	// Filter keys >= 'b' / 'b' 이상 키 필터링
	result = Filter(data, func(k string, v int) bool {
		return k >= "b"
	})
	expected = map[string]int{"b": 2, "c": 3, "d": 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter(k >= 'b') = %v, want %v", result, expected)
	}

	// Test immutability / 불변성 테스트
	if len(data) != 4 {
		t.Error("Filter() modified original map")
	}
}

// TestFilterEmpty tests Filter with empty result / 빈 결과로 Filter 테스트
func TestFilterEmpty(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2}

	result := Filter(data, func(k string, v int) bool {
		return v > 10
	})

	if len(result) != 0 {
		t.Errorf("Filter(v > 10) = %v, want empty map", result)
	}
}

// TestFilterKeys tests the FilterKeys function / FilterKeys 함수 테스트
func TestFilterKeys(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	result := FilterKeys(data, func(k string) bool {
		return k >= "b"
	})
	expected := map[string]int{"b": 2, "c": 3, "d": 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FilterKeys() = %v, want %v", result, expected)
	}
}

// TestFilterValues tests the FilterValues function / FilterValues 함수 테스트
func TestFilterValues(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Filter even values / 짝수 값만 필터링
	result := FilterValues(data, func(v int) bool {
		return v%2 == 0
	})
	expected := map[string]int{"b": 2, "d": 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FilterValues(even) = %v, want %v", result, expected)
	}
}

// TestPick tests the Pick function / Pick 함수 테스트
func TestPick(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	result := Pick(data, "a", "c")
	expected := map[string]int{"a": 1, "c": 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Pick('a', 'c') = %v, want %v", result, expected)
	}

	// Test with non-existent keys / 존재하지 않는 키로 테스트
	result = Pick(data, "x", "y")
	if len(result) != 0 {
		t.Errorf("Pick(non-existent) = %v, want empty map", result)
	}

	// Test with mixed keys / 혼합 키로 테스트
	result = Pick(data, "a", "x", "c")
	expected = map[string]int{"a": 1, "c": 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Pick(mixed) = %v, want %v", result, expected)
	}
}

// TestOmit tests the Omit function / Omit 함수 테스트
func TestOmit(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	result := Omit(data, "b", "d")
	expected := map[string]int{"a": 1, "c": 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Omit('b', 'd') = %v, want %v", result, expected)
	}

	// Test with non-existent keys / 존재하지 않는 키로 테스트
	result = Omit(data, "x", "y")
	if !reflect.DeepEqual(result, data) {
		t.Errorf("Omit(non-existent) = %v, want %v", result, data)
	}
}

// TestPickBy tests the PickBy function / PickBy 함수 테스트
func TestPickBy(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Pick entries where value is even / 값이 짝수인 항목만 선택
	result := PickBy(data, func(k string, v int) bool {
		return v%2 == 0
	})
	expected := map[string]int{"b": 2, "d": 4}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PickBy(even) = %v, want %v", result, expected)
	}

	// Pick entries where key < 'c' / 키가 'c'보다 작은 항목만 선택
	result = PickBy(data, func(k string, v int) bool {
		return k < "c"
	})
	expected = map[string]int{"a": 1, "b": 2}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PickBy(k < 'c') = %v, want %v", result, expected)
	}
}

// TestOmitBy tests the OmitBy function / OmitBy 함수 테스트
func TestOmitBy(t *testing.T) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	// Omit entries where value is even / 값이 짝수인 항목 제외
	result := OmitBy(data, func(k string, v int) bool {
		return v%2 == 0
	})
	expected := map[string]int{"a": 1, "c": 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("OmitBy(even) = %v, want %v", result, expected)
	}
}

// Benchmarks / 벤치마크

func BenchmarkFilter(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string, v int) bool { return v > 2 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Filter(data, fn)
	}
}

func BenchmarkFilterKeys(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string) bool { return k >= "c" }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterKeys(data, fn)
	}
}

func BenchmarkFilterValues(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(v int) bool { return v%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterValues(data, fn)
	}
}

func BenchmarkPick(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Pick(data, "a", "c", "e")
	}
}

func BenchmarkOmit(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Omit(data, "b", "d")
	}
}

func BenchmarkPickBy(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string, v int) bool { return v%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PickBy(data, fn)
	}
}

func BenchmarkOmitBy(b *testing.B) {
	data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	fn := func(k string, v int) bool { return v%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = OmitBy(data, fn)
	}
}
